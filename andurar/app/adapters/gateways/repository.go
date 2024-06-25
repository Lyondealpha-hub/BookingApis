package gateways

import (
	"andurar.api/app/adapters/presenters"
	requestdto "andurar.api/app/domain/request_dto"
	"andurar.api/ent"
)

type (
	BookingRepo interface {
		Read(id int) (*ent.Bookings, error)
		ReadAll() (*presenters.PaginationResponse, error)
		Insert(request *requestdto.BookingRequest) (*ent.Bookings, error)
		Update(id int, request *requestdto.BookingRequest) (*ent.Bookings, error)
		Delete(id int) error
	}
)
