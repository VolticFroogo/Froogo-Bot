package help

import (
	"strings"

	"github.com/VolticFroogo/Finland-Bot/meme"
	"github.com/bwmarrin/discordgo"
)

var prefix string

// Init sets the variables needed for the help package and runs some initialization scripts.
func Init(setPrefix string) {
	prefix = setPrefix
}

// Run chooses the help function to run and then runs it.
func Run(s *discordgo.Session, m *discordgo.MessageCreate) {
	switch strings.ToLower(m.Content) {
	case prefix + "help meme":
		meme.Help(s, m)
		break

	default:
		help(s, m)
	}
}

func help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 0x003580,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Title:       "Help",
		Description: "If you want to get more information for a specific feature type in `" + prefix + "help FEATURE`.\nFor example, `" + prefix + "help meme`.",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "Meme",
				Value: meme.Description,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Finland Bot by Froogo.",
		},
	})
}
