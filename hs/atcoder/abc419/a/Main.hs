-- stack script --resolver lts-23.14

main :: IO ()
main = getLine >>= putStrLn.solve

solve :: String -> String
solve "red" = "SSS"
solve "blue" = "FFF"
solve "green" = "MMM"
solve _ = "Unknown"
