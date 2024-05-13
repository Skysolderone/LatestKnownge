package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type A struct {
	ent.Schema
}

func (A) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (A) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bs", B.Type),
	}
}
