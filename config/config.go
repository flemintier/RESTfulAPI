package config

type Config struct {
	Port string
}

// Structure des documents
type Doc struct {
	ID          int    `json:"iD"`
	Nom         string `json:"nom"`
	Description string `json:"description"`
}
