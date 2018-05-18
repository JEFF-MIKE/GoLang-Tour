package main

import (
	"fmt"
	b "bufio"
	"os"
)

func main(){

	var reader = b.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Print(text)
	fmt.Println("I created an object which reads from os standard input then reads as far") 
	fmt.Println("as the newline please appreciate this amazing task.")
	fmt.Println("There is a 72 word limit on commit messages,this is Van Dongen's doing")
}



