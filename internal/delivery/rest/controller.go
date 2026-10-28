package rest

type WebResponse[T any] struct {
	Data T `json:"data"`
}
