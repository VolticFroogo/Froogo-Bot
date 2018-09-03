package tron

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	tickerURL = "https://api.coinmarketcap.com/v2/ticker/1958"
)

var (
	prefix string
)

// TickerAPI is a JSON response from coinmarketcap.
type TickerAPI struct {
	Data struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Symbol            string      `json:"symbol"`
		WebsiteSlug       string      `json:"website_slug"`
		Rank              int         `json:"rank"`
		CirculatingSupply float64     `json:"circulating_supply"`
		TotalSupply       float64     `json:"total_supply"`
		MaxSupply         interface{} `json:"max_supply"`
		Quotes            struct {
			USD struct {
				Price            float64 `json:"price"`
				Volume24H        float64 `json:"volume_24h"`
				MarketCap        float64 `json:"market_cap"`
				PercentChange1H  float64 `json:"percent_change_1h"`
				PercentChange24H float64 `json:"percent_change_24h"`
				PercentChange7D  float64 `json:"percent_change_7d"`
			} `json:"USD"`
		} `json:"quotes"`
		LastUpdated int `json:"last_updated"`
	} `json:"data"`
	Metadata struct {
		Timestamp int         `json:"timestamp"`
		Error     interface{} `json:"error"`
	} `json:"metadata"`
}

// Init sets the variables needed for the help package and runs some initialization scripts.
func Init(setPrefix string) {
	prefix = setPrefix
}

// Run gets the Tron information and then sends it.
func Run(s *discordgo.Session, m *discordgo.MessageCreate) {
	client := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, tickerURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "froogo-bot")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	tron := TickerAPI{}
	err = json.Unmarshal(body, &tron)
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 0x003580,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Title:       "Tronix",
		Description: "Some information of Tronix in USD.",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "Market Cap Basic Information",
				Value: fmt.Sprintf("Rank: %v | Value: %v", tron.Data.Rank, tron.Data.Quotes.USD.Price),
			},
			&discordgo.MessageEmbedField{
				Name:  "Value Change",
				Value: fmt.Sprintf("Hour: %v | Day: %v | Week: %v", tron.Data.Quotes.USD.PercentChange1H, tron.Data.Quotes.USD.PercentChange24H, tron.Data.Quotes.USD.PercentChange7D),
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Froogo Bot by Harry.",
		},
	})

	if err != nil {
		log.Printf("Tron sending embed message error: %v", err)
	}
}
