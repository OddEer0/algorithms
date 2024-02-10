package ozoncontest

import (
	"fmt"
	"strconv"
)

const (
	Spades  = 'S'
	Clubs   = 'C'
	Diamons = 'D'
	Hearts  = 'H'
)

type (
	Card struct {
		Type  rune
		Value int
	}
	Player struct {
		Num   int
		Cards []*Card
	}
	Deck   map[rune]map[int]bool
	Result struct {
		count int
		cards []string
	}
)

func getDeck() Deck {
	return Deck{
		Spades:  {2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true},
		Clubs:   {2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true},
		Diamons: {2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true},
		Hearts:  {2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true},
	}
}

func parseCardToStr(card *Card) string {
	result := ""

	switch card.Value {
	case 10:
		result += "T"
	case 11:
		result += "J"
	case 12:
		result += "Q"
	case 13:
		result += "K"
	case 14:
		result += "A"
	default:
		result += strconv.Itoa(card.Value)
	}

	return result + string(card.Type)
}

func parseStrToCard(s string) *Card {
	value := 0
	switch s[0] {
	case 'T':
		value = 10
	case 'J':
		value = 11
	case 'Q':
		value = 12
	case 'K':
		value = 13
	case 'A':
		value = 14
	default:
		value, _ = strconv.Atoi(string(s[0]))
	}
	return &Card{
		Type:  rune(s[1]),
		Value: value,
	}
}

func threePokerInput() ([]*Player, Deck) {
	var playerCount int
	fmt.Scan(&playerCount)
	players := make([]*Player, 0, playerCount)
	deck := getDeck()
	for i := 0; i < playerCount; i++ {
		player := &Player{
			Cards: make([]*Card, 0, 2),
			Num:   i + 1,
		}
		var first, second string
		fmt.Scan(&first, &second)
		firstCard, secondCard := parseStrToCard(first), parseStrToCard(second)
		player.Cards = append(player.Cards, firstCard, secondCard)
		deck[firstCard.Type][firstCard.Value] = false
		deck[secondCard.Type][secondCard.Value] = false

		players = append(players, player)
	}

	return players, deck
}

func getMaxCardValue(players []*Player) int {
	result := 2

	for _, player := range players {
		if player.Cards[0].Value > result {
			result = player.Cards[0].Value
		}
		if player.Cards[1].Value > result {
			result = player.Cards[1].Value
		}
	}

	return result
}

func filterElderCards(players []*Player, card *Card) []*Player {
	result := make([]*Player, 0, len(players))
	maxValue := getMaxCardValue(players)
	if card.Value > maxValue {
		maxValue = card.Value
	}

	for _, player := range players {
		if maxValue == player.Cards[0].Value || maxValue == player.Cards[1].Value || card.Value == maxValue {
			result = append(result, player)
		}
	}

	return result
}

func filterPairElderCards(players []*Player, card *Card) []*Player {
	result := make([]*Player, 0, len(players))
	max := 2
	for _, player := range players {
		if player.Cards[0].Value == card.Value || player.Cards[0].Value == player.Cards[1].Value {
			if player.Cards[0].Value > max {
				max = player.Cards[0].Value
			}
		} else if player.Cards[1].Value == card.Value {
			if player.Cards[1].Value > max {
				max = player.Cards[1].Value
			}
		}
	}

	for _, player := range players {
		if max == player.Cards[0].Value && max == player.Cards[1].Value || max == player.Cards[0].Value && max == card.Value || max == player.Cards[1].Value && max == card.Value {
			result = append(result, player)
		}
	}

	return result
}

func getWinners(players []*Player, card *Card) []*Player {
	result := make([]*Player, 0, len(players))

	for _, player := range players {
		if player.Cards[0].Value == player.Cards[1].Value && player.Cards[0].Value == card.Value {
			result = append(result, player)
		}
	}

	if len(result) == 0 {
		for _, player := range players {
			if player.Cards[0].Value == card.Value || player.Cards[1].Value == card.Value || player.Cards[0].Value == player.Cards[1].Value {
				result = append(result, player)
			}
		}
		result = filterPairElderCards(result, card)
	} else {
		return result
	}

	if len(result) == 0 {
		result = filterElderCards(players, card)
	}

	return result
}

func threePoker(players []*Player, deck Deck) (int, []string) {
	winCount := 0
	winCards := make([]string, 0, 15)

	for deckType, deckCards := range deck {
		for cardValue, has := range deckCards {
			if has {
				card := &Card{Type: deckType, Value: cardValue}
				winners := getWinners(players, card)
				for _, player := range winners {
					if player.Num == 1 {
						winCount++
						winCards = append(winCards, parseCardToStr(card))
					}
				}
			}
		}
	}

	return winCount, winCards
}

func Task8() {
	var count int
	fmt.Scan(&count)

	results := make([]*Result, count)

	for i := 0; i < count; i++ {
		players, deck := threePokerInput()
		winCount, winCards := threePoker(players, deck)
		results[i] = &Result{count: winCount, cards: winCards}
	}

	for i := 0; i < count; i++ {
		fmt.Println(results[i].count)
		for k := 0; k < results[i].count; k++ {
			fmt.Print(results[i].cards[k])
			if k != results[i].count-1 {
				fmt.Print("\n")
			}
		}
		if i != count-1 && results[i].count > 0 {
			fmt.Print("\n")
		}
	}
}
