package garnethelpers

func (p *Prediction) Move(newX int64, newY int64, senderAddress string) {
	senderKey := p.addressToEntityKey(senderAddress)
	if !(p.PlayerGet(senderKey) == true) {
		panic("wallet is not registered")
	}
	if !(newX >= int64(0) && newY >= int64(0)) {
		panic("invalid X and Y")
	}
	p.PositionSet(senderKey, newX, newY)
}
