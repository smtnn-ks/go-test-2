package router

import (
	"encoding/base64"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func authMiddleware(c *fiber.Ctx) error {
	authString := c.Get("Authorization")
	if authString == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	strs := strings.Split(authString, " ")
	if len(strs) != 2 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	if strs[0] != "Basic" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	credsByte, err := base64.URLEncoding.DecodeString(strs[1])
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	creds := string(credsByte)
	creds = strings.Trim(creds, "\n")

	strs = strings.Split(creds, ":")

	if len(strs) != 2 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if err := checkAuth(strs[0], strs[1]); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Next()
}
