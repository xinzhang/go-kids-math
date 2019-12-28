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

var studentName string = "Ian"

func main() {
	fmt.Printf("Hi %s, there are %d questions to answer in your test: ", studentName)
	fmt.Println()

	startTime := time.Now()
	fmt.Println(startTime.Format("2006-01-02 15:04:05"))
	rand.Seed(time.Now().UnixNano())

  i := 0

  //y := rand.Intn(10)
  for y:= 2; y<=9; y++ {
	for x:= y; x <= 9; x++ {
    //for y:=1; y <= 9; y++ {
		answerRight := false
    wrongCnt := 0
    i++

		for !answerRight {
			
			fmt.Printf("%d) %d X %d = ", i, y, x)
			expected := x * y			

      buf := bufio.NewReader(os.Stdin)
      
      answer, err := buf.ReadString('\n')      
      if strings.TrimRight(answer, "\n") == "?" && wrongCnt >= 5 {
        fmt.Printf(" %d X %d = %d \n", y, x, expected)
        answerRight = true
        break
      }

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
        wrongCnt ++        
			}
    }
  }
	}

	duration := time.Since(startTime)
	fmt.Printf("Total %s \n", humanizeDuration(duration))
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
