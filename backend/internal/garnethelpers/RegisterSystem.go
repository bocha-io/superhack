package garnethelpers

func (p *Prediction) register(senderAddress string) {
	senderKey := p.addressToEntityKey(senderAddress)
	if !(p.PlayerGet(senderKey) == false) {
		panic("wallet already registered")
	}
	p.PlayerSet(senderKey, true)
	p.InventoryFirstMonSet(senderKey, Hanchon)
	p.InventorySecondMonSet(senderKey, Renzok)
	p.InventoryThirdMonSet(senderKey, Ramon)
	p.PositionSet(senderKey, int64(0), int64(0))
	p.StatusSet(senderKey, Walking)
}