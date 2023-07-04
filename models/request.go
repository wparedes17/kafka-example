package models

type MsgRequest struct {
	MsgKey  string `schema:"msg_key"`
	MsgText string `schema:"msg_text"`
}
