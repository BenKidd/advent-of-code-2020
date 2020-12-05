package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var isValidHairColor = regexp.MustCompile("([0-9a-f]){6}").MatchString
var validEyeColours = map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}

func main() {
	passportFile, err := os.Open("passportData.txt")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer passportFile.Close()

	scanner := bufio.NewScanner(passportFile)

	numValidPassports := 0
	var passport map[string]string
	passport = make(map[string]string)

	for {
		if !scanner.Scan() {
			if isValidPassport(passport) {
				numValidPassports++
			}

			break
		}

		currentLine := scanner.Text()

		if len(currentLine) > 0 {
			addPassengerDetails(passport, currentLine)
			continue
		}

		if isValidPassport(passport) {
			numValidPassports++
		}

		passport = make(map[string]string)
	}

	fmt.Println("Number of valid passports:", numValidPassports)
}

func addPassengerDetails(details map[string]string, data string) {
	segments := strings.Split(data, " ")

	for i := 0; i < len(segments); i++ {
		dataBits := strings.Split(segments[i], ":")

		details[dataBits[0]] = dataBits[1]
	}
}

func isValidPassport(details map[string]string) bool {
	if len(details) < 7 {
		return false
	}

	if byr, ok := details["byr"]; ok {
		i, err := strconv.Atoi(byr)

		if err != nil {
			return false
		}

		if i < 1920 || i > 2002 {
			return false
		}
	} else {
		return false
	}

	if iyr, ok := details["iyr"]; ok {
		i, err := strconv.Atoi(iyr)

		if err != nil {
			return false
		}

		if i < 2010 || i > 2020 {
			return false
		}
	} else {
		return false
	}

	if eyr, ok := details["eyr"]; ok {
		i, err := strconv.Atoi(eyr)

		if err != nil {
			return false
		}

		if i < 2020 || i > 2030 {
			return false
		}
	} else {
		return false
	}

	if hgt, ok := details["hgt"]; ok {
		if len(hgt) < 2 {
			return false
		}

		measurementStr := hgt[:len(hgt)-2]

		measurement, err := strconv.Atoi(measurementStr)

		if err != nil {
			return false
		}

		if strings.HasSuffix(hgt, "cm") {
			if measurement < 150 || measurement > 193 {
				return false
			}
		} else if strings.HasSuffix(hgt, "in") {
			if measurement < 59 || measurement > 76 {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}

	if hcl, ok := details["hcl"]; ok {
		if len(hcl) < 7 {
			return false
		}

		if !strings.HasPrefix(hcl, "#") {
			return false
		}

		if !isValidHairColor(hcl[1:]) {
			return false
		}
	} else {
		return false
	}

	if ecl, ok := details["ecl"]; ok {
		if !validEyeColours[ecl] {
			return false
		}
	} else {
		return false
	}

	if pid, ok := details["pid"]; ok {
		if len(pid) != 9 {
			return false
		}

		_, err := strconv.Atoi(pid)

		if err != nil {
			return false
		}
	} else {
		return false
	}

	if _, ok := details["cid"]; !ok {
		// do nothing with this
	}

	return true
}
