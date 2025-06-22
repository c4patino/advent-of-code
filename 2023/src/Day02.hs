module Day02 (part1, part2, run) where

import Text.Read (readMaybe)

data Game = Game {id :: Int, draw :: Draw} deriving (Show)

data Draw = Draw {red :: Int, green :: Int, blue :: Int} deriving (Show)

part1 :: [Game] -> Int
part1 = sum . map gid . filter possible
  where
    gid (Game i _) = i

part2 :: [Game] -> Int
part2 = sum . map power
  where
    power (Game _ (Draw r g b)) = r * g * b

run :: FilePath -> IO (Int, Int)
run filename = do
  input <- readFile filename
  let games = parseGames input
  return (part1 games, part2 games)

instance Semigroup Draw where
  (Draw a b c) <> (Draw x y z) = Draw (max a x) (max b y) (max c z)

instance Monoid Draw where
  mempty = Draw 0 0 0

possible :: Game -> Bool
possible = valid . draw
  where
    valid (Draw r g b) = r <= 12 && g <= 13 && b <= 14

parseGames :: String -> [Game]
parseGames input = map parseGame $ lines input

parseGame :: String -> Game
parseGame s =
  let (header, rest) = break (== ':') s
      gameId = read (drop 5 header) :: Int
      drawsStr = drop 2 rest
      draws = parseDraws drawsStr
   in Game gameId draws

parseDraws :: String -> Draw
parseDraws s =
  let drawParts = splitOn ';' s
   in foldl (<>) mempty (map parseDraw drawParts)

parseDraw :: String -> Draw
parseDraw s =
  let counts = splitOn ',' s
   in foldl (<>) mempty (map parseCount counts)

parseCount :: String -> Draw
parseCount s =
  case words s of
    [nStr, color] ->
      case (readMaybe nStr :: Maybe Int) of
        Just i -> case color of
          "red" -> Draw i 0 0
          "green" -> Draw 0 i 0
          "blue" -> Draw 0 0 i
          _ -> mempty
        Nothing -> mempty
    _ -> mempty

splitOn :: Char -> String -> [String]
splitOn _ [] = []
splitOn delim str =
  let (first, rest) = break (== delim) str
   in first : case rest of
        [] -> []
        (_ : rest') -> splitOn delim rest'
