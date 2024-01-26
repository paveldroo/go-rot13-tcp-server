package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	s := bufio.NewScanner(conn)
	for s.Scan() {
		ln := strings.ToLower(s.Text())
		bs := []byte(ln)
		r := rot13(bs)

		fmt.Fprintf(conn, "%s - %s\n", ln, r)
	}
}

func rot13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}
