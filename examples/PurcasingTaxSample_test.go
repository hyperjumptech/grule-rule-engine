package examples

import (
	"testing"
)

func TestCashFlowCalculator_CalculatePurchases(t *testing.T) {
	// logrus.SetLevel(logrus.DebugLevel)
	calc := &CashFlowCalculator{}
	calc.CalculatePurchases()
}
