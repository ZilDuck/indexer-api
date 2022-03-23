package entity

import "fmt"

type Attributes map[string]map[string]int64

func (a Attributes) HasAttribute(traitType string) bool {
	if _, ok := a[traitType]; ok {
		return true
	}
	return false
}

func (a Attributes) HasTraitValue(traitType string, value interface{}) bool {
	if _, ok := a[traitType][fmt.Sprintf("%v", value)]; ok {
		return true
	}
	return false
}