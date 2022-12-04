import System.IO

main = do
    l0 <- getLine
    let ws0 = words l0
    let g = read $ head ws0 :: Int
    let t = read $ ws0 !! 1 :: Int
    l1 <- getLine
    let items = add 0 (words l1)
    putStrLn $ show $ round $ ((fromIntegral (g-t)) * 0.9) - fromIntegral items

add :: Int -> [String] -> Int
add a [] = a
add a (w:ws) = add (a + n) ws
    where n = read w :: Int

