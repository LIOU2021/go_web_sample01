package person

type Struct1 struct {
	name string
}

func (s *Struct1) SetName(n string) {
	s.name = n + " modified from method"
}

func (s Struct1) GetName() string {
	return s.name
}
