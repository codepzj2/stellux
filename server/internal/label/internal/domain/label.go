package domain

type LabelType string

type Label struct {
	ID        string `bson:"_id"`
	LabelType string `bson:"label_type"`
	Name      string `bson:"name"`
}