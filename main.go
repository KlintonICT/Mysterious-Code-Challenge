package main

import (
	"encoding/base64"
	"fmt"
)

func getAnswer(str string) (result string) {
	for _, char := range str {
		if char == ':' {
			result = " " + result
		} else {
			result = string(char) + result
		}
	}
	return
}

func main() {
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)

	answer := getAnswer(string(sd))

	fmt.Println(answer)
}
