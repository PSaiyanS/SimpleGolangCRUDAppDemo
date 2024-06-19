package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty(),
		field.String("password").Unique().NotEmpty(),
		field.String("email").NotEmpty(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
	}
}
