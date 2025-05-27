package person

import "mage-study-go/model"

type Arthas struct {
	name  string
	items []item
}

func NewArthas(name string) *Arthas {
	return &Arthas{
		name: name,
	}
}

func (a *Arthas) Name() string {
	return a.name
}

func (a *Arthas) PickUpItem(item item) {
	a.items = append(a.items, item)
}

func (a *Arthas) BagItens() []model.Item {
	bag := make([]model.Item, 15)
	for i, item := range a.items {
		bag[i] = model.Item{
			Name: item.Name(),
			Level: item.Level(),
		}
	}
	return bag
}