package dto

type Event struct {
	Topic string `json:"topic"`
	Key   string `json:"key"`
	Data  any    `json:"data"`
}
