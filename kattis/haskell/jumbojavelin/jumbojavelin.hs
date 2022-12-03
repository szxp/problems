import System.IO

main = do
    z <- getLine
    let n = read z :: Int
    a <- add 0
    putStrLn $ show (a - n + 1)

add a = do
    done <- isEOF
    if done then
        return a
    else do
        z <- getLine
        let l = read z :: Int
        add (a + l)

