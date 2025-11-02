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

var total int = 10
var studentName string = "Eason"

var useTwo bool = true
var numberSeed int = 30
var timesSeed int = 5
var plusMinusSeed int = 2

func getPlusMinus() string {
	var rNum = rand.Intn(2) + 1
	if rNum == 1 {
		return "+"
	} else {
		return "-"
	}
}

func checkTimes(times int) string {
	if (times == 1) {
		return ""
	} else { 
		return strconv.Itoa(times)
	}
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

func getTotalEquation(x int, y int) (int, int, string) {
	var plusMinus = getPlusMinus()

	var times1 = rand.Intn(timesSeed) + 1
	var times2 = rand.Intn(timesSeed) + 1

	var total int = 0

	if (plusMinus == "+") {
		total = times1 * x + times2 * y
	} else {
		total = times1 * x - times2 * y
	}

	var line = checkTimes(times1) + "X " + plusMinus + " " + checkTimes(times2) + "Y = " + strconv.Itoa(total)
	return times1, total, line
}

func main() {
	// Load config
	config := loadConfig("config")
	if val, ok := config["algebra_total"]; ok {
		if num, err := strconv.Atoi(val); err == nil {
			total = num
		}
	}

	fmt.Println(total)
	fmt.Printf("Hi %s, there are %d questions to answer in your test: ", studentName, total)
	fmt.Println()

	startTime := time.Now()
	fmt.Println(startTime.Format("2006-01-02 15:04:05"))
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < total; i++ {
		
		var x, y int = 0, 0
		x = rand.Intn(numberSeed) + 1
		y = rand.Intn(numberSeed) + 1

		var times1, total1, line1 = getTotalEquation(x, y);
		var times2, total2, line2 = getTotalEquation(x, y);

		for (total2 % total1 == 0 && times2 / times1 == total2 / total1) {
			times2, total2, line2 = getTotalEquation(x, y)
		}

		var answerRight bool = false		

		for !answerRight {
			fmt.Println(line1)
			fmt.Println(line2)

			buf := bufio.NewReader(os.Stdin)
			answer, err := buf.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				continue
			}

			r := strings.Split(strings.TrimSuffix(answer, "\n"), ",")
			if (len(r) != 2) {
				fmt.Println("you need to enter like this x, y")
				continue
			}

			myAnswerX, err1 := strconv.Atoi(strings.Trim(r[0], " "))
			if err1 != nil {
				fmt.Println(err1)
				continue
			}

			myAnswerY, err2 := strconv.Atoi(strings.Trim(r[1], " "))
			if err2 != nil {
				fmt.Println(err2)
				continue
			}

			if myAnswerX == x && myAnswerY == y {
				fmt.Println("Right")
				answerRight = true
			} else {
				fmt.Println("Wrong")
			}

			continue
		}
	}

	duration := time.Since(startTime)
	fmt.Printf("Total %s \n", utils.HumanizeDuration(duration))

	f, _ := os.OpenFile("records", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	f.WriteString(startTime.Format("2006-01-02") + ": " + studentName + ": " + utils.HumanizeDuration(duration) + "\n")
}
