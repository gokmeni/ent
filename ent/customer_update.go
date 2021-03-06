// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/gokmeni/ent/ent/customer"
	"github.com/gokmeni/ent/ent/predicate"
)

// CustomerUpdate is the builder for updating Customer entities.
type CustomerUpdate struct {
	config
	hooks      []Hook
	mutation   *CustomerMutation
	predicates []predicate.Customer
}

// Where adds a new predicate for the builder.
func (cu *CustomerUpdate) Where(ps ...predicate.Customer) *CustomerUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetCompanyName sets the companyName field.
func (cu *CustomerUpdate) SetCompanyName(s string) *CustomerUpdate {
	cu.mutation.SetCompanyName(s)
	return cu
}

// SetContactName sets the contactName field.
func (cu *CustomerUpdate) SetContactName(s string) *CustomerUpdate {
	cu.mutation.SetContactName(s)
	return cu
}

// SetContactTitle sets the contactTitle field.
func (cu *CustomerUpdate) SetContactTitle(s string) *CustomerUpdate {
	cu.mutation.SetContactTitle(s)
	return cu
}

// SetAddress sets the address field.
func (cu *CustomerUpdate) SetAddress(s string) *CustomerUpdate {
	cu.mutation.SetAddress(s)
	return cu
}

// SetCity sets the city field.
func (cu *CustomerUpdate) SetCity(s string) *CustomerUpdate {
	cu.mutation.SetCity(s)
	return cu
}

// SetRegion sets the region field.
func (cu *CustomerUpdate) SetRegion(s string) *CustomerUpdate {
	cu.mutation.SetRegion(s)
	return cu
}

// SetPostalCode sets the postalCode field.
func (cu *CustomerUpdate) SetPostalCode(s string) *CustomerUpdate {
	cu.mutation.SetPostalCode(s)
	return cu
}

// SetCountry sets the country field.
func (cu *CustomerUpdate) SetCountry(s string) *CustomerUpdate {
	cu.mutation.SetCountry(s)
	return cu
}

// SetPhone sets the phone field.
func (cu *CustomerUpdate) SetPhone(s string) *CustomerUpdate {
	cu.mutation.SetPhone(s)
	return cu
}

// SetFax sets the fax field.
func (cu *CustomerUpdate) SetFax(s string) *CustomerUpdate {
	cu.mutation.SetFax(s)
	return cu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CustomerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CustomerUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CustomerUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: customer.FieldID,
			},
		},
	}
	if ps := cu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CompanyName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldCompanyName,
		})
	}
	if value, ok := cu.mutation.ContactName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldContactName,
		})
	}
	if value, ok := cu.mutation.ContactTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldContactTitle,
		})
	}
	if value, ok := cu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAddress,
		})
	}
	if value, ok := cu.mutation.City(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldCity,
		})
	}
	if value, ok := cu.mutation.Region(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldRegion,
		})
	}
	if value, ok := cu.mutation.PostalCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPostalCode,
		})
	}
	if value, ok := cu.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldCountry,
		})
	}
	if value, ok := cu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPhone,
		})
	}
	if value, ok := cu.mutation.Fax(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldFax,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CustomerUpdateOne is the builder for updating a single Customer entity.
type CustomerUpdateOne struct {
	config
	hooks    []Hook
	mutation *CustomerMutation
}

// SetCompanyName sets the companyName field.
func (cuo *CustomerUpdateOne) SetCompanyName(s string) *CustomerUpdateOne {
	cuo.mutation.SetCompanyName(s)
	return cuo
}

// SetContactName sets the contactName field.
func (cuo *CustomerUpdateOne) SetContactName(s string) *CustomerUpdateOne {
	cuo.mutation.SetContactName(s)
	return cuo
}

// SetContactTitle sets the contactTitle field.
func (cuo *CustomerUpdateOne) SetContactTitle(s string) *CustomerUpdateOne {
	cuo.mutation.SetContactTitle(s)
	return cuo
}

// SetAddress sets the address field.
func (cuo *CustomerUpdateOne) SetAddress(s string) *CustomerUpdateOne {
	cuo.mutation.SetAddress(s)
	return cuo
}

// SetCity sets the city field.
func (cuo *CustomerUpdateOne) SetCity(s string) *CustomerUpdateOne {
	cuo.mutation.SetCity(s)
	return cuo
}

// SetRegion sets the region field.
func (cuo *CustomerUpdateOne) SetRegion(s string) *CustomerUpdateOne {
	cuo.mutation.SetRegion(s)
	return cuo
}

// SetPostalCode sets the postalCode field.
func (cuo *CustomerUpdateOne) SetPostalCode(s string) *CustomerUpdateOne {
	cuo.mutation.SetPostalCode(s)
	return cuo
}

// SetCountry sets the country field.
func (cuo *CustomerUpdateOne) SetCountry(s string) *CustomerUpdateOne {
	cuo.mutation.SetCountry(s)
	return cuo
}

// SetPhone sets the phone field.
func (cuo *CustomerUpdateOne) SetPhone(s string) *CustomerUpdateOne {
	cuo.mutation.SetPhone(s)
	return cuo
}

// SetFax sets the fax field.
func (cuo *CustomerUpdateOne) SetFax(s string) *CustomerUpdateOne {
	cuo.mutation.SetFax(s)
	return cuo
}

// Save executes the query and returns the updated entity.
func (cuo *CustomerUpdateOne) Save(ctx context.Context) (*Customer, error) {
	var (
		err  error
		node *Customer
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CustomerUpdateOne) SaveX(ctx context.Context) *Customer {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CustomerUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CustomerUpdateOne) sqlSave(ctx context.Context) (c *Customer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: customer.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Customer.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.CompanyName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldCompanyName,
		})
	}
	if value, ok := cuo.mutation.ContactName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldContactName,
		})
	}
	if value, ok := cuo.mutation.ContactTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldContactTitle,
		})
	}
	if value, ok := cuo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAddress,
		})
	}
	if value, ok := cuo.mutation.City(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldCity,
		})
	}
	if value, ok := cuo.mutation.Region(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldRegion,
		})
	}
	if value, ok := cuo.mutation.PostalCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPostalCode,
		})
	}
	if value, ok := cuo.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldCountry,
		})
	}
	if value, ok := cuo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPhone,
		})
	}
	if value, ok := cuo.mutation.Fax(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldFax,
		})
	}
	c = &Customer{config: cuo.config}
	_spec.Assign = c.assignValues
	_spec.ScanValues = c.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return c, nil
}
