module Common.Tests where

import Test.Hspec

-- | Represents a single test case with an input file and its expected output.
data TestCase a = TestCase
  { -- | Path to the input file for the test case.
    filename :: FilePath,
    -- | Expected output value for the test case, of any type 'a'.
    expectedOutput :: a
  }

-- | Run a suite of test cases by applying a runner function to input files,
--   then projecting a part of the result for comparison against expected outputs.
--
-- Parameters:
--
-- * @label@: parent label for the specification.
-- * @projector@: function to extract a single value from a tuple (e.g., part 1 or part 2).
-- * @runner@: function under test, which takes a filename and returns an IO action producing a tuple of outputs.
-- * @testCases@: list of test cases with input files and expected results.
--
-- Returns an Hspec 'Spec' running all test cases with descriptive labels and assertions.
mkPartSpec ::
  (Eq a, Show a) =>
  -- | parent label for the specification
  String ->
  -- | projector from a tuple to the value to test
  ((a, a) -> a) ->
  -- | function under test (runner)
  (FilePath -> IO (a, a)) ->
  -- | list of test cases with input files and expected outputs
  [TestCase a] ->
  -- | resulting Hspec Spec with all tests
  Spec
mkPartSpec label projector runner testCases =
  describe label $
    mapM_
      ( \(TestCase file expected) ->
          it ("run " ++ file ++ " returns " ++ show expected) $ do
            result <- runner file
            projector result `shouldBe` expected
      )
      testCases
