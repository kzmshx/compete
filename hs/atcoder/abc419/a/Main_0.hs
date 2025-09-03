-- stack script --resolver lts-23.14

main :: IO ()
main = do
  s <- getLine
  putStrLn $ case s of
    "red"   -> "SSS"
    "blue"  -> "FFF"
    "green" -> "MMM"
    _       -> "Unknown"
