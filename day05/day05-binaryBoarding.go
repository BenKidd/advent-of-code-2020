package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	boardingPassFile, err := os.Open("boardingPasses.txt")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer boardingPassFile.Close()

	scanner := bufio.NewScanner(boardingPassFile)
	maxSeatID := -1

	var seatIDs [868]int
	passNumber := 0

	for {
		if !scanner.Scan() {
			break
		}

		var pass = scanner.Text()

		row := findRow(pass[:7])
		column := findColumn(pass[7:])
		seatID := findSeatID(row, column)

		if seatID < 0 {
			log.Fatal("Critical error with seat ID of:", seatID)
		}

		if seatID > maxSeatID {
			maxSeatID = seatID
		}

		seatIDs[passNumber] = seatID
		passNumber++
	}

	fmt.Println("Max Seat ID:", maxSeatID)

	sort.Ints(seatIDs[:])

	mySeatID := -1

	for i := 0; i < len(seatIDs)-1; i++ {
		if seatIDs[i+1]-seatIDs[i] == 2 {
			mySeatID = seatIDs[i] + 1
			break
		}
	}

	fmt.Println("My seat ID:", mySeatID)
}

func findRow(rowData string) int {
	minValue := 1
	maxValue := 128
	change := maxValue / 2
	letter := ""

	for i := 0; i < len(rowData); i++ {
		letter = rowData[i : i+1]

		if letter == "F" {
			maxValue -= change
		} else if letter == "B" {
			minValue += change
		} else {
			log.Fatal("Could not match:", letter)
			return -1
		}

		change /= 2
	}

	return minValue - 1
}

func findColumn(columnData string) int {
	minValue := 1
	maxValue := 8
	change := maxValue / 2
	letter := ""

	for i := 0; i < len(columnData); i++ {
		letter = columnData[i : i+1]

		if letter == "L" {
			maxValue -= change
		} else if letter == "R" {
			minValue += change
		} else {
			log.Fatal("Could not match:", letter)
			return -1
		}

		change /= 2
	}

	return minValue - 1
}

func findSeatID(row int, column int) int {
	seatID := row * 8
	seatID += column

	return seatID
}
