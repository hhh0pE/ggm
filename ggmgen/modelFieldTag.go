package main

import (
	"strconv"
	"strings"
)

type modelFieldTag struct {
	Name, Value string
}

func ParseFieldTags(tags string) []modelFieldTag {
	if len(tags) == 0 {
		return nil
	}
	var fieldTags []modelFieldTag
	fields := strings.Fields(tags)
	for _, field := range fields {
		var newTag modelFieldTag
		if splited := strings.Split(field, ":"); len(splited) == 2 {
			var unquoting_err error
			if newTag.Name, unquoting_err = strconv.Unquote(splited[0]); unquoting_err != nil {
				continue
			}
			if newTag.Value, unquoting_err = strconv.Unquote(splited[1]); unquoting_err != nil {
				continue
			}
			fieldTags = append(fieldTags, newTag)
		}
	}
	return fieldTags
}
