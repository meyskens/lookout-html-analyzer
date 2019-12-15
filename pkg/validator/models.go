package validator

// Response is the struct that contains the output of a validator service
type Response struct {
	URL      string            `json:"url"`
	Messages []ResponseMessage `json:"messages"`
	Source   struct {
		Code     string `json:"code"`
		Type     string `json:"type"`
		Encoding string `json:"encoding"`
	} `json:"source"`
	Language string `json:"language"`
}

// ResponseMessage is a message sent in Response.Messages
type ResponseMessage struct {
	Type         string `json:"type"`
	Subtype      string `json:"subtype"`
	LastLine     int    `json:"lastLine"`
	LastColumn   int    `json:"lastColumn"`
	URL          string `json:"url"`
	Message      string `json:"message"`
	Extract      string `json:"extract,omitempty"`
	HiliteStart  int    `json:"hiliteStart,omitempty"`
	HiliteLength int    `json:"hiliteLength,omitempty"`
}
