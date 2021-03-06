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

var studentName string = "Ian"

func randomNumber() int {
	return rand.Intn(8) + 2
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
	fmt.Printf("Hi %s, there are %d questions to answer in your test: ", studentName, total)
	fmt.Println()

	startTime := time.Now()
	fmt.Println(startTime.Format("2006-01-02 15:04:05"))
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < total; i++ {
		answerRight := false
		x := randomNumber()
		y := randomNumber()

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
	f.WriteString(startTime.Format("2006-01-02") + ": " + myName + ": " + strconv.Itoa(total) + " questions: " + utils.H,ml  umanizeDuration(duration) + "\n")

}
