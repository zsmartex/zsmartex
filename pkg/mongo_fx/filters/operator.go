package filters

import (
	"github.com/zsmartex/zsmartex/pkg/mongo_fx"
	"go.mongodb.org/mongo-driver/bson"
)

func applyFilter(filters ...mongo_fx.Filter) []bson.M {
	value := []bson.M{}
	for _, filter := range filters {
		k, v := filter()
		value = append(value, bson.M{k: v})
	}

	return value
}

func WithAnd(filters ...mongo_fx.Filter) mongo_fx.Filter {
	return func() (k string, v interface{}) {
		return "$and", applyFilter(filters...)
	}
}

func WithOr(filters ...mongo_fx.Filter) mongo_fx.Filter {
	return func() (k string, v interface{}) {
		return "$or", applyFilter(filters...)
	}
}

func WithNot(filters ...mongo_fx.Filter) mongo_fx.Filter {
	return func() (k string, v interface{}) {
		return "$not", applyFilter(filters...)
	}
}

func WithNor(filters ...mongo_fx.Filter) mongo_fx.Filter {
	return func() (k string, v interface{}) {
		return "$nor", applyFilter(filters...)
	}
}
