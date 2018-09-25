package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

//PluginInfo Standard plugin information layout
type PluginInfo struct {
	PluginName string `json:"plugin_name,omitempty"`
	PluginType string `json:"plugin_type,omitempty"`
	PluginDesc string `json:"plugin_desc,omitempty"`
}

func main() {
	fmt.Println("starting plugin testing application")
	fmt.Println("Loading plugins")

	allPlugins, err := filepath.Glob("plugins/*.so")
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("Plugins Loaded")

	for _, filename := range allPlugins {
		p, err := plugin.Open(filename)
		if err != nil {
			panic(err)
		}

		symbol, err := p.Lookup("About")
		if err != nil {
			panic(err)
		}

		aboutFunc, ok := symbol.(func() PluginInfo)
		if !ok {
			panic("Plugin has no 'About() PluginInfo' function")
		}

		fmt.Println("Plugin " + aboutFunc().PluginName + " loaded")
	}

	log.Println("Testing application is no running. Send 'shutdown' or 'ctrl + c' to stop the bot.")

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("cannot read from stdin: %s", err)
		}

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if line == "shutdown" {
			log.Println("Shutting down the bot.")
			return
		}
		log.Println("inbound command is " + line)
	}
}
