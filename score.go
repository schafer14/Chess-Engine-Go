package maurice

type Score struct {
	opening int
	midGame int
	endGame int
}

func (s *Score) add(score Score) *Score {
	s.opening += score.opening
	s.midGame += score.midGame
	s.endGame += score.endGame

	return s
}

func (s *Score) sub(score Score) *Score {
	s.opening += score.opening
	s.midGame += score.midGame
	s.endGame += score.endGame

	return s
}
