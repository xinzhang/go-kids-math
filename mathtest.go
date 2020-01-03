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

var useMultiply bool = true
var plusMinusSeed int = 999
var multiplySeed int = 99

func main() {
	fmt.Printf("Hi %s, there are %d questions to answer in your test: ", studentName, total)
	fmt.Println()

	startTime := time.Now()
	fmt.Println(startTime.Format("2006-01-02 15:04:05"))
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < total; i++ {
		answerRight := false
		defaultMethodOptions := 2
		if useMultiply {
			defaultMethodOptions = 3
		}
		method := rand.Intn(defaultMethodOptions)
		var expected, x, y int = 0, 0, 0

		if method == 0 {
			x = rand.Intn(plusMinusSeed) + 1
			y = rand.Intn(plusMinusSeed) + 1
			expected = x + y
		} else if method == 1 {
			x = rand.Intn(plusMinusSeed) + 1
			y = rand.Intn(plusMinusSeed) + 1
			if y > x {
				x, y = y, x
			}
			expected = x - y
		} else {
			x = rand.Intn(multiplySeed) + 1
			y = rand.Intn(multiplySeed) + 1
			expected = x * y
		}

		for !answerRight {
			if method == 0 {
				fmt.Printf("%d) %d + %d = ", i+1, x, y)
			} else if method == 1 {
				fmt.Printf("%d) %d - %d = ", i+1, x, y)
			} else {
				fmt.Printf("%d) %d * %d = ", i+1, x, y)
			}

			buf := bufio.NewReader(os.Stdin)
			answer, err := buf.ReadString('\n')
			myanswer, err := strconv.Atoi(strings.TrimSuffix(answer, "\n"))
			if err != nil {
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

	f, _ := os.OpenFile("records", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	f.WriteString(startTime.Format("2006-01-02") + ": " + studentName + ": " + humanizeDuration(duration) + "\n")

}
