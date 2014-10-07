package main

import (
	"flag"
	"fmt"
	"os"

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
		if pathErr, ok := err.(*os.PathError); ok {
			fmt.Printf("PathError: %v\n", pathErr)
			fmt.Printf("Op: %v\nPath: %v\n", pathErr.Op, pathErr.Path)
		} else {
			fmt.Printf("Error Type: %T", err)
			fmt.Println("Error: ", err)
		}
	} else {
		character := NewCharacter(cfg)
		fmt.Printf("Character Info...\n%v\n", character)
	}
}
