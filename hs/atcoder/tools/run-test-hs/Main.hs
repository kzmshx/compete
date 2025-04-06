import Control.Monad (forM_)
import Data.List (intercalate, isPrefixOf, isSuffixOf, sort)
import Data.Text (unpack)
import Data.Text.IO (readFile)
import System.Directory (listDirectory)
import System.Environment (getArgs)
import System.Exit (exitFailure)
import System.FilePath (takeDirectory, (</>))
import System.Process (readProcess, system)
import Text.XHtml (action)

-- Pipe operator
(|>) :: a -> (a -> b) -> b
x |> f = f x

infixl 1 |>

-- Check if a file has a given prefix and suffix
isAffix :: String -> String -> String -> Bool
isAffix pre suf f = pre `isPrefixOf` f && suf `isSuffixOf` f

-- Create a path to a test case file
joinPath :: String -> String -> String
joinPath dir file = dir ++ "/" ++ file

-- Find test case files
findTestCaseFiles :: String -> String -> IO [(String, String)]
findTestCaseFiles contest problem = do
  files <- listDirectory (contest </> problem)
  let inputFiles = filter (isAffix "testcase_" "_input.txt") files |> map (joinPath (contest </> problem)) |> sort
  let outputFiles = map (\input -> take (length input - 10) input ++ "_output.txt") inputFiles
  return $ zip inputFiles outputFiles

-- テストケースを実行
runTestCase :: String -> String -> String -> String -> IO (Bool, String, String, String)
runTestCase contest problem inputFile expectedFile = do
  input <- Data.Text.unpack <$> Data.Text.IO.readFile inputFile
  expected <- Data.Text.unpack <$> Data.Text.IO.readFile expectedFile
  let buildCommand = "cd " ++ contest ++ " && cabal v2-build --offline " ++ problem
  _ <- readProcess "sh" ["-c", buildCommand] ""
  let executable = "cd " ++ contest ++ " && $(cabal list-bin " ++ problem ++ ")"
  actual <- readProcess "sh" ["-c", executable] input
  return (actual == expected, input, expected, actual)

-- テスト結果を表示
printResult :: String -> (Bool, String, String, String) -> IO ()
printResult testCaseFile (result, input, expected, actual) = do
  putStrLn "----------------------------------------"
  putStrLn $ testCaseFile ++ "\n"
  putStrLn $ intercalate "\n" ["[Input]", input]
  putStrLn $ intercalate "\n" ["[Expected]", expected]
  putStrLn $ intercalate "\n" ["[Got]", actual]
  putStrLn $ if result then "✅ PASSED" else "❌ FAILED"
  putStrLn "----------------------------------------"

main :: IO ()
main = do
  args <- getArgs
  case args of
    [contest, problem] -> do
      testCaseFiles <- findTestCaseFiles contest problem
      forM_ testCaseFiles $ \(inputFile, outputFile) -> do
        result <- runTestCase contest problem inputFile outputFile
        printResult inputFile result
    _ -> do
      putStrLn "Usage: ./tools/run-test <contest> <problem>"
      exitFailure
