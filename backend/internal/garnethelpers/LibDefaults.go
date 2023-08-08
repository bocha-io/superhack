package garnethelpers

func (p *Prediction) getHp(mon int64) int64 {
	if mon == Renzok {
		return int64(100)
	}
	if mon == Ramon {
		return int64(6)
	}
	if mon == Hanchon {
		return int64(50)
	}
	return int64(0)
}
