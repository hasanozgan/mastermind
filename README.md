## Mastermind game

Mastermind is a board game. The first player invents a code consisting of 4 balls of six different colors and the second player tries to break it. For the task simplicity, we'll use capital letters from A to F instead of colors. Note that the letters might be repeated in the code.

The second player makes consecutive attempts to guess the code: he suggests the new combination of letters on each try until his guess is correct. The first player evaluates each combination by comparing it to her secret and says how many letters are guessed correctly.

First, she says the number of letters which are guessed right with their positions. Then she compares the remaining letters and says how many of them are guessed properly: the letters are present both in the guess and in the secret, but stay not in their correct positions. If there are duplicate letters in the guess, they all are counted as "wrongPosition" letters if they correspond to the same number of duplicated letters in the hidden code.

