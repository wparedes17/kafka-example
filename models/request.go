package models

type MsgRequest struct {
	MsgKey   string `schema:"msg_key"`
	JobID    int64  `schema:"job_id"`
	UserID   string `schema:"user_id"`
	ResumeID int64  `schema:"resume_id"`
}
