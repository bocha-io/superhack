package garnethelpers

func (p *Prediction) getMonFromInventory(player string, pos int64) int64 {
	if pos == int64(0) {
		return p.InventoryFirstMonGet(player)
	}
	if pos == int64(1) {
		return p.InventorySecondMonGet(player)
	}
	if pos == int64(2) {
		return p.InventoryThirdMonGet(player)
	}
	if !(false) {
		panic("invalid pos")
	}
	return Ramon
}

func (p *Prediction) setPlayerCurrentMon(gameKey string, player int64, monID string) {
	if player == int64(0) {
		p.PlayerOneCurrentMonSet(gameKey, monID)
		return
	}
	if player == int64(1) {
		p.PlayerTwoCurrentMonSet(gameKey, monID)
		return
	}
	if !(false) {
		panic("invalid player")
	}
}

func (p *Prediction) setGameMon(player string, pos int64, monID string) {
	if pos == int64(0) {
		p.PlayerFirstMonSet(player, monID)
		return
	}
	if pos == int64(1) {
		p.PlayerSecondMonSet(player, monID)
		return
	}
	if pos == int64(2) {
		p.PlayerThirdMonSet(player, monID)
		return
	}
	if !(false) {
		panic("invalid pos")
	}
}

func (p *Prediction) CreateMatch(playerA string, playerB string) {
	if !(p.PlayerGet(playerA) == true) {
		panic("player a is not registered")
	}
	if !(p.PlayerGet(playerB) == true) {
		panic("player b is not registered")
	}
	if !(p.StatusGet(playerA) == Walking) {
		panic("player a is already in a match")
	}
	if !(p.StatusGet(playerB) == Walking) {
		panic("player b is already in a match")
	}
	gameKey := playerA
	p.MatchSet(gameKey, true)
	p.PlayerOneSet(gameKey, playerA)
	p.PlayerTwoSet(gameKey, playerB)
	p.StatusSet(playerA, Fighting)
	p.StatusSet(playerB, Fighting)
	for player := int64(0); player < int64(2); player++ {
		playerKey := playerA
		if player == int64(1) {
			playerKey = playerB
		}
		for i := int64(0); i < int64(3); i++ {
			key := p.monKey(playerKey, i)
			p.MonSet(key, true)
			p.MonSpecieSet(key, p.getMonFromInventory(playerKey, i))
			p.MonHpSet(key, p.getHp(p.getMonFromInventory(playerKey, i)))
			p.setGameMon(playerKey, i, key)
			if i == int64(0) {
				p.setPlayerCurrentMon(gameKey, player, key)
			}
		}
	}
}
