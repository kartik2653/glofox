package class

import (
	apimodel "glofox/internal/model/api_model"
	"glofox/pkg"
	"glofox/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ClassHandler struct {
	service ClassService
}

func NewClassHandler(service ClassService) *ClassHandler {
	return &ClassHandler{service: service}
}

func (h *ClassHandler) CreateClass(c *fiber.Ctx) error {
	var class apimodel.Class

	if err := c.BodyParser(&class); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.ErrorResponse{Error: pkg.ErrBadRequest})
	}

	err := h.service.CreateClass(&class)
	if err != nil {

		switch err {
		case pkg.ErrClassNameRequired,
			pkg.ErrStartDateRequired,
			pkg.ErrEndDateRequired,
			pkg.ErrInvalidCapacity,
			pkg.ErrInstructorRequired,
			pkg.ErrInvalidDateFormat:

			return c.Status(fiber.StatusBadRequest).
				JSON(response.ErrorResponse{Error: err.Error()})
		}

		return c.Status(fiber.StatusInternalServerError).
			JSON(response.ErrorResponse{Error: pkg.ErrInternalServer})
	}

	return c.Status(fiber.StatusCreated).
		JSON(response.SuccessResponse{
			Message: "Class created successfully",
		})
}

func (h *ClassHandler) ListClasses(c *fiber.Ctx) error {
	query := make(map[string]interface{})

	if className := c.Query("class_name"); className != "" {
		query["class_name"] = className
	}
	if classID := c.Query("id"); classID != "" {
		id, err := strconv.Atoi(classID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.ErrorResponse{Error: pkg.ErrBadRequest})
		}
		query["id"] = id
	}

	skipStr := c.Query("skip", "0")
	limitStr := c.Query("limit", "10")

	skip, err := strconv.Atoi(skipStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.ErrorResponse{Error: pkg.ErrBadRequest})
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.ErrorResponse{Error: pkg.ErrBadRequest})
	}

	classes, total, limit, offset, err := h.service.ListClasses(query, skip, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(response.ErrorResponse{Error: pkg.ErrInternalServer})
	}

	return c.Status(fiber.StatusOK).
		JSON(response.SuccessResponse{
			Data: map[string]interface{}{
				"classes": classes,
				"total":   total,
				"limit":   limit,
				"skip":    offset,
			},
			Message: "Classes retrieved successfully",
		})
}
