package unit

type UnitLabel struct {
	long string
	slug string
}

var (
	LabelGramm = UnitLabel{long: "gramm", slug: "g"}
	LabelKcal  = UnitLabel{long: "kcal", slug: "kcal"}
)

// Unit represents a unit of measurements
// which describes one property of a food
// it has been mapped to
type Unit interface {
	Label() string
	Slug() string
}

type Gramm struct {
	label UnitLabel
	value float64
}

func NewGramm(value float64) *Gramm {
	return &Gramm{
		label: LabelGramm,
		value: value,
	}
}

// Label returns the full label of the UnitLabel
func (g Gramm) Label() string {
	return g.label.long
}

// Slug returns the short-hand (slug) of the UnitLabel
func (g Gramm) Slug() string {
	return g.label.slug
}

type Kcal struct {
	label UnitLabel
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
