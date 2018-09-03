package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/VolticFroogo/Froogo-Bot/cat"
	"github.com/VolticFroogo/Froogo-Bot/help"
	"github.com/VolticFroogo/Froogo-Bot/meme"
	"github.com/VolticFroogo/Froogo-Bot/rage"
	"github.com/VolticFroogo/Froogo-Bot/tron"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	token, prefix string
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.StringVar(&prefix, "p", "", "Bot Prefix")
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
	err = dg.UpdateStatus(0, prefix+"help")
	if err != nil {
		log.Printf("Setting status error: %v", err)
	}

	help.Init(prefix)
	err = meme.Init(prefix)
	if err != nil {
		log.Printf("Initializing meme error: %v", err)
	}
	err = cat.Init(prefix)
	if err != nil {
		log.Printf("Initializing cat error: %v", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Froogo Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// Run when the messageCreate handler is called.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return // Ignore the message if it was sent by the bot.
	}

	if m.Author.Bot {
		return // Ignore the message if it was sent by a bot.
	}

	if strings.HasPrefix(m.Content, prefix) {
		err := s.ChannelMessageDelete(m.ChannelID, m.ID)
		if err != nil {
			log.Printf("Deleting command message error: %v", err)
		}

		if strings.HasPrefix(strings.ToLower(m.Content), prefix+"meme") {
			meme.Run(s, m)
		} else if strings.HasPrefix(strings.ToLower(m.Content), prefix+"help") {
			help.Run(s, m)
		} else if strings.HasPrefix(strings.ToLower(m.Content), prefix+"cat") {
			cat.Run(s, m)
		} else if strings.HasPrefix(strings.ToLower(m.Content), prefix+"tron") {
			tron.Run(s, m)
		} else if strings.HasPrefix(strings.ToLower(m.Content), prefix+"rage") {
			rage.Run(s, m)
		} else {
			help.UnknownCommand(s, m)
		}
	}
}

func guildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	var err error

	if m.GuildID == "428661403935571981" { // If we're in Finland.
		err = s.GuildMemberRoleAdd(m.GuildID, m.User.ID, "428665344173932561") // Gives user the peasant role.
	} else if m.GuildID == "371742166432743444" { // If we're in Sloth Stream.
		err = s.GuildMemberRoleAdd(m.GuildID, m.User.ID, "371743291311521803") // Gives user the member role.
	} else if m.GuildID == "428993299496435732" { // If we're in OJH Designs.
		err = s.GuildMemberRoleAdd(m.GuildID, m.User.ID, "429005223294402561") // Gives user the member role.
	}

	if err != nil {
		log.Printf("Adding role on join error: %v", err)
	}
}
