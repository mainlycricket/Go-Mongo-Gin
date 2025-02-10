package responses

type DefaultApiResponse struct {
	Successs bool   `json:"success"`
	Message  string `json:"message"`
	Data     any    `json:"data"`
}
