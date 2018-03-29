package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/VolticFroogo/Finland-Bot/meme"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	token  string
	prefix = ">"
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.AddHandler(guildMemberAdd)
	dg.UpdateStatus(0, prefix+"help")
	meme.Init(prefix)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Finland Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// Run when the messageCreate handler is called.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return // Ignore the message if it was send by the bot.
	}

	if strings.HasPrefix(m.Content, prefix) {
		if strings.HasPrefix(m.Content, prefix+"meme") {
			meme.Run(s, m)
		} else if strings.HasPrefix(m.Content, prefix+"help") {
			meme.Help(s, m)
		}
	}
}

func guildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	s.GuildMemberRoleAdd(m.GuildID, m.User.ID, "428665344173932561") // Gives user the peasant role.
}
