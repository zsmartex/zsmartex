package filters

import (
	"time"

	"github.com/zsmartex/zsmartex/pkg/mongo_fx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func WithID(id string) mongo_fx.Filter {
	return func() (k string, v interface{}) {
		_id, _ := primitive.ObjectIDFromHex(id)

		return "_id", _id
	}
}

func WithIDs(ids ...string) mongo_fx.Filter {
	return func() (k string, v interface{}) {
		var _ids []primitive.ObjectID

		for _, id := range ids {
			_id, _ := primitive.ObjectIDFromHex(id)
			_ids = append(_ids, _id)
		}

		return "_id", bson.M{"$in": _ids}
	}
}

func WithCreatedAtBy(created_at time.Time) mongo_fx.Filter {
	return func() (k string, v interface{}) {
		return "created_at", created_at
	}
}

func WithUpdatedAtBy(updated_at time.Time) mongo_fx.Filter {
	return func() (k string, v interface{}) {
		return "updated_at", updated_at
	}
}

func WithCreatedAtAfter(created_at time.Time) mongo_fx.Filter {
	return func() (k string, v interface{}) {
		return "created_at", bson.M{"$gt": created_at}
	}
}

func WithCreatedAtBefore(created_at time.Time) mongo_fx.Filter {
	return func() (k string, v interface{}) {
		return "created_at", bson.M{"$lt": created_at}
	}
}

func WithUpdatedAtAfter(updated_at time.Time) mongo_fx.Filter {
	return func() (k string, v interface{}) {
		return "updated_at", bson.M{"$gt": updated_at}
	}
}

func WithUpdatedAtBefore(updated_at time.Time) mongo_fx.Filter {
	return func() (k string, v interface{}) {
		return "updated_at", bson.M{"$lt": updated_at}
	}
}
