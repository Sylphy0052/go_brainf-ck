package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	INCR  byte = 62 // +
	DECR  byte = 60 // -
	NEXT  byte = 43 // >
	PREV  byte = 45 // <
	READ  byte = 44 // ,
	WRITE byte = 46 // .
	OPEN  byte = 91 // [
	CLOSE byte = 93 // ]
)

var MAX_BUFFER_SIZE = 64 * 8

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[1])
		return
	}
	src, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	b := make([]byte, MAX_BUFFER_SIZE, MAX_BUFFER_SIZE)
	i := 0
	si := 0
	buf := make([]byte, 1)
	for {
		switch src[si] {
		case INCR:
			i++
		case DECR:
			i--
		case NEXT:
			b[i]++
		case PREV:
			b[i]--
		case WRITE:
			buf[0] = b[i]
			os.Stdout.Write(buf)
		case READ:
			os.Stdin.Read(buf)
			b[i] = buf[0]
		case OPEN:
			if b[i] == 0 {
				n := 0
				for {
					si++
					if src[si] == OPEN {
						n++
					} else if src[si] == CLOSE {
						n--
						if n < 0 {
							break
						}
					}
				}
			}
		case CLOSE:
			if b[i] != 0 {
				n := 0
				for {
					si--
					if src[si] == CLOSE {
						n++
					} else if src[si] == OPEN {
						n--
						if n < 0 {
							break
						}
					}
				}
			}
		}
		si++
		if si >= len(src) {
			break
		}
	}
}
