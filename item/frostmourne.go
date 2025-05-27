package item

type Frostmourne struct {
	name string
	level int
}

func NewFroustmourne(level int) *Frostmourne {
	return &Frostmourne{
		name: "Froustmourne",
		level: level,
	}
}

func (f *Frostmourne) Name() string {
	return f.name
}

func (f *Frostmourne) Level() int {
	return f.level
}

func (f *Frostmourne) Use() {

}