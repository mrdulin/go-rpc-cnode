package models

type Topic struct {
	ID          string  `json:"id"`
	AuthorID    string  `json:"author_id"`
	Tab         *string `json:"tab"`
	Content     *string `json:"content"`
	Title       string  `json:"title"`
	LastReplyAt *string `json:"last_reply_at"`
	Good        *bool   `json:"good"`
	Top         *bool   `json:"top"`
	ReplyCount  *int    `json:"reply_count"`
	VisitCount  *int    `json:"visit_count"`
	CreateAt    *string `json:"create_at"`
	IsCollect   *bool   `json:"is_collect"`
	Author      *User   `json:"author"`
}

type TopicDetail struct {
	ID          string  `json:"id"`
	AuthorID    string  `json:"author_id"`
	Tab         *string `json:"tab"`
	Content     *string `json:"content"`
	Title       string  `json:"title"`
	LastReplyAt *string `json:"last_reply_at"`
	Good        *bool   `json:"good"`
	Top         *bool   `json:"top"`
	ReplyCount  *int    `json:"reply_count"`
	VisitCount  *int    `json:"visit_count"`
	CreateAt    *string `json:"create_at"`
	IsCollect   *bool   `json:"is_collect"`
	//Replies     []*Reply `json:"replies"`
	Author *User `json:"author"`
}

type TopicForMessage struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	LastReplyAt *string `json:"last_reply_at"`
}

type TopicRecent struct {
	ID          string  `json:"id"`
	Title       *string `json:"title"`
	LastReplyAt *string `json:"last_reply_at"`
	Author      *User   `json:"author"`
}

type TopicTab string

const (
	TopicTabAsk   TopicTab = "ask"
	TopicTabShare TopicTab = "share"
	TopicTabJob   TopicTab = "job"
	TopicTabGood  TopicTab = "good"
)
