package models

type MessageToProcess struct {
	RetryNumber  int          `json:"retryNumber"`
	JobID        int64        `json:"jobId" example:"12345"`
	UserID       string       `json:"userId" example:"steve.wozniak"`
	ResumeID     int64        `json:"resumeId" example:"12345"`
	BoardID      *int64       `json:"boardId,omitempty" example:"12345"`
	Origin       int          `json:"origin" example:"12345"`
	Rank         int          `json:"rank" example:"100"`
	CreationDate string       `json:"creationDate" example:"2022-08-03"`
	Skills       []UserSkills `json:"skills"`
}

type UserSkills struct {
	ID     int64  `json:"skillId" example:"1"`
	Title  string `json:"skillTitle" example:"Excel"`
	Rating int    `json:"rating" example:"2"`
}
