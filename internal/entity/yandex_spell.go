package entity

type TextInput struct {
	Text    string `json:"text" binding:"required"`
	Lang    string `form:"lang,omitempty"`
	Options int    `form:"options,omitempty"`
	Format  bool   `form:"format,omitempty"`
}

type TextsInput struct {
	Texts   []string `json:"texts" binding:"required"`
	Lang    string   `form:"lang,omitempty"`
	Options int      `form:"options,omitempty"`
	Format  bool     `form:"format,omitempty"`
}

type Response [][]Misspell

type Misspell struct {
	Code        int      `json:"code"`
	Pos         int      `json:"pos"`
	Row         int      `json:"row"`
	Col         int      `json:"col"`
	Len         int      `json:"len"`
	Word        string   `json:"word"`
	Suggestions []string `json:"s"`
}

type CorrectorResponse struct {
	Texts []string `json:"texts"`
}
