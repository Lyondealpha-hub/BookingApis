package requestdto

type (
	BookingRequest struct {
		FullName       string `json:"firstName"`
		Email          string `json:"email"`
		Phone          string `json:"phone"`
		Room           string `json:"room"`
		NoOfGuests     string `json:"noOfGuests"`
		CheckIn        string `json:"checkIn"`
		CheckOut       string `json:"checkOut"`
		PickUp         string `json:"pickUp"`
		SpecialRequest string `json:"specialRequest"`
	}
)
