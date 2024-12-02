package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/yahya-lafhal/steganography/decoding"
	"github.com/yahya-lafhal/steganography/encoding"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select an option:")
	fmt.Println("1. Encode")
	fmt.Println("2. Decode")
	fmt.Print("Enter your choice (1 or 2): ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice != "1" && choice != "2" {
		fmt.Println("Invalid choice! Exiting.")
		return
	}

	action := "encode"
	if choice == "2" {
		action = "decode"
	}

	fmt.Printf("Enter the filename to %s: ", action)
	filename, _ := reader.ReadString('\n')
	filename = strings.TrimSpace(filename)

	switch action {
	case "encode":
		fmt.Printf("Enter the message to encode ")
		message, _ := reader.ReadString('\n')
		encoding.Encode(filename, message)
	case "decode":
		result := decoding.Decode(filename)
		fmt.Printf("Result : %s", result)
	}

}
