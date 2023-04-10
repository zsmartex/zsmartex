package projections

import (
	"context"
	"errors"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/eventhandler/projector"
	"github.com/zsmartex/pkg/v3/utils"
	"github.com/zsmartex/zsmartex/internal/users/domain"
	"github.com/zsmartex/zsmartex/internal/users/domain/events"
	"github.com/zsmartex/zsmartex/internal/users/infras/repo"
	"github.com/zsmartex/zsmartex/pkg/encryption"
	userv1 "github.com/zsmartex/zsmartex/proto/common/user/v1"
)

var _ = projector.Projector(&UserProjector{})

const userProjectorType projector.Type = "user"

type UserProjector struct {
	userRepository repo.UserRepository
}

func NewUserProjector(userRepository repo.UserRepository) *UserProjector {
	return &UserProjector{
		userRepository,
	}
}

func (p *UserProjector) ProjectorType() projector.Type {
	return userProjectorType
}

func (p *UserProjector) Project(ctx context.Context, event eh.Event, entity eh.Entity) (eh.Entity, error) {
	user, ok := entity.(*domain.User)
	if !ok {
		return nil, errors.New("model is of incorrect type")
	}

	switch event.EventType() {
	case domain.UserCreatedEvent:
		data, ok := event.Data().(*events.UserCreatedEvent)
		if !ok {
			return nil, errors.New("invalid event data")
		}

		if len(data.Email) > 0 {
			applyCredentials(user, userv1.UserCredentials_EMAIL, data.Email)
		}

		if len(data.Phone) > 0 {
			applyCredentials(user, userv1.UserCredentials_PHONE, data.Phone)
		}

		user.ID = event.AggregateID()
		user.PasswordDigest = data.PasswordDigest
		user.Role = userv1.UserRole_MEMBER
		user.State = userv1.UserState_PENDING
		user.Level = 0
		user.CreatedAt = event.Timestamp()
		user.UpdatedAt = event.Timestamp()
	case domain.UserUpdatedEvent:
		data, ok := event.Data().(*events.UserUpdatedEvent)
		if !ok {
			return nil, errors.New("invalid event data")
		}

		if len(data.Email) > 0 {
			applyCredentials(user, userv1.UserCredentials_EMAIL, data.Email)
		}

		if len(data.Phone) > 0 {
			applyCredentials(user, userv1.UserCredentials_PHONE, data.Phone)
		}

		if len(data.PasswordDigest) > 0 {
			user.PasswordDigest = data.PasswordDigest
		}

		if len(userv1.UserRole_name[int32(data.Role)]) > 0 {
			user.Role = data.Role
		}

		if len(userv1.UserState_name[int32(data.State)]) > 0 {
			user.State = data.State
		}

		user.UpdatedAt = event.Timestamp()
	case domain.UserLabelAppliedEvent:
		data, ok := event.Data().(*events.UserLabelAppliedEvent)
		if !ok {
			return nil, errors.New("invalid event data")
		}

		applyLabel(user, data.Key, data.Value, data.Description)
	case domain.UserLabelDestroyedEvent:
		data, ok := event.Data().(*events.UserLabelDestroyedEvent)
		if !ok {
			return nil, errors.New("invalid event data")
		}

		if err := destroyLabel(user, data.Key); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func applyCredentials(
	user *domain.User,
	credentialsType userv1.UserCredentials_Type,
	value string,
) {
	credentials := &domain.UserCredentials{
		Type:           credentialsType,
		ValueIndex:     utils.HashStringCRC32(value),
		ValueEncrypted: encryption.Encrypt(value),
	}

	for i, c := range user.Credentials {
		if c.Type == credentialsType {
			user.Credentials[i] = credentials
			return
		}
	}

	user.Credentials = append(user.Credentials, credentials)
}

func applyLabel(
	user *domain.User,
	key string,
	value string,
	description string,
) {
	label := &domain.UserLabel{
		Key:         key,
		Value:       value,
		Description: description,
	}

	for i, c := range user.Labels {
		if c.Key == key {
			user.Labels[i] = label
			return
		}
	}

	user.Labels = append(user.Labels, label)
}

func destroyLabel(user *domain.User, key string) error {
	for i, label := range user.Labels {
		if key == label.Key {
			user.Labels[i] = user.Labels[len(user.Labels)-1]
			user.Labels[len(user.Labels)-1] = &domain.UserLabel{}
			user.Labels = user.Labels[:len(user.Labels)-1]

			return nil
		}
	}

	return errors.New("label is not exists")
}
