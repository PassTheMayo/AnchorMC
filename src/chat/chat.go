package chat

type Chat struct {
	Text          string `json:"text"`
	Color         Color  `json:"color"`
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Underlined    bool   `json:"underlined"`
	Strikethrough bool   `json:"strikethrough"`
	Obfuscated    bool   `json:"obfuscated"`
}
