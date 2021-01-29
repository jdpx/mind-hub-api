package graphcms

// Course is a GraphCMS representation of a Course
type Course struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Sessions    []*Session `json:"sessions"`
}

// Session is a GraphCMS representation of a Session
type Session struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Steps       []*Step `json:"steps"`
	Course      *Course `json:"course"`
}

// Step is a GraphCMS representation of a Step
type Step struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	VideoURL    *string  `json:"videoUrl"`
	Audio       *Audio   `json:"audio"`
	Question    *string  `json:"question"`
	Session     *Session `json:"session"`
}

type Audio struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}
