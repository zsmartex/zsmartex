package mongo_fx

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Filter func() (string, interface{})

func ApplyFilters(fs ...Filter) bson.M {
	result := make(bson.M)

	for _, f := range fs {
		key, value := f()
		result[key] = value
	}

	return result
}
