module Day03Spec (spec) where

import Common.Tests (TestCase (..), mkPartSpec)
import Day03 (run)
import Test.Hspec

part1TestCases :: [TestCase Int]
part1TestCases =
  [ TestCase "inputs/day03/example.txt" 4361,
    TestCase "inputs/day03/input.txt" 520019
  ]

part2TestCases :: [TestCase Int]
part2TestCases =
  [ TestCase "inputs/day03/example.txt" 467835,
    TestCase "inputs/day03/input.txt" 75519888
  ]

spec :: Spec
spec = do
  mkPartSpec "Day 3/1 tests" fst run part1TestCases
  mkPartSpec "Day 3/2 tests" snd run part2TestCases
