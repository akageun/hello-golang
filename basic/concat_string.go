package main

import (
	"fmt"
	"bytes"
)

func main() {
	fmt.Println("Hello" + "," + " World" + "!")

	var buffer bytes.Buffer
	buffer.WriteString("Hello")

	fmt.Println(buffer.String())

	buffer.WriteString(",")
	buffer.WriteString(" World!")

	fmt.Println(buffer.String())
}
