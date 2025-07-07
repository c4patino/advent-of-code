module Main where

import qualified Day01
import qualified Day02
import qualified Day03
import System.Console.GetOpt
import System.Environment (getArgs)
import System.Exit (exitFailure)

data Options = Options {optFile :: Maybe FilePath, optHelp :: Bool} deriving (Show)

defaultOptions :: Options
defaultOptions = Options {optFile = Nothing, optHelp = False}

options :: [OptDescr (Options -> Options)]
options =
  [ Option
      ['f']
      ["file"]
      (ReqArg (\f opts -> opts {optFile = Just f}) "FILE")
      "input file to use",
    Option
      ['h']
      ["help"]
      (NoArg (\opts -> opts {optHelp = True}))
      "show help"
  ]

runDay :: Int -> FilePath -> IO (String, String)
runDay 1 filename = do
  (a, b) <- Day01.run filename
  return (show a, show b)
runDay 2 filename = do
  (a, b) <- Day02.run filename
  return (show a, show b)
runDay 3 filename = do
  (a, b) <- Day03.run filename
  return (show a, show b)
runDay _ _ = return ("Not implemented", "Not implemented")

defaultFilename :: Int -> FilePath
defaultFilename day = "inputs/day" ++ (if day < 10 then "0" else "") ++ show day ++ "/input.txt"

main :: IO ()
main = do
  args <- getArgs
  case getOpt Permute options args of
    (opts, [dayStr], []) -> do
      let day = read dayStr
          finalOpts = foldl (flip id) defaultOptions opts
          filename = case optFile finalOpts of
            Just f -> f
            Nothing -> defaultFilename day
      (part1Answer, part2Answer) <- runDay day filename
      putStrLn $ "Part 1: " ++ part1Answer
      putStrLn $ "Part 2: " ++ part2Answer
    (_, [], []) -> do
      putStrLn "Error: Day number is required"
      putStrLn $ usageInfo "Usage: aoc2023 <day> [OPTIONS]" options
      exitFailure
    (_, _, errs) -> do
      putStrLn $ concat errs
      putStrLn $ usageInfo "Usage: aoc2023 <day> [OPTIONS]" options
      exitFailure
