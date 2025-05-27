package person

import "mage-study/model"

type Nazgul struct {
	name  string
	items []item // bag de items
}

func NewNazgul(name string) *Nazgul {
	return &Nazgul{
		name: name,
	}
}

func (n *Nazgul) Name() string {
	return n.name
}

func (n *Nazgul) PickUpItem(item item) {
	n.items = append(n.items, item)
}

func (n *Nazgul) BagItens() []model.Item {
	bag := make([]model.Item, 15)
	for i, item := range n.items {
		bag[i] = model.Item{
			Name: item.Name(),
			Level: item.Level(),
		}
	}
	return bag
}
