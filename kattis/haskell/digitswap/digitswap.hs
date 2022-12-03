-- Digit Swap
-- Ann Britt-Caroline has a safe with a 2-digit code. Occasionally, she types in the code too fast, accidentally swapping the positions of the two digits. She has asked if you could program her safe to check if not only the 2-digit code she entered was correct, but if the code with the two digits swapped is also correct.

-- The first step in programming her safe in this manner is to write a program to swap the digits of a 2-digit code.

-- Input
-- The input consists of a single 2-digit code with only digits  to , without any space between the digits.

-- Output
-- Output a single line with the swapped 2-digit code, without any space between the digits.

import System.IO

main = do
    z <- getLine
    putStrLn $ reverse z

