package bf_interpreter

import "fmt"

func main() {
	input := ""
	buf := [30]byte{}
	ptr := 0
	rt_pc_stk := NewStack()
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
			rt_pc_stk.push(pc)
		case ']':
			if buf[ptr] != 0 {
				pc = rt_pc_stk.top()
			} else {
				rt_pc_stk.pop()
				pc++
			}
		}
	}
}
