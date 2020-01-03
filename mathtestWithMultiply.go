package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var total int = 10
var studentName string = "Ian"
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
		method := rand.Intn(3)
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
	fmt.Printf("Total %s \n", humanizeDuration(duration))

	f, _ := os.OpenFile("records", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	f.WriteString(startTime.Format("2006-01-02") + ": " + studentName + ": " + humanizeDuration(duration) + "\n")

}

func humanizeDuration(duration time.Duration) string {
	days := int64(duration.Hours() / 24)
	hours := int64(math.Mod(duration.Hours(), 24))
	minutes := int64(math.Mod(duration.Minutes(), 60))
	seconds := int64(math.Mod(duration.Seconds(), 60))

	chunks := []struct {
		singularName string
		amount       int64
	}{
		{"day", days},
		{"hour", hours},
		{"minute", minutes},
		{"second", seconds},
	}

	parts := []string{}

	for _, chunk := range chunks {
		switch chunk.amount {
		case 0:
			continue
		case 1:
			parts = append(parts, fmt.Sprintf("%d %s", chunk.amount, chunk.singularName))
		default:
			parts = append(parts, fmt.Sprintf("%d %ss", chunk.amount, chunk.singularName))
		}
	}

	return strings.Join(parts, " ")
}
