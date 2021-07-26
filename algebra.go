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

var totalQuestions int = 8
var studentName string = "Ian"

var useTwo bool = true
var numberSeed int = 40
var timesSeed int = 3
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

func getTotalEquation(x int, y int) string {
	var plusMinus = getPlusMinus()

	var times1 = rand.Intn(timesSeed) + 1
	var times2 = rand.Intn(timesSeed) + 1

	var total int = 0

	if (plusMinus == "+") {
		total = times1 * x + times2 * y
	} else {
		total = times1 * x - times2 * y
	}

	return checkTimes(times1) + "X " + plusMinus + " " + checkTimes(times2) + "Y = " + strconv.Itoa(total)
}

func main() {
	fmt.Println(totalQuestions)
	fmt.Printf("Hi %s, there are %d questions to answer in your test: ", studentName, totalQuestions)
	fmt.Println()

	startTime := time.Now()
	fmt.Println(startTime.Format("2006-01-02 15:04:05"))
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < totalQuestions; i++ {
		
		var x, y int = 0, 0
		x = rand.Intn(numberSeed) + 1
		y = rand.Intn(numberSeed) + 1

		var line1 = getTotalEquation(x, y);
		var line2 = getTotalEquation(x, y);

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

			r := strings.Split(strings.TrimSuffix(answer, "\r\n"), ",")
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
