package strategy

//- Strategies knows the possibles context.
type Strategies map[string]Invoke

//- Strategy know instantiate a map of strategies.
type Strategy struct {
	strategies Strategies
}

func NewStrategy() *Strategy {

	strategies := make(Strategies)

	NewA := func() *All {
		return &All{code: 1}
	}

	NewB := func() *All {
		return &All{code: 5}
	}

	strategies["A"] = NewA()
	strategies["B"] = NewB()

	return &Strategy{
		strategies,
	}
}

func (s *Strategy) GetAll() Strategies {
	return s.strategies
}
