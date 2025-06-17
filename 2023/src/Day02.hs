module Day02 (part1, part2, run) where

part1 :: [String] -> Int
part1 = const 42

part2 :: [String] -> Int
part2 = const 42

run :: FilePath -> IO (Int, Int)
run filename = do
  input <- readFile filename
  let parsed = lines input
  return (part1 parsed, part2 parsed)
