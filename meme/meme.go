package meme

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Meme is a meme.
type Meme struct {
	Command, Title, Description, Image, EmbedDescription string
}

var (
	prefix string
	// Memes is a slice of meme.
	Memes map[string]Meme
)

// Init sets the variables needed for the meme package.
func Init(setPrefix string) {
	prefix = setPrefix

	Memes = make(map[string]Meme)

	Memes[">meme no u"] = Meme{
		Command:          ">meme no u",
		Title:            "No U",
		Description:      "The No U trap card.",
		EmbedDescription: "The No U trap card has been activated.",
		Image:            "https://pics.me.me/trap-no-u-trap-card-ysii-en014-negate-the-effect-29426981.png",
	}
}

// RunMeme runs the meme's embed message.
func (meme Meme) RunMeme(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 0xB23C84,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Title:       meme.Title,
		Description: meme.Description,
		Image: &discordgo.MessageEmbedImage{
			URL: meme.Image,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Finland Bot by Froogo.",
		},
	})
}

// Help runs the help command.
func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 0x003580,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Title:       "Meme Help",
		Description: "The meme feature allows you to easily send dank memes in Finland.",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "Find memes",
				Value: "To find some memes just run `" + prefix + "meme list` and it will respond with a list of memes.",
			},
			&discordgo.MessageEmbedField{
				Name:  "Example",
				Value: "Here's how you'd run the no u meme: `" + prefix + "meme no u`",
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Finland Bot by Froogo.",
		},
	})
}

// Run chooses the meme function to run and then runs it.
func Run(s *discordgo.Session, m *discordgo.MessageCreate) {
	if meme, ok := Memes[m.Content]; ok {
		meme.RunMeme(s, m)
		return
	}

	if strings.HasPrefix(m.Content, prefix+"meme list") {
		list(s, m)
	}
}

func list(s *discordgo.Session, m *discordgo.MessageCreate) {
	var list string
	for meme := range Memes {
		list += Memes[meme].Title + ": " + Memes[meme].Description + "\n"
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 0x003580,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Title:       "List of Memes",
		Description: list,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Finland Bot by Froogo.",
		},
	})
}
