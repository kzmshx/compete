-- stack script --resolver lts-23.14 --package bytestring --package containers --package heaps

import Control.Monad (replicateM)
import qualified Data.ByteString.Char8 as BS
import Data.Maybe (fromJust)
import qualified Data.Heap as H

readInt :: IO Int
readInt = readLn

readInts :: IO [Int]
readInts = map (fst . fromJust . BS.readInt) . BS.words <$> BS.getLine

data Query = Type1 Int | Type2

parseQuery :: [Int] -> Query
parseQuery [1, x] = Type1 x
parseQuery [2] = Type2
parseQuery _ = error "Invalid query format"

readQuery :: IO Query
readQuery = parseQuery <$> readInts

processQueries :: [Query] -> (Int, [Int])
processQueries queries = go [] queries []
  where
    go :: [Int] -> [Query] -> [Int] -> (Int, [Int])
    go state [] outputs =
      let minVal = if null state then maxBound else minimum state
      in (minVal, reverse outputs)
    go state (Type1 x : qs) outputs =
      go (state ++ [x]) qs outputs
    go state (Type2 : qs) outputs =
      if null state
        then go [] qs outputs
        else
          let minVal = minimum state
              deleteFirst _ [] = []
              deleteFirst x (y:ys)
                | x == y = ys
                | otherwise = y : deleteFirst x ys
              newState = deleteFirst minVal state
          in go newState qs (minVal : outputs)

main :: IO ()
main = do
  q <- readInt
  queries <- replicateM q readQuery
  let (minValue, outputs) = processQueries queries
  mapM_ print outputs
