package generator

import (
	"strings"
)

var isJsonOmitEmpty bool = true
var jsonOmitEmptyMap map[string]bool = map[string]bool{"float64": true,"float32":true,"int32":true,"int64":true}

// Name returns the name of this plugin, "grpcserial".
func (g *Generator) Name() string {
	return "qeelyn"
}

func (g *Generator) SetJsonOmitEmptyMap(list string)  {
	val := strings.Split(list,",")
	for _, v := range val {
		jsonOmitEmptyMap[v] = true
	}
}

func (g *Generator) SetIsJsonOmitEmpty(isEmtpy string)  {
	if isEmtpy == "false" || isEmtpy == "0" {
		isJsonOmitEmpty = false
	}
}
