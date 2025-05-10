package utilities

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

const (
	StatusOK                  = 200
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusConflict            = 409
	StatusInternalServerError = 500
	StatusTooManyRequests     = 429
)

type AppError struct {
	Message string
	Code    int
	Details interface{}
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	if appError, ok := err.(*AppError); ok {
		if appError.Details != nil {
			return c.Status(appError.Code).JSON(fiber.Map{
				"error":   appError.Message,
				"details": appError.Details,
			})
		}
		return c.Status(appError.Code).JSON(fiber.Map{
			"error": appError.Message,
		})
	}

	log.Printf("Internal server error: %v", err)

	return c.Status(StatusInternalServerError).JSON(fiber.Map{
		"error": "Internal Server Error",
	})
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

type NotFoundError struct {
	AppError
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{AppError{Message: message, Code: StatusNotFound}}
}

type BadRequestError struct {
	AppError
}

func NewBadRequestError(message string, details interface{}) *BadRequestError {
	return &BadRequestError{AppError{Message: message, Code: StatusBadRequest, Details: details}}
}

type UnauthorizedError struct {
	AppError
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{AppError{Message: message, Code: StatusUnauthorized}}
}

type InternalServerError struct {
	AppError
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{AppError{Message: message, Code: StatusInternalServerError}}
}

type TooManyRequestsError struct {
	AppError
}

func NewTooManyRequestsError(message string) *TooManyRequestsError {
	return &TooManyRequestsError{AppError{Message: message, Code: StatusTooManyRequests}}
}

type ForbiddenError struct {
	AppError
}

func NewForbiddenError(message string) *ForbiddenError {
	return &ForbiddenError{AppError{Message: message, Code: StatusForbidden}}
}

type ConflictError struct {
	AppError
}

func NewConflictError(message string) *ConflictError {
	return &ConflictError{AppError{Message: message, Code: StatusConflict}}
}
