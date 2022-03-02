package bd

import (
	"context"
	"time"

	"github.com/raulpressel/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertoTweet graba el tweet en la BD*/

func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err //para devolver un string vacio se puede usar String("") o simplemente "" ojo que esta ultima opcion no es compatible con versiones viejas de go
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil //aca lo mismo se puede usar objID.String para transformarlo en stringo o usar objID.hex()

}
