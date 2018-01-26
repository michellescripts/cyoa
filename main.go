package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

var currentArc = "intro"

func getChapter() {
	for i, c := range stories {
		if i == currentArc {
			var answer string
			fmt.Println(c.Title)
			for _, p := range c.Paragraphs {
				fmt.Println(p)
				fmt.Println()
			}
			if len(c.Options) != 0 {
				fmt.Println("What do you do?")
				optionMap := make(map[int]string, 3) // TODO: fix the map to reset for each option
				for i, o := range c.Options {
					// new map of options to int
					fmt.Println(i+1, ":", o.Text)
					optionMap[i+1] = o.Chapter
				}
				fmt.Scanln(&answer)
				index, _ := strconv.Atoi(answer)
				fmt.Println("Heading to ", optionMap[index])
				fmt.Println()
				fmt.Println()
				fmt.Println()
				currentArc = optionMap[index]
				getChapter()
			}
			if len(c.Options) == 0 {
				return
			}
		}
	}
}

var stories Story

func getStory() {
	story, _ := os.Open("./gopher.json")
	byteValue, _ := ioutil.ReadAll(story)
	json.Unmarshal(byteValue, &stories)

	getChapter()
	defer story.Close()
}

func main() {
	getStory()
}
