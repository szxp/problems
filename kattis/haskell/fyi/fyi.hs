import System.IO

main = do
    z <- getLine
    putStrLn $ op z

op ('5':'5':'5':_) = "1"
op _ = "0"

