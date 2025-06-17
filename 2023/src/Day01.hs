module Day01 (part1, part2, run) where

import Data.Char (isDigit)
import Data.List (isPrefixOf)
import Data.Maybe (maybeToList)

part1 :: [String] -> Int
part1 = sum . map (extractValue . filter isDigit)

part2 :: [String] -> Int
part2 = sum . map (extractValue . collectDigits)

run :: FilePath -> IO (Int, Int)
run filename =
  readFile filename >>= \contents ->
    let parsed = lines contents
     in return (part1 parsed, part2 parsed)

collectDigits :: String -> String
collectDigits [] = []
collectDigits s@(c : cs)
  | isDigit c = c : collectDigits cs
  | otherwise = maybeToList (spelledDigitToChar s) ++ collectDigits cs

extractValue :: String -> Int
extractValue [] = 0
extractValue s@(d : _) = read [d, last s]

spelledDigitToChar :: String -> Maybe Char
spelledDigitToChar s = lookup True matches
  where
    matches =
      [ ("one" `isPrefixOf` s, '1'),
        ("two" `isPrefixOf` s, '2'),
        ("three" `isPrefixOf` s, '3'),
        ("four" `isPrefixOf` s, '4'),
        ("five" `isPrefixOf` s, '5'),
        ("six" `isPrefixOf` s, '6'),
        ("seven" `isPrefixOf` s, '7'),
        ("eight" `isPrefixOf` s, '8'),
        ("nine" `isPrefixOf` s, '9')
      ]
