package main

import(
	f "fmt"
)

func yes(x,y int)(z int){
	z = x - y
	return z
}

func main(){
	f.Println(yes(10,8))
}