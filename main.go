package main
import (
	"fmt"
	"math/rand"
	"time"
)
//Cardクラス
type Card struct{
	suit string
	value string
	intValue int
}
//Card/setter
func NewCard(suit string, value string, intValue int) *Card {
	return &Card{suit, value, intValue}
}
//Card/getter
func (c Card) String() string{
	return c.suit + c.value ;
}

type Deck struct{
	deck []Card
}
//Deck/setter
func NewDeck() *Deck{
	d := &Deck{}
	d.deck = d.createDeck()
	d.shuffleDeck()

	return d
}

func (d Deck) createDeck() []Card{
	newDeck := []Card{}
	suits := []string{"♣", "♦", "♥", "♠"}
	values := []string{"A","2","3","4","5","6","7","8","9","10","J","Q","K"}

	for i := 0; i < len(suits); i++{
		for j := 0; j < len(values); j++{
			newDeck = append(newDeck, *NewCard(suits[i], values[j], j + 1))
		}
	}
	return newDeck
}

func (d *Deck) shuffleDeck(){
	rand.Seed(time.Now().UnixNano())

	for i := len(d.deck) - 1; i >= 0; i--{
		j := rand.Intn(i + 1)
		d.deck[i], d.deck[j] = d.deck[j], d.deck[i]
	}
}

func (d *Deck) drawCard() Card{
	card := d.deck[len(d.deck) - 1]
	d.deck = d.deck[:len(d.deck) - 1]
	return card
}

type Table struct{
	gameMode string
	amountOfPlayers int
}

func createTable(gameMode string, amountOfPlayers int) *Table{
	return &Table{gameMode: gameMode, amountOfPlayers: amountOfPlayers}
}

func startGame(table *Table) [][]Card{
	deck := NewDeck()

	newTable := [][]Card{}

	for i := 0; i < table.amountOfPlayers; i++{
		playerHand := []Card{}

		for j := 0; j < numOfCards(table); j++{
			card := deck.drawCard()
			playerHand = append(playerHand, card)
		}
		newTable = append(newTable, playerHand);
	}
	return newTable
}

func numOfCards(table *Table) int{
	if table.gameMode == "Pair of Cards"{
		return 5
	}else{
		return 2
	}
}

func printTable(playerCards [][]Card, table *Table){
	fmt.Println("Amount of players:", table.amountOfPlayers, "Game mode:", table.gameMode)
	for i := 0; i < table.amountOfPlayers; i++{
		fmt.Println("player",(i+1), "hand is")
		for _, v := range playerCards[i] {
			fmt.Println(v)
		}
	}
}

func main(){
	table := createTable("Pair of Cards",5)
	game := startGame(table)
	printTable(game, table)
}