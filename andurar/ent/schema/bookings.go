package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Bookings holds the schema definition for the Bookings entity.
type Bookings struct {
	ent.Schema
}

// // Time mixin.
// func (Bookings) Mixin() []ent.Mixin {
// 	return []ent.Mixin{
// 		TimeMixin{},
// 	}
// }

// Fields of the Bookings.
func (Bookings) Fields() []ent.Field {
	return []ent.Field{
		field.String("full_name").Optional(),
		field.String("email").Optional(),
		field.String("phone").Optional(),
		field.String("room").Optional(),
		field.String("no_of_guests").Optional(),
		field.String("check_in").Optional(),
		field.String("check_out").Optional(),
		field.Enum("pickup").Values("yes", "no").Optional(),
		field.Enum("status").Values("confirmed", "pending").Default("pending"),
		field.Enum("payment_status").Values("paid", "pending").Default("pending"),
		field.String("special_requests").Optional(),
	}
}

// Edges of the Bookings.
func (Bookings) Edges() []ent.Edge {
	return nil
}
