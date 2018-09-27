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

//AboutPlug is the about plugin interface
type AboutPlug interface {
	About()
	StartConnection()
}

//PluginInfo Standard plugin information layout
type PluginInfo struct {
	PluginName string `json:"plugin_name,omitempty"`
	PluginType string `json:"plugin_type,omitempty"`
	PluginDesc string `json:"plugin_desc,omitempty"`
}

func main() {
	fmt.Println()
	fmt.Println("starting plugin testing application")
	fmt.Println("Loading plugins")

	allPlugins, err := filepath.Glob("plugins/*.so")
	if err != nil {
		log.Panicln(err)
	}

	for _, filename := range allPlugins {
		plug, err := plugin.Open(filename)
		if err != nil {
			panic(err)
		}

		fmt.Println(plug)

		// 2. look up a symbol (an exported function or variable)
		// in this case, variable Greeter
		symAbout, err := plug.Lookup("About")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(symAbout)

		// 3. Assert that loaded symbol is of a desired type
		// in this case interface type Greeter (defined above)
		var about AboutPlug
		about, ok := symAbout.(AboutPlug)
		if !ok {
			fmt.Println("unexpected type from module symbol")
			os.Exit(1)
		}

		var info PluginInfo

		// 4. use the module
		about.About()

		fmt.Println(info.PluginName)
	}

	fmt.Println("Plugins Loaded")

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
