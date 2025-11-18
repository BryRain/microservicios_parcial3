package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Libro struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Titulo      string            `bson:"titulo" json:"titulo"`
	Autor       string            `bson:"autor" json:"autor"`
	Anio        int               `bson:"anio" json:"anio"`
	Editorial   string            `bson:"editorial" json:"editorial"`
}
