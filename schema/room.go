package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Room struct {
	ent.Schema
}

func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").NotEmpty(),
	}
}

func (Room) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
	}
}
