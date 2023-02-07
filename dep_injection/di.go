package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	// Fprintf takes a writer where to send the string to instead of stdout
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	// Now we can accept a http.ResponseWriter as well
	Greet(w, "world")
}

func main() {
	// "Injecting" os.Stdout, which satisifes io.Writer interface
	// Greet(os.Stdout, "Elodie")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
