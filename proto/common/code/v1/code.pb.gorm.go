package codev1

import (
	context "context"
	fmt "fmt"
	gateway "github.com/infobloxopen/atlas-app-toolkit/gateway"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	resource "github.com/infobloxopen/atlas-app-toolkit/gorm/resource"
	null "github.com/mbahjadol/null"
	errors "github.com/zsmartex/protoc-gen-gorm/errors"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	gorm "gorm.io/gorm"
	strings "strings"
	time "time"
)

type UserORM struct {
	ValidatedAt    null.Time   `json:"validated_at,omitempty"`
	PhoneIndex     null.Int64  `json:"phone_index,omitempty"`
	Data           null.Bytes  `json:"data,omitempty"`
	User           *UserORM    `gorm:"foreignKey:UserId;references:Id" json:"user,omitempty"`
	Code           string      `json:"code,omitempty"`
	Type           string      `json:"type,omitempty"`
	EmailEncrypted null.String `json:"email_encrypted,omitempty"`
	PhoneEncrypted null.String `json:"phone_encrypted,omitempty"`
	UserId         *int64      `json:"user_id,omitempty"`
	AttemptCount   int64       `json:"attempt_count,omitempty"`
	ExpiredAt      *time.Time  `json:"expired_at,omitempty"`
	Id             *int64      `gorm:"type:bigint;primaryKey;autoIncrement" json:"id,omitempty"`
	Category       string      `json:"category,omitempty"`
	EmailIndex     null.Int64  `json:"email_index,omitempty"`
}

// TableName overrides the default tablename generated by GORM
func (UserORM) TableName() string {
	return "users"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *User) ToORM(ctx context.Context) (UserORM, error) {
	to := UserORM{}
	var err error
	if prehook, ok := interface{}(m).(UserWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if m.Id != nil {
		if v, err := resource.DecodeInt64(&User{}, m.Id); err != nil {
			return to, err
		} else {
			to.Id = &v
		}
	}
	if m.UserId != nil {
		if v, err := resource.DecodeInt64(&User{}, m.UserId); err != nil {
			return to, err
		} else {
			to.UserId = &v
		}
	}
	to.Code = m.Code
	to.Type = CodeType_name[int32(m.Type)]
	to.Category = CodeCategory_name[int32(m.Category)]
	err = to.EmailEncrypted.Scan(m.EmailEncrypted)
	if err != nil {
		return to, err
	}
	err = to.EmailIndex.Scan(m.EmailIndex)
	if err != nil {
		return to, err
	}
	err = to.PhoneEncrypted.Scan(m.PhoneEncrypted)
	if err != nil {
		return to, err
	}
	err = to.PhoneIndex.Scan(m.PhoneIndex)
	if err != nil {
		return to, err
	}
	to.AttemptCount = m.AttemptCount
	err = to.Data.Scan(m.Data)
	if err != nil {
		return to, err
	}
	t := m.ValidatedAt.AsTime()
	to.ValidatedAt.Scan(t)
	if m.ExpiredAt != nil {
		t := m.ExpiredAt.AsTime()
		to.ExpiredAt = &t
	}
	if m.User != nil {
		tempUser, err := m.User.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.User = &tempUser
	}
	if posthook, ok := interface{}(m).(UserWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *UserORM) ToPB(ctx context.Context) (User, error) {
	to := User{}
	var err error
	if prehook, ok := interface{}(m).(UserWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	if m.Id != nil {
		if v, err := resource.Encode(&User{}, *m.Id); err != nil {
			return to, err
		} else {
			to.Id = v
		}
	}
	if m.UserId != nil {
		if v, err := resource.Encode(&User{}, *m.UserId); err != nil {
			return to, err
		} else {
			to.UserId = v
		}
	}
	to.Code = m.Code
	to.Type = CodeType(CodeType_value[m.Type])
	to.Category = CodeCategory(CodeCategory_value[m.Category])
	to.EmailEncrypted = m.EmailEncrypted.String
	to.EmailIndex = m.EmailIndex.Int64
	to.PhoneEncrypted = m.PhoneEncrypted.String
	to.PhoneIndex = m.PhoneIndex.Int64
	to.AttemptCount = m.AttemptCount
	to.Data = m.Data.Bytes
	if m.ValidatedAt.IsValid() {
		to.ValidatedAt = timestamppb.New(m.ValidatedAt.Time)
	}
	if m.ExpiredAt != nil {
		to.ExpiredAt = timestamppb.New(*m.ExpiredAt)
	}
	if m.User != nil {
		tempUser, err := m.User.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.User = &tempUser
	}
	if posthook, ok := interface{}(m).(UserWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type User the arg will be the target, the caller the one being converted from

// UserBeforeToORM called before default ToORM code
type UserWithBeforeToORM interface {
	BeforeToORM(context.Context, *UserORM) error
}

// UserAfterToORM called after default ToORM code
type UserWithAfterToORM interface {
	AfterToORM(context.Context, *UserORM) error
}

// UserBeforeToPB called before default ToPB code
type UserWithBeforeToPB interface {
	BeforeToPB(context.Context, *User) error
}

// UserAfterToPB called after default ToPB code
type UserWithAfterToPB interface {
	AfterToPB(context.Context, *User) error
}

// DefaultCreateUser executes a basic gorm create call
func DefaultCreateUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type UserORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == nil || *ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &UserORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := UserORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(UserORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type UserORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteUser(ctx context.Context, in *User, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == nil || *ormObj.Id == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&UserORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type UserORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteUserSet(ctx context.Context, in []*User, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []*int64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == nil || *ormObj.Id == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&UserORM{})).(UserORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&UserORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&UserORM{})).(UserORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type UserORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*User, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*User, *gorm.DB) error
}

// DefaultStrictUpdateUser clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateUser")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	var count int64
	lockedRow := &UserORM{}
	count = db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow).RowsAffected
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	filterUser := UserORM{}
	if ormObj.Id == nil || *ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	filterUser.UserId = new(int64)
	*filterUser.UserId = *ormObj.Id
	if err = db.Where(filterUser).Delete(UserORM{}).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		err = gateway.SetCreated(ctx, "")
	}
	return &pbResponse, err
}

type UserORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchUser executes a basic gorm update call with patch behavior
func DefaultPatchUser(ctx context.Context, in *User, updateMask *field_mask.FieldMask, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj User
	var err error
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadUser(ctx, &User{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskUser(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateUser(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(UserWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type UserWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *User, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *User, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *User, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *User, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetUser executes a bulk gorm update call with patch behavior
func DefaultPatchSetUser(ctx context.Context, objects []*User, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*User, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*User, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchUser(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskUser patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskUser(ctx context.Context, patchee *User, patcher *User, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*User, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedValidatedAt bool
	var updatedExpiredAt bool
	var updatedUser bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"UserId" {
			patchee.UserId = patcher.UserId
			continue
		}
		if f == prefix+"Code" {
			patchee.Code = patcher.Code
			continue
		}
		if f == prefix+"Type" {
			patchee.Type = patcher.Type
			continue
		}
		if f == prefix+"Category" {
			patchee.Category = patcher.Category
			continue
		}
		if f == prefix+"EmailEncrypted" {
			patchee.EmailEncrypted = patcher.EmailEncrypted
			continue
		}
		if f == prefix+"EmailIndex" {
			patchee.EmailIndex = patcher.EmailIndex
			continue
		}
		if f == prefix+"PhoneEncrypted" {
			patchee.PhoneEncrypted = patcher.PhoneEncrypted
			continue
		}
		if f == prefix+"PhoneIndex" {
			patchee.PhoneIndex = patcher.PhoneIndex
			continue
		}
		if f == prefix+"AttemptCount" {
			patchee.AttemptCount = patcher.AttemptCount
			continue
		}
		if f == prefix+"Data" {
			patchee.Data = patcher.Data
			continue
		}
		if !updatedValidatedAt && strings.HasPrefix(f, prefix+"ValidatedAt.") {
			if patcher.ValidatedAt == nil {
				patchee.ValidatedAt = nil
				continue
			}
			if patchee.ValidatedAt == nil {
				patchee.ValidatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"ValidatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.ValidatedAt, patchee.ValidatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"ValidatedAt" {
			updatedValidatedAt = true
			patchee.ValidatedAt = patcher.ValidatedAt
			continue
		}
		if !updatedExpiredAt && strings.HasPrefix(f, prefix+"ExpiredAt.") {
			if patcher.ExpiredAt == nil {
				patchee.ExpiredAt = nil
				continue
			}
			if patchee.ExpiredAt == nil {
				patchee.ExpiredAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"ExpiredAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.ExpiredAt, patchee.ExpiredAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"ExpiredAt" {
			updatedExpiredAt = true
			patchee.ExpiredAt = patcher.ExpiredAt
			continue
		}
		if !updatedUser && strings.HasPrefix(f, prefix+"User.") {
			updatedUser = true
			if patcher.User == nil {
				patchee.User = nil
				continue
			}
			if patchee.User == nil {
				patchee.User = &User{}
			}
			if o, err := DefaultApplyFieldMaskUser(ctx, patchee.User, patcher.User, &field_mask.FieldMask{Paths: updateMask.Paths[i:]}, prefix+"User.", db); err != nil {
				return nil, err
			} else {
				patchee.User = o
			}
			continue
		}
		if f == prefix+"User" {
			updatedUser = true
			patchee.User = patcher.User
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListUser executes a gorm list call
func DefaultListUser(ctx context.Context, db *gorm.DB) ([]*User, error) {
	in := User{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &UserORM{}, &User{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []UserORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*User{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type UserORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]UserORM) error
}