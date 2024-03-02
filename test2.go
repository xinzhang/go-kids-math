package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// generateQuestion creates a math question, sometimes including parentheses.
func generateQuestion() (question string, answer int) {
	// Randomly decide the structure of the question
	hasParentheses := rand.Intn(2) // 0 or 1, decides if parentheses will be used

	operand1 := rand.Intn(10) + 1
	operand2 := rand.Intn(10) + 1
	operand3 := rand.Intn(10) + 1

	operations := []string{"+", "-", "*", "/"}
	opIndex1 := rand.Intn(len(operations))
	opIndex2 := rand.Intn(len(operations))

	op1 := operations[opIndex1]
	op2 := operations[opIndex2]

	if hasParentheses == 1 && opIndex1 < 2 { // Use parentheses for addition/subtraction
		question = fmt.Sprintf("(%d %s %d) %s %d", operand1, op1, operand2, op2, operand3)
		switch op1 {
		case "+":
			operand1 += operand2
		case "-":
			operand1 -= operand2
		}
	} else {
		question = fmt.Sprintf("%d %s %d %s %d", operand1, op1, operand2, op2, operand3)
	}

	// Calculate the answer based on the operation order, assuming no parentheses or after parentheses are applied
	switch op2 {
	case "*":
		operand2 *= operand3
	case "/":
		operand2 /= operand3
	case "+":
		if op1 == "*" || op1 == "/" {
			switch op1 {
			case "*":
				operand1 *= operand2
			case "/":
				operand1 /= operand2
			}
			answer = operand1 + operand3
			return
		}
		operand2 += operand3
	case "-":
		if op1 == "*" || op1 == "/" {
			switch op1 {
			case "*":
				operand1 *= operand2
			case "/":
				operand1 /= operand2
			}
			answer = operand1 - operand3
			return
		}
		operand2 -= operand3
	}

	// Apply the first operation if it hasn't been applied yet
	switch op1 {
	case "+":
		answer = operand1 + operand2
	case "-":
		answer = operand1 - operand2
	case "*":
		answer = operand1 * operand2
	case "/":
		answer = operand1 / operand2
	}

	return
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator.
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 20; i++ {
		question, answer := generateQuestion()

		// Ask the question
		fmt.Printf("Question %d: %s = ", i+1, question)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		userAnswer, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please enter a valid number.")
			i-- // Decrement i to repeat the question
			continue
		}

		// Check the answer
		if userAnswer == answer {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("Incorrect. The correct answer is %d.\n", answer)
		}
	}
}
