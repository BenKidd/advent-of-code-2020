package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instructionSet struct {
	instruction string
	value       int
	visited     bool
}

type checkpoint struct {
	position         int
	accumulatedValue int
}

type accReturn struct {
	foundEnd         bool
	accumulatedValue int
}

func main() {
	inputFile := "bootInstructions.txt"
	//inputFile = "sampleBootInstructions.txt"

	instructions := processInputFile(inputFile)

	accumulatedValue := 0
	toggleableInstructions := findInstructionsToToggle(instructions)

	for i := 0; i < len(toggleableInstructions); i++ {
		toggleInstruction(instructions, toggleableInstructions[i].position)
		instructions[toggleableInstructions[i].position].visited = false

		accumulatedValue = 0

		resetVisitedInstructions(instructions)

		result := findValidAccumulator2(instructions, toggleableInstructions[i])

		if result.foundEnd {
			accumulatedValue = result.accumulatedValue
			break
		}

		toggleInstruction(instructions, toggleableInstructions[i].position)
	}

	fmt.Println("Accumulated total:", accumulatedValue)
}

func processInputFile(inputFile string) []instructionSet {
	fileInput, err := os.Open(inputFile)

	instructions := make([]instructionSet, 0)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	scanner := bufio.NewScanner(fileInput)

	for {
		if !scanner.Scan() {
			break
		}

		instructionData := strings.Split(scanner.Text(), " ")

		var newInstruction instructionSet
		newInstruction.instruction = instructionData[0]
		value, err := strconv.Atoi(instructionData[1])

		if err != nil {
			log.Fatal(err)
			return nil
		}

		newInstruction.value = value

		instructions = append(instructions, newInstruction)
	}

	return instructions
}

func findValidAccumulator2(instructions []instructionSet, startingPoint checkpoint) accReturn {
	curPos := startingPoint.position
	oldPos := curPos

	var resultSet accReturn
	resultSet.accumulatedValue = startingPoint.accumulatedValue

	for {
		if curPos >= len(instructions) {
			resultSet.foundEnd = true
			break
		}

		curInstruction := instructions[curPos]

		if curInstruction.visited {
			break
		}

		curInstruction.visited = true
		oldPos = curPos

		if curInstruction.instruction == "nop" {
			curPos++
		}

		if curInstruction.instruction == "acc" {
			resultSet.accumulatedValue += curInstruction.value
			curPos++
		}

		if curInstruction.instruction == "jmp" {
			curPos += curInstruction.value
		}

		instructions[oldPos] = curInstruction
	}

	return resultSet
}

func findValidAccumulator(instructions []instructionSet) accReturn {
	curPos := 0
	oldPos := 0

	var resultSet accReturn

	for {
		if curPos >= len(instructions) {
			resultSet.foundEnd = true
			break
		}

		curInstruction := instructions[curPos]

		if curInstruction.visited {
			break
		}

		curInstruction.visited = true
		oldPos = curPos

		if curInstruction.instruction == "nop" {
			curPos++
		}

		if curInstruction.instruction == "acc" {
			resultSet.accumulatedValue += curInstruction.value
			curPos++
		}

		if curInstruction.instruction == "jmp" {
			curPos += curInstruction.value
		}

		instructions[oldPos] = curInstruction
	}

	return resultSet
}

func findInstructionsToToggle(instructions []instructionSet) []checkpoint {
	curPos := 0
	oldPos := 0
	accumulatedValue := 0

	instructionsToToggle := make([]checkpoint, 0)

	for {
		curInstruction := instructions[curPos]

		if curInstruction.visited {
			break
		}

		curInstruction.visited = true
		oldPos = curPos

		if curInstruction.instruction == "nop" {
			newCheckpoint := checkpoint{position: curPos, accumulatedValue: accumulatedValue}
			instructionsToToggle = append(instructionsToToggle, newCheckpoint)
			curPos++
		}

		if curInstruction.instruction == "acc" {
			accumulatedValue += curInstruction.value
			curPos++
		}

		if curInstruction.instruction == "jmp" {
			newCheckpoint := checkpoint{position: curPos, accumulatedValue: accumulatedValue}
			instructionsToToggle = append(instructionsToToggle, newCheckpoint)
			curPos += curInstruction.value
		}

		instructions[oldPos] = curInstruction
	}

	return instructionsToToggle
}

func resetVisitedInstructions(instructions []instructionSet) {
	for i := 0; i < len(instructions); i++ {
		curInstruction := instructions[i]
		curInstruction.visited = false
		instructions[i] = curInstruction
	}
}

func toggleInstruction(instructions []instructionSet, positionToSwap int) {
	curInstruction := instructions[positionToSwap]

	if curInstruction.instruction == "nop" {
		curInstruction.instruction = "jmp"
	} else if curInstruction.instruction == "jmp" {
		curInstruction.instruction = "nop"
	}

	instructions[positionToSwap] = curInstruction
}
