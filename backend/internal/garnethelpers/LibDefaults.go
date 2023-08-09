package garnethelpers

func (p *Prediction) getHp(mon int64) int64 {
	if mon == Cobarett {
		return int64(350)
	}
	if mon == Flarezael {
		return int64(450)
	}
	if mon == Firomenis {
		return int64(300)
	}
	if mon == Baobaffe {
		return int64(500)
	}
	if mon == Howliage {
		return int64(350)
	}
	if mon == Sunnydra {
		return int64(420)
	}
	if mon == Tobishimi {
		return int64(400)
	}
	if mon == Mobiusk {
		return int64(300)
	}
	if mon == Ramon {
		return int64(450)
	}
	return int64(0)
}

func (p *Prediction) getMonSpeedAndType(mon int64) (int64, int64) {
	if mon == Cobarett {
		return int64(100), Fire
	}
	if mon == Flarezael {
		return int64(60), Fire
	}
	if mon == Firomenis {
		return int64(120), Fire
	}
	if mon == Baobaffe {
		return int64(50), Grass
	}
	if mon == Howliage {
		return int64(100), Grass
	}
	if mon == Sunnydra {
		return int64(60), Grass
	}
	if mon == Tobishimi {
		return int64(80), Water
	}
	if mon == Mobiusk {
		return int64(120), Water
	}
	if mon == Ramon {
		return int64(70), Water
	}
	return int64(0), Normal
}

func (p *Prediction) getAttack(mon int64, pos int64) (int64, int64, int64) {
	if mon == Cobarett {
		if pos == int64(0) {
			return p.getAttackValues(LavaPlume)
		}
		if pos == int64(1) {
			return p.getAttackValues(FireLash)
		}
		if pos == int64(2) {
			return p.getAttackValues(Bite)
		}
		if pos == int64(3) {
			return p.getAttackValues(Tackle)
		}
	}
	if mon == Flarezael {
		if pos == int64(0) {
			return p.getAttackValues(FieryDance)
		}
		if pos == int64(1) {
			return p.getAttackValues(FireLash)
		}
		if pos == int64(2) {
			return p.getAttackValues(CrushGrip)
		}
		if pos == int64(3) {
			return p.getAttackValues(Flail)
		}
	}
	if mon == Firomenis {
		if pos == int64(0) {
			return p.getAttackValues(FieryDance)
		}
		if pos == int64(1) {
			return p.getAttackValues(LavaPlume)
		}
		if pos == int64(2) {
			return p.getAttackValues(Tackle)
		}
		if pos == int64(3) {
			return p.getAttackValues(Flail)
		}
	}
	if mon == Baobaffe {
		if pos == int64(0) {
			return p.getAttackValues(LeafTornado)
		}
		if pos == int64(1) {
			return p.getAttackValues(RazorLeaf)
		}
		if pos == int64(2) {
			return p.getAttackValues(Tackle)
		}
		if pos == int64(3) {
			return p.getAttackValues(Flail)
		}
	}
	if mon == Howliage {
		if pos == int64(0) {
			return p.getAttackValues(SolarBlade)
		}
		if pos == int64(1) {
			return p.getAttackValues(RazorLeaf)
		}
		if pos == int64(2) {
			return p.getAttackValues(Tackle)
		}
		if pos == int64(3) {
			return p.getAttackValues(Bite)
		}
	}
	if mon == Sunnydra {
		if pos == int64(0) {
			return p.getAttackValues(SolarBlade)
		}
		if pos == int64(1) {
			return p.getAttackValues(LeafTornado)
		}
		if pos == int64(2) {
			return p.getAttackValues(Flail)
		}
		if pos == int64(3) {
			return p.getAttackValues(Bite)
		}
	}
	if mon == Tobishimi {
		if pos == int64(0) {
			return p.getAttackValues(AquaTail)
		}
		if pos == int64(1) {
			return p.getAttackValues(BubbleBeam)
		}
		if pos == int64(2) {
			return p.getAttackValues(Flail)
		}
		if pos == int64(3) {
			return p.getAttackValues(Bite)
		}
	}
	if mon == Mobiusk {
		if pos == int64(0) {
			return p.getAttackValues(AquaTail)
		}
		if pos == int64(1) {
			return p.getAttackValues(HydroVortex)
		}
		if pos == int64(2) {
			return p.getAttackValues(Tackle)
		}
		if pos == int64(3) {
			return p.getAttackValues(Bite)
		}
	}
	if mon == Ramon {
		if pos == int64(0) {
			return p.getAttackValues(BubbleBeam)
		}
		if pos == int64(1) {
			return p.getAttackValues(HydroVortex)
		}
		if pos == int64(2) {
			return p.getAttackValues(Tackle)
		}
		if pos == int64(3) {
			return p.getAttackValues(Flail)
		}
	}
	return int64(0), int64(0), Normal
}

func (p *Prediction) getAttackValues(attack int64) (int64, int64, int64) {
	if attack == LavaPlume {
		return int64(80), int64(30), Fire
	}
	if attack == FireLash {
		return int64(50), int64(50), Fire
	}
	if attack == FieryDance {
		return int64(110), int64(20), Fire
	}
	if attack == LeafTornado {
		return int64(120), int64(30), Grass
	}
	if attack == RazorLeaf {
		return int64(50), int64(50), Grass
	}
	if attack == SolarBlade {
		return int64(80), int64(40), Grass
	}
	if attack == AquaTail {
		return int64(60), int64(50), Water
	}
	if attack == BubbleBeam {
		return int64(120), int64(30), Water
	}
	if attack == HydroVortex {
		return int64(90), int64(60), Water
	}
	if attack == Bite {
		return int64(30), int64(80), Normal
	}
	if attack == Tackle {
		return int64(40), int64(60), Normal
	}
	if attack == CrushGrip {
		return int64(80), int64(30), Normal
	}
	if attack == Flail {
		return int64(20), int64(100), Normal
	}
	return int64(0), int64(0), Normal
}

func (p *Prediction) getAttackDamage(dmg int64, attack int64, mon int64) int64 {
	if attack == Fire {
		if mon == Grass {
			return dmg * int64(2)
		}
		if mon == Water {
			return dmg / int64(2)
		}
	} else {
		if attack == Grass {
			if mon == Water {
				return dmg * int64(2)
			}
			if mon == Fire {
				return dmg / int64(2)
			}
		} else {
			if attack == Water {
				if mon == Fire {
					return dmg * int64(2)
				}
				if mon == Grass {
					return dmg / int64(2)
				}
			}
		}
	}
	return dmg
}
