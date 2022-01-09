package unit

type Label struct {
	long string
	slug string
}

var (
	LabelGramm = Label{long: "gramm", slug: "g"}
	LabelKcal  = Label{long: "kcal", slug: "kcal"}
)

// Unit represents a unit of measurements
// which describes one property of a food
// it has been mapped to
type Unit interface {
	Label() string
	Slug() string
	Value() float64
}

type Gramm struct {
	label Label
	value float64
}

func NewGramm(value float64) *Gramm {
	return &Gramm{
		label: LabelGramm,
		value: value,
	}
}

// Label returns the full label of the Label
func (g Gramm) Label() string {
	return g.label.long
}

// Slug returns the short-hand (slug) of the Label
func (g Gramm) Slug() string {
	return g.label.slug
}

func (g Gramm) Value() float64 {
	return g.value
}

type Kcal struct {
	label Label
	value float64
}

func NewKcal(value float64) Kcal {
	return Kcal{
		label: LabelKcal,
		value: value,
	}
}

func (k Kcal) Label() string {
	return k.label.long
}

func (k Kcal) Slug() string {
	return k.label.slug
}

func (k Kcal) Value() float64 {
	return k.value
}
