-- Pipe operator
(|>) :: a -> (a -> b) -> b
x |> f = f x

infixl 1 |>

main :: IO ()
main = do
  print "Hello, World!"
