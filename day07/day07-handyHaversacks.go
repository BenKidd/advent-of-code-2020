package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type containedBag struct {
	description string
	number      int
}

type bagData struct {
	outerBags []string
	// description string // might not be needed?
	innerBags []containedBag
}

func main() {
	rulesFile, err := os.Open("packingRules.txt")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer rulesFile.Close()

	rulesInput := bufio.NewScanner(rulesFile)
	var bagRules = make(map[string]bagData)

	for {
		if !rulesInput.Scan() {
			break
		}

		newRule := rulesInput.Text()

		ruleData := strings.Split(newRule, " bags contain ")
		newBagDesc := strings.TrimSpace(ruleData[0])
		newBagContents := strings.Split(ruleData[1], ",")

		newBag := bagRules[newBagDesc]

		if strings.HasPrefix(newBagContents[0], "no") {
			bagRules[newBagDesc] = newBag
			continue
		}

		newBag.innerBags = make([]containedBag, len(newBagContents))

		outerBag := []string{newBagDesc}

		for i := 0; i < len(newBagContents); i++ {
			var innerBag containedBag
			innerBagDetails := strings.Split(strings.TrimSpace(newBagContents[i]), " ")

			innerBagName := innerBagDetails[1] + " " + innerBagDetails[2]

			innerBag.description = innerBagName
			innerBag.number, err = strconv.Atoi(innerBagDetails[0])

			if err != nil {
				log.Fatal(err)
				return
			}

			newBag.innerBags[i] = innerBag

			// Ensures that we add the new reference to the outer bag
			// without overwriting existing data
			existingBag := bagRules[innerBagName]
			existingBag.outerBags = append(existingBag.outerBags, outerBag...)
			bagRules[innerBagName] = existingBag
		}

		bagRules[newBagDesc] = newBag
	}

	bagsContainingGoldBag := make(map[string]bool)
	goldBag := bagRules["shiny gold"]

	searchOutsideBags(bagsContainingGoldBag, bagRules, goldBag.outerBags)

	fmt.Println("Number of bags containing 1 or more gold bags:", len(bagsContainingGoldBag))

	bagsInsideGoldBag := searchInsideBags(bagRules, goldBag.innerBags)

	// We deduct one so we don't count the shiny gold bag itself
	bagsInsideGoldBag--

	fmt.Println("Number of bags contained inside 1 gold bag:", bagsInsideGoldBag)

}

func searchOutsideBags(bagsVisited map[string]bool, bagRules map[string]bagData, outerBags []string) {
	for i := 0; i < len(outerBags); i++ {
		outerBagName := outerBags[i]
		bagsVisited[outerBagName] = true

		outerBag := bagRules[outerBagName]
		searchOutsideBags(bagsVisited, bagRules, outerBag.outerBags)
	}
}

func searchInsideBags(bagRules map[string]bagData, innerBags []containedBag) int {
	totalInnerBags := 1

	if len(innerBags) == 0 {
		return 1
	}

	for i := 0; i < len(innerBags); i++ {
		innerBagDetails := bagRules[innerBags[i].description]

		numContainedBags := searchInsideBags(bagRules, innerBagDetails.innerBags)
		totalInnerBags += numContainedBags * innerBags[i].number
	}

	return totalInnerBags
}
