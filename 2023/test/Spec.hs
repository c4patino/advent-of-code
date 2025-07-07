module Main (main) where

import qualified Day01Spec
import qualified Day02Spec
import qualified Day03Spec
import qualified Day04Spec
import Test.Hspec

main :: IO ()
main = hspec $ do
  Day01Spec.spec
  Day02Spec.spec
  Day03Spec.spec
  Day04Spec.spec
