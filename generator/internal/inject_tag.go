package internal

import (
	"fmt"
	"strings"
	"regexp"
)

var (
	injectRule = regexp.MustCompile(`^//\s*@inject_tag:\s*(.*)$`)
	rTags    = regexp.MustCompile(`[\w_]+:"[^"]+"`)
)

type TagItem struct {
	Key   string
	Value string
}

type TagItems []TagItem

func (ti TagItems) Format() string {
	tags := []string{}
	for _, item := range ti {
		tags = append(tags, fmt.Sprintf(`%s:%s`, item.Key, item.Value))
	}
	return strings.Join(tags, " ")
}

func tagFromComment(comment string) (tag string) {
	match := injectRule.FindStringSubmatch(comment)
	if len(match) == 2 {
		tag = match[1]
	}
	return
}

func NewTagItems(tag string) TagItems {
	items := []TagItem{}
	splitted := rTags.FindAllString(tag, -1)

	for _, t := range splitted {
		t = strings.Replace(t, " ", "", -1)
		sepPos := strings.Index(t, ":")
		items = append(items, TagItem{
			Key:   t[:sepPos],
			Value: t[sepPos+1:],
		})
	}
	return items
}
