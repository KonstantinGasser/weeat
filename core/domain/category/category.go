package category

type Type int

const (
	Food Type = iota + 1
	Recipe
)

type Category struct {
	ID    int
	Type  Type
	Label string
	Emoji string
}
