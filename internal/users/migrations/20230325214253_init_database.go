package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/mbahjadol/null"
	"gorm.io/gorm"
)

var initDatabase = gormigrate.Migration{
	ID: "20230325214253",
	Migrate: func(db *gorm.DB) error {
		type User struct {
			ID             uint64      `gorm:"primaryKey;autoIncrement;not null"`
			UID            string      `gorm:"type:character varying;not null;uniqueIndex:index_users_on_uid"`
			Username       null.String `gorm:"type:character varying;uniqueIndex:index_users_on_username"`
			PasswordDigest string      `gorm:"type:character varying;not null"`
			Level          int32       `gorm:"type:integer;not null;default:0"`
			OTP            bool        `gorm:"type:boolean;not null;default:false"`
			Role           string      `gorm:"type:character varying;not null;default:member"`
			State          string      `gorm:"type:character varying;not null;default:pending"`
			ReferralUID    null.String `gorm:"type:character varying"`
			Data           []byte      `gorm:"type:jsonb"`
			CreatedAt      time.Time   `gorm:"type:timestamp;not null"`
			UpdatedAt      time.Time   `gorm:"type:timestamp;not null"`
		}

		type UserCredentials struct {
			UserID         uint64    `gorm:"type:bigint;not null;uniqueIndex:idx_users_on_user_id_and_type"`
			ValueIndex     int64     `gorm:"type:bigint;not null;uniqueIndex:index_user_credentials_on_value_index"`
			ValueEncrypted string    `gorm:"type:character varying;not null"`
			Type           string    `gorm:"type:character varying(8);not null;uniqueIndex:idx_users_on_user_id_and_type"`
			CreatedAt      time.Time `gorm:"type:timestamp;not null"`
			UpdatedAt      time.Time `gorm:"type:timestamp;not null"`
			User           *User     `gorm:"constraint:OnDelete:CASCADE"`
		}

		return db.AutoMigrate(
			User{},
			UserCredentials{},
		)
	},
	Rollback: func(db *gorm.DB) error {
		return db.Migrator().DropTable("users", "user_credentials")
	},
}
