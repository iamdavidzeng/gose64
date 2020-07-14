package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

const encode = "--encode"
const decode = "--decode"
const hint = `For example1: gose64 --encode 'hello, world'.
For exmaple2: gose64 --decode 'aGVsbG8sIHdvcmxk'.`

func main() {
	osArgs := os.Args
	if len(osArgs) != 3 {
		fmt.Println("gose64 need two arguments to execute.")
		fmt.Println(hint)
		return
	}

	firstArg := osArgs[1]

	validFirstArg := []string{encode, decode}
	isFirstArgValid := false
	for _, item := range validFirstArg {
		if item == firstArg {
			isFirstArgValid = true
		}
	}
	if !isFirstArgValid {
		fmt.Printf("Invalid command arguments: %v\n", firstArg)
		fmt.Println("gose64 first argument only support --decode or --encode.")
		fmt.Println(hint)
		return
	}

	stringWaitToHandle := osArgs[2]
	if stringWaitToHandle == "" {
		fmt.Println("gose64 second argument can't be empty.")
		fmt.Println(hint)
		return
	}
	if firstArg == encode {
		encoded := base64.StdEncoding.EncodeToString([]byte(stringWaitToHandle))
		fmt.Println(encoded)
		return
	}
	decoded, err := base64.StdEncoding.DecodeString(stringWaitToHandle)
	if err != nil {
		fmt.Println("decode error:", err)
	}
	fmt.Println(string(decoded))
	return
}
