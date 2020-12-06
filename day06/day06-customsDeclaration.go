package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var currentForm map[string]int
	currentForm = make(map[string]int)

	formFile, err := os.Open("formData.txt")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer formFile.Close()

	formData := bufio.NewScanner(formFile)
	formTotal := 0
	expandedFormTotal := 0
	peopleInGroup := 0
	eofReached := false

	for {
		if !formData.Scan() {
			eofReached = true
		}

		if eofReached || len(formData.Text()) == 0 {
			formTotal += len(currentForm)

			for _, numAnswers := range currentForm {
				if numAnswers == peopleInGroup {
					expandedFormTotal++
				}
			}

			currentForm = make(map[string]int)
			peopleInGroup = 0

			if eofReached {
				break
			} else {
				continue
			}
		}

		formLine := formData.Text()

		peopleInGroup++

		for i := 0; i < len(formLine); i++ {
			answeredQuestion := formLine[i : i+1]
			currentForm[answeredQuestion]++
		}
	}

	fmt.Println("Sum of form data:", formTotal)
	fmt.Println("Sum of expanded form data:", expandedFormTotal)
}
