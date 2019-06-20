package main

import (
	"fmt"
	//"github.com/01-edu/Z01"
	"os"
	)

/*func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}*/

func isEven(nbr int) bool {
	if even(nbr){
		return true
	} else {
		return false
	}
}

func even(nbr int) bool{
	if (nbr%2)==0{
		return true	
	}
	return false
}

func main() {
	lengthOfArg:=len(os.Args[1:])
	EvenMsg:="I have an even number of arguments"
	OddMsg:="I have an odd number of arguments"
	if isEven(lengthOfArg) {
		fmt.Println(EvenMsg)
	} else {
		fmt.Println(OddMsg)
	}
}
