package main

import (
	"fmt"
	"time"

	"github.com/warthog618/go-gpiocdev"
)

const (
	pin = 23
)

func main() {
	in, err := gpiocdev.RequestLine("gpiochip0", pin, gpiocdev.AsOutput(1))
	defer in.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(in.Chip())

	info, err := in.Info()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", info)
	time.Sleep(3 * time.Second)
	err = in.SetValue(0)
	if err != nil {
		panic(err)
	}
}
