import System.IO

main = do
    z <- getLine
    let w = words z
    let n = read $ head w :: Int
    let m = read $ last w :: Int
    putStrLn $ show (cmp n m)
    where

cmp :: Int -> Int -> Int
cmp n m
    | n > m = 1
    | otherwise = 0


