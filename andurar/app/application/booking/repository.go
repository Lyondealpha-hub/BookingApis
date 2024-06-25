package booking

import (
	"context"

	"andurar.api/app/adapters/gateways"
	"andurar.api/app/adapters/presenters"
	"andurar.api/app/application"
	requestdto "andurar.api/app/domain/request_dto"
	"andurar.api/ent"
)

type repository struct {
	db  *ent.Client
	ctx context.Context
}

func NewRepository(db *ent.Client) gateways.BookingRepo {
	return &repository{
		db:  db,
		ctx: context.Background(),
	}
}

// Read implements gateways.BookingRepo.
func (r *repository) Read(id int) (*ent.Bookings, error) {
	panic("unimplemented")
}

// Read implements gateways.BookingRepo.
func (r *repository) ReadAll() (*presenters.PaginationResponse, error) {
	count := r.db.Bookings.Query().CountX(r.ctx)
	result, err := r.db.Bookings.Query().All(r.ctx)

	if err != nil {
		return nil, err
	}

	return application.Paginate(count, result)
}

// Insert implements gateways.BookingRepo.
func (r *repository) Insert(request *requestdto.BookingRequest) (*ent.Bookings, error) {

	result, err := r.db.Bookings.Create().
		SetFullName(request.FullName).
		SetEmail(request.Email).
		SetPhone(request.Phone).
		SetRoom(request.Room).
		SetCheckIn(request.CheckIn).
		SetCheckOut(request.CheckOut).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete implements gateways.BookingRepo.
func (r *repository) Delete(id int) error {
	panic("unimplemented")
}

// Update implements gateways.BookingRepo.
func (r *repository) Update(id int, request *requestdto.BookingRequest) (*ent.Bookings, error) {
	panic("unimplemented")
}
