package router

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func registerHandler(c *fiber.Ctx) error {
	var dto registerDto_t
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	if err := validator.New().Struct(dto); err != nil {
		c.SendString(err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}
	log.Infof("Register dto: %v\n", dto)

	if err := register(dto); err != nil {
		return err
	}

	return nil
}

func createHandler(c *fiber.Ctx) error {
	var dto newsDto_t
	if err := c.BodyParser(&dto); err != nil {
		log.Error(err)
		return fiber.ErrBadRequest
	}

	if err := validator.New().Struct(dto); err != nil {
		c.SendString(err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}

	log.Infof("create dto: %v\n", dto)

	if err := create(dto); err != nil {
		return err
	}

	return nil
}

func editHandler(c *fiber.Ctx) error {
	newsIdString := c.Params("id", "")
	if newsIdString == "" {
		return fiber.ErrBadRequest
	}
	newsId, err := strconv.ParseInt(newsIdString, 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}

	var dto newsDto_t
	if err := c.BodyParser(&dto); err != nil {
		log.Error(err)
		return fiber.ErrBadRequest
	}

	if err := validator.New().Struct(dto); err != nil {
		c.SendString(err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}

	log.Infof("Edit dto: %v\n", dto)

	if err := edit(newsId, dto); err != nil {
		return err
	}

	return nil
}

func listHandler(c *fiber.Ctx) error {
	limit, err := strconv.ParseInt(c.Query("limit", "10"), 10, 0)
	if err != nil {
		return fiber.ErrBadRequest
	}
	offset, err := strconv.ParseInt(c.Query("offset", "0"), 10, 0)
	if err != nil {
		return fiber.ErrBadRequest
	}

	list, err := list(limit, offset)
	if err != nil {
		return err
	}
	listOut := listOut_t{
		News: list,
	}
	if list == nil || len(list) == 0 {
		listOut.Success = false
	} else {
		listOut.Success = true
	}
	return c.JSON(listOut)
}
