package editor

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	mux "github.com/hyperjumptech/hyper-mux"
	"io/ioutil"
	"net/http"
)

type JSONData struct {
	Name     string
	JSONData string `json:"jsonText"`
}

type EvaluateRequest struct {
	GrlText string      `json:"grlText"`
	Input   []*JSONData `json:"jsonInput"`
}

func InitializeEvaluationRoute(router *mux.HyperMux) {
	router.AddRoute("/evaluate", http.MethodPost, func(writer http.ResponseWriter, reader *http.Request) {
		bodyBytes, err := ioutil.ReadAll(reader.Body)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(fmt.Sprintf("error while reading body stream. got %v", err)))

			return
		}
		evReq := &EvaluateRequest{}
		err = json.Unmarshal(bodyBytes, evReq)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = writer.Write([]byte(fmt.Sprintf("wrong json format. got %v \n\n Json : %s", err, string(bodyBytes))))

			return
		}

		dataContext := ast.NewDataContext()

		for _, JSONDat := range evReq.Input {
			jsonByte, err := base64.StdEncoding.DecodeString(JSONDat.JSONData)
			if err != nil {
				writer.WriteHeader(http.StatusBadRequest)
				_, _ = writer.Write([]byte(fmt.Sprintf("json data named %s should be sent using base64. got %v", JSONDat.Name, err)))

				return
			}
			err = dataContext.AddJSON(JSONDat.Name, jsonByte)
			if err != nil {
				writer.WriteHeader(http.StatusBadRequest)
				_, _ = writer.Write([]byte(fmt.Sprintf("invalid JSON data named %s when add json to context got %v", JSONDat.Name, err)))

				return
			}
		}

		knowledgeLibrary := ast.NewKnowledgeLibrary()
		ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

		grlByte, err := base64.StdEncoding.DecodeString(evReq.GrlText)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = writer.Write([]byte(fmt.Sprintf("GRL data should be sent using base64. got %v", err)))

			return
		}

		err = ruleBuilder.BuildRuleFromResource("Evaluator", "0.0.1", pkg.NewBytesResource(grlByte))
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = writer.Write([]byte(fmt.Sprintf("invalid GRL : %s", err.Error())))

			return
		}

		eng1 := &engine.GruleEngine{MaxCycle: 5}
		kb := knowledgeLibrary.NewKnowledgeBaseInstance("Evaluator", "0.0.1")
		err = eng1.Execute(dataContext, kb)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = writer.Write([]byte(fmt.Sprintf("Grule Error : %s", err.Error())))

			return
		}

		respData := make(map[string]interface{})
		for _, keyName := range dataContext.GetKeys() {
			respData[keyName] = dataContext.Get(keyName)
		}

		resultBytes, err := json.Marshal(respData)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(fmt.Sprintf("Error marshaling result : %s", err.Error())))

			return
		}

		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(resultBytes)
	})
}
