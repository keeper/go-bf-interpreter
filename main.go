package main

import "fmt"

func main() {
	var input = ""
	var buf [1024]byte
	var ptr = 0
	fmt.Scanf("%s", &input)
	for _, c := range input {
		switch c {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			buf[ptr]++
		case '-':
			buf[ptr]--
		case '.':
			fmt.Printf("%c", buf[ptr])
		case ',':
			fmt.Scanf("%c", &buf[ptr])
		case '[', ']':
			fmt.Println("Not implemented")
		}
	}
}
