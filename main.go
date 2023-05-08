package main
import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
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
	for i := 0; i <  table.amountOfPlayers; i++{
		fmt.Println("player",(i+1), "hand is")
		for _, v := range playerCards[i] {
			fmt.Println(v)
		}
	}
}

func winnerOfPairOfCards(playerCards [][]Card, table *Table) string{
	var orderOfStrength =  []int{1,13,12,11,10,9,8,7,6,5,4,3,2}
	playerHands := [][]int{}
	for i := 0; i < table.amountOfPlayers; i++{
		numArr := convertToNumbers(playerCards[i])
		playerHands = append(playerHands,numArr)
	}
	mapSlice := []map[int]int{}
	for _,v := range playerHands{
		hashmap := createHashmap(v,orderOfStrength)
		mapSlice = append(mapSlice, hashmap)
	}
	winPlayer := 1
	winner := "draw"
	flag := false
	for i := 1; i < len(playerHands); i++{
		pairOfCards := 0
		for j := 0; j < len(orderOfStrength); j++{
			if mapSlice[winPlayer - 1][orderOfStrength[j]] > mapSlice[i][orderOfStrength[j]]{
				if pairOfCards < mapSlice[winPlayer - 1][orderOfStrength[j]]{
					pairOfCards = mapSlice[winPlayer - 1][orderOfStrength[j]]
					winner = "player"+ strconv.Itoa(winPlayer)
					flag = false;
				}
			}else if mapSlice[winPlayer - 1][orderOfStrength[j]] < mapSlice[i][orderOfStrength[j]]{
				if pairOfCards < mapSlice[i][orderOfStrength[j]]{
					pairOfCards = mapSlice[i][orderOfStrength[j]]
					winner = "player"+ strconv.Itoa(i+1)
					flag = true;
				}
			}
		}
		if flag{
			winPlayer = i+1
		}  
	}
	fmt.Println("Winner of the game is ")
	fmt.Println(winPlayer)
	return winner
}

func convertToNumbers(playerHands []Card) []int{
	arr := []int{}
	for i,_ := range playerHands{
		arr = append(arr, playerHands[i].intValue)
	}
	return arr
}

func createHashmap(numArr []int,orderOfStrength []int) map[int]int{
	hashmap := map[int]int{}
	for _,v := range orderOfStrength{
		hashmap[v] = 0
	}
	for _,v := range numArr{
		hashmap[v] = (hashmap[v] + 1) 
	}
	return hashmap

}

func main(){
	table := createTable("Pair of Cards",5)
	game := startGame(table)
	printTable(game, table)
	fmt.Println(winnerOfPairOfCards(game, table))
}