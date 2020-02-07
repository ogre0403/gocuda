package main

import (
	"fmt"
	"github.com/ogre0403/gocuda/pkg/cuda"
)

func main() {
	fmt.Print("Hello")
	cuda.Init(0)
	fmt.Printf("There is %d GPU\n", cuda.DeviceGetCount())
}
