package main

import (
	"bytes"
	"fmt"
	"go-tamboon/cipher"
	"os"
)

func main() {
	dat, _ := os.ReadFile("data/fng.1000.csv.rot128")
	// r := strings.NewReader(string(dat))
	rot128reader, err := cipher.NewRot128Reader(bytes.NewBuffer(dat))
	if err != nil {
		fmt.Println(err)
	}
	buff := make([]byte, len(dat))
	rot128reader.Read(buff)
	f, err := os.OpenFile("data/fng.1000.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w, _ := cipher.NewRot128Writer(f)
	w.Write(buff)
	fmt.Println(string(buff))
}
