package item

type OneRing struct {
	name string
	level int
}

func NewOneRing(level int) *OneRing {
	return &OneRing{
		level: level,
		name: "One Ring",
	}
}

func (o *OneRing) Name() string {
	return o.name
}

func (o *OneRing) Use() {

}

func (o *OneRing) Level() int {
	return o.level
}