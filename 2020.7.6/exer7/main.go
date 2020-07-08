package main

import (
	"bytes"
	"fmt"
)

func main()  {
	//var buf1 = bytes.NewBufferString("hello world")

	var buf2 = bytes.NewBuffer([]byte("hello"))

	buf2.WriteString(" world")

	fmt.Println(buf2)
}