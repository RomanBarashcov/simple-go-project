package utils

import "fmt"

func WriteConsoleLog(value string) {
	yellow := "\033[33m"
	reset := "\033[0m"
	fmt.Println(yellow + "===========================================================================")
	fmt.Printf(" 			%v \n", value)
	fmt.Println("===========================================================================" + reset)
}
