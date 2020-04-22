package v1

import (
	"log"
	"strconv"
)

type Poker struct {}

func (p *Poker) compareHands(p1 string, p2 string) {
	log.Printf("Comparing '%v' with '%v'", p1, p2)
	p.isStraightFlush(p1)
}


func (p *Poker) isStraightFlush(hand string) {
	log.Print("Checking for straight flush")
}

func (p *Poker) getCardValue(card string) (int, int, error) {
	val, err := strconv.Atoi(string(card[0]))
	if err!=nil {
		switch string(card[0]) {
		case "T":
			val = 10
		case "J":
			val = 11;
		case "Q":
			val = 12;
		case "K":
			val = 13;
		case "A":
			val = 14;
		default:
			log.Fatalf("Unable to parse card: %v", card)
		}
	}

	log.Print(val)

	return 0, 0, nil
}