package main
import(
	f "fmt"
	m "math/rand"
	b "bufio"
	"os"
	s "strconv"
	"strings"
)
func begin(){
	f.Println("We're gonna play a game!")
	f.Println("enter 0 to end the game!")
	f.Println("Pick a number between 1 to 10.")
}

func main(){
	begin()
	for {
		reader := b.NewReader(os.Stdin)
		// I think the above is an object which we can call
		// particular methods on. Not 100% sure really...
		f.Println("Enter number(1-10): ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text,"\n","",-1)
		// gotta remove the end of line character somehow.
		if x,err := s.Atoi(text); err == nil {
			if (x == 0){
				f.Println("//////////////////////////////////////")
				f.Println("/ Exiting, thanks for trying it out! /")
				f.Println("//////////////////////////////////////")
				return
				// if you are actually reading this instead of running the code
				// you have to think about what you are doing for a moment.
			} else if (x > 10) || (x < 0){
				f.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
				f.Println("Well you entered an integer, but it was not between 1 to 10 inclusive")
				f.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
			} else {
				g := true
				for (g){
					p := m.Intn(10)
					if p == 0 {
						continue
						// if we get a zero, keep rolling
					} else {
						g = false
						if p == x {
							f.Println("---------------------------------------------------")
							f.Println("Congrats! you got a matching number!")
							f.Println("Remember, press 0 to stop running the program")
							f.Println("---------------------------------------------------")							
						} else {
							f.Println("***************************************************")
							f.Println("I am sorry, but you did not guess correctly!")
							f.Println("Your number was:", x)
							f.Println("The correct number was: ", p)
							f.Println("Keep guessing!")
							f.Println("***************************************************")
						}
					}
				}
			}
		} else {
			f.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
			f.Println("You must enter an integer, anything else will just not do!")
			f.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
		}
	}
}