package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

var rot13 = map[byte]byte{
	'A': 'N',
	'B': 'O',
	'C': 'P',
	'D': 'Q',
	'E': 'R',
	'F': 'S',
	'G': 'T',
	'H': 'U',
	'I': 'V',
	'J': 'W',
	'K': 'X',
	'L': 'Y',
	'M': 'Z',
	'N': 'A',
	'O': 'B',
	'P': 'C',
	'Q': 'D',
	'R': 'E',
	'S': 'F',
	'T': 'G',
	'U': 'H',
	'V': 'I',
	'W': 'J',
	'X': 'K',
	'Y': 'L',
	'Z': 'M',
	'a': 'n',
	'b': 'o',
	'c': 'p',
	'd': 'q',
	'e': 'r',
	'f': 's',
	'g': 't',
	'h': 'u',
	'i': 'v',
	'j': 'w',
	'k': 'x',
	'l': 'y',
	'm': 'z',
	'n': 'a',
	'o': 'b',
	'p': 'c',
	'q': 'd',
	'r': 'e',
	's': 'f',
	't': 'g',
	'u': 'h',
	'v': 'i',
	'w': 'j',
	'x': 'k',
	'y': 'l',
	'z': 'm',
	'0': '9',
	'1': '8',
	'2': '7',
	'3': '6',
	'4': '5',
	'5': '4',
	'6': '3',
	'7': '2',
	'8': '1',
	'9': '0',
	' ': ' ',
	'~': '~',
	'!': '!',
	'@': '@',
	'#': '#',
	'$': '$',
	'%': '%',
	'^': '^',
	'&': '&',
	'*': '*',
	'(': '(',
	')': ')',
	'_': '_',
	'+': '+',
	'`': '`',
	'-': '-',
	'=': '=',
	'?': '?',
	',': ',',
	'.': '.',
}

// 从 rot.r 中读出字节切片，经过转换后重新填充 []byte
func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i, v := range b {
		b[i] = rot13[v]
	}
	return n, nil
}

func main() {
	s1 := strings.NewReader("Lbh penpxrq gur pbqr!")
	r1 := rot13Reader{s1}
	io.Copy(os.Stdout, &r1)

	s2 := strings.NewReader("You cracked the code!")
	r2 := rot13Reader{s2}
	// io.Copy(os.Stdout, &r2)
	// 或者这样输出
	b := make([]byte, 8)
	for {
		n, err := r2.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
