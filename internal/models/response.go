package models

type Response struct {
	IsError bool        `json:"isError"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}
