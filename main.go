package main

import "fmt"

func main() {
	g := NewGenerator()
	uid := g.Generate()
	fmt.Println(uid.ToString())
}
