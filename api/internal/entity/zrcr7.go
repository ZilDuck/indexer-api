package entity

import "encoding/json"

type Zrc7Attribute struct {
	TraitType   string      `json:"trait_type"`
	Value       interface{} `json:"value"`
	MimeType    *string     `json:"mime_type"`
	Integrity   *string     `json:"integrity"`
	DisplayType *string     `json:"display_type"`
}

func MapToZrc7Attributes(data interface{}) ([]Zrc7Attribute, error) {
	byteData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var attributes []Zrc7Attribute
	if err := json.Unmarshal(byteData, &attributes); err != nil {
		return nil, err
	}

	return attributes, nil
}