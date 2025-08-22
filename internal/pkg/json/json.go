package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type Jsoniter struct{}

func (j *Jsoniter) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := jsoniter.NewEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

func (j *Jsoniter) Deserialize(c echo.Context, i interface{}) error {
	return jsoniter.NewDecoder(c.Request().Body).Decode(i)
}
