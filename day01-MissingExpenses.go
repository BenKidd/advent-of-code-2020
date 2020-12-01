package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello World!")

	expensesReport := "expenses.txt"

	expensesFile, err := os.Open(expensesReport)

	if err != nil {
		log.Fatal(err)
		return
	}

	var expensesList [200]int
	scanner := bufio.NewScanner(expensesFile)

	for i := 0; i < 200; i++ {
		scanner.Scan()
		var newText = scanner.Text()
		expense, err := strconv.Atoi(newText)

		if err != nil {
			log.Fatal(err)
			return
		}

		expensesList[i] = expense
	}

	for i := 0; i < 200; i++ {
		for j := i + 1; j < 200; j++ {
			var result = expensesList[i] + expensesList[j]

			if result == 2020 {
				fmt.Println("Result 1:", expensesList[i])
				fmt.Println("Result 2:", expensesList[j])

				var finalExpense = expensesList[i] * expensesList[j]

				fmt.Println("Final expense 1 is:", finalExpense)
			}
		}
	}

	for i := 0; i < 200; i++ {
		for j := i + 1; j < 200; j++ {
			for k := j + 1; k < 200; k++ {
				var result = expensesList[i] + expensesList[j] + expensesList[k]

				if result == 2020 {
					fmt.Println("Result 1:", expensesList[i])
					fmt.Println("Result 2:", expensesList[j])
					fmt.Println("Result 3:", expensesList[k])

					var finalExpense = expensesList[i] * expensesList[j] * expensesList[k]

					fmt.Println("Final expense 2 is:", finalExpense)
				}
			}
		}
	}
}
