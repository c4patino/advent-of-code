module Main (main) where

import qualified Day01Spec
import qualified Day02Spec
import Test.Hspec

main :: IO ()
main = hspec $ do
  Day01Spec.spec
  Day02Spec.spec
