package ast

// NewSalience create new Salience AST object
func NewSalience(val int) *Salience {
	return &Salience{
		SalienceValue: val,
	}
}

// Salience is a simple AST object that stores salience
type Salience struct {
	SalienceValue int
}

// SalienceReceiver must be implemented by any AST object that stores salience
type SalienceReceiver interface {
	AcceptSalience(salience *Salience) error
}
