package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("No input")
		return
	}

	var (
		text = args[0]
		size = len(text)
		buf  = make([]byte, 0, size)
		in   bool
	)

	const (
		link  = "http://"
		nlink = len(link)
	)

	fmt.Println("text size: ", size)
	for i := 0; i < size; i++ {
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			in = true

			buf = append(buf,link...)
			i += nlink
		}
		c := text[i]
		switch c {
			case ' ', '\t', '\n'
			in = false
		}
		if in {
			c = '*'
		}
		buf = append(buf, c)

	}
	fmt.Println(string(buf))
}
