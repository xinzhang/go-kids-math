package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"kids/utils"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var totalQuestions int = 30
var studentName string = "Eason"
var apiKey string = ""

// ChatGPT API structures
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatChoice struct {
	Message ChatMessage `json:"message"`
}

type ChatResponse struct {
	Choices []ChatChoice `json:"choices"`
}

// Question structure
type Question struct {
	Text    string
	OptionA string
	OptionB string
	OptionC string
	OptionD string
	Answer  string
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

func loadPromptTemplate(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading prompt file: %v\n", err)
		return ""
	}
	return string(content)
}

// Loading indicator in a separate goroutine
func showLoading(done chan bool) {
	spinner := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	i := 0
	for {
		select {
		case <-done:
			fmt.Print("\r")
			return
		default:
			fmt.Printf("\r%s Generating %d questions... Please wait...", spinner[i], totalQuestions)
			i = (i + 1) % len(spinner)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func callChatGPTBatch(prompt string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	messages := []ChatMessage{
		{
			Role:    "system",
			Content: "You are a helpful assistant that generates thinking skills questions for students. You must respond with ONLY a JSON array containing exactly 30 questions. Each question must have this exact format: {\"text\": \"question text\", \"a\": \"option A\", \"b\": \"option B\", \"c\": \"option C\", \"d\": \"option D\", \"answer\": \"a/b/c/d\"}. Return ONLY valid JSON array with 30 questions, no other text.",
		},
		{
			Role:    "user",
			Content: prompt + "\n\nBased on the examples above, generate 30 similar thinking skills questions with 4 options (A, B, C, D) each and indicate the correct answer. Return ONLY a valid JSON array of 30 questions.",
		},
	}

	reqBody := ChatRequest{
		Model:    "gpt-4.1",
		Messages: messages,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{Timeout: 120 * time.Second} // Longer timeout for batch generation
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var chatResp ChatResponse
	err = json.Unmarshal(body, &chatResp)
	if err != nil {
		return "", fmt.Errorf("error parsing response: %v\nResponse: %s", err, string(body))
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("no response from ChatGPT")
	}

	return chatResp.Choices[0].Message.Content, nil
}

func parseQuestions(jsonStr string) ([]Question, error) {
	// Clean up the response - sometimes ChatGPT wraps JSON in code blocks
	jsonStr = strings.TrimSpace(jsonStr)
	if strings.HasPrefix(jsonStr, "```json") {
		jsonStr = strings.TrimPrefix(jsonStr, "```json")
		jsonStr = strings.TrimSuffix(jsonStr, "```")
		jsonStr = strings.TrimSpace(jsonStr)
	} else if strings.HasPrefix(jsonStr, "```") {
		jsonStr = strings.TrimPrefix(jsonStr, "```")
		jsonStr = strings.TrimSuffix(jsonStr, "```")
		jsonStr = strings.TrimSpace(jsonStr)
	}

	var dataArray []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &dataArray)
	if err != nil {
		return nil, fmt.Errorf("error parsing questions JSON: %v\nJSON: %s", err, jsonStr)
	}

	questions := make([]Question, 0, len(dataArray))
	for _, data := range dataArray {
		question := Question{
			Text:    fmt.Sprintf("%v", data["text"]),
			OptionA: fmt.Sprintf("%v", data["a"]),
			OptionB: fmt.Sprintf("%v", data["b"]),
			OptionC: fmt.Sprintf("%v", data["c"]),
			OptionD: fmt.Sprintf("%v", data["d"]),
			Answer:  strings.ToLower(strings.TrimSpace(fmt.Sprintf("%v", data["answer"]))),
		}
		questions = append(questions, question)
	}

	return questions, nil
}

func displayQuestion(q *Question, questionNum int) {
	fmt.Printf("\n%d) %s\n", questionNum, q.Text)
	fmt.Printf("   A. %s\n", q.OptionA)
	fmt.Printf("   B. %s\n", q.OptionB)
	fmt.Printf("   C. %s\n", q.OptionC)
	fmt.Printf("   D. %s\n", q.OptionD)
	fmt.Print("\nYour answer (a/b/c/d): ")
}

func getStudentName(defaultName string) string {
	fmt.Printf("Current student name is %s, what is your name? Press enter to continue: ", defaultName)
	buf := bufio.NewReader(os.Stdin)
	myInput, _ := buf.ReadString('\n')

	if strings.TrimSpace(myInput) == "" {
		return defaultName
	}
	return strings.TrimSpace(myInput)
}

func main() {
	// Load config
	config := loadConfig("config")
	if val, ok := config["thinkingskills_total"]; ok {
		if num, err := strconv.Atoi(val); err == nil {
			totalQuestions = num
		}
	}

	// Get API key from environment variable
	apiKey = os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OpenAI API key not found in environment variable.")
		fmt.Println("Please set the OPENAI_API_KEY environment variable.")
		os.Exit(1)
	}

	// Load prompt template
	promptTemplate := loadPromptTemplate("OC_Thinking_Skills_Full_QA.md")
	if promptTemplate == "" {
		fmt.Println("Error: Could not load OC_Thinking_Skills_Full_QA.md file.")
		os.Exit(1)
	}

	fmt.Printf("Hi %s, there are %d thinking skills questions to answer.\n", studentName, totalQuestions)
	fmt.Println("Each question has 4 options (A, B, C, D).")
	fmt.Println("If you answer correctly, you'll move to the next question.")
	fmt.Println("If you answer incorrectly, you'll need to try again (no points awarded).")

	// Generate all questions at once with loading indicator
	done := make(chan bool)
	go showLoading(done)

	response, err := callChatGPTBatch(promptTemplate)
	done <- true // Stop the loading indicator

	if err != nil {
		fmt.Printf("\nError calling ChatGPT: %v\n", err)
		fmt.Println("Please try again later.")
		os.Exit(1)
	}

	questions, err := parseQuestions(response)
	if err != nil {
		fmt.Printf("\nError parsing questions: %v\n", err)
		fmt.Println("Please try again later.")
		os.Exit(1)
	}

	if len(questions) != totalQuestions {
		fmt.Printf("\nWarning: Expected %d questions but got %d\n", totalQuestions, len(questions))
		totalQuestions = len(questions)
	}

	fmt.Printf("\n✓ Successfully generated %d questions!\n", totalQuestions)
	fmt.Println("\nStarting test...")

	startTime := time.Now()
	fmt.Println("Start time:", startTime.Format("2006-01-02 15:04:05"))

	score := 0
	reader := bufio.NewReader(os.Stdin)

	// Present questions one by one
	for i := 0; i < totalQuestions; i++ {
		answerRight := false
		firstAttempt := true
		currentQuestion := &questions[i]

		// Keep asking until they get it right
		for !answerRight {
			displayQuestion(currentQuestion, i+1)

			answer, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			answer = strings.ToLower(strings.TrimSpace(answer))

			// Validate input
			if answer != "a" && answer != "b" && answer != "c" && answer != "d" {
				fmt.Println("Invalid input. Please enter a, b, c, or d.")
				continue
			}

			if answer == currentQuestion.Answer {
				fmt.Println("✓ Correct!")
				answerRight = true
				// Only award points on first attempt
				if firstAttempt {
					score++
				}
			} else {
				fmt.Println("✗ Wrong! Please try again...")
				firstAttempt = false
			}
		}
	}

	duration := time.Since(startTime)
	fmt.Printf("\n═══════════════════════════════════════\n")
	fmt.Printf("Test Complete!\n")
	fmt.Printf("Final Score: %d/%d\n", score, totalQuestions)
	fmt.Printf("Time taken: %s\n", utils.HumanizeDuration(duration))
	fmt.Printf("═══════════════════════════════════════\n\n")

	myName := getStudentName(studentName)

	// Save results
	f, _ := os.OpenFile("ts_trial_test_records", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	defer f.Close()
	f.WriteString(fmt.Sprintf("%s: %s: %d/%d questions: %s\n",
		startTime.Format("2006-01-02"),
		myName,
		score,
		totalQuestions,
		utils.HumanizeDuration(duration)))
}
