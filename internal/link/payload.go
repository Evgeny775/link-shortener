package link

type LinkCreateRequest struct {
	URL string `json:"url" validate:"requaered,url"`
}
