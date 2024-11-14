package entity

import "go.mongodb.org/mongo-driver/v2/bson"

type UserEntity struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Email    string        `bson:"email,omitempty"`
	Password string        `bson:"password,omitempty"`
	Name     string        `bson:"name,omitempty"`
	NickName string        `bson:"nick_name,omitempty"`
	QtyWins  uint          `bson:"qty_wins,omitempty"`
}
