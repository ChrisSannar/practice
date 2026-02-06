package main

import "fmt"

func main() {
	var tempString string = "I am a temp string: racecar"
	str := []byte(tempString)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	fmt.Println(string(str))
}
