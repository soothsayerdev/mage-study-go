package dungeon

import "mage-study-go/model"

type LothricCastle struct {
	person person
	name   string
	Type   string
}

func NewLothricCastle(name string, person person) *LothricCastle {
	return &LothricCastle{
		name:   name,
		person: person,
		Type:   "Lothric Castle",
	}
}

func (l *LothricCastle) Name() string {
	return l.name
}

func (l *LothricCastle) PersonName() string {
	return l.person.Name()
}

func (l *LothricCastle) ListDungItems() []model.Item {
	return l.person.BagItens()
}