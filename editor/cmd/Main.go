package main

import (
	"github.com/sirupsen/logrus"
	"grule-rule-engine/editor"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	editor.Start()
}
