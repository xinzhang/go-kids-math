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
var studentName string = "Ian"

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
	fmt.Printf("Hi %s, there are %d questions to answer in your test: ", studentName, total)
	fmt.Println()

	startTime := time.Now()
	fmt.Println(startTime.Format("2006-01-02 15:04:05"))
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < total; i++ {
		
		var x, y int = 0, 0
		x = rand.Intn(numberSeed) + 1
		y = rand.Intn(numberSeed) + 1

		var line1 = getTotalEquation(x, y);
		var line2 = getTotalEquation(x, y);

		var answerRight bool = false		

		for !answerRight {
			fmt.Println(line1)
			fmt.Println(line2)

			//var myAnswerX, myAnswerY int = 0, 0
			buf := bufio.NewReader(os.Stdin)
			answer, err := buf.ReadString('\n')

			r := strings.Split(strings.TrimSuffix(answer, "\n"), ",")
			myAnswerX, err := strconv.Atoi(strings.Trim(r[0], " "))
			myAnswerY, err := strconv.Atoi(strings.Trim(r[1], " "))

			if err != nil {
				fmt.Println(err)
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
