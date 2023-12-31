package mongo_fx

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReadRepository[T any] interface {
	Collection(ctx context.Context, filters ...Filter) *mongo.Collection
	Count(ctx context.Context, filters ...Filter) (int, error)
	Find(ctx context.Context, opts *options.FindOptions, filters ...Filter) (models []*T, err error)
	First(ctx context.Context, filters ...Filter) (model *T, err error)
	Last(ctx context.Context, filters ...Filter) (model *T, err error)
}

type WriteRepository[T any] interface {
	Collection(ctx context.Context, filters ...Filter) *mongo.Collection
	FirstOrCreate(ctx context.Context, model *T, create *T, filters ...Filter) error
	Create(ctx context.Context, model *T, filters ...Filter) error
	Updates(ctx context.Context, model interface{}, value map[string]interface{}, filters ...Filter) error
	Delete(ctx context.Context, filters ...Filter) error
}

func NewReadRepository[T any](collection *mongo.Collection) ReadRepository[T] {
	return newRepository[T](collection)
}

func NewWriteRepository[T any](collection *mongo.Collection) WriteRepository[T] {
	return newRepository[T](collection)
}

type repository[T any] struct {
	collection *mongo.Collection
}

func newRepository[T any](collection *mongo.Collection) repository[T] {
	return repository[T]{
		collection,
	}
}

func (r repository[T]) Collection(ctx context.Context, filters ...Filter) *mongo.Collection {
	return r.collection
}

func (r repository[T]) Count(ctx context.Context, filters ...Filter) (int, error) {
	result, err := r.collection.CountDocuments(ctx, ApplyFilters(filters...))
	return int(result), err
}

func (r repository[T]) Find(ctx context.Context, opts *options.FindOptions, filters ...Filter) (models []*T, err error) {
	cursor, err := r.collection.Find(ctx, ApplyFilters(filters...), opts)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &models); err != nil {
		return nil, err
	}

	return models, nil
}

func (r repository[T]) First(ctx context.Context, filters ...Filter) (model *T, err error) {
	if err := r.collection.FindOne(ctx, ApplyFilters(filters...)).Decode(&model); err != nil {
		return nil, err
	}

	return
}

func (r repository[T]) Last(ctx context.Context, filters ...Filter) (model *T, err error) {
	if err := r.collection.FindOne(
		ctx,
		ApplyFilters(filters...),
		options.FindOne().SetSort(bson.D{{Key: "_id", Value: -1}}),
	).Decode(&model); err != nil {
		return nil, err
	}

	return model, nil
}

func (r repository[T]) FirstOrCreate(ctx context.Context, model *T, create *T, filters ...Filter) error {
	err := r.collection.FindOne(ctx, ApplyFilters(filters...)).Decode(model)
	if err == nil {
		return nil
	}

	_, err = r.collection.InsertOne(ctx, create)
	if err != nil {
		return err
	}

	return nil
}

func (r repository[T]) Create(ctx context.Context, model *T, filters ...Filter) error {
	if _, err := r.collection.InsertOne(ctx, &model); err != nil {
		return err
	}

	return nil
}

func (r repository[T]) Updates(ctx context.Context, model interface{}, value map[string]interface{}, filters ...Filter) error {
	result, err := r.collection.UpdateOne(ctx, model, map[string]interface{}{"$set": value})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (r repository[T]) Delete(ctx context.Context, filters ...Filter) error {
	_, err := r.collection.DeleteOne(ctx, ApplyFilters(filters...))
	if err != nil {
		return err
	}

	return nil
}
