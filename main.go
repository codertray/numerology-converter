package main

import (
	"example.com/numerology/conversions"

	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	text, err := getUserInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	letters := conversions.Format(text)

	fmt.Println(letters)

	soulNumber := conversions.Calculate(letters)

	fmt.Println(soulNumber)

	for soulNumber >= 10 {
		num := conversions.Merger(soulNumber)
		soulNumber = num
	}
	fmt.Printf("Your Number is %v!\n", soulNumber)
}

func getUserInput() (string, error) {
	fmt.Println("What's Your Full Name?")
	fmt.Println("If your name has changed use your birth name.")
	fmt.Print("Name:  ")
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil || text == "" {
		err = errors.New("input failed")
		return "", err
	}

	err = nil
	return text, err
}
