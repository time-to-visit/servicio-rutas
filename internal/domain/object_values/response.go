package objectValues

type responseWithData struct {
	StatusCode int         `json:"status_code"`
	Title      string      `json:"title"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func NewResponseWithData(StatusCode int, Title string, Message string, Data interface{}) responseWithData {
	return responseWithData{
		StatusCode,
		Title,
		Message,
		Data,
	}
}
