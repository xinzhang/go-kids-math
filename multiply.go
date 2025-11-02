package main

import (
	"bufio"
	"fmt"
	"kids/utils"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var total int = 50
var studentName string = "Eason"

var xSeed int = 14
var ySeed int = 10

func randomXNumber() int {
	return rand.Intn(xSeed) + 2
}

func randomYNumber() int {
	return rand.Intn(ySeed) + 2
}

func loadConfig(filename string) map[string]string {
	config := make(map[string]string)
	file, err := os.Open(filename)
	if err != nil {
		return config
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			config[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	return config
}

func getStudentName(defaultName string) string {
	fmt.Printf("Current student name is %s, what is your name? press enter to Continue  ", defaultName)
	buf := bufio.NewReader(os.Stdin)
	myInput, _ := buf.ReadString('\n')

	if strings.TrimSuffix(myInput, "\n") == "" {
		return defaultName
	} else {
		return strings.TrimSuffix(myInput, "\n")
	}
}

func main() {
	// Load config
	config := loadConfig("config")
	if val, ok := config["multiply_total"]; ok {
		if num, err := strconv.Atoi(val); err == nil {
			total = num
		}
	}
	if val, ok := config["multiply_X_seed"]; ok {
		if num, err := strconv.Atoi(val); err == nil {
			xSeed = num
		}
	}
	if val, ok := config["multiply_Y_seed"]; ok {
		if num, err := strconv.Atoi(val); err == nil {
			ySeed = num
		}
	}

	fmt.Printf("Hi %s, there are %d questions to answer in your test: ", studentName, total)
	fmt.Println()

	startTime := time.Now()
	fmt.Println(startTime.Format("2006-01-02 15:04:05"))
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < total; i++ {
		answerRight := false
		x := randomXNumber()
		y := randomYNumber()

		for !answerRight {
			expected := 0

			fmt.Printf("%d) %d X %d = ", i+1, x, y)
			expected = x * y

			buf := bufio.NewReader(os.Stdin)
			answer, err := buf.ReadString('\n')
			myanswer, err := strconv.Atoi(strings.TrimSuffix(answer, "\n"))
			if err != nil {
				// handle error
				fmt.Println(err)
				continue
			}
			if myanswer == expected {
				fmt.Println("Right")
				answerRight = true
			} else {
				fmt.Println("Wrong")
			}
		}
	}

	duration := time.Since(startTime)
	fmt.Printf("Total %s \n", utils.HumanizeDuration(duration))
	myName := getStudentName(studentName)

	f, _ := os.OpenFile("multiply_records", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	f.WriteString(startTime.Format("2006-01-02") + ": " + myName + ": " + strconv.Itoa(total) + " questions: " + utils.HumanizeDuration(duration) + "\n")

}
