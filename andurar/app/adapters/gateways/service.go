package gateways

import (
	"andurar.api/app/adapters/presenters"
	requestdto "andurar.api/app/domain/request_dto"
	"andurar.api/ent"
)

type (
	BookingService interface {
		Fetch(id int) (*ent.Bookings, error)
		FetchAll() (*presenters.PaginationResponse, error)
		Create(request *requestdto.BookingRequest) (*ent.Bookings, error)
		Update(id int, request *requestdto.BookingRequest) (*ent.Bookings, error)
		Remove(id int) error
	}
)
