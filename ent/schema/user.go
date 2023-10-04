package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).
			Unique().
			Annotations(entproto.Field(1)),
		field.String("uid").
			Unique().
			Annotations(entproto.Field(2)),
		field.String("email").
			Optional().Unique().
			Annotations(entproto.Field(3)),
		field.String("password").
			Annotations(entproto.Field(4)),
		field.Enum("role").
			Values("member", "admin").
			Default("member").
			Annotations(entproto.Field(5), entproto.Enum(map[string]int32{
				"member": 0,
				"admin":  1,
			})),
		field.Time("created_at").
			Immutable().
			Annotations(entproto.Field(6)),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(entproto.Field(7)),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uid").
			Unique(),
		index.Fields("email").
			Unique(),
		index.Fields("role").
			Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
