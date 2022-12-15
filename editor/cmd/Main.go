package main

import (
	"github.com/hyperjumptech/grule-rule-engine/editor"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	editor.Start()
}
