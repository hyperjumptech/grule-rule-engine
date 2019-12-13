package examples

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestCashFlowCalculator_CalculatePurchases(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	calc := &CashFlowCalculator{}
	calc.CalculatePurchases()
}
