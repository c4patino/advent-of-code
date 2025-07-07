module Day03 (part1, part2, run) where

import Data.Char (isDigit)
import Data.Maybe (mapMaybe)
import qualified Data.Set as Set

part1 :: Grid -> Int
part1 grid =
  sum . map (read . digits) . filter (numberHasAdjacentSymbol grid) $ extractNumbers grid

part2 :: Grid -> Int
part2 grid =
  let numbers = extractNumbers grid
      gears = extractGears grid
      validGears = filter (\g -> length (findAdjacentNumbersToGear numbers g) == 2) gears
      gearRatios = [product (map (read . digits) (findAdjacentNumbersToGear numbers g)) | g <- validGears]
   in sum gearRatios

run :: FilePath -> IO (Int, Int)
run filename = do
  input <- readFile filename
  let grid = mkGrid input
  return (part1 grid, part2 grid)

newtype Grid = Grid [[Char]]
  deriving (Show, Eq)

data Number = Number {digits :: String, positions :: [(Int, Int)]}
  deriving (Show, Eq, Ord)

mkGrid :: String -> Grid
mkGrid = Grid . lines

neighbors :: [(Int, Int)]
neighbors = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]

getCharAt :: Grid -> (Int, Int) -> Maybe Char
getCharAt (Grid rows) (r, c) =
  if r >= 0 && r < length rows && c >= 0 && c < length (rows !! r)
    then Just ((rows !! r) !! c)
    else Nothing

isSymbol :: Char -> Bool
isSymbol c = not (isDigit c) && c /= '.'

extractNumbers :: Grid -> [Number]
extractNumbers (Grid rows) = concatMap extractRowNumbers (zip [0 ..] rows)
  where
    extractRowNumbers (r, row) = go 0 row []
      where
        go _ [] acc = reverse acc
        go c xs acc =
          case span isDigit xs of
            ("", _ : rest) -> go (c + 1) rest acc
            (num, rest) ->
              let len = length num
                  pos = [(r, c + i) | i <- [0 .. len - 1]]
               in go (c + len) rest (Number num pos : acc)

extractGears :: Grid -> [(Int, Int)]
extractGears (Grid rows) =
  [ (r, c)
    | r <- [0 .. length rows - 1],
      c <- [0 .. length (rows !! r) - 1],
      getCharAt (Grid rows) (r, c) == Just '*'
  ]

findNumberAtPosition :: [Number] -> (Int, Int) -> Maybe Number
findNumberAtPosition numbers pos =
  let filtered = filter (\n -> pos `elem` positions n) numbers
   in case filtered of
        (x : _) -> Just x
        [] -> Nothing

findAdjacentNumbersToGear :: [Number] -> (Int, Int) -> [Number]
findAdjacentNumbersToGear numbers gearPos =
  let adjPos = adjacentPositions gearPos
      adjacentNums = mapMaybe (findNumberAtPosition numbers) adjPos
   in Set.toList (Set.fromList adjacentNums)

numberHasAdjacentSymbol :: Grid -> Number -> Bool
numberHasAdjacentSymbol grid (Number _ pos) = any (hasSymbolAdjacentToPosition grid) pos

hasSymbolAdjacentToPosition :: Grid -> (Int, Int) -> Bool
hasSymbolAdjacentToPosition grid (r, c) =
  any
    (maybe False isSymbol . getCharAt grid)
    (adjacentPositions (r, c))

adjacentPositions :: (Int, Int) -> [(Int, Int)]
adjacentPositions (r, c) = [(r + dr, c + dc) | (dr, dc) <- neighbors]
