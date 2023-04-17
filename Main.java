package pairOfCards;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
//カードクラス
class Card{
  public String suit;
  public String value;
  public int intValue;
  public Card(String suit, String value, int intValue){
      this.suit = suit;
      this.value = value;
      this.intValue = intValue;
  }
  //カード表示
  public String getCardString(){
      return this.suit + this.value + "(" + this.intValue + ")";
  }
}

class Deck{
  public ArrayList<Card> deck;

  public Deck(){
      this.deck = this.createDeck();
      this.shuffleDeck();
  }
  //デッキ作成
  public  ArrayList<Card> createDeck(){
      ArrayList<Card> newdeck = new ArrayList<>();
      String[] suit = new String[]{"♣", "♦", "♥", "♠"};
      String[] value = new String[]{"A","2","3","4","5","6","7","8","9","10","J","Q","K"};
      for(int i = 0; i < suit.length; i++){
          for(int j = 0; j < value.length; j++ ){
              newdeck.add(new Card(suit[i], value[j], j+1));
          }
      }

      return newdeck;
  }
  //カードをシャッフル
  public void shuffleDeck(){
      for(int i = this.deck.size()-1; i >= 0; i--){
          int j = (int)(Math.floor(Math.random() * (i + 1)));
          Card tmp = this.deck.get(i);
          this.deck.set(i,this.deck.get(j));
          this.deck.set(j,tmp);
      }
  }
  //カードを1枚引く
  public Card drawCard(){
      return this.deck.remove(this.deck.size()-1);
  }
}

class Table{
  public String gameMode;
  public int amountOfPlayers;

  public Table(String gameMode, int amountOfPlayers){
      this.gameMode = gameMode;
      this.amountOfPlayers = amountOfPlayers;
  }
}

class Dealer{

  public static ArrayList<ArrayList<Card>> startGame(Table table){
      Deck deck = new Deck();

      ArrayList<ArrayList<Card>> newTable = new ArrayList<>();

      for(int i = 0; i < table.amountOfPlayers; i++){
          ArrayList<Card> playerHand = new ArrayList<>();
          
          for(int j = 0; j < Dealer.numOfCards(table); j++){
              Card card = deck.drawCard();
              playerHand.add(card);
          }
          newTable.add(playerHand);
      }
      return newTable;
  }

  public static int numOfCards(Table table){
      if(table.gameMode.equals("Pair of Cards"))return 5;
      return 0;
  }
  public static void printTable(ArrayList<ArrayList<Card>> playerCards,Table table){
      System.out.println("Amount of players: " + table.amountOfPlayers + " Game mode: " + table.gameMode);
      for(int i = 0; i < table.amountOfPlayers; i++){
          System.out.println("player"+ (i+1) + " hand is ");
          for(int j = 0; j < playerCards.get(i).size(); j++){
              System.out.println(playerCards.get(i).get(j).getCardString());
          }
      }
  }
  public static String winnerOfPairOfCards(ArrayList<ArrayList<Card>> playerCards,Table table){
	  final int[] orderOfStrength = new int[] {1,13,12,11,10,9,8,7,6,5,4,3,2};
	  ArrayList<int[]> playerHands = new ArrayList<int[]>();
	  for(int i = 0; i < table.amountOfPlayers;i++) {
		  int[] numArr = Helper.convertToNumbers(playerCards.get(i));
		  playerHands.add(numArr);
	  }
	  ArrayList<Map<Integer,Integer>> mapArray = new ArrayList<Map<Integer,Integer>>();
	  for(int i = 0; i < playerHands.size();i++) {
		  Map<Integer,Integer> hashmap = Helper.createHashmap(playerHands.get(i),orderOfStrength);
		  mapArray.add(hashmap);
	  }
	  int winPlayer = 1;
	  String winner = "draw";
	  boolean flag = false;
	  //勝者判定
	  //初期はplayer1を固定、勝利したplayerをwinPlayerに代入
	  for(int i = 1; i < playerHands.size();i++) {
		  int pairOfCards = 0;
		  for(int j = 0; j < orderOfStrength.length;j++) {
			  if(mapArray.get(winPlayer -1).get(orderOfStrength[j]) > mapArray.get(i).get(orderOfStrength[j])) {
				  if(pairOfCards < mapArray.get(winPlayer -1).get(orderOfStrength[j])) {
					  pairOfCards = mapArray.get(winPlayer -1).get(orderOfStrength[j]);
					  winner = "player" + (winPlayer);
					  flag = false;
				  }
			  }else if(mapArray.get(winPlayer -1).get(orderOfStrength[j]) < mapArray.get(i).get(orderOfStrength[j])) {
				  if(pairOfCards < mapArray.get(i).get(orderOfStrength[j])) {
					  pairOfCards = mapArray.get(i).get(orderOfStrength[j]);
					  winner = "player" + (i+1);
					  flag = true;
				  }
			  }
		  }
		  if(flag) winPlayer = i+1;
	  }
	  System.out.println("Winner of the game is ");
	  System.out.println(winPlayer);
	  return winner;
  }
}

class Helper{
	//数列に変換
	public static int[] convertToNumbers(ArrayList<Card> playerHands) {
		int[] arr = new int[playerHands.size()];
		for(int i = 0; i < playerHands.size(); i++) {
			arr[i] = playerHands.get(i).intValue;			
		}
		return arr;
	}
	//Hashmap作製
	public static Map<Integer,Integer> createHashmap(int[] numArr, int[] orderOfStrength){
		Map<Integer,Integer> hashmap = new HashMap<Integer,Integer>();
		for(int i = 0; i < orderOfStrength.length;i++) {
			hashmap.put(orderOfStrength[i], 0);
		}
		for(int i = 0; i < numArr.length;i++) {
			hashmap.replace(numArr[i], hashmap.get(numArr[i]) + 1);
		}
		return hashmap;
	}
}

class Main{

  public static void main(String[] args){
      Table table = new Table("Pair of Cards",5);
      ArrayList<ArrayList<Card>> game = Dealer.startGame(table);
      Dealer.printTable(game, table);
      System.out.println(Dealer.winnerOfPairOfCards(game, table));


  }
}