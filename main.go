package main

func main()  {
	// A few plays of deck functionalities. Find more in deck.go

	cards := newDeck()
	cards.shuffle()
	cards.print()

	hand, remainingCards := deal(cards, 4)

	hand.print()
	remainingCards.print()
}
