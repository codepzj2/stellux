package apiwrap

import (
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func ConvertBsonID(id string) bson.ObjectID {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		slog.Warn("ConvertBsonID", "error", err.Error())
		return bson.ObjectID{}
	}
	return objID
}
