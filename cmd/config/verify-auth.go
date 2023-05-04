package config

import (
	"encoding/json"
	"net/http"
	objectValues "service-routes/internal/domain/object_values"
	"service-routes/internal/infra/rest"

	"github.com/labstack/echo/v4"
)

func AuthVerify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		headers := make(map[string]string)
		headers["Authorization"] = c.Request().Header.Get("Authorization")
		client := rest.NewRequest(GetConfig().Microservices.Auth, headers)
		response, err := client.Post("", nil)
		if err != nil {
			return c.NoContent(http.StatusBadGateway)
		}
		if response.StatusCode == 200 {
			auth := objectValues.Auth{}
			err := json.Unmarshal([]byte(response.Body), &auth)
			if err != nil {
				return c.NoContent(http.StatusBadGateway)
			}
			c.Set("auth", auth)
			return next(c)
		}
		return c.NoContent(http.StatusUnauthorized)
	}
}
