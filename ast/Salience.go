package ast

func NewSalience(val int) *Salience {
	return &Salience{
		SalienceValue: val,
	}
}

type Salience struct {
	SalienceValue int
}

type SalienceReceiver interface {
	AcceptSalience(salience *Salience) error
}
