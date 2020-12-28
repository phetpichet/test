// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Piichet/app/ent/office"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// OfficeCreate is the builder for creating a Office entity.
type OfficeCreate struct {
	config
	mutation *OfficeMutation
	hooks    []Hook
}

// SetFname sets the fname field.
func (oc *OfficeCreate) SetFname(s string) *OfficeCreate {
	oc.mutation.SetFname(s)
	return oc
}

// Mutation returns the OfficeMutation object of the builder.
func (oc *OfficeCreate) Mutation() *OfficeMutation {
	return oc.mutation
}

// Save creates the Office in the database.
func (oc *OfficeCreate) Save(ctx context.Context) (*Office, error) {
	if _, ok := oc.mutation.Fname(); !ok {
		return nil, &ValidationError{Name: "fname", err: errors.New("ent: missing required field \"fname\"")}
	}
	if v, ok := oc.mutation.Fname(); ok {
		if err := office.FnameValidator(v); err != nil {
			return nil, &ValidationError{Name: "fname", err: fmt.Errorf("ent: validator failed for field \"fname\": %w", err)}
		}
	}
	var (
		err  error
		node *Office
	)
	if len(oc.hooks) == 0 {
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OfficeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oc.mutation = mutation
			node, err = oc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OfficeCreate) SaveX(ctx context.Context) *Office {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oc *OfficeCreate) sqlSave(ctx context.Context) (*Office, error) {
	o, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	o.ID = int(id)
	return o, nil
}

func (oc *OfficeCreate) createSpec() (*Office, *sqlgraph.CreateSpec) {
	var (
		o     = &Office{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: office.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: office.FieldID,
			},
		}
	)
	if value, ok := oc.mutation.Fname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: office.FieldFname,
		})
		o.Fname = value
	}
	return o, _spec
}
