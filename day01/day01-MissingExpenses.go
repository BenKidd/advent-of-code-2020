package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	expensesReport := "expenses.txt"

	expensesFile, err := os.Open(expensesReport)

	if err != nil {
		log.Fatal(err)
		return
	}
	defer expensesFile.Close()

	// Need to figure out how to "split" this so the length doesn't need to be hard-coded
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

	resultForTwo := calculateExpenses(expensesList[:], 2020, 0, 0, 2)
	fmt.Println("Result for 2 expenses:", resultForTwo)

	resultForThree := calculateExpenses(expensesList[:], 2020, 0, 0, 3)
	fmt.Println("Result for 3 expenses:", resultForThree)
}

func calculateExpenses(expensesList []int, target int, sum int, pos int, depth int) int {
	// No point continuing if you've blown past the target
	if sum > target {
		return 0
	}

	// Boundary check
	if pos >= len(expensesList) {
		return 0
	}

	// At depth of 0, we don't want to iterate any more, just return out after the final check
	if depth == 0 {
		// The return value will be used to multiply the matched values together.  Hence 0 and 1 to get a correct or zeroed out result
		if sum != target {
			return 0
		}

		return 1
	}

	for i := pos + 1; i < len(expensesList); i++ {
		tempSum := sum + expensesList[i]

		result := calculateExpenses(expensesList, target, tempSum, pos+1, depth-1)

		if result != 0 {
			return result * expensesList[i]
		}
	}

	return 0
}
