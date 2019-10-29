package main

import (
	v1 "github.com/yuswift/gomod_demo/swift"
	v2 "github.com/yuswift/gomod_demo/v2/swift"
)

func main() {
	println(v1.Getmodversion())
	println(v2.Getmodversion())
}
