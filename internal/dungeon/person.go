package dungeon

import "mage-study/model"

type person interface {
	Name() string
	BagItens() []model.Item
}