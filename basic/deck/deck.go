package deck

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"math/rand"
)

type deck []string

func Deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}

func NewDeck() deck {
	cards := deck{}
	cardSuits := []string{"Diamonds", "Hearts", "Spades", "Clubs"}
	cardValues := []string{
		"Ace",
		"Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0777)
}

func (d deck) NewDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) Shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		position := r.Intn(len(d) - 1)
		temp := d[i]
		d[i] = d[position]
		d[position] = temp
	}
	return d
}