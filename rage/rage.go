package rage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	videoURL = "https://www.googleapis.com/youtube/v3/videos?part=statistics&id=XRvEnFiaLm8&key=AIzaSyBX--xoiGeC56rX_Lq4YqF1joXJLWkQpSE"
)

// VideoStatistics is the statistics struct from YouTube v3 API.
type VideoStatistics struct {
	Kind     string `json:"kind"`
	Etag     string `json:"etag"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind       string `json:"kind"`
		Etag       string `json:"etag"`
		ID         string `json:"id"`
		Statistics struct {
			ViewCount     string `json:"viewCount"`
			LikeCount     string `json:"likeCount"`
			DislikeCount  string `json:"dislikeCount"`
			FavoriteCount string `json:"favoriteCount"`
			CommentCount  string `json:"commentCount"`
		} `json:"statistics"`
	} `json:"items"`
}

// Run gets the Rage video statistics and then sends it.
func Run(s *discordgo.Session, m *discordgo.MessageCreate) {
	client := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, videoURL, nil)
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

	rage := VideoStatistics{}
	err = json.Unmarshal(body, &rage)
	if err != nil {
		log.Fatal(err)
	}

	likes, _ := strconv.Atoi(rage.Items[0].Statistics.LikeCount)
	dislikes, _ := strconv.Atoi(rage.Items[0].Statistics.DislikeCount)

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 0x003580,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Title:       "Rage 2 E3 Trailer Theme Earrape",
		Description: "Some statistics from Shad0w's terrible earrape video.",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "Important Statistics",
				Value: fmt.Sprintf("Views: %v | Likes: %v | Dislikes: %v | Like Ratio: %v", rage.Items[0].Statistics.ViewCount, likes, dislikes, likes/dislikes),
			},
			&discordgo.MessageEmbedField{
				Name:  "Not so importants statistics",
				Value: fmt.Sprintf("Favourites: %v | Comments: %v", rage.Items[0].Statistics.FavoriteCount, rage.Items[0].Statistics.CommentCount),
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Froogo Bot by Harry.",
		},
	})

	if err != nil {
		log.Printf("Rage sending embed message error: %v", err)
	}
}
