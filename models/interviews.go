package models

import (
	"time"
)

type Interview struct {
	Company     string      `bson:"company" json:"company" validate:"required"`
	Role        string      `bson:"role" json:"role" validate:"required"`
	Level       string      `bson:"level" json:"level" validate:"required,oneof=intenship fresher experienced"`
	Questions   []Questions `bson:"questions" json:"questions" default:""`
	Experience  string      `bson:"experience" json:"experience" validate:"required"`
	Tips        string      `bson:"tips" json:"tips" default:""`
	Tags        []string    `bson:"tags" json:"tags" validate:"required"`
	IsAnonymous bool        `bson:"isanonymous" json:"isanonymous" default:"true"`
	AuthorId    string      `bson:"authorId" json:"authorId"`
	Status      string      `bson:"status" json:"status" validation:"required, oneof=draft published rejected flagged pending" default:"pending"`
	Authorname  string      `bson:"authorname" json:"authorname" validation:"required"`
	Likes       int         `bson:"likes" json:"likes" default:"0"`
	Comments    int         `bson:"comments" json:"Comments" default:"0"`
	Views       int         `bson:"views" json:"views" default:"0"`
	CreatedAt   time.Time   `bson:"createdat" json:"createdat" default:"time.Now()"`
	UpdateAt    time.Time   `bson:"updateat" json:"updateat" default:"time.Now()"`
	LikedBy     []string    `bson:"likedby" json:"likedby" default:"[]"`
	Flags       []Flags     `bson:"flags" json:"flags" default:"[]"`
}

type Questions struct {
	Question string `bson:"question" json:"question"`
	Answer   string `bson:"answer" json:"answer"`
}

type Flags struct {
	Reason     string `bson:"reason" json:"reason" validate:"required"`
	ReportedBy string `bson:"reportedby" json:"reportedby" validate:"required"`
	Date       string `bson:"date" json:"date" default:"time.Now()"`
}
