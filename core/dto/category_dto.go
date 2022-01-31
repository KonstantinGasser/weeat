package dto

type Category struct {
	ID    int    `json:"id"`
	Type  int    `json:"type"`
	Label string `json:"label"`
	Emoji string `json:"emoji"`
}
