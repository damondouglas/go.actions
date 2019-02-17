package google

import "encoding/json"

const (
	// OptionValueSpec is assigned to SystemIntent.Data.Type for CarouselSelect.
	OptionValueSpec = "type.googleapis.com/google.actions.v2.OptionValueSpec"
)

const (
	// OptionIntent is assigned to SystemIntent.Intent for CarouselSelect.
	OptionIntent = "actions.intent.OPTION"
)

// ExpectedIntent the app is asking the assistant to provide.
type ExpectedIntent struct {
	Intent         string
	InputValueData json.RawMessage
	ParameterName  string
}

// SystemIntent is a system intent.
type SystemIntent struct {
	Intent string `json:"intent,omitempty"`
	Data   *Data  `json:"data,omitempty"`
}

// Data of a system intent.
type Data struct {
	Type           string      `json:"@type,omitempty"`
	CarouselSelect *Select     `json:"carouselSelect,omitempty"`
	ListSelect     *Select     `json:"listSelect,omitempty"`
	DialogSpec     *DialogSpec `json:"dialogSpec,omitempty"`
}

// Select presents options to user.
type Select struct {
	Items []*SelectItem `json:"items,omitempty"`
}

// SelectItem is the option to the user.
type SelectItem struct {
	OptionInfo  *OptionInfo `json:"optionInfo,omitempty"`
	Description string      `json:"description,omitempty"`
	Image       *Image      `json:"image,omitempty"`
	Title       string      `json:"title,omitempty"`
}

// OptionInfo defines data in the carousel select item.
type OptionInfo struct {
	Key      string   `json:"key,omitempty"`
	Synonyms []string `json:"synonyms,omitempty"`
}

// DialogSpec defines specifications for user feedback.
type DialogSpec struct {
	RequestConfirmationText string `json:"requestConfirmationText,omitempty"`
	RequestDatetimeText     string `json:"requestDatetimeText,omitempty"`
	RequestDateText         string `json:"requestDateText,omitempty"`
	RequestTimeText         string `json:"requestTimeText,omitempty"`
}
