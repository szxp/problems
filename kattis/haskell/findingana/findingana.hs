import System.IO

main = do
    z <- getLine
    putStrLn $ dropWhile (/='a') z

