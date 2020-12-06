package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type movement struct {
	xMove    int
	yMove    int
	treesHit int
}

func main() {
	treeMap := "forestMap.txt"

	movementMaps := []movement{
		movement{xMove: 1, yMove: 1},
		movement{xMove: 3, yMove: 1},
		movement{xMove: 5, yMove: 1},
		movement{xMove: 7, yMove: 1},
		movement{xMove: 1, yMove: 2},
	}

	checkTrees(treeMap, movementMaps)

	fmt.Println("Number of trees hit 1.5:", movementMaps[0].treesHit)
	fmt.Println("Number of trees hit 1.5:", movementMaps[1].treesHit)
	fmt.Println("Number of trees hit 1.5:", movementMaps[2].treesHit)
	fmt.Println("Number of trees hit 1.5:", movementMaps[3].treesHit)
	fmt.Println("Number of trees hit 1.5:", movementMaps[4].treesHit)

	combinedTotal := 1

	for i := 0; i < len(movementMaps); i++ {
		combinedTotal *= movementMaps[i].treesHit
	}

	fmt.Println("Final result of all trees hit multiplied:", combinedTotal)
}

func checkTrees(treeMap string, movementMaps []movement) {
	treeFile, err := os.Open(treeMap)

	if err != nil {
		log.Fatal(err)
		return
	}
	defer treeFile.Close()

	scanner := bufio.NewScanner(treeFile)
	endReached := false
	currentDepth := 0

	// Ensure that you're always looking at the first row initially
	if !scanner.Scan() {
		return
	}

	for {
		currentLine := scanner.Text()

		for i := 0; i < len(movementMaps); i++ {
			horizontalToCheck := movementMaps[i].xMove * (currentDepth / movementMaps[i].yMove) % len(currentLine)
			skipLine := currentDepth%movementMaps[i].yMove != 0

			if skipLine {
				continue
			}

			if currentLine[horizontalToCheck:horizontalToCheck+1] == "#" {
				movementMaps[i].treesHit++
			}
		}

		currentDepth++

		if !scanner.Scan() {
			endReached = true
			break
		}

		if endReached {
			break
		}
	}
}
