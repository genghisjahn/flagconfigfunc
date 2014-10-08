package main

import (
	"flag"
	"fmt"

	"code.google.com/p/gcfg"
)

var configPath = flag.String("class", "fighter", "Value of class to build.  Default is: fighter")

//You can put the flag definitions in the init function as well.
//Flag values are pulled in until you call flag.Parse()

func main() {
	var cfg Config
	flag.Parse()
	configpath := fmt.Sprintf("config/%v.gcfg", *configPath)
	if err := gcfg.ReadFileInto(&cfg, configpath); err != nil {
		fmt.Println("Error: ", err)
	} else {
		character := NewCharacter(cfg)
		fmt.Printf("Character Info...\n%v\n", character)
	}
}
