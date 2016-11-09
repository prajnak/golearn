package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var r io.Reader
	r = os.Stdin
	fmt.Println(reflect.TypeOf(r))
	r = bufio.NewReader(r)
	fmt.Println(reflect.TypeOf(r))
	r = new(bytes.Buffer)
	fmt.Println(reflect.TypeOf(r))

	tty, _ := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	r = tty
	fmt.Println(r, reflect.TypeOf(r))
	var w io.Writer
	w = r.(io.Writer)
	fmt.Println("type of w is: ", reflect.TypeOf(w), w)
}
