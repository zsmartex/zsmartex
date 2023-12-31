package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

var CollectionModule = fx.Module(
	"repo.Collection",
	fx.Provide(mapCollection),
)

type collectionList struct {
	fx.Out

	UsersCollection *mongo.Collection `name:"users_collection"`
}

func mapCollection(mongodb *mongo.Database) collectionList {
	return collectionList{
		UsersCollection: mongodb.Collection("users"),
	}
}
