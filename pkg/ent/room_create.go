// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"test10/pkg/ent/message"
	"test10/pkg/ent/room"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoomCreate is the builder for creating a Room entity.
type RoomCreate struct {
	config
	mutation *RoomMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (rc *RoomCreate) SetCreatedAt(t time.Time) *RoomCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RoomCreate) SetNillableCreatedAt(t *time.Time) *RoomCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RoomCreate) SetUpdatedAt(t time.Time) *RoomCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RoomCreate) SetNillableUpdatedAt(t *time.Time) *RoomCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetName sets the "name" field.
func (rc *RoomCreate) SetName(s string) *RoomCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetID sets the "id" field.
func (rc *RoomCreate) SetID(i int64) *RoomCreate {
	rc.mutation.SetID(i)
	return rc
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (rc *RoomCreate) AddMessageIDs(ids ...int64) *RoomCreate {
	rc.mutation.AddMessageIDs(ids...)
	return rc
}

// AddMessages adds the "messages" edges to the Message entity.
func (rc *RoomCreate) AddMessages(m ...*Message) *RoomCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return rc.AddMessageIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (rc *RoomCreate) Mutation() *RoomMutation {
	return rc.mutation
}

// Save creates the Room in the database.
func (rc *RoomCreate) Save(ctx context.Context) (*Room, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoomCreate) SaveX(ctx context.Context) *Room {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoomCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoomCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoomCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := room.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := room.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoomCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Room.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Room.updated_at"`)}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Room.name"`)}
	}
	if v, ok := rc.mutation.Name(); ok {
		if err := room.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Room.name": %w`, err)}
		}
	}
	return nil
}

func (rc *RoomCreate) sqlSave(ctx context.Context) (*Room, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoomCreate) createSpec() (*Room, *sqlgraph.CreateSpec) {
	var (
		_node = &Room{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(room.Table, sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = rc.conflict
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(room.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(room.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(room.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := rc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.MessagesTable,
			Columns: []string{room.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Room.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoomUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (rc *RoomCreate) OnConflict(opts ...sql.ConflictOption) *RoomUpsertOne {
	rc.conflict = opts
	return &RoomUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Room.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RoomCreate) OnConflictColumns(columns ...string) *RoomUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RoomUpsertOne{
		create: rc,
	}
}

type (
	// RoomUpsertOne is the builder for "upsert"-ing
	//  one Room node.
	RoomUpsertOne struct {
		create *RoomCreate
	}

	// RoomUpsert is the "OnConflict" setter.
	RoomUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *RoomUpsert) SetUpdatedAt(v time.Time) *RoomUpsert {
	u.Set(room.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RoomUpsert) UpdateUpdatedAt() *RoomUpsert {
	u.SetExcluded(room.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *RoomUpsert) SetName(v string) *RoomUpsert {
	u.Set(room.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RoomUpsert) UpdateName() *RoomUpsert {
	u.SetExcluded(room.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Room.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(room.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RoomUpsertOne) UpdateNewValues() *RoomUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(room.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(room.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Room.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RoomUpsertOne) Ignore() *RoomUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoomUpsertOne) DoNothing() *RoomUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoomCreate.OnConflict
// documentation for more info.
func (u *RoomUpsertOne) Update(set func(*RoomUpsert)) *RoomUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoomUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RoomUpsertOne) SetUpdatedAt(v time.Time) *RoomUpsertOne {
	return u.Update(func(s *RoomUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RoomUpsertOne) UpdateUpdatedAt() *RoomUpsertOne {
	return u.Update(func(s *RoomUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *RoomUpsertOne) SetName(v string) *RoomUpsertOne {
	return u.Update(func(s *RoomUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RoomUpsertOne) UpdateName() *RoomUpsertOne {
	return u.Update(func(s *RoomUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *RoomUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RoomCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoomUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RoomUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RoomUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RoomCreateBulk is the builder for creating many Room entities in bulk.
type RoomCreateBulk struct {
	config
	err      error
	builders []*RoomCreate
	conflict []sql.ConflictOption
}

// Save creates the Room entities in the database.
func (rcb *RoomCreateBulk) Save(ctx context.Context) ([]*Room, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Room, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoomMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoomCreateBulk) SaveX(ctx context.Context) []*Room {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoomCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoomCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Room.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoomUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (rcb *RoomCreateBulk) OnConflict(opts ...sql.ConflictOption) *RoomUpsertBulk {
	rcb.conflict = opts
	return &RoomUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Room.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RoomCreateBulk) OnConflictColumns(columns ...string) *RoomUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RoomUpsertBulk{
		create: rcb,
	}
}

// RoomUpsertBulk is the builder for "upsert"-ing
// a bulk of Room nodes.
type RoomUpsertBulk struct {
	create *RoomCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Room.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(room.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RoomUpsertBulk) UpdateNewValues() *RoomUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(room.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(room.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Room.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RoomUpsertBulk) Ignore() *RoomUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoomUpsertBulk) DoNothing() *RoomUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoomCreateBulk.OnConflict
// documentation for more info.
func (u *RoomUpsertBulk) Update(set func(*RoomUpsert)) *RoomUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoomUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RoomUpsertBulk) SetUpdatedAt(v time.Time) *RoomUpsertBulk {
	return u.Update(func(s *RoomUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RoomUpsertBulk) UpdateUpdatedAt() *RoomUpsertBulk {
	return u.Update(func(s *RoomUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *RoomUpsertBulk) SetName(v string) *RoomUpsertBulk {
	return u.Update(func(s *RoomUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RoomUpsertBulk) UpdateName() *RoomUpsertBulk {
	return u.Update(func(s *RoomUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *RoomUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RoomCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RoomCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoomUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
