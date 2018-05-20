/*
The FBI run countless quantum computers in order to solve the 
ominous question which is: Is my number fizz, buzz,fizzbuzz,
or boring. Rob Pike was faced with the most challenging task he has 
ever been faced with: is his age fizz,buzz,fizzbuzz,or none of them.
My code will solve this dilemma once and for all.

Also I am totally not writing fizzbuzz for the meme or for deep
lore reasons, this is a very serious task for those of us in 3rd or 
fourth class.
*/

package main

import (
	f "fmt"
	b "bytes"
)
func Fizz(a ...int){
	f.Println("These are the numbers:\n",a)
	for _, number := range a{
		var writer b.Buffer // I can't do := here for some reason
		banana := true
		if number % 3 == 0{
			writer.WriteString("fizz")
			banana = false
		}
		if number % 5 == 0{
			writer.WriteString("buzz")
			banana = false
		}
		if banana{
			writer.WriteString("Boring")
		}
		f.Println(number, writer.String())
	}
}

func main(){
	Fizz(1,2,3,4,5,6,7,8,9,10,15)
}

// this was totally not testing numerous arguments and buffer nonononono

