package main

import(
	f "fmt"
)

func uselessCountToTen(){
	f.Println("As an advanced coder, I shall now count from 1 to 10 in go!")
	for i := 1; i <= 10; i++ {
		defer f.Println(i)
	}
}

func main(){
	uselessCountToTen()
	f.Println("Go is such a useless language...")
}