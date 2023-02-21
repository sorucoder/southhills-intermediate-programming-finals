# Chapter 1, Exercise 9

## Prompt

Research current rates of monetary exchange. Draw a flowchart or write pseudo-
code to represent the logic of a program that allows the user to enter a number of
dollars and convert it to Euros and Japanese yen.

**Bonus:**

Use an API to provide real-time curency exchange rates.

## Usage

1. First, create an account on [ExchangeRate-API](https://www.exchangerate-api.com/) and get a valid API key.

2. Second, set the environment variable `EXCHANGERATE_API_KEY` to your API key:
```bash
export EXCHANGERATE_API_KEY=xxxxxxxxx
```

3. Finally, run the program:
```bash
go run main.go
```