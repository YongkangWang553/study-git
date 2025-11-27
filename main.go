package main

import "log"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("hello world")

	log.Println("I am learing git")
}
