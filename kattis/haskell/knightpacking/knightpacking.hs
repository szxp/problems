import System.IO

main = do
    z <- getLine
    let n = read z :: Int
    putStrLn $ win n
    
win n | n `mod` 2 == 1 = "first"
      | otherwise      = "second"

