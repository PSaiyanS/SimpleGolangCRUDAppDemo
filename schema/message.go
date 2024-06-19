package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Message struct {
	ent.Schema
}

func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("content").NotEmpty(),
		field.Int64("user_id").Optional(),
		field.Int64("room_id").Optional(),
	}
}

func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("messages").
			Field("user_id").
			Unique(),
		edge.From("room", Room.Type).
			Ref("messages").
			Field("room_id").
			Unique(),
	}
}
