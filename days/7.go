package days

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type hand struct {
	cards []string
	bet   uint64
}

var strength = []string{
	"A",
	"K",
	"Q",
	"J",
	"T",
	"9",
	"8",
	"7",
	"6",
	"5",
	"4",
	"3",
	"2",
}

func Run7() {
	lines := convertInputToSlices(7)
	hands := []hand{}

	for _, line := range lines {
		hands = append(hands, splitBets(line))
	}

	fiveOfAKind := []hand{}
	fourOfAKind := []hand{}
	fullHouse := []hand{}
	threeOfAKind := []hand{}
	twoPair := hand{}
	pair := []hand{}
	highcard := []hand{}

	for _, str := range strength {
		for _, hand := range hands {
			if ok, _ := regexp.Match(fmt.Sprintf("[%s]{5}", str), []byte(hand.cards[0])); ok {
				fiveOfAKind = append(fiveOfAKind, hand)
			}
			if ok, _ := regexp.Match(fmt.Sprintf("[%s]{4}", str), []byte(hand.cards[0])); ok {
				fiveOfAKind = append(fiveOfAKind, hand)
			}
			if ok, _ := regexp.Match(fmt.Sprintf("[%s,%s]{3}{2}", str, str[+1]), []byte(hand.cards[0])); ok {
				fiveOfAKind = append(fiveOfAKind, hand)
			}
			if ok, _ := regexp.Match(fmt.Sprintf("[%s]{5}", str), []byte(hand.cards[0])); ok {
				fiveOfAKind = append(fiveOfAKind, hand)
			}
			if ok, _ := regexp.Match(fmt.Sprintf("[%s]{5}", str), []byte(hand.cards[0])); ok {
				fiveOfAKind = append(fiveOfAKind, hand)
			}
			if ok, _ := regexp.Match(fmt.Sprintf("[%s]{5}", str), []byte(hand.cards[0])); ok {
				fiveOfAKind = append(fiveOfAKind, hand)
			}
		}
	}
}

func splitBets(line string) hand {
	card := strings.Split(line, " ")
	h := hand{}
	for i, c := range card {
		newString := strings.Trim(c, "\n")
		if i == 0 {
			h.cards = []string{newString}
			// for _, letter := range newString {
			// 	h.cards = append(h.cards, string(letter))
			// }
		} else {
			bet, err := strconv.ParseUint(newString, 10, 64)
			if err != nil {
				fmt.Println(err)
			}

			h.bet = bet
		}
	}

	return h
}
