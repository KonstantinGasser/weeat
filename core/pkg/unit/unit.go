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
	Scale(scaler int) Unit
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

func (g Gramm) Scale(scaler int) Unit {
	if g.value <= 0 {
		return &Gramm{label: g.label}
	}
	return &Gramm{
		label: g.label,
		value: (g.value / 100.0) * float64(scaler),
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

func (k Kcal) Scale(scaler int) Unit {
	if k.value <= 0 {
		return &Kcal{label: k.label}
	}

	return &Kcal{
		label: k.label,
		value: (k.value / 100.0) * float64(scaler),
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
