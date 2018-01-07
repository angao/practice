package main

import (
	"fmt"
	"bytes"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma1(str string) string {
	if len(str) <= 3 {
		return str
	}
	var buf bytes.Buffer
	var tmp []byte
	var s = []byte(str)

	if comm := strings.LastIndex(str, "."); comm > -1 {
		tmp = s[comm:]
		s = s[:comm]
	}

	if strings.HasPrefix(str, "-") || strings.HasPrefix(str, "+") {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	for n := len(s); n > 0; {
		m := n % 3
		if m > 0 {
			buf.Write(s[:m])
			s = s[m:n]
			n -= m
		}
		fmt.Fprintf(&buf, ",%s", s[:3])
		s = s[3:]
		n -= 3
	}
	buf.Write(tmp)
	return buf.String()
}

func main() {
	fmt.Println(comma("12345"))
	fmt.Println(comma1("12345354769045.8806"))
}
