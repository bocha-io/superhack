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

func (p *Prediction) fight(
	matchID string,
	monOne string,
	monTwo string,
	posOne int64,
	posTwo int64,
) bool {
	playerOneSpeed := int64(0)
	speedMonPlayerOne, _ := p.getMonSpeedAndType(p.MonSpecieGet(monOne))
	_, atkSpeedMonOne, _ := p.getAttack(p.MonSpecieGet(monOne), posOne)
	playerOneSpeed = speedMonPlayerOne + atkSpeedMonOne
	playerTwoSpeed := int64(0)
	speedMonPlayerTwo, _ := p.getMonSpeedAndType(p.MonSpecieGet(monTwo))
	_, atkSpeedMonTwo, _ := p.getAttack(p.MonSpecieGet(monTwo), posTwo)
	playerTwoSpeed = speedMonPlayerTwo + atkSpeedMonTwo
	if playerOneSpeed >= playerTwoSpeed {
		if p.attack(monOne, monTwo, posOne) == int64(0) {
			if p.checkIfGameHasEnded(p.PlayerTwoGet(matchID), monTwo) {
				p.endGame(matchID, p.PlayerOneGet(matchID), p.PlayerTwoGet(matchID))
				return true
			}
		} else {
			if p.attack(monTwo, monOne, posTwo) == int64(0) {
				if p.checkIfGameHasEnded(p.PlayerOneGet(matchID), monOne) {
					p.endGame(matchID, p.PlayerTwoGet(matchID), p.PlayerOneGet(matchID))
					return true
				}
			}
		}
	} else {
		if p.attack(monTwo, monOne, posTwo) == int64(0) {
			if p.checkIfGameHasEnded(p.PlayerOneGet(matchID), monOne) {
				p.endGame(matchID, p.PlayerTwoGet(matchID), p.PlayerOneGet(matchID))
				return true
			}
		} else {
			if p.attack(monOne, monTwo, posOne) == int64(0) {
				if p.checkIfGameHasEnded(p.PlayerTwoGet(matchID), monTwo) {
					p.endGame(matchID, p.PlayerOneGet(matchID), p.PlayerTwoGet(matchID))
					return true
				}
			}
		}
	}
	return false
}

func (p *Prediction) checkIfGameHasEnded(playerID string, monDead string) bool {
	res := int64(0)
	firstMon := p.PlayerFirstMonGet(playerID)
	if firstMon != monDead {
		res = res + p.MonHpGet(firstMon)
	}
	secondMon := p.PlayerSecondMonGet(playerID)
	if secondMon != monDead {
		res = res + p.MonHpGet(secondMon)
	}
	thirdMon := p.PlayerThirdMonGet(playerID)
	if thirdMon != monDead {
		res = res + p.MonHpGet(thirdMon)
	}
	return res == int64(0)
}

func (p *Prediction) endGame(matchID string, winner string, loser string) {
	p.MatchDeleterecord(matchID)
	p.PlayerOneDeleterecord(matchID)
	p.PlayerTwoDeleterecord(matchID)
	p.MatchResultSet(matchID, winner, loser)
	p.StatusSet(winner, Walking)
	p.StatusSet(loser, Walking)
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
	if playerOneAction == Surrender {
		p.endGame(matchID, p.PlayerTwoGet(matchID), p.PlayerOneGet(matchID))
		return
	}
	if playerTwoAction == Surrender {
		p.endGame(matchID, p.PlayerOneGet(matchID), p.PlayerTwoGet(matchID))
		return
	}
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
	if p.MonHpGet(p1Mon) == int64(0) {
		p.endGame(matchID, p.PlayerTwoGet(matchID), p.PlayerOneGet(matchID))
		return
	} else {
		if p.MonHpGet(p2Mon) == int64(0) {
			p.endGame(matchID, p.PlayerOneGet(matchID), p.PlayerTwoGet(matchID))
			return
		}
	}
	if !p1Executed {
		if !p2Executed {
			if p.fight(matchID, p1Mon, p2Mon, posOne, posTwo) {
				return
			}
		} else {
			if p.attack(p1Mon, p2Mon, posOne) == int64(0) {
				if p.checkIfGameHasEnded(p.PlayerTwoGet(matchID), p2Mon) {
					p.endGame(matchID, p.PlayerOneGet(matchID), p.PlayerTwoGet(matchID))
					return
				}
			}
		}
	} else {
		if p2Executed == false {
			if p.attack(p2Mon, p1Mon, posTwo) == int64(0) {
				if p.checkIfGameHasEnded(p.PlayerOneGet(matchID), p1Mon) {
					p.endGame(matchID, p.PlayerTwoGet(matchID), p.PlayerOneGet(matchID))
					return
				}
			}
		}
	}
}
