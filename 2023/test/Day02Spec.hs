module Day02Spec (spec) where

import Common.Tests (TestCase (..), mkPartSpec)
import Day02 (run)
import Test.Hspec

part1TestCases :: [TestCase Int]
part1TestCases =
  [ TestCase "inputs/day02/example.txt" 8,
    TestCase "inputs/day02/input.txt" 2617
  ]

part2TestCases :: [TestCase Int]
part2TestCases =
  [ TestCase "inputs/day02/example.txt" 2286,
    TestCase "inputs/day02/input.txt" 59795
  ]

spec :: Spec
spec = do
  mkPartSpec "Day 2/1 tests" fst run part1TestCases
  mkPartSpec "Day 2/2 tests" snd run part2TestCases
