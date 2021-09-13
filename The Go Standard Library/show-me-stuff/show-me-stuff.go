package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	//printVersion()
	//getFromTerminal()
	doSomething()
}

func doSomething() {
	args := os.Args[1:]

	if len(args) == 0  {
		fmt.Println("You must need to add some args")
		return
	}


	}

}

func getFromTerminal() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Tell me your name")
	text, _ := reader.ReadString('\n')
	fmt.Printf("Hello %v", text)
}

func printVersion() {
	fmt.Println("current version of go" + runtime.Version())
	fmt.Printf("current version of go %v", runtime.Version())
}
