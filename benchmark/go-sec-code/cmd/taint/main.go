package main

import (
	"log"

	"github.com/cokeBeer/goot/pkg/example/dataflow/taint"
)

func main() {
	runner := taint.NewRunner("go-sec-code")
	runner.ModuleName = "go-sec-code"
	runner.PassThroughSrcPath = []string{}
	runner.UsePointerAnalysis = true
	runner.PassThroughDstPath = "passthrough.json"
	runner.TaintGraphDstPath = "taintgraph.json"
	runner.PassThroughOnly = false
	runner.InitOnly = false
	runner.TargetFunc = ""
	runner.Debug = true
	runner.PersistToNeo4j = true
	runner.Neo4jURI = "bolt://localhost:7687"
	runner.Neo4jUsername = "neo4j"
	runner.Neo4jPassword = "password"
	err := runner.Run()
	if err != nil {
		log.Fatal(err)
	}
}
