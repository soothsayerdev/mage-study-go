package dungeon

import "mage-study-go/model"

type person interface {
	Name() string
	BagItens() []model.Item
}