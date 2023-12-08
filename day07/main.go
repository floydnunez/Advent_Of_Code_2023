package main

import (
	"2023/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards           string
	translatedCards string
	bid             int
	kind            int
	joker           bool
}

func main() {
	fileLines := util.ReadFileIntoArray("day07/input.txt")
	var hands []hand
	for _, line := range fileLines {
		thisHand := parseHand(line, false)
		hands = append(hands, thisHand)
	}
	total := calcTotal(hands)
	fmt.Println("Part 1:", total) //253933213

	var jokersHands []hand
	for _, line := range fileLines {
		thisHand := parseHand(line, true)
		jokersHands = append(jokersHands, thisHand)
	}
	total2 := calcTotal(jokersHands)
	//253148077 too low
	fmt.Println("Part 2:", total2) //253473930

}

func calcTotal(hands []hand) int {
	sort.SliceStable(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j])
	})
	total := 0
	for i, thisHand := range hands {
		rank := i + 1
		partial := rank * thisHand.bid
		//if thisHand.joker {
		//fmt.Println("hand", thisHand.cards, thisHand.translatedCards, rank, thisHand.bid, partial, "kind:", thisHand.kind, "translated:", translateKind(thisHand.kind))
		//}
		total += partial
	}
	return total
}

func translateKind(kind int) string {
	switch kind {
	case 1:
		return "Five of a Kind"
	case 2:
		return "Four of a Kind"
	case 3:
		return "Full House"
	case 4:
		return "Three of a Kind"
	case 5:
		return "Two Pairs"
	case 6:
		return "Pair"
	case 7:
		return "High Card"
	default:
		return "uuuuuh"
	}
}

func compareHands(a hand, b hand) bool {
	if a.kind == b.kind {
		for i := range a.translatedCards {
			if a.translatedCards[i] != b.translatedCards[i] {
				return a.translatedCards[i] < b.translatedCards[i]
			}
		}
		return true
	} else {
		return a.kind > b.kind
	}
}

func parseHand(line string, jokers bool) hand {
	parts := strings.Split(line, " ")
	cards := parts[0]
	bid, _ := strconv.Atoi(parts[1])
	translatedCards := translateToOrderable(cards, jokers)
	kind := getKindOfHand(cards, jokers)
	result := hand{cards: cards, translatedCards: translatedCards, bid: bid, kind: kind, joker: false}
	if strings.Contains(cards, "J") {
		result.joker = true
	}
	return result
}

func getKindOfHand(cards string, jokers bool) int {
	cards = orderHand(cards)
	diffCards := countDiffCards(cards, jokers)
	if isFiveOfAKind(diffCards) {
		return 1
	} else if isFourOfAKind(diffCards) {
		return 2
	} else if isFullHouse(diffCards) {
		return 3
	} else if isThreeOfAKind(diffCards) {
		return 4
	} else if isTwoPair(diffCards) {
		return 5
	} else if isOnePair(diffCards) {
		return 6
	} else if isHighCard(diffCards) {
		return 7
	}
	return 999
}

func isHighCard(cards map[string]int) bool {
	return len(cards) == 5
}

func isOnePair(diffCards map[string]int) bool {
	if len(diffCards) == 4 {
		maxCount := 0
		for _, count := range diffCards {
			if count > maxCount {
				maxCount = count
			}
		}
		if maxCount == 2 {
			return true
		}
	}
	return false
}

func isTwoPair(diffCards map[string]int) bool {
	if len(diffCards) == 3 {
		maxCount := 0
		for _, count := range diffCards {
			if count > maxCount {
				maxCount = count
			}
		}
		if maxCount == 2 {
			return true
		}
	}
	return false
}

func isThreeOfAKind(diffCards map[string]int) bool {
	if len(diffCards) == 3 {
		for _, count := range diffCards {
			if count == 3 {
				return true
			}
		}
	}
	return false
}

func isFullHouse(diffCards map[string]int) bool {
	if len(diffCards) == 2 {
		for _, count := range diffCards {
			if count == 3 {
				return true
			}
		}
	}
	return false
}

func isFourOfAKind(diffCards map[string]int) bool {
	if len(diffCards) == 2 {
		for _, count := range diffCards {
			if count == 4 {
				return true
			}
		}
	}
	return false
}

func isFiveOfAKind(diffCards map[string]int) bool {
	return len(diffCards) == 1
}

func countDiffCards(cards string, jokers bool) map[string]int {
	eachCard := strings.Split(cards, "")
	cardCount := make(map[string]int)
	for _, card := range eachCard {
		if counting, ok := cardCount[card]; ok {
			counting++
			cardCount[card] = counting
		} else {
			cardCount[card] = 1
		}
	}
	if jokers && strings.Contains(cards, "J") {
		var maxCard string
		maxCount := 0
		for card, count := range cardCount {
			if count > maxCount && card != "J" {
				maxCount = count
				maxCard = card
			}
		}
		if maxCount == 0 {
			maxCount = 1
			maxCard = "J"
		}
		//fmt.Println("maxcount", maxCount, "card:", maxCard)
		jokerizedHand := strings.Replace(cards, "J", maxCard, -1)
		//fmt.Println("jokerized:", jokerizedHand)
		return countDiffCards(jokerizedHand, false)
	}
	return cardCount
}

func orderHand(cards string) string {
	translatedCards := translateToOrderable(cards, false)
	eachCard := strings.Split(translatedCards, "")
	sort.SliceStable(eachCard, func(i, j int) bool {
		return eachCard[i] > eachCard[j]
	})
	sortedCards := ""
	for _, card := range eachCard {
		sortedCards += card
	}
	untranslatedCards := untranslate(sortedCards)
	return untranslatedCards
}

func untranslate(cards string) string {
	cards = strings.Replace(cards, "0", "J", -1)
	cards = strings.Replace(cards, "U", "J", -1)
	cards = strings.Replace(cards, "0", "J", -1)
	cards = strings.Replace(cards, "V", "Q", -1)
	cards = strings.Replace(cards, "W", "K", -1)
	cards = strings.Replace(cards, "X", "A", -1)
	return cards
}

func translateToOrderable(cards string, jokers bool) string {
	if jokers {
		cards = strings.Replace(cards, "J", "0", -1)
	} else {
		cards = strings.Replace(cards, "J", "U", -1)
	}
	cards = strings.Replace(cards, "Q", "V", -1)
	cards = strings.Replace(cards, "K", "W", -1)
	cards = strings.Replace(cards, "A", "X", -1)
	return cards
}
