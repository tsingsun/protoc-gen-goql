package generator

import (
	"strings"
	"github.com/tsingsun/protoc-gen-goql/generator/internal"
	"fmt"
)

var jsonOmitEmptyMap map[string]bool = map[string]bool{}

// Name returns the name of this plugin, "grpcserial".
func (g *Generator) Name() string {
	return "qeelyn"
}

func (g *Generator) SetJsonOmitEmptyMap(list string) {
	val := strings.Split(list, ",")
	for _, v := range val {
		jsonOmitEmptyMap[v] = true
	}
}

func (g *Generator) ParseComments(path string) (internal.TagItems, bool) {
	if !g.writeOutput {
		return nil, false
	}
	var tags internal.TagItems
	if loc, ok := g.file.comments[path]; ok {
		text := strings.TrimSuffix(loc.GetLeadingComments(), "\n")
		for _, line := range strings.Split(text, "\n") {
			lineTag := internal.NewTagItems(line)
			for _, v := range lineTag {
				if v.Key != "" {
					tags = append(tags, v)
				}
			}
		}
		return tags, true
	}
	return nil, false
}

func (g *Generator) JoinInjectTagTo(path string, tag string, typename string, jsonName string) string {
	injectTags, ok := g.ParseComments(path)
	var hasTagsJson bool = false
	if ok {
		for _, v := range injectTags {
			if v.Key == "json" {
				hasTagsJson = true
			}
		}
	}
	if hasTagsJson {
		tag += " " + injectTags.Format()
	} else {
		if _, ok := jsonOmitEmptyMap[typename]; !ok {
			tag += " " + fmt.Sprintf("json:%q", jsonName+",omitempty") + " " + injectTags.Format()
		} else {
			tag += " " + fmt.Sprintf("json:%q", jsonName) + " " + injectTags.Format()
		}
	}
	return tag
}
