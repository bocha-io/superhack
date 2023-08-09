package garnethelpers

func (p *Prediction) getPlayerFightingMon(player string, pos int64) string {
	if pos == int64(0) {
		return p.PlayerFirstMonGet(player)
	}
	if pos == int64(1) {
		return p.PlayerSecondMonGet(player)
	}
	if pos == int64(2) {
		return p.PlayerThirdMonGet(player)
	}
	if !(false) {
		panic("invalid pos")
	}
	return string(int64(0))
}

func (p *Prediction) attack(attackingMon string, attackedMon string, pos int64) int64 {
	atkDmg, _, atkElement := p.getAttack(p.MonSpecieGet(attackingMon), pos)
	_, attackedMonElement := p.getMonSpeedAndType(p.MonSpecieGet(attackedMon))
	dmg := p.getAttackDamage(atkDmg, atkElement, attackedMonElement)
	hp := p.MonHpGet(attackedMon)
	if dmg > hp {
		p.MonHpSet(attackedMon, int64(0))
		return int64(0)
	} else {
		p.MonHpSet(attackedMon, hp-dmg)
		return hp - dmg
	}
}

func (p *Prediction) bothAttacks(monOne string, monTwo string, posOne int64, posTwo int64) {
	playerOneSpeed := int64(0)
	speedMonPlayerOne, _ := p.getMonSpeedAndType(p.MonSpecieGet(monOne))
	atkSpeedMonOne, _, _ := p.getAttack(p.MonSpecieGet(monOne), posOne)
	playerOneSpeed = speedMonPlayerOne + atkSpeedMonOne
	playerTwoSpeed := int64(0)
	speedMonPlayerTwo, _ := p.getMonSpeedAndType(p.MonSpecieGet(monTwo))
	atkSpeedMonTwo, _, _ := p.getAttack(p.MonSpecieGet(monTwo), posTwo)
	playerTwoSpeed = speedMonPlayerTwo + atkSpeedMonTwo
	if playerOneSpeed >= playerTwoSpeed {
		if p.attack(monOne, monTwo, posOne) > int64(0) {
			p.attack(monTwo, monOne, posTwo)
		}
	} else {
		if p.attack(monTwo, monOne, posTwo) > int64(0) {
			p.attack(monOne, monTwo, posOne)
		}
	}
}

func (p *Prediction) Battle(
	matchID string,
	playerOneAction int64,
	posOne int64,
	playerTwoAction int64,
	posTwo int64,
) {
	if !(p.MatchGet(matchID)) {
		panic("match is not created")
	}
	p1Executed := false
	p1Mon := p.PlayerOneCurrentMonGet(matchID)
	p2Executed := false
	p2Mon := p.PlayerTwoCurrentMonGet(matchID)
	if playerOneAction == Swap {
		p1Executed = true
		p1Mon = p.getPlayerFightingMon(p.PlayerOneGet(matchID), posOne)
		p.PlayerOneCurrentMonSet(matchID, p1Mon)
	}
	if playerTwoAction == Swap {
		p2Executed = true
		p2Mon = p.getPlayerFightingMon(p.PlayerTwoGet(matchID), posTwo)
		p.PlayerTwoCurrentMonSet(matchID, p2Mon)
	}
	if !p1Executed {
		if !p2Executed {
			p.bothAttacks(p1Mon, p2Mon, posOne, posTwo)
		} else {
			p.attack(p1Mon, p2Mon, posOne)
		}
	} else {
		if p2Executed == false {
			p.attack(p2Mon, p1Mon, posTwo)
		}
	}
}
