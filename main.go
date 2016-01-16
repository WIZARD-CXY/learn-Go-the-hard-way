package main

import (
	"encoding/json"
	"log"
	"os"
)

type Cheat struct {
	Name     string
	Contents []Content
}

//TODO:complete the struct to be encoded by encoding/json.
type Content struct {
	Comment string
	Command string
}

func CheatSheet(command string) string {
	file, err := os.Open("./commands.json")
	if err != nil {
		panic(err)
	}
	var cheats []Cheat
	dec := json.NewDecoder(file)
	dec.Decode(&cheats)
	var out string
	//TODO:find the name of which cheatsheet matchs command
	//and add to out.
	// ...
	for _, item := range cheats {
		if item.Name == command {
			for _, output := range item.Contents {
				out += output.Comment
				out += "\n"
				out += output.Command
			}
		}

		break
	}
	out += "\n"
	return out
}

func main() {
	args := os.Args
	if len(args) != 1 {
		log.Fatal("want one argument")
	}

	println(`Unix has a lot of commands to remember.
To help us search the command quickly,we will create a small cheatsheet command.
We will store the commands as json.In this exercise you can play with Go IO and json encoding.`)
}
