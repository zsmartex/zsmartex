package schema

import (
	"encoding/json"
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Code holds the schema definition for the Code entity.
type Code struct {
	ent.Schema
}

// Fields of the Code.
func (Code) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).
			Unique().
			Annotations(entproto.Field(1)),
		field.String("user_id").
			Annotations(entproto.Field(2)),
		field.String("code").Annotations(entproto.Field(3)),
		field.Enum("type").
			Values("unknown", "email", "phone").
			Default("unknown").
			Annotations(entproto.Field(4), entproto.Enum(map[string]int32{
				"unknown": 0,
				"email":   1,
				"phone":   2,
			})),
		field.Enum("category").
			Values("unknown", "register", "reset_password", "login").
			Default("unknown").
			Annotations(entproto.Field(5), entproto.Enum(map[string]int32{
				"unknown":        0,
				"register":       1,
				"reset_password": 2,
				"login":          3,
			})),
		field.Int64("email_index").
			Optional().
			Annotations(entproto.Field(6)),
		field.String("email_encrypted").
			Optional().
			Annotations(entproto.Field(7)),
		field.Int64("phone_index").
			Optional().
			Annotations(entproto.Field(8)),
		field.String("phone_encrypted").
			Optional().
			Annotations(entproto.Field(9)),
		field.JSON("data", json.RawMessage{}).
			Optional().
			Annotations(entproto.Field(10, entproto.Type(descriptorpb.FieldDescriptorProto_TYPE_BYTES))),
		field.Time("created_at").
			Immutable().
			Annotations(entproto.Field(11)),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(entproto.Field(12)),
	}
}

// Edges of the Code.
func (Code) Edges() []ent.Edge {
	return nil
}

func (Code) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
