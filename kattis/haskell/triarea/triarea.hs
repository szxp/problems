import System.IO

main = do
    z <- getLine
    let w = words z
    let n = read $ head w :: Double
    let m = read $ last w :: Double
    putStrLn $ show $ n * m / 2.0

