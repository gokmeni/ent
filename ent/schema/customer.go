package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

type Customer struct {
	ent.Schema
}

func (Customer) Config() ent.Config {
	return ent.Config{
		Table: "customers",
	}
}

func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("customer_id").Unique().Immutable(),
		field.String("companyName").StorageKey("company_name"),
		field.String("contactName").StorageKey("contact_name"),
		field.String("contactTitle").StorageKey("contact_title"),
		field.String("address").StorageKey("address"),
		field.String("city").StorageKey("city"),
		field.String("region").StorageKey("region"),
		field.String("postalCode").StorageKey("postal_code"),
		field.String("country").StorageKey("country"),
		field.String("phone").StorageKey("phone"),
		field.String("fax").StorageKey("fax"),
	}
}