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

var useMultiply bool = true
var useDivision bool = false

var plusMinusSeed int = 999
var multiplySeed int = 99
var divideSeedX int = 999
var divideSeedY int = 9

var lineReturn = "\n"

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

func main() {
	// Load config
	config := loadConfig("config")
	if val, ok := config["mathtest_total"]; ok {
		if num, err := strconv.Atoi(val); err == nil {
			total = num
		}
	}
	if val, ok := config["mathtest_multiply"]; ok {
		useMultiply = (val == "true" || val == "1")
	}
	if val, ok := config["mathtest_divide"]; ok {
		useDivision = (val == "true" || val == "1")
	}

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
		if useDivision {
			defaultMethodOptions = 4
		}
		method := rand.Intn(defaultMethodOptions)
		var expected, x, y int = 0, 0, 0
		var expectedReminder int = 0

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

		} else if method == 2 {
			x = rand.Intn(multiplySeed) + 1
			y = rand.Intn(multiplySeed) + 1
			expected = x * y
		} else {
			x = rand.Intn(divideSeedX) + 1
			y = rand.Intn(divideSeedY) + 1
			if y > x {
				x, y = y, x
			}
			expected = x / y
			expectedReminder = x % y
		}

		for !answerRight {
			if method == 0 {
				fmt.Printf("%d) %d + %d = ", i+1, x, y)
			} else if method == 1 {
				fmt.Printf("%d) %d - %d = ", i+1, x, y)
			} else if method == 2 {
				fmt.Printf("%d) %d * %d = ", i+1, x, y)
			} else {
				fmt.Printf("%d) %d / %d = ", i+1, x, y)
			}

			buf := bufio.NewReader(os.Stdin)
			answer, err := buf.ReadString('\n')

			// this is divided
			if method == 3 {
				myanswerReminder := 0
				r := strings.Split(strings.TrimSuffix(answer, lineReturn), ",")
				myanswer, err := strconv.Atoi(strings.Trim(r[0], " "))

				if len(r) == 1 {
					if expectedReminder > 0 {
						fmt.Println("Wrong")
						continue
					}
				} else {
					myanswerReminder, err = strconv.Atoi(strings.Trim(r[1], " "))
				}

				if err != nil {
					fmt.Println(err)
					continue
				}

				if myanswer == expected && myanswerReminder == expectedReminder {
					fmt.Println("Right")
					answerRight = true
				} else {
					fmt.Println("Wrong")
				}

				continue
			}

			myanswer, err := strconv.Atoi(strings.TrimSuffix(answer, lineReturn))
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
	fmt.Printf("Total %s %s", utils.HumanizeDuration(duration), lineReturn)

	f, _ := os.OpenFile("records", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	f.WriteString(startTime.Format("2006-01-02") + ": " + studentName + ": " + utils.HumanizeDuration(duration) + "\n")

}
