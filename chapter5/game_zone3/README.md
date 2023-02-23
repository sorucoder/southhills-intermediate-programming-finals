# Chapter 5, Game Zone 3

## Prompt

Write the code for the dice game Pig, in which a player can compete with the
computer. The object of the game is to be the first to score 100 points. The user
and computer take turns “rolling” a pair of dice following these rules:

- On a turn, each player rolls two dice. If no 1 appears, the dice values are added
to a running total for the turn, and the player can choose whether to roll again
or pass the turn to the other player. When a player passes, the accumulated
turn total is added to the player’s game total.
- If a 1 appears on one of the dice, the player’s turn total becomes 0; in other
words, nothing more is added to the player’s game total for that turn, and it
becomes the other player’s turn.
- If a 1 appears on both of the dice, not only is the player’s turn over, but the
player’s entire accumulated total is reset to 0.
- When the computer does not roll a 1 and can choose whether to roll again,
generate a random value from 1 to 2. The computer will then decide to
continue when the value is 1 and decide to quit and pass the turn to the player
when the value is not 1

## Usage

Run the program:
```bash
go run main.go