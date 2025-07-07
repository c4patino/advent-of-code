module Day04 (part1, part2, run) where

import Data.Char (isDigit, isSpace)
import Data.List (dropWhileEnd)
import qualified Data.Set as Set

part1 :: [Card] -> Int
part1 cards = sum (map scoreCard cards)

part2 :: [Card] -> Int
part2 = const 42

run :: FilePath -> IO (Int, Int)
run filename = do
  contents <- readFile filename
  let cardLines = lines contents
      cards = map parseCardLine cardLines
  return (part1 cards, part2 cards)

data Card = Card {index :: Int, wins :: Set.Set Int, nums :: Set.Set Int}
  deriving (Show)

trim :: String -> String
trim = dropWhileEnd isSpace . dropWhile isSpace

parseCardLine :: String -> Card
parseCardLine line =
  case stripPrefix "Card" (trim line) of
    Nothing -> error $ "Line does not start with 'Card': " ++ line
    Just rest1 ->
      let rest2 = trim rest1
          (numStr, rest3) = span isDigit rest2
          rest4 = trim rest3
       in case rest4 of
            (':' : rest5) -> parseRest numStr (trim rest5)
            _ -> error $ "Expected ':' after card number in line: " ++ line

parseRest :: String -> String -> Card
parseRest numStr rest =
  case reads numStr of
    [(num, "")] ->
      let (left, rightWithBar) = break (== '|') rest
          right = drop 1 rightWithBar -- remove '|'
          winningSet = Set.fromList $ map read $ words (trim left)
          yourSet = Set.fromList $ map read $ words (trim right)
       in Card num winningSet yourSet
    _ -> error $ "Invalid card number: " ++ numStr

stripPrefix :: String -> String -> Maybe String
stripPrefix [] ys = Just ys
stripPrefix (x : xs) (y : ys)
  | x == y = stripPrefix xs ys
  | otherwise = Nothing
stripPrefix _ _ = Nothing

scoreCard :: Card -> Int
scoreCard (Card _ winningSet yourSet) =
  case matchCount of
    0 -> 0
    n -> 2 ^ (n - 1)
  where
    matchCount = Set.size $ Set.intersection winningSet yourSet
