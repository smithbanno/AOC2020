package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// CheckValidity1 is a function to verify the validity of passwords in file
func CheckValidity1() {
	file, err := os.Open("day2.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer file.Close()

	var validPasswords int

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		var (
			lowerLimit, upperLimit int
			char                   rune
			password               string
		)
		fmt.Sscanf(sc.Text(), "%d-%d %c: %s", &lowerLimit, &upperLimit, &char, &password) //Formatted scan of the input string

		if strings.Count(password, string(char)) >= lowerLimit && strings.Count(password, string(char)) <= upperLimit {
			validPasswords++
		}
	}
	fmt.Printf("There are %d valid passowords.\n", validPasswords)
}

// CheckValidity2 is a function to verify the validity of passwords in file
func CheckValidity2() {
	//read in file
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	reader := bufio.NewReader(file)

	var line string

	validPasswordCount := 0

	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		split := strings.Split(line, ":")
		policy := split[0]
		password := strings.TrimSpace(split[1])

		tmp := strings.Split(policy, " ")
		rules := tmp[0]
		char := tmp[1]
		tmp = strings.Split(rules, "-")
		pos1, err := strconv.Atoi(tmp[0])
		if err != nil {
			log.Fatal(err)
		}
		pos2, err := strconv.Atoi(tmp[1])
		if err != nil {
			log.Fatal(err)
		}

		index1 := pos1 - 1
		index2 := pos2 - 1
		char1 := string(password[index1])
		char2 := string(password[index2])

		if (char1 != char2) && (char1 == char || char2 == char) {
			//fmt.Printf("Password %s matches policy %s\n", password, policy)
			validPasswordCount++
		} else {
			//fmt.Printf("Password %s does not match policy %s\n", password, policy)
		}
	}
	fmt.Printf("There are %d valid passowords.\n", validPasswordCount)
}
