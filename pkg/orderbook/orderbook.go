package orderbook

type Orderbook struct {
	Asks *Tree
	Bids *Tree
}

func (b *Orderbook) Add(orde *Order) {}

func (b *Orderbook) Match(order *Order) {
	offers := b.Asks
	if order.Side == SideSell {
		offers = b.Bids
	}

	iter := offers.Iterator()

	for iter.Next() {
		price := iter.Key
		priceLevel := iter.Value
	}
}
