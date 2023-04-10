package postgres

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm/clause"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// ErrCouldNotDialDB is when the database could not be dialed.
	ErrCouldNotDialDB = errors.New("could not dial database")
	// ErrNoDBClient is when no database client is set.
	ErrNoDBClient = errors.New("no database client")
	// ErrCouldNotClearDB is when the database could not be cleared.
	ErrCouldNotClearDB = errors.New("could not clear database")
	// ErrModelNotSet is when an model factory is not set on the Repo.
	ErrModelNotSet = errors.New("model not set")
	// ErrInvalidQuery is when a query was not returned from the callback to FindCustom.
	ErrInvalidQuery = errors.New("invalid query")
)

// Repo implements a PostgreSQL repo for entities.
type Repo[T eh.Entity] struct {
	client *gorm.DB
}

func NewRepo[T eh.Entity](host, port, user, password, database string) (*Repo[T], error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, database, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not connect to DB: %w", err)
	}

	return NewRepoWithClient[T](db)
}

func NewRepoWithClient[T eh.Entity](client *gorm.DB) (*Repo[T], error) {
	if client == nil {
		return nil, ErrNoDBClient
	}

	return &Repo[T]{client: client}, nil
}

func (r *Repo[T]) MigrateTables(entity interface{}) error {
	return r.client.AutoMigrate(entity)
}

func (r *Repo[T]) Parent() eh.ReadRepo {
	return nil
}

func (r *Repo[T]) InnerRepo(ctx context.Context) eh.ReadRepo {
	return nil
}

func IntoRepo[T eh.Entity](ctx context.Context, repo eh.ReadRepo) *Repo[T] {
	if repo == nil {
		return nil
	}

	if r, ok := repo.(*Repo[T]); ok {
		return r
	}

	return IntoRepo[T](ctx, repo.InnerRepo(ctx))
}

func (r *Repo[T]) Find(ctx context.Context, id uuid.UUID) (eh.Entity, error) {
	var entity T

	if err := r.client.WithContext(ctx).Preload(clause.Associations).First(&entity, "id = ?", id.String()).Error; err == gorm.ErrRecordNotFound {
		return nil, &eh.RepoError{
			Err:      eh.ErrEntityNotFound,
			Op:       eh.RepoOpFind,
			EntityID: id,
		}
	} else if err != nil {
		return nil, &eh.RepoError{
			Err:      err,
			Op:       eh.RepoOpFind,
			EntityID: id,
		}
	}

	return entity, nil
}

func (r *Repo[T]) FindAll(ctx context.Context) ([]eh.Entity, error) {
	var records []T

	if err := r.client.WithContext(ctx).Preload(clause.Associations).Find(&records).Error; err != nil {
		return nil, &eh.RepoError{
			Err: err,
			Op:  eh.RepoOpFindAll,
		}
	}

	entities := make([]eh.Entity, 0)
	for _, v := range records {
		entities = append(entities, v)
	}

	return entities, nil
}

func (r *Repo[T]) Save(ctx context.Context, entity eh.Entity) error {
	id := entity.EntityID()
	if id == uuid.Nil {
		return &eh.RepoError{
			Err: fmt.Errorf("missing entity ID"),
			Op:  eh.RepoOpSave,
		}
	}

	if err := r.client.Session(&gorm.Session{Context: ctx, FullSaveAssociations: true}).Save(entity).Error; err != nil {
		return &eh.RepoError{
			Err:      fmt.Errorf("could not save/update: %w", err),
			Op:       eh.RepoOpSave,
			EntityID: id,
		}
	}

	return nil
}

func (r *Repo[T]) Remove(ctx context.Context, id uuid.UUID) error {
	var entity T

	result := r.client.WithContext(ctx).Delete(entity, "id = ?", id.String())
	if err := result.Error; err != nil {
		return &eh.RepoError{
			Err: err,
			Op:  eh.RepoOpRemove,
		}
	} else if result.RowsAffected == 0 {
		return &eh.RepoError{
			Err: eh.ErrEntityNotFound,
			Op:  eh.RepoOpRemove,
		}
	}

	return nil
}

func (r *Repo[T]) Close() error {
	return nil
}
