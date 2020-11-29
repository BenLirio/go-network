package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	con, err := net.Dial("tcp", "192.168.86.41:3306")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer con.Close()
	_, err = con.Write([]byte("test"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
