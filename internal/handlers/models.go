package handlers

type ShortenRequest struct {
	URL string `json:"url" valid:"url,required"`
}

type ShortenReply struct {
	ShortenURLResult string `json:"result"`
}

type ResponseEntity struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}
