module Day04Spec (spec) where

import Common.Tests (TestCase (..), mkPartSpec)
import Day04 (run)
import Test.Hspec

part1TestCases :: [TestCase Int]
part1TestCases =
  [ TestCase "inputs/day04/example.txt" 13,
    TestCase "inputs/day04/input.txt" 25571
  ]

part2TestCases :: [TestCase Int]
part2TestCases =
  [ TestCase "inputs/day04/example.txt" 30,
    TestCase "inputs/day04/input.txt" 8805731
  ]

spec :: Spec
spec = do
  mkPartSpec "Day 4/1 tests" fst run part1TestCases
  mkPartSpec "Day 4/2 tests" snd run part2TestCases
