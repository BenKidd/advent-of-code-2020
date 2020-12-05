package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	treeMap := "forestMap.txt"

	// Right 1, down 1
	treesHit1 := checkTrees(treeMap, 1, 1)

	// Right 2, down 1
	treesHit2 := checkTrees(treeMap, 3, 1)

	// Right 5, down 1.
	treesHit3 := checkTrees(treeMap, 5, 1)

	// Right 7, down 1.
	treesHit4 := checkTrees(treeMap, 7, 1)

	// Right 1, down 2.
	treesHit5 := checkTrees(treeMap, 1, 2)

	// could replace this with an array of structs which holds the horizontal/vertical movement
	// for each, and track the number of collisions made along each route in another array.
	// Multiply the final result before returning.  Keeps it independent AND efficient

	fmt.Println("Number of trees hit 1:", treesHit1)
	fmt.Println("Number of trees hit 2:", treesHit2)
	fmt.Println("Number of trees hit 3:", treesHit3)
	fmt.Println("Number of trees hit 4:", treesHit4)
	fmt.Println("Number of trees hit 5:", treesHit5)

	combinedTotal := treesHit1 * treesHit2 * treesHit3 * treesHit4 * treesHit5
	fmt.Println("Final result of all trees hit multiplied:", combinedTotal)
}

// For this function I had the option of either making it work for any values input, or
// preventing multiple filereads by checking for collisions for each 5 cases on every
// line read.  I opted for keeping it keeping it open, but realise it could be refactored
// the other way to make it faster and more efficient
func checkTrees(treeMap string, horizontalMovement int, verticalMovement int) int {
	treeFile, err := os.Open(treeMap)

	if err != nil {
		log.Fatal(err)
		return -1
	}
	defer treeFile.Close()

	scanner := bufio.NewScanner(treeFile)
	endReached := false
	horizontalPosition := 0
	treesHit := 0

	// Ensure that you're always looking at the first row initially
	if !scanner.Scan() {
		return -1
	}

	for {
		currentLine := scanner.Text()

		if currentLine[horizontalPosition:horizontalPosition+1] == "#" {
			treesHit++
		}

		horizontalPosition += horizontalMovement

		if horizontalPosition >= len(currentLine) {
			horizontalPosition = horizontalPosition - len(currentLine)
		}

		for i := 0; i < verticalMovement; i++ {
			if !scanner.Scan() {
				endReached = true
				break
			}
		}

		if endReached {
			break
		}
	}

	return treesHit
}
