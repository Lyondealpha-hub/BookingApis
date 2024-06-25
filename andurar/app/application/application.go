package application

import (
	"log"
	"time"

	"andurar.api/app/adapters/presenters"
	"andurar.api/utils/bodyparser"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func BodyParser(c fiber.Ctx, v any) error {
	httpReq, err := adaptor.ConvertRequest(c, false)
	if err != nil {
		return err
	}
	return bodyparser.Parse(httpReq, v)
}

func ParseRFC3339Datetime(rfc3339Datetime ...string) time.Time {
	if rfc3339Datetime == nil || rfc3339Datetime[0] == "" {
		return time.Now()
	}
	rfc3339Time, err := time.Parse(time.RFC3339, rfc3339Datetime[0])
	if err != nil {
		log.Panicln("Error parsing RFC3339 datetime:", err)
	}
	return rfc3339Time
}

func Paginate(count int, results any) (*presenters.PaginationResponse, error) {
	return &presenters.PaginationResponse{
		Count: count,
		Data:  results,
	}, nil
}
