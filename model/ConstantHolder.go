package model

// ConstantHolder define all graphs that should be able to hold a constants.
type ConstantHolder interface {
	AcceptConstant(cons *Constant) error
}
