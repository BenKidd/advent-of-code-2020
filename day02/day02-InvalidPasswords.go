package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read in password list
	passwordsReport := "passwords.txt"

	passwordsFile, err := os.Open(passwordsReport)

	if err != nil {
		log.Fatal(err)
		return
	}
	defer passwordsFile.Close()

	scanner := bufio.NewScanner(passwordsFile)

	validSledPasswords := 0
	validTobogganPasswords := 0

	for i := 0; i < 1000; i++ {
		scanner.Scan()
		var password = scanner.Text()

		if validateSledPassword(password) {
			validSledPasswords++
		}

		if validateTobogganPassword(password) {
			validTobogganPasswords++
		}
	}

	fmt.Println("Number of valid Sled passwords:", validSledPasswords)
	fmt.Println("Number of valid Toboggon passwords:", validTobogganPasswords)
}

func validateSledPassword(passwordRaw string) bool {
	// parse each password into the following
	// - minimum number of characters
	// - maximum number of characters
	// - validChar = character to search for
	// - password
	separatorIndex := strings.Index(passwordRaw, ":")

	policy := passwordRaw[:separatorIndex]

	policyDashIndex := strings.Index(policy, "-")
	validChar := policy[separatorIndex-1 : separatorIndex]

	minCharCount := parseValueFromString(policy, 0, policyDashIndex)
	maxCharCount := parseValueFromString(policy, policyDashIndex+1, separatorIndex-2)

	password := passwordRaw[separatorIndex+1:]

	validCharCount := strings.Count(password, validChar)

	if validCharCount >= minCharCount && validCharCount <= maxCharCount {
		return true
	}

	return false
}

func validateTobogganPassword(passwordRaw string) bool {
	// parse each password into the following
	// - first Index for policy
	// - second Index for policy
	// - validChar = character to search for
	// - password
	separatorIndex := strings.Index(passwordRaw, ":")

	policy := passwordRaw[:separatorIndex]

	policyDashIndex := strings.Index(policy, "-")
	validChar := policy[separatorIndex-1 : separatorIndex]

	firstIndex := parseValueFromString(policy, 0, policyDashIndex)
	secondIndex := parseValueFromString(policy, policyDashIndex+1, separatorIndex-2)

	password := passwordRaw[separatorIndex+1:]

	firstIndexMatch := password[firstIndex:firstIndex+1] == validChar
	secondIndexMatch := password[secondIndex:secondIndex+1] == validChar

	if firstIndexMatch != secondIndexMatch {
		return true
	}

	return false
}

func parseValueFromString(textToSearch string, indexStart int, indexEnd int) int {
	minCharCount, err := strconv.Atoi(textToSearch[indexStart:indexEnd])

	if err != nil {
		log.Fatal(err)
		return -1
	}
	return minCharCount
}
