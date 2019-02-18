package database

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type DefTimeInfo interface {
	SetCreated(time.Time)
	SetModified(time.Time)
}

type Doc interface {
	GetId() bson.ObjectId
	SetId(bson.ObjectId)
}

type DocEspecial interface {
	GetId() string
	SetId(string)
}

type BaseDoc struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Created  time.Time     `bson:"_created" json:"_created"`
	Modified time.Time     `bson:"_modified" json:"_modified"`
}

// Satisfy the document interface
func (d *BaseDoc) GetId() bson.ObjectId {
	return d.Id
}

func (d *BaseDoc) SetId(id bson.ObjectId) {
	d.Id = id
}

func (d *BaseDoc) SetCreated(t time.Time) {
	d.Created = t
}

func (d *BaseDoc) SetModified(t time.Time) {
	d.Modified = t
}
