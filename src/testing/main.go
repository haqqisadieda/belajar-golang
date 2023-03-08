package main

import "log"

func HelloWorld(name string) string {
	log.Print("Hello " + name)
	return "Hello " + name
}

func main() {
	log.Print("Hello World!")
}
