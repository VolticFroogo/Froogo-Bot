package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
)

type catPicture struct {
	Title, EmbedDescription, Image string
}

type meme struct {
	Command, Title, Description, Image, EmbedDescription string
	Color                                                int
}

func main() {
	catPictures := []catPicture{
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/mkpq1Zo.jpg",
		},
		catPicture{
			Title:            "Rosie",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/Yp6ngUn.jpg",
		},
		catPicture{
			Title:            "Dori and Rosie",
			EmbedDescription: "Harry's Cats.",
			Image:            "https://i.imgur.com/zXD9Oxc.jpg",
		},
		catPicture{
			Title:            "Dori and Rosie",
			EmbedDescription: "Harry's Cats.",
			Image:            "https://i.imgur.com/zXD9Oxc.jpg",
		},
		catPicture{
			Title:            "https://i.imgur.com/QTRWBUs.jpg",
			EmbedDescription: "Harry's Cat.",
			Image:            "Dori",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/4YEo6Zj.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/coM4NmM.jpg",
		},
		catPicture{
			Title:            "Dori and Rosie",
			EmbedDescription: "Harry's Cats.",
			Image:            "https://i.imgur.com/Sjnz50X.jpg",
		},
		catPicture{
			Title:            "Dori and Rosie",
			EmbedDescription: "Harry's Cats.",
			Image:            "https://i.imgur.com/PUouD4G.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/H7W2ZUS.jpg",
		},
		catPicture{
			Title:            "Rosie",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/dDqMnJg.jpg",
		},
		catPicture{
			Title:            "Rosie",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/7I6VPmF.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/AdLeUvU.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/B4NJ0wF.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/mORrLTh.jpg",
		},
		catPicture{
			Title:            "Rosie",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/0EfZ1fr.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/pYHgmfr.jpg",
		},
		catPicture{
			Title:            "Dori and Rosie",
			EmbedDescription: "Harry's Cats.",
			Image:            "https://i.imgur.com/sUvY5D4.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/xEtnYby.jpg",
		},
		catPicture{
			Title:            "Dori and Rosie",
			EmbedDescription: "Harry's Cats.",
			Image:            "https://i.imgur.com/gCZ0EYn.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/xt9sIW1.jpg",
		},
		catPicture{
			Title:            "Dori and Rosie",
			EmbedDescription: "Harry's Cats.",
			Image:            "https://i.imgur.com/kc5AyjU.jpg",
		},
		catPicture{
			Title:            "Rosie",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/773kn2P.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/p6AiXno.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/wB00Eg7.jpg",
		},
		catPicture{
			Title:            "Rosie",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/2kJGcfI.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/hHLQ8pm.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/0qA1w13.jpg",
		},
		catPicture{
			Title:            "Rosie",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/8BTWx4T.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/lkj45Wz.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/uxAxAZz.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/pBYvwiX.jpg",
		},
		catPicture{
			Title:            "Dori",
			EmbedDescription: "Harry's Cat.",
			Image:            "https://i.imgur.com/yG0lFyX.jpg",
		},
		catPicture{
			Title:            "Mr. Pink",
			EmbedDescription: "Jolywog's Cat.",
			Image:            "https://i.imgur.com/g6mk02L.jpg",
		},
		catPicture{
			Title:            "Mr. Pink",
			EmbedDescription: "Jolywog's Cat.",
			Image:            "https://i.imgur.com/q0oVejo.jpg",
		},
		catPicture{
			Title:            "Mr. Pink",
			EmbedDescription: "Jolywog's Cat.",
			Image:            "https://i.imgur.com/0i1DqAu.jpg",
		},
		catPicture{
			Title:            "Ripple",
			EmbedDescription: "Jolywog's Cat.",
			Image:            "https://i.imgur.com/wxJd2xZ.jpg",
		},
		catPicture{
			Title:            "Ripple",
			EmbedDescription: "Jolywog's Cat.",
			Image:            "https://i.imgur.com/aCmnqHa.jpg",
		},
		catPicture{
			Title:            "Ripple",
			EmbedDescription: "Jolywog's Cat.",
			Image:            "https://i.imgur.com/mchYvlI.jpg",
		},
		catPicture{
			Title:            "Ripple",
			EmbedDescription: "Jolywog's Cat.",
			Image:            "https://i.imgur.com/2NRSrPo.jpg",
		},
		catPicture{
			Title:            "Mr. Pink",
			EmbedDescription: "Jolywog's Cat.",
			Image:            "https://i.imgur.com/XTSBJuV.jpg",
		},
		catPicture{
			Title:            "Mr. Pink",
			EmbedDescription: "Jolywog's Cat.",
			Image:            "https://i.imgur.com/aqrHh0L.jpg",
		},
		catPicture{
			Title:            "Ripple",
			EmbedDescription: "Jolywog's Cat.",
			Image:            "https://i.imgur.com/nNSRIAT.jpg",
		},
		catPicture{
			Title:            "Mr. Pink",
			EmbedDescription: "Jolywog's Cat.",
			Image:            "https://i.imgur.com/nHXBoAR.jpg",
		},
		catPicture{
			Title:            "Mr. Pink",
			EmbedDescription: "Jolywog's Cat.",
			Image:            "https://i.imgur.com/Zm6foFZ.jpg",
		},
		catPicture{
			Title:            "Ben",
			EmbedDescription: "Cat's Cat.",
			Image:            "https://i.imgur.com/DxJOujp.jpg",
		},
		catPicture{
			Title:            "Ben",
			EmbedDescription: "Cat's Cat.",
			Image:            "https://i.imgur.com/Uf2DTQ3.jpg",
		},
		catPicture{
			Title:            "Ben",
			EmbedDescription: "Cat's Cat.",
			Image:            "https://i.imgur.com/2FUcJjM.jpg",
		},
		catPicture{
			Title:            "Ben",
			EmbedDescription: "Cat's Cat.",
			Image:            "https://i.imgur.com/Xt9h0cr.jpg",
		},
		catPicture{
			Title:            "Ben",
			EmbedDescription: "Cat's Cat.",
			Image:            "https://i.imgur.com/U7Bo4uJ.jpg",
		},
		catPicture{
			Title:            "Ben",
			EmbedDescription: "Cat's Cat.",
			Image:            "https://i.imgur.com/5ZR191T.jpg",
		},
		catPicture{
			Title:            "Ben",
			EmbedDescription: "Cat's Cat.",
			Image:            "https://imgur.com/0h9z1Hq.jpg",
		},
		catPicture{
			Title:            "Ben",
			EmbedDescription: "Cat's Cat.",
			Image:            "https://imgur.com/8AzW9jf.jpg",
		},
		catPicture{
			Title:            "Ben",
			EmbedDescription: "Cat's Cat.",
			Image:            "https://imgur.com/RRjtYso.jpg",
		},
		catPicture{
			Title:            "Ben",
			EmbedDescription: "Cat's Cat.",
			Image:            "https://imgur.com/vFhaStN.jpg",
		},
		catPicture{
			Title:            "Ben",
			EmbedDescription: "Cat's Cat.",
			Image:            "https://i.imgur.com/UPjCcsr.jpg",
		},
		catPicture{
			Title:            "Snoop",
			EmbedDescription: "Connor's Cat.",
			Image:            "https://i.imgur.com/a81Gb6O.jpg",
		},
	}

	var catFile bytes.Buffer                            // Buffer that's saved to file.
	err := gob.NewEncoder(&catFile).Encode(catPictures) // Encode catPictures to buffer.
	if err != nil {
		log.Printf("Error encoding catPictures: %v", err)
		return
	}

	err = ioutil.WriteFile("../cat/cats.gob", catFile.Bytes(), 0666)
	if err != nil {
		log.Printf("Writing to file error: %v", err)
		return
	}

	memes := make(map[string]meme)

	memes["no u"] = meme{
		Command:          "no u",
		Title:            "No U",
		Description:      "The No U trap card.",
		EmbedDescription: "The No U trap card has been activated.",
		Image:            "https://i.imgur.com/R6z9LAz.png",
		Color:            0xB23C84,
	}

	var memeFile bytes.Buffer                     // Buffer that's saved to file.
	err = gob.NewEncoder(&memeFile).Encode(memes) // Encode meme to buffer.
	if err != nil {
		log.Printf("Error encoding memes: %v", err)
		return
	}

	err = ioutil.WriteFile("../meme/memes.gob", memeFile.Bytes(), 0666)
	if err != nil {
		log.Printf("Writing to file error: %v", err)
	}
}
