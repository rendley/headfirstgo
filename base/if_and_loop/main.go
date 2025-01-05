package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	//seconds := time.Now().Unix()
	//rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've choosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")
	fmt.Println(target) // target not need if we play this game

	reader := bufio.NewReader(os.Stdin)
	succsess := false
	for guesses := 0; guesses < 2; guesses++ {
		fmt.Println("You have", 2-guesses, "guesses left")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops, your guess was LOW")
		} else if guess > target {
			fmt.Println("Opp your guess was HIGH")
		} else {
			succsess = true
			fmt.Println("You guess it. Congratulations!!!")
			break
		}
	}
	if !succsess {
		fmt.Println("Sorry. You didn't guess my number.It was: ", target)
	}

}
