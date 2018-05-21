package main

import (
	f "fmt"
	r "runtime"
)
func OsDetector(){
	f.Print("Does your os support go")
	switch os := r.GOOS; os {
	case "darwin": // OS X
		f.Println("No.")
	case "linux":
		f.Println("No.")
	default:
		f.Println("No.")
	}
}

func main() {
	OsDetector()
	f.Println("SteamOS supports golang")

}

