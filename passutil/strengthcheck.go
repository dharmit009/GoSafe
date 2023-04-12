package passutil

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StrengthCheck(checkThisPass string) float32 {

	// fmt.Println("### passtest ###")

	// var checkThisPass string = "typepassword"

	var uppercase_letters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var lowercase_letters string = strings.ToLower(uppercase_letters)
	var digits string = "0123456789"
	var symbols string = "!@#$%&*()[]{}-_+=';:.,"
	var upperCount, lowerCount, digitCount, symbolCount, spaceCount = 0, 0, 0, 0, 0
	var score float32 = 0
	var common_pass_flag bool = false
	var invalidPassFlag bool = false

	//fmt.Println("Currently Checking: ", checkThisPass)

  if len(checkThisPass) < 8 {
    return 0
  }

	// checking if password is common or not !!
	file, err := os.Open("./rockyou.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == checkThisPass {
			common_pass_flag = true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	for i := 0; i < len(checkThisPass); i++ {
		currentChar := string(checkThisPass[i])
		if strings.Contains(uppercase_letters, currentChar) {
			upperCount += 1
		} else if strings.Contains(lowercase_letters, currentChar) {
			lowerCount += 1
		} else if strings.Contains(digits, currentChar) {
			digitCount += 1
		} else if strings.Contains(symbols, currentChar) {
			symbolCount += 1
    } else if strings.Contains(" ", currentChar) {
      spaceCount += 1
		} else {
			invalidPassFlag = true
		}
	}


	if upperCount > 0 {
		score += 2
	}
	if lowerCount > 0 {
		score += 1
	}
	if digitCount > 0 {
		score += 1
	}
	if symbolCount > 0 {
		score += 2
	}
	if upperCount > 0 && lowerCount > 0 && digitCount > 0 && symbolCount > 0 && len(checkThisPass) > 12 {
		score += 2
	} else {
		score += 1
	}
	if invalidPassFlag == true {
		// fmt.Println("Invalid Password !")
		score = 0
	}

  if common_pass_flag == true {
      // fmt.Println("Common Password: ", checkThisPass)
      score = 0
  } else {
      score += 2
  }


	// fmt.Println("Length Of Pass: ", len(checkThisPass))
	// fmt.Println("Upper Count: ", upperCount)
	// fmt.Println("Lower Count: ", lowerCount)
	// fmt.Println("Digit Count: ", digitCount)
	// fmt.Println("Symbol Count: ", symbolCount)
	// fmt.Println("Score: ", score)
	return score

}
