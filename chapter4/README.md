# Chapter 4, Exercise 11

## Prompt

Amanda Cho, a supervisor in a retail clothing store, wants to acknowledge high-
achieving salespeople. Write code for a program with the following criteria:

- Continuously accepts each salesperson’s first and last names, the
number of shifts worked in a month, number of transactions completed this
month, and the dollar value of those transactions.
- Display each salesperson’s name with a productivity score, which is computed by first dividing dollars by
transactions and dividing the result by shifts worked. Display three asterisks after
the productivity score if it is 50 or higher.
- Accepts each salesperson’s data and displays the name and a
bonus amount. The bonuses will be distributed as follows:
	- If the productivity score is 30 or less, the bonus is $25.
	- If the productivity score is 31 or more and less than 80, the bonus is $50.
	- If the productivity score is 80 or more and less than 200, the bonus is $100.
	- If the productivity score is 200 or higher, the bonus is $200
Also, it should be efficient accounting for the fact that sixty percent of employees have a productivity score greater than 200.

## Usage

Run the program:
```bash
go run main.go
```