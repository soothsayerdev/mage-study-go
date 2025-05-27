package item

type DaggerMorgul struct {
	name string
	level int
}

func NewDaggerMorgul(level int) *DaggerMorgul {
	return &DaggerMorgul{
		level: level,
		name: "Dagger Morgul",
	}
}

func (d *DaggerMorgul) Name() string {
	return d.name
}

func (d *DaggerMorgul) Use() {

}

func (d *DaggerMorgul) Level() int {
	return d.level
}

