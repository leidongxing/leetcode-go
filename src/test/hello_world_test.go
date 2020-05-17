package test

import (
	"fmt"
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	}
	fmt.Println("hello world")
	os.Exit(0)
}
