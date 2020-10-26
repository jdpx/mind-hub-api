package graphcms

type coursesResponse struct {
	Courses []*Course `json:"courses"`
}

type courseResponse struct {
	Course *Course `json:"course"`
}

type sessionsResponse struct {
	Sessions []*Session `json:"sessions"`
}

type sessionResponse struct {
	Session *Session `json:"session"`
}
