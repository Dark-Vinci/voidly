package models

type Message struct {
	Type       string   `json:"type"`
	Text       string   `json:"text"`
	Server     string   `json:"server"`
	FromUserID string   `json:"from_user_id"`
	ToUserID   *string  `json:"to_user_id"`
	ToGroupID  *string  `json:"to_group_id"`
	Content    string   `json:"content"`
	Links      []string `json:"links"`
}
