package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

func readBuffer(bufCh chan []byte) {
	buf := make([]byte, 1024)

	for {
		if n, err := syscall.Read(syscall.Stdin, buf); err == nil {
			bufCh <- buf[:n]
		}
	}
}

func Noncanonical() {
	fd := int(os.Stdin.Fd())

	val, err := unix.IoctlGetTermios(fd, unix.TIOCGETA)
	if err != nil {
		log.Fatal(err)
	}
	val.Lflag &^= syscall.ICANON
	val.Cc[syscall.VMIN] = 3
	if err := unix.IoctlSetTermios(fd, unix.TIOCSETA, val); err != nil {
		log.Fatal(err)
	}
	bufCh := make(chan []byte, 1)
	go readBuffer(bufCh)

	for {
		fmt.Printf("\ninput: %c\n", <-bufCh)
	}
}
