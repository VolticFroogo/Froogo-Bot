package cat

import (
	"encoding/gob"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Picture is a cat picture.
type Picture struct {
	Title, EmbedDescription, Image string
}

var (
	prefix string

	// Pictures is an array of all pictures.
	Pictures []Picture

	// Description of the meme package.
	Description = "Get a random picture of one our cats."
)

// Init sets the variables needed for the help package and runs some initialization scripts.
func Init(setPrefix string) (err error) {
	rand.Seed(time.Now().Unix())
	prefix = setPrefix

	catsFile, err := os.Open("cat/cats.gob")
	if err != nil {
		return
	}

	err = gob.NewDecoder(catsFile).Decode(&Pictures)
	return
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// Run chooses a random cat picture and then sends it.
func Run(s *discordgo.Session, m *discordgo.MessageCreate) {
	catPicture := Pictures[random(0, len(Pictures))]

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 0x00aaff,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Title:       catPicture.Title,
		Description: catPicture.EmbedDescription,
		Image: &discordgo.MessageEmbedImage{
			URL: catPicture.Image,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Froogo Bot by Harry.",
		},
	})

	if err != nil {
		log.Printf("Cat sending embed message error: %v", err)
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
		Title:       "Cat Help",
		Description: "The meme feature allows you to easily see our cat pictures.",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "Example",
				Value: "Here's how you'd send a cat picture: `" + prefix + "cat`",
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Froogo Bot by Harry.",
		},
	})

	if err != nil {
		log.Printf("Cat sending help embed message error: %v", err)
	}
}
