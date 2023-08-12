package garnethelpers

func (p *Prediction) SwapMon(mon int64, pos int64, senderAddress string) {
	senderKey := p.addressToEntityKey(senderAddress)
	if !(p.PlayerGet(senderKey)) {
		panic("user is not registered")
	}
	if !(pos == int64(0) || pos == int64(1) || pos == int64(2)) {
		panic("invalid pos")
	}
	if pos == int64(0) {
		p.InventoryFirstMonSet(senderKey, mon)
	} else {
		if pos == int64(1) {
			p.InventorySecondMonSet(senderKey, mon)
		} else {
			if pos == int64(2) {
				p.InventoryThirdMonSet(senderKey, mon)
			}
		}
	}
}
