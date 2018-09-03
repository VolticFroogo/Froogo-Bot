package meme

import (
	"encoding/gob"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Meme is a meme.
type Meme struct {
	Command, Title, Description, Image, EmbedDescription string
	Color                                                int
}

var (
	prefix string

	// Memes is a slice of meme.
	Memes map[string]Meme

	// Description of the meme package.
	Description = "An easy way to show some dank memes."
)

// Init sets the variables needed for the meme package and runs some initialization scripts.
func Init(setPrefix string) (err error) {
	prefix = setPrefix

	Memes = make(map[string]Meme)

	memesFile, err := os.Open("meme/memes.gob")
	if err != nil {
		return
	}

	err = gob.NewDecoder(memesFile).Decode(&Memes)

	return
}

// RunMeme runs the meme's embed message.
func (meme Meme) RunMeme(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: meme.Color,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Title:       meme.Title,
		Description: meme.EmbedDescription,
		Image: &discordgo.MessageEmbedImage{
			URL: meme.Image,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Froogo Bot by Harry.",
		},
	})

	if err != nil {
		log.Printf("Sending meme embed message error: %v", err)
	}
}

// Help runs the help command.
func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 0x003580,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Title:       "Meme Help",
		Description: "The meme feature allows you to easily send dank memes in Finland.",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "Find Memes",
				Value: "To find some memes just run `" + prefix + "meme list` and it will respond with a list of memes.",
			},
			&discordgo.MessageEmbedField{
				Name:  "Example",
				Value: "Here's how you'd run the no u meme: `" + prefix + "meme no u`",
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Froogo Bot by Harry.",
		},
	})

	if err != nil {
		log.Printf("Sending help meme embed message error: %v", err)
	}
}

// Run chooses the meme function to run and then runs it.
func Run(s *discordgo.Session, m *discordgo.MessageCreate) {
	if meme, ok := Memes[strings.Trim(strings.ToLower(m.Content), prefix+"meme ")]; ok {
		meme.RunMeme(s, m) // Run the meme they asked for.
		return
	}

	if strings.HasPrefix(strings.ToLower(m.Content), prefix+"meme list") {
		list(s, m) // Show the list of memes.
		return
	}

	if strings.ToLower(m.Content) == prefix+"meme" {
		Help(s, m) // They ran meme show some help.
		return
	}

	unknownMeme(s, m) // Don't know what they want to see.
}

func list(s *discordgo.Session, m *discordgo.MessageCreate) {
	var list string
	for meme := range Memes {
		list += Memes[meme].Title + ": " + Memes[meme].Description + "\n"
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 0x003580,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Title:       "List of Memes",
		Description: list,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Froogo Bot by Harry.",
		},
	})

	if err != nil {
		log.Printf("Listing meme embed message error: %v", err)
	}
}

func unknownMeme(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 0x003580,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Title:       "I don't know that meme",
		Description: "I haven't heard of that meme before, if you want it adding message Froogo and if he can be bothered he'll add it for you.\n\nFor a list of memes type in: `" + prefix + "meme list`.",
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Froogo Bot by Harry.",
		},
	})

	if err != nil {
		log.Printf("Unknown meme embed message error: %v", err)
	}
}
