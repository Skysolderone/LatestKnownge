package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type B struct {
	ent.Schema
}

func (B) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (B) Edges() []ent.Edge {
	return nil
}
