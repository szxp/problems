import System.IO

main = do
    _ <- getLine
    echo 1

echo i = do done <- isEOF
            if done
               then return ()
               else do z <- getLine
                       if i `mod` 2 == 1
                          then do putStrLn z
                                  echo (i+1)
                          else echo (i+1)

