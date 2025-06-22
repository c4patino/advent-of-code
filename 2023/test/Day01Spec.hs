module Day01Spec (spec) where

import Common.Tests (TestCase (..), mkPartSpec)
import Day01 (run)
import Test.Hspec

part1TestCases :: [TestCase Int]
part1TestCases =
  [ TestCase "inputs/day01/part1.txt" 142,
    TestCase "inputs/day01/input.txt" 55108
  ]

part2TestCases :: [TestCase Int]
part2TestCases =
  [ TestCase "inputs/day01/part2.txt" 281,
    TestCase "inputs/day01/input.txt" 56324
  ]

spec :: Spec
spec = do
  mkPartSpec "Day 1/1 tests" fst run part1TestCases
  mkPartSpec "Day 1/2 tests" snd run part2TestCases
