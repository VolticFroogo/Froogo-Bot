package help

import (
	"strings"

	"github.com/VolticFroogo/Froogo-Bot/cat"
	"github.com/VolticFroogo/Froogo-Bot/meme"
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

	case prefix + "help cat":
		cat.Help(s, m)
		break

	case prefix + "help":
		help(s, m)
		break

	default:
		unknownHelp(s, m)
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
			&discordgo.MessageEmbedField{
				Name:  "Cat",
				Value: cat.Description,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Froogo Bot by Harry.",
		},
	})
}

func unknownHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 0x003580,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Title:       "I don't know that feature",
		Description: "You must've mistyped a feature to get help on so I don't know what it is. To learn how to use help type in: `" + prefix + "help`.",
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Froogo Bot by Harry.",
		},
	})
}

// UnknownCommand is for a command that we don't know.
func UnknownCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 0x003580,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Title:       "I don't know that command",
		Description: "I haven't heard of that command before if if you want to get some help for commands type in: `" + prefix + "help`.\n\nIf you think this is a command that should be added message Froogo and he might add it for you.",
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Froogo Bot by Harry.",
		},
	})
}
