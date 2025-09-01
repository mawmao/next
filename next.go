package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	args := os.Args
	var cmd string

	if len(args) < 2 {
		cmd = "all"
	} else {
		cmd = args[1]
	}

	store, err := os.OpenFile("./store.next", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("store can't be accessed: %v", err)
	}
	defer store.Close()

	switch cmd {
	case "all":
		fmt.Println("next: showing all todos")
		for {
			line, err := readStoreLine(store)
			if err == io.EOF {
				break
			}
			fmt.Println(line)
		}
	case "add":
		fmt.Println("next: add new todo")
	case "del":
		fmt.Println("next: delete todo by id")
	case "help":
		usage()
	default:
		log.Printf("next: invalid argument %q", cmd)
		usage()
		os.Exit(2)
	}
}

func parseStoreLine(s string) string {

	return ""
}

func readStoreLine(f *os.File) (string, error) {
	l := make([]byte, 0, 64)
	buf := make([]byte, 8)

	for {
		n, err := f.Read(buf)
		if n > 0 {
			chunk := buf[:n]
			if i := bytes.IndexByte(chunk, '\n'); i >= 0 {
				l = append(l, chunk[:i]...)
				if i+1 < n {
					f.Seek(int64(i-n+1), io.SeekCurrent)
				}
				return string(l), nil
			}
			l = append(l, chunk...)
		}
		if err != nil {
			if err == io.EOF {
				if len(l) > 0 {
					return string(l), nil
				}
				return "", io.EOF
			}
			return "", err
		}
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "next: usage")
	fmt.Fprintln(os.Stderr, "\tall\tshow all current todos")
	fmt.Fprintln(os.Stderr, "\tadd\tadd a todo")
}
