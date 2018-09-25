package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

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
func About() PluginInfo {
	Info := PluginInfo{
		PluginName: "irc",
		PluginType: "service",
		PluginDesc: "Service for IRC",
	}
	return Info
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
