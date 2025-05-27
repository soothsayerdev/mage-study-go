package person

type item interface {
	Use()
	Name() string
	Level() int
}