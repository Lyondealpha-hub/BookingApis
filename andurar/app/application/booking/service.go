package booking

import (
	"andurar.api/app/adapters/gateways"
	"andurar.api/app/adapters/presenters"
	requestdto "andurar.api/app/domain/request_dto"
	"andurar.api/ent"
)

type service struct {
	repo gateways.BookingRepo
}

func NewService(repo gateways.BookingRepo) gateways.BookingService {
	return &service{
		repo: repo,
	}
}

// Fetch implements gateways.BookingService.
func (s *service) Fetch(id int) (*ent.Bookings, error) {
	return s.repo.Read(id)
}

// FetchAll implements gateways.BookingService.
func (s *service) FetchAll() (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll()
}

// Create implements gateways.BookingService.
func (s *service) Create(request *requestdto.BookingRequest) (*ent.Bookings, error) {
	return s.repo.Insert(request)
}

// Remove implements gateways.BookingService.
func (s *service) Remove(id int) error {
	panic("unimplemented")
}

// Update implements gateways.BookingService.
func (s *service) Update(id int, request *requestdto.BookingRequest) (*ent.Bookings, error) {
	panic("unimplemented")
}
