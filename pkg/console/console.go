package console

import (
	"log"
	"os"
)

var console = log.New(os.Stdout, "", 0)

func Printf(format string, v ...interface{}) {
	console.Printf(format, v...)
}

func Print(v ...interface{}) {
	console.Print(v...)
}

func Println(v ...interface{}) {
	console.Println(v...)
}

func Fatal(v ...interface{}) {
	console.Fatal(v...)
}
