package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type about PluginInfo

type config struct {
	Token  string `json:"token,omitempty"`
	Prefix string `json:"prefix,omitempty"`
}

//PluginInfo struct to manage  plugin information
type PluginInfo struct {
	PluginName string `json:"plugin_name,omitempty"`
	PluginType string `json:"plugin_type,omitempty"`
	PluginDesc string `json:"plugin_desc,omitempty"`
}

//About this plugin
func (a about) About() {
	about := PluginInfo{
		PluginName: "discord",
		PluginType: "service",
		PluginDesc: "Service for Discord",
	}

	fmt.Println(about.PluginName)
}

//StartConnection is the starting command to set up and manage connections
func StartConnection() bool {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot token")
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
		return false
	}

	err = dg.Open()
	if err != nil {
		log.Fatalln("error opening connection,", err)
		return false
	}

	return true
}
