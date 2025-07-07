module Day04 (part1, part2, run) where

import Data.IntMap.Strict (IntMap)
import qualified Data.IntMap.Strict as IntMap
import Data.List (stripPrefix)
import Data.Maybe (fromMaybe)
import qualified Data.Set as Set

part1 :: [Card] -> Int
part1 cards = sum $ scoreCard <$> cards

part2 :: [Card] -> Int
part2 cards = sum finalCounts
  where
    cardMap = buildCardMap cards
    maxIdx = getMaxIndex cards
    initial = initializeCounts maxIdx
    finalCounts = propagateCards cardMap initial

run :: FilePath -> IO (Int, Int)
run filename = do
  contents <- readFile filename
  let cardLines = lines contents
      cards = map parseCardLine cardLines
  return (part1 cards, part2 cards)

data Card = Card {index :: Int, wins :: Set.Set Int, nums :: Set.Set Int}
  deriving (Show)

parseCardLine :: String -> Card
parseCardLine line = Card idx winSet numSet
  where
    withoutCard = fromMaybe line (stripPrefix "Card" line)
    trimmed = dropWhile (== ' ') withoutCard
    (indexPart, rest) = break (== ':') trimmed
    idx = read indexPart

    withoutColon = drop 1 rest
    (winPart, numPart) = break (== '|') withoutColon
    winSet = Set.fromList $ map read $ words winPart
    numSet = Set.fromList $ map read $ words $ drop 1 numPart

scoreCard :: Card -> Int
scoreCard card =
  case getMatches card of
    0 -> 0
    n -> 2 ^ (n - 1)

buildCardMap :: [Card] -> IntMap Card
buildCardMap cards = IntMap.fromList [(index c, c) | c <- cards]

getMaxIndex :: [Card] -> Int
getMaxIndex = maximum . map index

initializeCounts :: Int -> IntMap Int
initializeCounts maxIdx = IntMap.fromList [(i, 1) | i <- [1 .. maxIdx]]

processCard :: IntMap Card -> IntMap Int -> Int -> IntMap Int
processCard cardMap counts i =
  case IntMap.lookup i cardMap of
    Nothing -> counts
    Just card ->
      let cardCount = IntMap.findWithDefault 0 i counts
          matchCount = getMatches card
          targetIndices = take matchCount [i + 1 .. IntMap.size counts]
       in foldl (\m j -> IntMap.insertWith (+) j cardCount m) counts targetIndices

getMatches :: Card -> Int
getMatches card = Set.size $ Set.intersection (wins card) (nums card)

propagateCards :: IntMap Card -> IntMap Int -> IntMap Int
propagateCards cardMap initial =
  foldl (processCard cardMap) initial [1 .. IntMap.size initial]
