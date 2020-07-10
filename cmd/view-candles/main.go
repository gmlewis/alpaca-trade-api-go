// view-candles copies the provided JSON candlestick data file to a /tmp dir
// and installs the trading-vue-js app, then runs "go-wasm" on it to display
// the chart.
//
// This relies on the files located in srcDir below.
//
// See:
//   https://github.com/mfrachet/go-wasm-cli
//   https://github.com/tvjsx/trading-vue-demo
//   https://github.com/tvjsx/trading-vue-js
//
// Usage:
//   view-candles daily/acb.json
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gmlewis/alpaca-trade-api-go/alpaca"
	jsoniter "github.com/json-iterator/go"
)

const (
	srcDir = "/home/glenn/src/github.com/tvjsx/trading-vue-demo"
)

var (
	outTemp = template.Must(template.New("out").Funcs(funcMap).Parse(templateStr))
	funcMap = template.FuncMap{
		"toMillis": toMillis,
	}
)

func toMillis(v int64) int64 {
	return v * 1000
}

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatalf("Usage: %v file.json", filepath.Base(os.Args[0]))
	}

	tempDir, err := ioutil.TempDir("", "view-candles")
	check("TempDir: %v", err)
	// defer os.RemoveAll(tempDir)

	// Copy the trading-vue-js files...
	for _, filename := range []string{"index.html", "trading-vue.min.js"} {
		copyFile(tempDir, filename)
	}

	// Create the data file...
	buf, err := ioutil.ReadFile(flag.Arg(0))
	check("ReadFile: %v", err)
	var bars []alpaca.Bar
	check("Unmarshal: %v", jsoniter.Unmarshal(buf, &bars))

	var bbuf bytes.Buffer
	err = outTemp.Execute(&bbuf, bars)
	check("Execute: %v", err)
	err = ioutil.WriteFile(filepath.Join(tempDir, "data.json"), bbuf.Bytes(), 0644)
	check("WriteFile: %v", err)

	// Run go-wasm in the tempDir
	os.Chdir(tempDir)
	err = ioutil.WriteFile(filepath.Join(tempDir, "run.sh"), []byte(bashScript), 0755)
	check("WriteFile: %v", err)
	fmt.Printf("pushd %v && ./run.sh\n", tempDir)
}

func copyFile(dstDir, filename string) {
	buf, err := ioutil.ReadFile(filepath.Join(srcDir, filename))
	check("ReadFile: %v", err)
	err = ioutil.WriteFile(filepath.Join(dstDir, filename), buf, 0644)
	check("WriteFile: %v", err)
}

func check(fmtStr string, args ...interface{}) {
	if err := args[len(args)-1]; err != nil {
		log.Fatalf(fmtStr, args...)
	}
}

var templateStr = `Data = {
"ohlcv": [
{{- range . }}
  [ {{ .Time | toMillis }}, {{ .Open }}, {{ .High }}, {{ .Low }}, {{ .Close }}, {{ .Volume }} ],
{{- end }}
]
}
`

var bashScript = `#!/bin/bash -x
go-wasm start
`
