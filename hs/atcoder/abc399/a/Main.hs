(|>) :: a -> (a -> b) -> b
x |> f = f x

infixl 1 |>

main :: IO ()
main = do
  _ :: Int <- readLn
  s <- getLine
  t <- getLine
  zip s t |> filter (uncurry (/=)) |> length |> print
