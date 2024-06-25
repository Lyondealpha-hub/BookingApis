// Code generated by ent, DO NOT EDIT.

package bookings

import (
	"andurar.api/ent/predicate"
	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Bookings {
	return predicate.Bookings(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Bookings {
	return predicate.Bookings(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Bookings {
	return predicate.Bookings(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Bookings {
	return predicate.Bookings(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Bookings {
	return predicate.Bookings(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Bookings {
	return predicate.Bookings(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Bookings {
	return predicate.Bookings(sql.FieldLTE(FieldID, id))
}

// FullName applies equality check predicate on the "full_name" field. It's identical to FullNameEQ.
func FullName(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldFullName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldEmail, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldPhone, v))
}

// Room applies equality check predicate on the "room" field. It's identical to RoomEQ.
func Room(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldRoom, v))
}

// NoOfGuests applies equality check predicate on the "no_of_guests" field. It's identical to NoOfGuestsEQ.
func NoOfGuests(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldNoOfGuests, v))
}

// CheckIn applies equality check predicate on the "check_in" field. It's identical to CheckInEQ.
func CheckIn(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldCheckIn, v))
}

// CheckOut applies equality check predicate on the "check_out" field. It's identical to CheckOutEQ.
func CheckOut(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldCheckOut, v))
}

// SpecialRequests applies equality check predicate on the "special_requests" field. It's identical to SpecialRequestsEQ.
func SpecialRequests(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldSpecialRequests, v))
}

// FullNameEQ applies the EQ predicate on the "full_name" field.
func FullNameEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldFullName, v))
}

// FullNameNEQ applies the NEQ predicate on the "full_name" field.
func FullNameNEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNEQ(FieldFullName, v))
}

// FullNameIn applies the In predicate on the "full_name" field.
func FullNameIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldIn(FieldFullName, vs...))
}

// FullNameNotIn applies the NotIn predicate on the "full_name" field.
func FullNameNotIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNotIn(FieldFullName, vs...))
}

// FullNameGT applies the GT predicate on the "full_name" field.
func FullNameGT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGT(FieldFullName, v))
}

// FullNameGTE applies the GTE predicate on the "full_name" field.
func FullNameGTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGTE(FieldFullName, v))
}

// FullNameLT applies the LT predicate on the "full_name" field.
func FullNameLT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLT(FieldFullName, v))
}

// FullNameLTE applies the LTE predicate on the "full_name" field.
func FullNameLTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLTE(FieldFullName, v))
}

// FullNameContains applies the Contains predicate on the "full_name" field.
func FullNameContains(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContains(FieldFullName, v))
}

// FullNameHasPrefix applies the HasPrefix predicate on the "full_name" field.
func FullNameHasPrefix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasPrefix(FieldFullName, v))
}

// FullNameHasSuffix applies the HasSuffix predicate on the "full_name" field.
func FullNameHasSuffix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasSuffix(FieldFullName, v))
}

// FullNameIsNil applies the IsNil predicate on the "full_name" field.
func FullNameIsNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldIsNull(FieldFullName))
}

// FullNameNotNil applies the NotNil predicate on the "full_name" field.
func FullNameNotNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldNotNull(FieldFullName))
}

// FullNameEqualFold applies the EqualFold predicate on the "full_name" field.
func FullNameEqualFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEqualFold(FieldFullName, v))
}

// FullNameContainsFold applies the ContainsFold predicate on the "full_name" field.
func FullNameContainsFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContainsFold(FieldFullName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContainsFold(FieldEmail, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneIsNil applies the IsNil predicate on the "phone" field.
func PhoneIsNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldIsNull(FieldPhone))
}

// PhoneNotNil applies the NotNil predicate on the "phone" field.
func PhoneNotNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldNotNull(FieldPhone))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContainsFold(FieldPhone, v))
}

// RoomEQ applies the EQ predicate on the "room" field.
func RoomEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldRoom, v))
}

// RoomNEQ applies the NEQ predicate on the "room" field.
func RoomNEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNEQ(FieldRoom, v))
}

// RoomIn applies the In predicate on the "room" field.
func RoomIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldIn(FieldRoom, vs...))
}

// RoomNotIn applies the NotIn predicate on the "room" field.
func RoomNotIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNotIn(FieldRoom, vs...))
}

// RoomGT applies the GT predicate on the "room" field.
func RoomGT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGT(FieldRoom, v))
}

// RoomGTE applies the GTE predicate on the "room" field.
func RoomGTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGTE(FieldRoom, v))
}

// RoomLT applies the LT predicate on the "room" field.
func RoomLT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLT(FieldRoom, v))
}

// RoomLTE applies the LTE predicate on the "room" field.
func RoomLTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLTE(FieldRoom, v))
}

// RoomContains applies the Contains predicate on the "room" field.
func RoomContains(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContains(FieldRoom, v))
}

// RoomHasPrefix applies the HasPrefix predicate on the "room" field.
func RoomHasPrefix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasPrefix(FieldRoom, v))
}

// RoomHasSuffix applies the HasSuffix predicate on the "room" field.
func RoomHasSuffix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasSuffix(FieldRoom, v))
}

// RoomIsNil applies the IsNil predicate on the "room" field.
func RoomIsNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldIsNull(FieldRoom))
}

// RoomNotNil applies the NotNil predicate on the "room" field.
func RoomNotNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldNotNull(FieldRoom))
}

// RoomEqualFold applies the EqualFold predicate on the "room" field.
func RoomEqualFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEqualFold(FieldRoom, v))
}

// RoomContainsFold applies the ContainsFold predicate on the "room" field.
func RoomContainsFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContainsFold(FieldRoom, v))
}

// NoOfGuestsEQ applies the EQ predicate on the "no_of_guests" field.
func NoOfGuestsEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldNoOfGuests, v))
}

// NoOfGuestsNEQ applies the NEQ predicate on the "no_of_guests" field.
func NoOfGuestsNEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNEQ(FieldNoOfGuests, v))
}

// NoOfGuestsIn applies the In predicate on the "no_of_guests" field.
func NoOfGuestsIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldIn(FieldNoOfGuests, vs...))
}

// NoOfGuestsNotIn applies the NotIn predicate on the "no_of_guests" field.
func NoOfGuestsNotIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNotIn(FieldNoOfGuests, vs...))
}

// NoOfGuestsGT applies the GT predicate on the "no_of_guests" field.
func NoOfGuestsGT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGT(FieldNoOfGuests, v))
}

// NoOfGuestsGTE applies the GTE predicate on the "no_of_guests" field.
func NoOfGuestsGTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGTE(FieldNoOfGuests, v))
}

// NoOfGuestsLT applies the LT predicate on the "no_of_guests" field.
func NoOfGuestsLT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLT(FieldNoOfGuests, v))
}

// NoOfGuestsLTE applies the LTE predicate on the "no_of_guests" field.
func NoOfGuestsLTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLTE(FieldNoOfGuests, v))
}

// NoOfGuestsContains applies the Contains predicate on the "no_of_guests" field.
func NoOfGuestsContains(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContains(FieldNoOfGuests, v))
}

// NoOfGuestsHasPrefix applies the HasPrefix predicate on the "no_of_guests" field.
func NoOfGuestsHasPrefix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasPrefix(FieldNoOfGuests, v))
}

// NoOfGuestsHasSuffix applies the HasSuffix predicate on the "no_of_guests" field.
func NoOfGuestsHasSuffix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasSuffix(FieldNoOfGuests, v))
}

// NoOfGuestsIsNil applies the IsNil predicate on the "no_of_guests" field.
func NoOfGuestsIsNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldIsNull(FieldNoOfGuests))
}

// NoOfGuestsNotNil applies the NotNil predicate on the "no_of_guests" field.
func NoOfGuestsNotNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldNotNull(FieldNoOfGuests))
}

// NoOfGuestsEqualFold applies the EqualFold predicate on the "no_of_guests" field.
func NoOfGuestsEqualFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEqualFold(FieldNoOfGuests, v))
}

// NoOfGuestsContainsFold applies the ContainsFold predicate on the "no_of_guests" field.
func NoOfGuestsContainsFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContainsFold(FieldNoOfGuests, v))
}

// CheckInEQ applies the EQ predicate on the "check_in" field.
func CheckInEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldCheckIn, v))
}

// CheckInNEQ applies the NEQ predicate on the "check_in" field.
func CheckInNEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNEQ(FieldCheckIn, v))
}

// CheckInIn applies the In predicate on the "check_in" field.
func CheckInIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldIn(FieldCheckIn, vs...))
}

// CheckInNotIn applies the NotIn predicate on the "check_in" field.
func CheckInNotIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNotIn(FieldCheckIn, vs...))
}

// CheckInGT applies the GT predicate on the "check_in" field.
func CheckInGT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGT(FieldCheckIn, v))
}

// CheckInGTE applies the GTE predicate on the "check_in" field.
func CheckInGTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGTE(FieldCheckIn, v))
}

// CheckInLT applies the LT predicate on the "check_in" field.
func CheckInLT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLT(FieldCheckIn, v))
}

// CheckInLTE applies the LTE predicate on the "check_in" field.
func CheckInLTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLTE(FieldCheckIn, v))
}

// CheckInContains applies the Contains predicate on the "check_in" field.
func CheckInContains(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContains(FieldCheckIn, v))
}

// CheckInHasPrefix applies the HasPrefix predicate on the "check_in" field.
func CheckInHasPrefix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasPrefix(FieldCheckIn, v))
}

// CheckInHasSuffix applies the HasSuffix predicate on the "check_in" field.
func CheckInHasSuffix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasSuffix(FieldCheckIn, v))
}

// CheckInIsNil applies the IsNil predicate on the "check_in" field.
func CheckInIsNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldIsNull(FieldCheckIn))
}

// CheckInNotNil applies the NotNil predicate on the "check_in" field.
func CheckInNotNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldNotNull(FieldCheckIn))
}

// CheckInEqualFold applies the EqualFold predicate on the "check_in" field.
func CheckInEqualFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEqualFold(FieldCheckIn, v))
}

// CheckInContainsFold applies the ContainsFold predicate on the "check_in" field.
func CheckInContainsFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContainsFold(FieldCheckIn, v))
}

// CheckOutEQ applies the EQ predicate on the "check_out" field.
func CheckOutEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldCheckOut, v))
}

// CheckOutNEQ applies the NEQ predicate on the "check_out" field.
func CheckOutNEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNEQ(FieldCheckOut, v))
}

// CheckOutIn applies the In predicate on the "check_out" field.
func CheckOutIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldIn(FieldCheckOut, vs...))
}

// CheckOutNotIn applies the NotIn predicate on the "check_out" field.
func CheckOutNotIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNotIn(FieldCheckOut, vs...))
}

// CheckOutGT applies the GT predicate on the "check_out" field.
func CheckOutGT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGT(FieldCheckOut, v))
}

// CheckOutGTE applies the GTE predicate on the "check_out" field.
func CheckOutGTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGTE(FieldCheckOut, v))
}

// CheckOutLT applies the LT predicate on the "check_out" field.
func CheckOutLT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLT(FieldCheckOut, v))
}

// CheckOutLTE applies the LTE predicate on the "check_out" field.
func CheckOutLTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLTE(FieldCheckOut, v))
}

// CheckOutContains applies the Contains predicate on the "check_out" field.
func CheckOutContains(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContains(FieldCheckOut, v))
}

// CheckOutHasPrefix applies the HasPrefix predicate on the "check_out" field.
func CheckOutHasPrefix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasPrefix(FieldCheckOut, v))
}

// CheckOutHasSuffix applies the HasSuffix predicate on the "check_out" field.
func CheckOutHasSuffix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasSuffix(FieldCheckOut, v))
}

// CheckOutIsNil applies the IsNil predicate on the "check_out" field.
func CheckOutIsNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldIsNull(FieldCheckOut))
}

// CheckOutNotNil applies the NotNil predicate on the "check_out" field.
func CheckOutNotNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldNotNull(FieldCheckOut))
}

// CheckOutEqualFold applies the EqualFold predicate on the "check_out" field.
func CheckOutEqualFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEqualFold(FieldCheckOut, v))
}

// CheckOutContainsFold applies the ContainsFold predicate on the "check_out" field.
func CheckOutContainsFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContainsFold(FieldCheckOut, v))
}

// PickupEQ applies the EQ predicate on the "pickup" field.
func PickupEQ(v Pickup) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldPickup, v))
}

// PickupNEQ applies the NEQ predicate on the "pickup" field.
func PickupNEQ(v Pickup) predicate.Bookings {
	return predicate.Bookings(sql.FieldNEQ(FieldPickup, v))
}

// PickupIn applies the In predicate on the "pickup" field.
func PickupIn(vs ...Pickup) predicate.Bookings {
	return predicate.Bookings(sql.FieldIn(FieldPickup, vs...))
}

// PickupNotIn applies the NotIn predicate on the "pickup" field.
func PickupNotIn(vs ...Pickup) predicate.Bookings {
	return predicate.Bookings(sql.FieldNotIn(FieldPickup, vs...))
}

// PickupIsNil applies the IsNil predicate on the "pickup" field.
func PickupIsNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldIsNull(FieldPickup))
}

// PickupNotNil applies the NotNil predicate on the "pickup" field.
func PickupNotNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldNotNull(FieldPickup))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Bookings {
	return predicate.Bookings(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Bookings {
	return predicate.Bookings(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Bookings {
	return predicate.Bookings(sql.FieldNotIn(FieldStatus, vs...))
}

// PaymentStatusEQ applies the EQ predicate on the "payment_status" field.
func PaymentStatusEQ(v PaymentStatus) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldPaymentStatus, v))
}

// PaymentStatusNEQ applies the NEQ predicate on the "payment_status" field.
func PaymentStatusNEQ(v PaymentStatus) predicate.Bookings {
	return predicate.Bookings(sql.FieldNEQ(FieldPaymentStatus, v))
}

// PaymentStatusIn applies the In predicate on the "payment_status" field.
func PaymentStatusIn(vs ...PaymentStatus) predicate.Bookings {
	return predicate.Bookings(sql.FieldIn(FieldPaymentStatus, vs...))
}

// PaymentStatusNotIn applies the NotIn predicate on the "payment_status" field.
func PaymentStatusNotIn(vs ...PaymentStatus) predicate.Bookings {
	return predicate.Bookings(sql.FieldNotIn(FieldPaymentStatus, vs...))
}

// SpecialRequestsEQ applies the EQ predicate on the "special_requests" field.
func SpecialRequestsEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEQ(FieldSpecialRequests, v))
}

// SpecialRequestsNEQ applies the NEQ predicate on the "special_requests" field.
func SpecialRequestsNEQ(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNEQ(FieldSpecialRequests, v))
}

// SpecialRequestsIn applies the In predicate on the "special_requests" field.
func SpecialRequestsIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldIn(FieldSpecialRequests, vs...))
}

// SpecialRequestsNotIn applies the NotIn predicate on the "special_requests" field.
func SpecialRequestsNotIn(vs ...string) predicate.Bookings {
	return predicate.Bookings(sql.FieldNotIn(FieldSpecialRequests, vs...))
}

// SpecialRequestsGT applies the GT predicate on the "special_requests" field.
func SpecialRequestsGT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGT(FieldSpecialRequests, v))
}

// SpecialRequestsGTE applies the GTE predicate on the "special_requests" field.
func SpecialRequestsGTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldGTE(FieldSpecialRequests, v))
}

// SpecialRequestsLT applies the LT predicate on the "special_requests" field.
func SpecialRequestsLT(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLT(FieldSpecialRequests, v))
}

// SpecialRequestsLTE applies the LTE predicate on the "special_requests" field.
func SpecialRequestsLTE(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldLTE(FieldSpecialRequests, v))
}

// SpecialRequestsContains applies the Contains predicate on the "special_requests" field.
func SpecialRequestsContains(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContains(FieldSpecialRequests, v))
}

// SpecialRequestsHasPrefix applies the HasPrefix predicate on the "special_requests" field.
func SpecialRequestsHasPrefix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasPrefix(FieldSpecialRequests, v))
}

// SpecialRequestsHasSuffix applies the HasSuffix predicate on the "special_requests" field.
func SpecialRequestsHasSuffix(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldHasSuffix(FieldSpecialRequests, v))
}

// SpecialRequestsIsNil applies the IsNil predicate on the "special_requests" field.
func SpecialRequestsIsNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldIsNull(FieldSpecialRequests))
}

// SpecialRequestsNotNil applies the NotNil predicate on the "special_requests" field.
func SpecialRequestsNotNil() predicate.Bookings {
	return predicate.Bookings(sql.FieldNotNull(FieldSpecialRequests))
}

// SpecialRequestsEqualFold applies the EqualFold predicate on the "special_requests" field.
func SpecialRequestsEqualFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldEqualFold(FieldSpecialRequests, v))
}

// SpecialRequestsContainsFold applies the ContainsFold predicate on the "special_requests" field.
func SpecialRequestsContainsFold(v string) predicate.Bookings {
	return predicate.Bookings(sql.FieldContainsFold(FieldSpecialRequests, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Bookings) predicate.Bookings {
	return predicate.Bookings(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Bookings) predicate.Bookings {
	return predicate.Bookings(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Bookings) predicate.Bookings {
	return predicate.Bookings(sql.NotPredicates(p))
}