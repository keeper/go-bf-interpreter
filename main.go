package main

import "fmt"

func main() {
	input := ""
	buf := [30]byte{}
	ptr := 0
	rt_pc_stk := []int{}
	pc := 0
	fmt.Scanf("%s", &input)
	for pc != len(input) {
		c := input[pc]
		switch c {
		case '>':
			ptr++
			pc++
		case '<':
			ptr--
			pc++
		case '+':
			buf[ptr]++
			pc++
		case '-':
			buf[ptr]--
			pc++
		case '.':
			fmt.Printf("%c", buf[ptr])
			pc++
		case ',':
			fmt.Scanf("%c", &buf[ptr])
			pc++
		case '[':
			pc++
			rt_pc_stk = append(rt_pc_stk, pc)
		case ']':
			rt_to := rt_pc_stk[len(rt_pc_stk)-1] // get top of the stack
			if buf[ptr] != 0 {
				pc = rt_to
			} else {
				rt_pc_stk = rt_pc_stk[:len(rt_pc_stk)-1] // pop the return pointer
				pc++
			}
		}
	}
}
