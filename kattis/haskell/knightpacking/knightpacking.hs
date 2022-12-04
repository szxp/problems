import System.IO

main = do
    z <- getLine
    let n = read z :: Int
    putStrLn $ win n
    
win n | odd n = "first"
      | otherwise = "second"
