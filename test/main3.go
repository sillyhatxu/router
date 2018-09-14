package main

import (
	"os"
	"flag"
	"fmt"
)

func init() {

}

func main() {
	os.Exit(realMain())
}

func realMain() int {
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
	return 0
}
//-word=opt -numb=7 -fork -svar=flag