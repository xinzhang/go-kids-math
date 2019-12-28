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

var total int = 20
var studentName string = "Ian"

func main() {
	fmt.Printf("Hi %s, there are %d questions to answer in your test: ", studentName, total)
	fmt.Println()

	startTime := time.Now()
	fmt.Println(startTime.Format("2006-01-02 15:04:05"))
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < total; i++ {
		answerRight := false
		x := rand.Intn(50) + 1
		y := rand.Intn(50) + 1
		method := rand.Intn(2)

		for !answerRight {
			expected := 0

			if method == 0 {
				fmt.Printf("%d) %d + %d = ", i+1, x, y)
				expected = x + y
			} else {
				if y > x {
					x, y = y, x
				}
				fmt.Printf("%d) %d - %d = ", i+1, x, y)
				expected = x - y
			}

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
