// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"playground/ronen-bootcamp/ent/audit"
	"playground/ronen-bootcamp/ent/organization"
	"playground/ronen-bootcamp/ent/predicate"
	"playground/ronen-bootcamp/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuditQuery is the builder for querying Audit entities.
type AuditQuery struct {
	config
	limit             *int
	offset            *int
	unique            *bool
	order             []OrderFunc
	fields            []string
	predicates        []predicate.Audit
	withUser          *UserQuery
	withOrganizations *OrganizationQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuditQuery builder.
func (aq *AuditQuery) Where(ps ...predicate.Audit) *AuditQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AuditQuery) Limit(limit int) *AuditQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AuditQuery) Offset(offset int) *AuditQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AuditQuery) Unique(unique bool) *AuditQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AuditQuery) Order(o ...OrderFunc) *AuditQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryUser chains the current query on the "user" edge.
func (aq *AuditQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(audit.Table, audit.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, audit.UserTable, audit.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganizations chains the current query on the "organizations" edge.
func (aq *AuditQuery) QueryOrganizations() *OrganizationQuery {
	query := &OrganizationQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(audit.Table, audit.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, audit.OrganizationsTable, audit.OrganizationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Audit entity from the query.
// Returns a *NotFoundError when no Audit was found.
func (aq *AuditQuery) First(ctx context.Context) (*Audit, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{audit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AuditQuery) FirstX(ctx context.Context) *Audit {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Audit ID from the query.
// Returns a *NotFoundError when no Audit ID was found.
func (aq *AuditQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{audit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AuditQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Audit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Audit entity is found.
// Returns a *NotFoundError when no Audit entities are found.
func (aq *AuditQuery) Only(ctx context.Context) (*Audit, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{audit.Label}
	default:
		return nil, &NotSingularError{audit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AuditQuery) OnlyX(ctx context.Context) *Audit {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Audit ID in the query.
// Returns a *NotSingularError when more than one Audit ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AuditQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{audit.Label}
	default:
		err = &NotSingularError{audit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AuditQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Audits.
func (aq *AuditQuery) All(ctx context.Context) ([]*Audit, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AuditQuery) AllX(ctx context.Context) []*Audit {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Audit IDs.
func (aq *AuditQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := aq.Select(audit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AuditQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AuditQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AuditQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AuditQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AuditQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuditQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AuditQuery) Clone() *AuditQuery {
	if aq == nil {
		return nil
	}
	return &AuditQuery{
		config:            aq.config,
		limit:             aq.limit,
		offset:            aq.offset,
		order:             append([]OrderFunc{}, aq.order...),
		predicates:        append([]predicate.Audit{}, aq.predicates...),
		withUser:          aq.withUser.Clone(),
		withOrganizations: aq.withOrganizations.Clone(),
		// clone intermediate query.
		sql:    aq.sql.Clone(),
		path:   aq.path,
		unique: aq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AuditQuery) WithUser(opts ...func(*UserQuery)) *AuditQuery {
	query := &UserQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withUser = query
	return aq
}

// WithOrganizations tells the query-builder to eager-load the nodes that are connected to
// the "organizations" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AuditQuery) WithOrganizations(opts ...func(*OrganizationQuery)) *AuditQuery {
	query := &OrganizationQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withOrganizations = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Identity string `json:"identity,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Audit.Query().
//		GroupBy(audit.FieldIdentity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AuditQuery) GroupBy(field string, fields ...string) *AuditGroupBy {
	grbuild := &AuditGroupBy{config: aq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	grbuild.label = audit.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Identity string `json:"identity,omitempty"`
//	}
//
//	client.Audit.Query().
//		Select(audit.FieldIdentity).
//		Scan(ctx, &v)
func (aq *AuditQuery) Select(fields ...string) *AuditSelect {
	aq.fields = append(aq.fields, fields...)
	selbuild := &AuditSelect{AuditQuery: aq}
	selbuild.label = audit.Label
	selbuild.flds, selbuild.scan = &aq.fields, selbuild.Scan
	return selbuild
}

func (aq *AuditQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !audit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	if audit.Policy == nil {
		return errors.New("ent: uninitialized audit.Policy (forgotten import ent/runtime?)")
	}
	if err := audit.Policy.EvalQuery(ctx, aq); err != nil {
		return err
	}
	return nil
}

func (aq *AuditQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Audit, error) {
	var (
		nodes       = []*Audit{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withUser != nil,
			aq.withOrganizations != nil,
		}
	)
	if aq.withUser != nil || aq.withOrganizations != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, audit.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Audit).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Audit{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withUser; query != nil {
		if err := aq.loadUser(ctx, query, nodes, nil,
			func(n *Audit, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withOrganizations; query != nil {
		if err := aq.loadOrganizations(ctx, query, nodes, nil,
			func(n *Audit, e *Organization) { n.Edges.Organizations = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AuditQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Audit, init func(*Audit), assign func(*Audit, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Audit)
	for i := range nodes {
		if nodes[i].user_audits == nil {
			continue
		}
		fk := *nodes[i].user_audits
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_audits" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AuditQuery) loadOrganizations(ctx context.Context, query *OrganizationQuery, nodes []*Audit, init func(*Audit), assign func(*Audit, *Organization)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Audit)
	for i := range nodes {
		fk := nodes[i].OrgID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "org_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aq *AuditQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.fields
	if len(aq.fields) > 0 {
		_spec.Unique = aq.unique != nil && *aq.unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AuditQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (aq *AuditQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   audit.Table,
			Columns: audit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: audit.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, audit.FieldID)
		for i := range fields {
			if fields[i] != audit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AuditQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(audit.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = audit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.unique != nil && *aq.unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuditGroupBy is the group-by builder for Audit entities.
type AuditGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AuditGroupBy) Aggregate(fns ...AggregateFunc) *AuditGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AuditGroupBy) Scan(ctx context.Context, v any) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

func (agb *AuditGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range agb.fields {
		if !audit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AuditGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql.Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agb.fields)+len(agb.fns))
		for _, f := range agb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agb.fields...)...)
}

// AuditSelect is the builder for selecting fields of Audit entities.
type AuditSelect struct {
	*AuditQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (as *AuditSelect) Scan(ctx context.Context, v any) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AuditQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

func (as *AuditSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
