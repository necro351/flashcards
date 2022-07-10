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
	_, err = fmt.Scanf("%d", &number)
	if err != nil {
		log.Fatal(err)
	}
	if number >= len(topics) || number < 0 {
		log.Fatalf("Invalid topic number")
	}
	topic := topics[number]
	fmt.Printf("🍿 REHEARSE: %s\n", topic.Title)

	// Ask the user which stage they want to practice at. Depending on the stage
	// present many small sets of cards or fewer large sets of cards.
	fmt.Printf("Pick group size. There are %d cards. Size 8 is EASY, 32+ is HARD:\nGroup Size: ", len(topic.Cards))
	size := 0
	_, err = fmt.Scanf("%d", &size)
	if size < 1 {
		log.Fatalf("Invalid group size")
	}
	numgroups := len(topic.Cards) / size
	for i := 0; i < numgroups; i++ {
		fmt.Printf("✌️ REHEARSE STAGE %d/%d\n", i+1, numgroups)
		subtopic := topic
		start, end := i*size, (i+1)*size
		if end >= len(subtopic.Cards) {
			end = len(subtopic.Cards)
		}
		subtopic.Cards = subtopic.Cards[start:end]
		rehearse(subtopic)
	}

	// Congratulate the user for practicing a topic and exit.
	fmt.Printf("👏👏👏 GOOD JOB PRACTICING. Hope to see you again soon!\n")
}

// rehearse a topic. First shuffle the cards, then ask the user to type the
// answers to the given questions.  Show them their average score. Put any
// incorrect cards into a sepearate pile to go through again. Do this until
// there are no incorrect cards set aside.
func rehearse(topic Topic) {
	cards, incorrect := []permutedCard{}, []permutedCard{}
	for _, c := range topic.Cards {
		incorrect = append(incorrect, permutedCard{c, rand.Int()})
	}
	sort.Sort(PermutedCards(cards))
	right, total := 0, 0
	for len(incorrect) > 0 {
		cards = incorrect
		incorrect = nil
		for i, c := range cards {
			fmt.Printf("Q: %s\nA? ", c.card.Question)
			answer := ""
			_, err := fmt.Scanf("%s", &answer)
			if err != nil {
				log.Fatal(err)
			}
			total++
			exact, nontonal := match(answer, c.card.Answer)
			if exact || nontonal {
				right++
				if exact {
					fmt.Printf("✅ %d/%d CORRECT %d+%d cards left\n",
						right, total, len(cards)-i-1, len(incorrect))
				} else {
					fmt.Printf("☑️  %s: CLOSE: %d/%d CORRECT %d+%d cards left\n",
						c.card.Answer, right, total, len(cards)-i-1, len(incorrect))
				}
			} else {
				incorrect = append(incorrect, permutedCard{c.card, rand.Int()})
				fmt.Printf("⾮ %s: MISTAKEN: %d/%d CORRECT %d+%d cards left\n",
					c.card.Answer, right, total, len(cards)-i-1, len(incorrect))
			}
		}
		sort.Sort(PermutedCards(incorrect))
	}
}

// match returns the result of two equality tests. First if the input exactly
// matches the expected, and second if it would match if the Chinese tonal
// numbers were ignored.
func match(input, expected string) (bool, bool) {
	exact := input == expected
	nontonal := flatten(input) == flatten(expected)
	return exact, nontonal
}

// flatten all Chinese tonal numbers by changing them to '_'.
func flatten(s string) string {
	var res strings.Builder
	for i := range s {
		switch s[i : i+1] {
		case "1", "2", "3", "4":
			res.WriteString("_")
		default:
			res.WriteString(s[i : i+1])
		}
	}
	return res.String()
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
