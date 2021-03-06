// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/Piichet/app/ent/predicate"
	"github.com/Piichet/app/ent/working"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// WorkingQuery is the builder for querying Working entities.
type WorkingQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Working
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (wq *WorkingQuery) Where(ps ...predicate.Working) *WorkingQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit adds a limit step to the query.
func (wq *WorkingQuery) Limit(limit int) *WorkingQuery {
	wq.limit = &limit
	return wq
}

// Offset adds an offset step to the query.
func (wq *WorkingQuery) Offset(offset int) *WorkingQuery {
	wq.offset = &offset
	return wq
}

// Order adds an order step to the query.
func (wq *WorkingQuery) Order(o ...OrderFunc) *WorkingQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// First returns the first Working entity in the query. Returns *NotFoundError when no working was found.
func (wq *WorkingQuery) First(ctx context.Context) (*Working, error) {
	ws, err := wq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ws) == 0 {
		return nil, &NotFoundError{working.Label}
	}
	return ws[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WorkingQuery) FirstX(ctx context.Context) *Working {
	w, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return w
}

// FirstID returns the first Working id in the query. Returns *NotFoundError when no id was found.
func (wq *WorkingQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{working.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (wq *WorkingQuery) FirstXID(ctx context.Context) int {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Working entity in the query, returns an error if not exactly one entity was returned.
func (wq *WorkingQuery) Only(ctx context.Context) (*Working, error) {
	ws, err := wq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ws) {
	case 1:
		return ws[0], nil
	case 0:
		return nil, &NotFoundError{working.Label}
	default:
		return nil, &NotSingularError{working.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WorkingQuery) OnlyX(ctx context.Context) *Working {
	w, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return w
}

// OnlyID returns the only Working id in the query, returns an error if not exactly one id was returned.
func (wq *WorkingQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{working.Label}
	default:
		err = &NotSingularError{working.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WorkingQuery) OnlyIDX(ctx context.Context) int {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Workings.
func (wq *WorkingQuery) All(ctx context.Context) ([]*Working, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wq *WorkingQuery) AllX(ctx context.Context) []*Working {
	ws, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ws
}

// IDs executes the query and returns a list of Working ids.
func (wq *WorkingQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wq.Select(working.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WorkingQuery) IDsX(ctx context.Context) []int {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WorkingQuery) Count(ctx context.Context) (int, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WorkingQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WorkingQuery) Exist(ctx context.Context) (bool, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WorkingQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WorkingQuery) Clone() *WorkingQuery {
	return &WorkingQuery{
		config:     wq.config,
		limit:      wq.limit,
		offset:     wq.offset,
		order:      append([]OrderFunc{}, wq.order...),
		unique:     append([]string{}, wq.unique...),
		predicates: append([]predicate.Working{}, wq.predicates...),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AddedTime time.Time `json:"added_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Working.Query().
//		GroupBy(working.FieldAddedTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wq *WorkingQuery) GroupBy(field string, fields ...string) *WorkingGroupBy {
	group := &WorkingGroupBy{config: wq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		AddedTime time.Time `json:"added_time,omitempty"`
//	}
//
//	client.Working.Query().
//		Select(working.FieldAddedTime).
//		Scan(ctx, &v)
//
func (wq *WorkingQuery) Select(field string, fields ...string) *WorkingSelect {
	selector := &WorkingSelect{config: wq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wq.sqlQuery(), nil
	}
	return selector
}

func (wq *WorkingQuery) prepareQuery(ctx context.Context) error {
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WorkingQuery) sqlAll(ctx context.Context) ([]*Working, error) {
	var (
		nodes = []*Working{}
		_spec = wq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &Working{config: wq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wq *WorkingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WorkingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (wq *WorkingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   working.Table,
			Columns: working.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: working.FieldID,
			},
		},
		From:   wq.sql,
		Unique: true,
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WorkingQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(working.Table)
	selector := builder.Select(t1.Columns(working.Columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(working.Columns...)...)
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WorkingGroupBy is the builder for group-by Working entities.
type WorkingGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WorkingGroupBy) Aggregate(fns ...AggregateFunc) *WorkingGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the group-by query and scan the result into the given value.
func (wgb *WorkingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wgb.path(ctx)
	if err != nil {
		return err
	}
	wgb.sql = query
	return wgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wgb *WorkingGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (wgb *WorkingGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WorkingGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wgb *WorkingGroupBy) StringsX(ctx context.Context) []string {
	v, err := wgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (wgb *WorkingGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{working.Label}
	default:
		err = fmt.Errorf("ent: WorkingGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wgb *WorkingGroupBy) StringX(ctx context.Context) string {
	v, err := wgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (wgb *WorkingGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WorkingGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wgb *WorkingGroupBy) IntsX(ctx context.Context) []int {
	v, err := wgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (wgb *WorkingGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{working.Label}
	default:
		err = fmt.Errorf("ent: WorkingGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wgb *WorkingGroupBy) IntX(ctx context.Context) int {
	v, err := wgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (wgb *WorkingGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WorkingGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wgb *WorkingGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (wgb *WorkingGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{working.Label}
	default:
		err = fmt.Errorf("ent: WorkingGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wgb *WorkingGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (wgb *WorkingGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WorkingGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wgb *WorkingGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (wgb *WorkingGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{working.Label}
	default:
		err = fmt.Errorf("ent: WorkingGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wgb *WorkingGroupBy) BoolX(ctx context.Context) bool {
	v, err := wgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wgb *WorkingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wgb.sqlQuery().Query()
	if err := wgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wgb *WorkingGroupBy) sqlQuery() *sql.Selector {
	selector := wgb.sql
	columns := make([]string, 0, len(wgb.fields)+len(wgb.fns))
	columns = append(columns, wgb.fields...)
	for _, fn := range wgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(wgb.fields...)
}

// WorkingSelect is the builder for select fields of Working entities.
type WorkingSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ws *WorkingSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ws.path(ctx)
	if err != nil {
		return err
	}
	ws.sql = query
	return ws.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ws *WorkingSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ws.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ws *WorkingSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WorkingSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ws *WorkingSelect) StringsX(ctx context.Context) []string {
	v, err := ws.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ws *WorkingSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ws.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{working.Label}
	default:
		err = fmt.Errorf("ent: WorkingSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ws *WorkingSelect) StringX(ctx context.Context) string {
	v, err := ws.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ws *WorkingSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WorkingSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ws *WorkingSelect) IntsX(ctx context.Context) []int {
	v, err := ws.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ws *WorkingSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ws.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{working.Label}
	default:
		err = fmt.Errorf("ent: WorkingSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ws *WorkingSelect) IntX(ctx context.Context) int {
	v, err := ws.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ws *WorkingSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WorkingSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ws *WorkingSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ws.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ws *WorkingSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ws.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{working.Label}
	default:
		err = fmt.Errorf("ent: WorkingSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ws *WorkingSelect) Float64X(ctx context.Context) float64 {
	v, err := ws.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ws *WorkingSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WorkingSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ws *WorkingSelect) BoolsX(ctx context.Context) []bool {
	v, err := ws.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ws *WorkingSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ws.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{working.Label}
	default:
		err = fmt.Errorf("ent: WorkingSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ws *WorkingSelect) BoolX(ctx context.Context) bool {
	v, err := ws.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ws *WorkingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ws.sqlQuery().Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ws *WorkingSelect) sqlQuery() sql.Querier {
	selector := ws.sql
	selector.Select(selector.Columns(ws.fields...)...)
	return selector
}
