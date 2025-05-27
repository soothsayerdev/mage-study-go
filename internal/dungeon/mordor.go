package dungeon

import "mage-study/model"

type Mordor struct {
	person person // dependencia injetada via interface, desacoplada de implementacoes concretas
	name   string
	Type   string
}

func NewMordorDung(name string, person person) *Mordor {
	return &Mordor{
		name:   name,
		Type:   "Mordor",
		person: person,
	}
}

// Name retorna o nome da dungeon
func (m *Mordor) Name() string {
	return m.name
}

// PersonName retorna o nome do personagem que esta na dungeon
func (m *Mordor) PersonName() string {
	return m.person.Name()
}

func (m *Mordor) ListDungItems() []model.Item {
	return m.person.BagItens()
}
