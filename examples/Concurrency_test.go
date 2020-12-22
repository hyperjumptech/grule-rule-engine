//  Copyright hyperjumptech/grule-rule-engine Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package examples

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/logger"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

const (
	VibonaciRule = `
rule PrepareVibonaci "Preparing Vibonaci" salience 10 {
	when 
		Vibo.A == 0 || Vibo.B == 0
	then
		Vibo.A = 1;
		Vibo.B = 1;
}

rule ExecuteVibonaci "Vibonaci execution" salience 10 {
	when 
		Vibo.Count < 80 && Vibo.A > 0 && Vibo.B > 0
	then
		Vibo.Count = Vibo.Count + 1;
		Vibo.C = Vibo.A + Vibo.B;
		Vibo.A = Vibo.B;
		Vibo.B = Vibo.C;
}
`
)

var (
	// threadFinishMap used to monitor the status of all thread
	threadFinishMap = make(map[string]bool)

	// syncD is a mutex object to protect threadFinishMap from concurrent map read/write
	syncD = sync.Mutex{}

	concurrencyTestlog = logger.Log.WithFields(logrus.Fields{
		"lib":  "grule",
		"file": "examples/Concurrency_test.go",
	})
)

// Vibonaci our model
type Vibonaci struct {
	fmt.Stringer
	Count int
	A     uint
	B     uint
	C     uint
}

func startThread(threadName string) {
	syncD.Lock()
	defer syncD.Unlock()
	threadFinishMap[threadName] = false
}

func finishThread(threadName string) {
	syncD.Lock()
	defer syncD.Unlock()
	threadFinishMap[threadName] = true
}

func isAllThreadFinished() bool {
	syncD.Lock()
	defer syncD.Unlock()
	fin := true
	for _, b := range threadFinishMap {
		fin = fin && b
		if !fin {
			return false
		}
	}
	return true
}

func beginThread(threadName string, lib *ast.KnowledgeLibrary, t *testing.T) {
	concurrencyTestlog.Tracef("Beginning thread %s", threadName)

	// Register this thread into our simple finish map check
	startThread(threadName)
	defer finishThread(threadName)

	// Prepare new DataContext to be used in this Thread
	dataContext := ast.NewDataContext()

	// Create our model and add into data context
	vibo := &Vibonaci{}
	err := dataContext.Add("Vibo", vibo)
	if err != nil {
		t.Fatalf("DataContext error on thread %s, got %s", threadName, err)
	}

	// Create a new engine for this thread
	engine := &engine.GruleEngine{MaxCycle: 100}

	// Get an instance of our KnowledgeBase from KnowledgeLibrary
	kb := lib.NewKnowledgeBaseInstance("VibonaciTest", "0.0.1")

	// Execute the KnowledgeBase against DataContext
	err = engine.Execute(dataContext, kb)
	if err != nil {
		t.Fatalf("Engine execution error on thread %s, got %s", threadName, err)
	}

	// Make sure this thread result a consistent result
	if vibo.Count != 80 {
		t.Errorf("%s Expected vibo.Count == 80 but %d", threadName, vibo.Count)
	}
	if vibo.A != 37889062373143906 {
		t.Errorf("%s Expected vibo.A == 37889062373143906 but %d", threadName, vibo.Count)
	}
	if vibo.B != 61305790721611591 {
		t.Errorf("%s Expected vibo.B == 61305790721611591 but %d", threadName, vibo.Count)
	}
	if vibo.C != 61305790721611591 {
		t.Errorf("%s Expected vibo.C == 61305790721611591 but %d", threadName, vibo.Count)
	}

	// We are done.
	concurrencyTestlog.Tracef("Finishing thread %s, Count:%d, A:%d,B:%d,C:%d", threadName, vibo.Count, vibo.A, vibo.B, vibo.C)
}

func TestConcurrency(t *testing.T) {
	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("VibonaciTest", "0.0.1", pkg.NewBytesResource([]byte(VibonaciRule)))
	assert.NoError(t, err)

	// Now start 500 thread that uses the KnowledgeLibrary
	for i := 1; i <= 500; i++ {
		go beginThread(fmt.Sprintf("T%d", i), lib, t)
	}

	// Wait until all thread finishes
	time.Sleep(1 * time.Second)
	for !isAllThreadFinished() {
		time.Sleep(500 * time.Millisecond)
	}
}
