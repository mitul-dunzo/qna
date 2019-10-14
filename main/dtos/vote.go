package dtos

const (
	Upvote   = "upvote"
	Downvote = "downvote"
)

type UserVote struct {
	QuestionId uint `json:"question"`
	AnswerId   uint `json:"answer"`
	Vote       int  `json:"vote"`
}

type Vote struct {
	UpvotedUserId uint   `json:"upvoted_user_id"`
	Type          string `json:"type"`
	EntityId      uint   `json:"entity_id"`
	Vote          string `json:"vote"`
}
