package deck

import (
	"fmt"
	// "io"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// Which is slice of strings
type deck []string

// So this function yhe receiver nahi banega because m koi apna banaya hua type use nahi krha even I am creating a new deck and returing it
func NewDeck() deck {
	card := deck{}

	cardName := []string{"Ace", "King", "Jack", "Queen"}
	cardNumber := []string{"One", "Two", "Three", "Four"}

	for i := 0; i < len(cardName); i++ {
		for j := 0; j < len(cardNumber); j++ {
			card = append(card, cardName[i]+" of "+cardNumber[j])
		}
	}

	return card
}

func (card deck) Print() {
	for i := 0; i < len(card); i++ {
		fmt.Println(card[i])
	}
}

func Deal(d deck, hand int) (deck, deck) {
	return d[:hand], d[hand:]
}

func (d deck) toString() string {
	ans := strings.Join(d, ",")
	return ans
}

func (d deck) SaveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func NewDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		//Option-1 - log the error and return a newdeck
		//Option-2 -log the error and quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	ans := strings.Split(string(bs), ",")
	// fmt.Println(ans)
	ans2 := deck(ans)

	return ans2
}

func (d deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	myrand := rand.New(source)

	for i := 0; i < len(d); i++ {
		pos := myrand.Intn(len(d) - 1)

		temp := d[i]
		d[i] = d[pos]
		d[pos] = temp
	}
}
