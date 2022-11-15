package main

import (
	"fmt"
	"math/rand"
)

//generate deck of cards
// dimond , heart, spade, clubs
//shuffle - make cards random
//Printl shuffled deck
//distribute
//n players
//m cards
//print the n players cards
// 13 * 4

//type - d, h, s, c
//Category - two to ten, A, J,Q,K

//shuffle, distribute,
//deal - give cards to players

// create cards deck
// shuffle funtion - randomize there position in deck
//

//One of heart, spade, club, dimond
type cardType struct {
	name string
}

type card struct {
	cardName string
	cType    cardType
}

type deck []card

const CARDCATEGORYS = 13
const CARDTYPES = 4

type pCards []card

var cardTypes = [4]string{"HEART", "SPADE", "CLUB", "DIAMOND"}
var cardCategories = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

func (d *deck) createDeck() {

	for cardTypeSeq := 0; cardTypeSeq < CARDTYPES; cardTypeSeq++ {
		name := cardTypes[cardTypeSeq]
		for cardCategoryIdx := 0; cardCategoryIdx < CARDCATEGORYS; cardCategoryIdx++ {
			c := card{
				cardName: cardCategories[cardCategoryIdx],
				cType: cardType{
					name: name,
				},
			}
			*d = append(*d, c)
		}
	}
}

func (d deck) printDeck() {
	for crd := range d {
		fmt.Println(d[crd])
	}
}

func (d *deck) shuffleDeck() {
	shuffleCnt := 200
	t := *d
	for i := 0; i < shuffleCnt; i++ {
		pos := rand.Intn(52)
		t[pos], t[0] = t[0], t[pos]
	}
	*d = t
}

func (d deck) distributeDeck() interface{} {

	var players []pCards

	cCnt := 0
	for i := 0; i < 4; i++ {

		var pcards pCards
		for c := 0; c < 3; c++ {
			pcards = append(pcards, d[cCnt])
			cCnt++
		}
		players = append(players, pcards)
	}
	return players
}

func deckDriver() {
	//	players := 4	//	cardsCnt := 3

	var d deck
	d.createDeck()
	d.printDeck()

	d.shuffleDeck()
	d.printDeck()

	// distributeDeck()	//distribute cards to players
	players := d.distributeDeck()

	// printPlayerCards()
	for i := 0; i < 4; i++ {
		fmt.Println(players.([]pCards)[i])
	}
}
