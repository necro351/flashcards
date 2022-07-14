package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
)

// Cards are grouped by topic. You want to learn a topic at a time
// until you have enough familiarity with all topics, that you can
// start mixing them together into a larger pool. Pick a topic and
// try to type in the corresponding answer for each card.

func main() {
	// Read all topics from the current directory into a sorted array of topics
	ents, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	topics := []Topic{}
	for _, ent := range ents {
		if !ent.IsDir() && strings.HasSuffix(ent.Name(), ".topic") {
			data, err := os.ReadFile(ent.Name())
			if err != nil {
				log.Print(err)
			} else {
				var t Topic
				err = json.Unmarshal(data, &t)
				if err != nil {
					log.Print(err)
				} else {
					topics = append(topics, t)
				}
			}
		}
	}
	sort.Sort(Topics(topics))

	// Ask the user which topic they want to practice
	fmt.Printf("Pick a topic to practice:\n")
	for i, t := range topics {
		fmt.Printf("%d: %s\n", i, t.Title)
	}
	fmt.Printf("Type the number of the topic to practice: ")
	number := 0
	_, err = fmt.Scanf("%d\n", &number)
	if err != nil {
		log.Fatal(err)
	}
	if number >= len(topics) || number < 0 {
		log.Fatalf("Invalid topic number")
	}
	topic := topics[number]

	// Shuffle, then ask the user to type the answers to the given questions.
	// Show them their average score.
	cards := []permutedCard{}
	for _, c := range topic.Cards {
		cards = append(cards, permutedCard{c, rand.Int()})
	}
	// wrongAnswers := cards
	sort.Sort(PermutedCards(cards))
	right, total := 0, 0
	for len(cards) != 0 {
		for _, c := range cards {
			// fmt.Printf("right:%d, card length:%d", right, len(cards))
			fmt.Printf("Q: %s\nA? ", c.card.Question)
			answer := ""
			_, err = fmt.Scanf("%s\n", &answer)
			if err != nil {
				log.Fatal(err)
			}
			total++
			if answer == c.card.Answer {
				right++
				for i := 0; i < len(cards); i++ {
					q := cards[i]
					if q.card.Question == c.card.Question {
						cards = append(cards[:i], cards[i+1:]...)
						i--
					}
				}
				fmt.Printf("✅ %d/%d CORRECT\n", right, total)
			} else {
				// wrongAnswers = append(cards, permutedCard{c.card, 1})
				fmt.Printf("⾮ Expected '%s'\n: %d/%d CORRECT\n", c.card.Answer, right, total)
			}
		}
		// cards := wrongAnswers
	}

	// Congratulate the user for practicing a topic and exit.
	fmt.Printf("👏👏👏 GOOD JOB PRACTICING. Hope to see you again soon!\n")
}

type Topic struct {
	Title string `json:"title"`
	Cards []Card `json:"cards"`
}

type Card struct {
	Question string `json:"q"`
	Answer   string `json:"a"`
}

type permutedCard struct {
	card   Card
	randid int
}

type Topics []Topic

func (t Topics) Len() int           { return len(t) }
func (t Topics) Less(i, j int) bool { return t[i].Title < t[j].Title }
func (t Topics) Swap(i, j int) {
	temp := t[i]
	t[i] = t[j]
	t[j] = temp
}

type PermutedCards []permutedCard

func (p PermutedCards) Len() int           { return len(p) }
func (p PermutedCards) Less(i, j int) bool { return p[i].randid < p[j].randid }
func (p PermutedCards) Swap(i, j int) {
	temp := p[i]
	p[i] = p[j]
	p[j] = temp
}
