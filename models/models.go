package models

// Item defines an item stored in the local store
type Item struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
