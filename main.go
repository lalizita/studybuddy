/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/lalizita/studybuddy/cmd"
	"github.com/lalizita/studybuddy/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
