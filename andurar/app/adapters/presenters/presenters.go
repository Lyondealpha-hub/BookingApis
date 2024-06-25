package presenters

import "github.com/gofiber/fiber/v3"

type (
	PaginationResponse struct {
		Count int `json:"count,omitempty"`
		Data  any `json:"data"`
	}
)

func ErrorResponse(e error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"error":  e.Error(),
	}
}

func SuccessResponse(data any) *fiber.Map {

	if data == nil {
		return &fiber.Map{
			"status": true,
			"data":   nil,
		}
	}

	d, ok := data.(*PaginationResponse)
	if ok && d.Count == 0 {
		return &fiber.Map{
			"status": true,
			"data":   nil,
		}
	}
	return &fiber.Map{
		"status": true,
		"data":   data,
	}
}
