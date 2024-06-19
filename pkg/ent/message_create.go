// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"test10/pkg/ent/message"
	"test10/pkg/ent/room"
	"test10/pkg/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	mutation *MessageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (mc *MessageCreate) SetCreatedAt(t time.Time) *MessageCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableCreatedAt(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MessageCreate) SetUpdatedAt(t time.Time) *MessageCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableUpdatedAt(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetContent sets the "content" field.
func (mc *MessageCreate) SetContent(s string) *MessageCreate {
	mc.mutation.SetContent(s)
	return mc
}

// SetUserID sets the "user_id" field.
func (mc *MessageCreate) SetUserID(i int64) *MessageCreate {
	mc.mutation.SetUserID(i)
	return mc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (mc *MessageCreate) SetNillableUserID(i *int64) *MessageCreate {
	if i != nil {
		mc.SetUserID(*i)
	}
	return mc
}

// SetRoomID sets the "room_id" field.
func (mc *MessageCreate) SetRoomID(i int64) *MessageCreate {
	mc.mutation.SetRoomID(i)
	return mc
}

// SetNillableRoomID sets the "room_id" field if the given value is not nil.
func (mc *MessageCreate) SetNillableRoomID(i *int64) *MessageCreate {
	if i != nil {
		mc.SetRoomID(*i)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MessageCreate) SetID(i int64) *MessageCreate {
	mc.mutation.SetID(i)
	return mc
}

// SetUser sets the "user" edge to the User entity.
func (mc *MessageCreate) SetUser(u *User) *MessageCreate {
	return mc.SetUserID(u.ID)
}

// SetRoom sets the "room" edge to the Room entity.
func (mc *MessageCreate) SetRoom(r *Room) *MessageCreate {
	return mc.SetRoomID(r.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mc *MessageCreate) Mutation() *MessageMutation {
	return mc.mutation
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MessageCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MessageCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MessageCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := message.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := message.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessageCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Message.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Message.updated_at"`)}
	}
	if _, ok := mc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Message.content"`)}
	}
	if v, ok := mc.mutation.Content(); ok {
		if err := message.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Message.content": %w`, err)}
		}
	}
	return nil
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MessageCreate) createSpec() (*Message, *sqlgraph.CreateSpec) {
	var (
		_node = &Message{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(message.Table, sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(message.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(message.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.Content(); ok {
		_spec.SetField(message.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if nodes := mc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.UserTable,
			Columns: []string{message.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.RoomTable,
			Columns: []string{message.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RoomID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mc *MessageCreate) OnConflict(opts ...sql.ConflictOption) *MessageUpsertOne {
	mc.conflict = opts
	return &MessageUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MessageCreate) OnConflictColumns(columns ...string) *MessageUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertOne{
		create: mc,
	}
}

type (
	// MessageUpsertOne is the builder for "upsert"-ing
	//  one Message node.
	MessageUpsertOne struct {
		create *MessageCreate
	}

	// MessageUpsert is the "OnConflict" setter.
	MessageUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *MessageUpsert) SetUpdatedAt(v time.Time) *MessageUpsert {
	u.Set(message.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MessageUpsert) UpdateUpdatedAt() *MessageUpsert {
	u.SetExcluded(message.FieldUpdatedAt)
	return u
}

// SetContent sets the "content" field.
func (u *MessageUpsert) SetContent(v string) *MessageUpsert {
	u.Set(message.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *MessageUpsert) UpdateContent() *MessageUpsert {
	u.SetExcluded(message.FieldContent)
	return u
}

// SetUserID sets the "user_id" field.
func (u *MessageUpsert) SetUserID(v int64) *MessageUpsert {
	u.Set(message.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *MessageUpsert) UpdateUserID() *MessageUpsert {
	u.SetExcluded(message.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *MessageUpsert) ClearUserID() *MessageUpsert {
	u.SetNull(message.FieldUserID)
	return u
}

// SetRoomID sets the "room_id" field.
func (u *MessageUpsert) SetRoomID(v int64) *MessageUpsert {
	u.Set(message.FieldRoomID, v)
	return u
}

// UpdateRoomID sets the "room_id" field to the value that was provided on create.
func (u *MessageUpsert) UpdateRoomID() *MessageUpsert {
	u.SetExcluded(message.FieldRoomID)
	return u
}

// ClearRoomID clears the value of the "room_id" field.
func (u *MessageUpsert) ClearRoomID() *MessageUpsert {
	u.SetNull(message.FieldRoomID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(message.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MessageUpsertOne) UpdateNewValues() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(message.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(message.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MessageUpsertOne) Ignore() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertOne) DoNothing() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreate.OnConflict
// documentation for more info.
func (u *MessageUpsertOne) Update(set func(*MessageUpsert)) *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MessageUpsertOne) SetUpdatedAt(v time.Time) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateUpdatedAt() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetContent sets the "content" field.
func (u *MessageUpsertOne) SetContent(v string) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateContent() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateContent()
	})
}

// SetUserID sets the "user_id" field.
func (u *MessageUpsertOne) SetUserID(v int64) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateUserID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *MessageUpsertOne) ClearUserID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.ClearUserID()
	})
}

// SetRoomID sets the "room_id" field.
func (u *MessageUpsertOne) SetRoomID(v int64) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetRoomID(v)
	})
}

// UpdateRoomID sets the "room_id" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateRoomID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateRoomID()
	})
}

// ClearRoomID clears the value of the "room_id" field.
func (u *MessageUpsertOne) ClearRoomID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.ClearRoomID()
	})
}

// Exec executes the query.
func (u *MessageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MessageUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MessageUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MessageCreateBulk is the builder for creating many Message entities in bulk.
type MessageCreateBulk struct {
	config
	err      error
	builders []*MessageCreate
	conflict []sql.ConflictOption
}

// Save creates the Message entities in the database.
func (mcb *MessageCreateBulk) Save(ctx context.Context) ([]*Message, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Message, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessageCreateBulk) SaveX(ctx context.Context) []*Message {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MessageCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MessageCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mcb *MessageCreateBulk) OnConflict(opts ...sql.ConflictOption) *MessageUpsertBulk {
	mcb.conflict = opts
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MessageCreateBulk) OnConflictColumns(columns ...string) *MessageUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// MessageUpsertBulk is the builder for "upsert"-ing
// a bulk of Message nodes.
type MessageUpsertBulk struct {
	create *MessageCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(message.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MessageUpsertBulk) UpdateNewValues() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(message.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(message.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MessageUpsertBulk) Ignore() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertBulk) DoNothing() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreateBulk.OnConflict
// documentation for more info.
func (u *MessageUpsertBulk) Update(set func(*MessageUpsert)) *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MessageUpsertBulk) SetUpdatedAt(v time.Time) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateUpdatedAt() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetContent sets the "content" field.
func (u *MessageUpsertBulk) SetContent(v string) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateContent() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateContent()
	})
}

// SetUserID sets the "user_id" field.
func (u *MessageUpsertBulk) SetUserID(v int64) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateUserID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *MessageUpsertBulk) ClearUserID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.ClearUserID()
	})
}

// SetRoomID sets the "room_id" field.
func (u *MessageUpsertBulk) SetRoomID(v int64) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetRoomID(v)
	})
}

// UpdateRoomID sets the "room_id" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateRoomID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateRoomID()
	})
}

// ClearRoomID clears the value of the "room_id" field.
func (u *MessageUpsertBulk) ClearRoomID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.ClearRoomID()
	})
}

// Exec executes the query.
func (u *MessageUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MessageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}