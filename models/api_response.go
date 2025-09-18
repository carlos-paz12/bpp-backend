package models

type APIResponse struct {
	Message  string `json:"message"`        //!<
	Error    string `json:"error"`          //!<
	HttpCode int    `json:"http_code"`      //!<
	Data     any    `json:"data,omitempty"` //!<
}
