package schema

import (
		"github.com/facebookincubator/ent"
		"github.com/facebookincubator/ent/schema/field"
)	

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("dname").NotEmpty(),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return nil
}
