package main

import (
	"fmt"
	"flag"
)

var name = flag.String("name","World!","A name to say hello to.")
var spanish bool

func init(){
	flag.BoolVar(&spanish,"spanish",false,"Use Spanish Language.")
	flag.BoolVar(&spanish,"s",false,"Use Spanish Language.")
}

func main(){
	flag.Parse()

	if spanish == true {
		fmt.Printf("Hola %s!\n",*name)
	} else {
		fmt.Printf("Hello %s!\n",*name)
}
}
