import System.IO
import Data.List

main = do
    z <- getLine
    putStrLn $ can z

can s | isSuffixOf "eh?" s = "Canadian!"
      | otherwise          = "Imposter!"

