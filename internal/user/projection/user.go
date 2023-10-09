package projection

import (
	"context"
	"sync"

	"github.com/modernice/goes/event"
	"github.com/modernice/goes/projection"
	"github.com/zsmartex/zsmartex/internal/user/domain"
	"github.com/zsmartex/zsmartex/internal/user/events"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	sync.RWMutex
	*projection.Base

	ctx             context.Context
	usersCollection *mongo.Collection
}

func NewUser(ctx context.Context, mongoClient *mongo.Client) *User {
	userProjection := &User{
		Base:            projection.New(),
		ctx:             ctx,
		usersCollection: mongoClient.Database("user").Collection("users"),
	}

	event.ApplyWith(userProjection, userProjection.userCreated, string(events.UserCreated))

	return userProjection
}

func (u *User) userCreated(evt event.Of[events.UserCreatedData]) {
	u.Lock()
	defer u.Unlock()

	data := evt.Data()

	u.usersCollection.InsertOne(u.ctx, domain.User{
		ID:             primitive.NewObjectID(),
		UID:            data.UID,
		Email:          data.Email,
		PasswordDigest: data.PasswordDigest,
		Role:           data.Role,
		State:          data.State,
		Labels:         data.Labels,
		CreatedAt:      evt.Time(),
		UpdatedAt:      evt.Time(),
	})
}
