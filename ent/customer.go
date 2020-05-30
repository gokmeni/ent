// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/gokmeni/ent/ent/customer"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CompanyName holds the value of the "companyName" field.
	CompanyName string `json:"companyName,omitempty"`
	// ContactName holds the value of the "contactName" field.
	ContactName string `json:"contactName,omitempty"`
	// ContactTitle holds the value of the "contactTitle" field.
	ContactTitle string `json:"contactTitle,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// Region holds the value of the "region" field.
	Region string `json:"region,omitempty"`
	// PostalCode holds the value of the "postalCode" field.
	PostalCode string `json:"postalCode,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Fax holds the value of the "fax" field.
	Fax string `json:"fax,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues() []interface{} {
	return []interface{}{
		&sql.NullString{}, // id
		&sql.NullString{}, // companyName
		&sql.NullString{}, // contactName
		&sql.NullString{}, // contactTitle
		&sql.NullString{}, // address
		&sql.NullString{}, // city
		&sql.NullString{}, // region
		&sql.NullString{}, // postalCode
		&sql.NullString{}, // country
		&sql.NullString{}, // phone
		&sql.NullString{}, // fax
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(values ...interface{}) error {
	if m, n := len(values), len(customer.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value.Valid {
		c.ID = value.String
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field companyName", values[0])
	} else if value.Valid {
		c.CompanyName = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field contactName", values[1])
	} else if value.Valid {
		c.ContactName = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field contactTitle", values[2])
	} else if value.Valid {
		c.ContactTitle = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field address", values[3])
	} else if value.Valid {
		c.Address = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field city", values[4])
	} else if value.Valid {
		c.City = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field region", values[5])
	} else if value.Valid {
		c.Region = value.String
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field postalCode", values[6])
	} else if value.Valid {
		c.PostalCode = value.String
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field country", values[7])
	} else if value.Valid {
		c.Country = value.String
	}
	if value, ok := values[8].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field phone", values[8])
	} else if value.Valid {
		c.Phone = value.String
	}
	if value, ok := values[9].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field fax", values[9])
	} else if value.Valid {
		c.Fax = value.String
	}
	return nil
}

// Update returns a builder for updating this Customer.
// Note that, you need to call Customer.Unwrap() before calling this method, if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return (&CustomerClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", companyName=")
	builder.WriteString(c.CompanyName)
	builder.WriteString(", contactName=")
	builder.WriteString(c.ContactName)
	builder.WriteString(", contactTitle=")
	builder.WriteString(c.ContactTitle)
	builder.WriteString(", address=")
	builder.WriteString(c.Address)
	builder.WriteString(", city=")
	builder.WriteString(c.City)
	builder.WriteString(", region=")
	builder.WriteString(c.Region)
	builder.WriteString(", postalCode=")
	builder.WriteString(c.PostalCode)
	builder.WriteString(", country=")
	builder.WriteString(c.Country)
	builder.WriteString(", phone=")
	builder.WriteString(c.Phone)
	builder.WriteString(", fax=")
	builder.WriteString(c.Fax)
	builder.WriteByte(')')
	return builder.String()
}

// Customers is a parsable slice of Customer.
type Customers []*Customer

func (c Customers) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
