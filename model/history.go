package model

type History struct {
	Name    string "bson:`name` json:`name`"
	Message string "bson:`message` json:`message`"
	Uid     string "bson:`uid` json:`uid`"
}
