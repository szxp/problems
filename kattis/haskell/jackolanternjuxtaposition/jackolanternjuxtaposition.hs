import System.IO

main = do
    z <- getLine
    let w = words z
    let n = read $ head w :: Int
    let m = read $ w !! 1 :: Int
    let o = read $ last w :: Int
    putStrLn $ show $ n * m * o

