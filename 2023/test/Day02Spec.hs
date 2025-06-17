module Day02Spec (spec) where

import Common.Tests (TestCase (..), mkPartSpec)
import Day02 (run)
import Test.Hspec

part1TestCases :: [TestCase Int]
part1TestCases = [TestCase "inputs/day02/part1.txt" 8]

part2TestCases :: [TestCase Int]
part2TestCases =
  []

spec :: Spec
spec = do
  mkPartSpec "Day 2/1 tests" fst run part1TestCases
  mkPartSpec "Day 2/2 tests" snd run part2TestCases
