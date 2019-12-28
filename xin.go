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

func sayHello() string {
	hours := time.Now().Hour()

	if hours > 5 && hours < 12 {
		return "good morning"
	} else if hours >= 12 && hours <= 20 {
		return "good afternoon"
	} else {
		return "good evening"
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(sayHello())

	x := rand.Intn(99)
	y := rand.Intn(30)

	fmt.Printf("what is %v * %v = ? \n", x, y)
	expectResult := (x * y)

	buf := bufio.NewReader(os.Stdin)
	myInput, _ := buf.ReadString('\n')

	if strings.TrimSuffix(myInput, "\n") == strconv.Itoa(expectResult) {
		fmt.Println("Correct answer !")
	} else {
		fmt.Println("Wrong answer ")
	}

}
