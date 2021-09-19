package endpointsControllers

import (
	"github.com/labstack/echo"
	"net/http"
)

// Response Struct
type Response struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
}

// Data Struct
type Data struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// PaginatedData Struct
type PaginatedData struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Count   int64       `json:"count"`
	Pages   int64       `json:"pages"`
}

// Error Struct
type Error struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
}

// MissingAuthorization struct
type MissingAuthorization struct {
	Message string `json:"message"`
}

// CompleteParams Struct
type CompleteParams struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Token           string `json:"token"`
}

func InternalError(c echo.Context, message string) error {
	return c.JSONPretty(http.StatusInternalServerError, Error{
		Message: message,
		Success: false,
	}, " ")
}

func ForbiddenError(c echo.Context) error {
	return c.JSONPretty(http.StatusForbidden, Error{
		Message: "You are not allowed to access this endpoint",
		Success: false,
	}, " ")
}

func BadRequestResponse(c echo.Context, message string) error {
	return c.JSONPretty(http.StatusBadRequest, Error{
		Message: message,
		Success: false,
	}, " ")
}

func MessageResponse(c echo.Context, status int, message string) error {
	return c.JSONPretty(status, Error{
		Message: message,
		Success: status <= 201,
	}, " ")
}
func MessageResponseForward(c echo.Context, status int, message string) error {
	return c.JSONPretty(status, Error{
		Message: message,
		Success: false,
	}, " ")
}

func DataResponse(c echo.Context, status int, data interface{}) error {
	return c.JSONPretty(status, Data{
		Data:    data,
		Success: true,
	}, " ")
}

func PaginatedDataResponse(c echo.Context, status int, data interface{}, count int64, pages int64) error {
	return c.JSONPretty(status, PaginatedData{
		Data:    data,
		Count:   count,
		Pages:   pages,
		Success: true,
	}, " ")
}

func JsonResponse(c echo.Context, status int, response interface{}) error {
	return c.JSONPretty(status, response, " ")
}

func InvalidResponse(c echo.Context) error {
	return BadRequestResponse(c, "Invalid Request")
}
