package model

// DecimalHolder defines all graph that should be able to store a decimal value.
type DecimalHolder interface {
	AcceptDecimal(val int64) error
}
