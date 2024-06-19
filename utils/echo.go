package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/labstack/echo/v4"

	"github.com/alifakhimi/simple-service-go/database"
	"github.com/alifakhimi/simple-service-go/utils/templates"
)

func IsDigit(s string) bool {
	if strings.TrimSpace(s) == "" {
		return false
	}

	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func ReplyTemplate(ctx echo.Context, httpStatus int, err error, template interface{}, meta interface{}) error {
	data, er := json.Marshal(template) // Convert to a json string
	if er != nil {
		return er
	}

	content := make(map[string]interface{})

	er = json.Unmarshal(data, &content) // Convert to a map
	if err != nil {
		return er
	}

	return Reply(ctx, httpStatus, err, content, meta)
}

// Reply ...
func Reply(ctx echo.Context, httpStatus int, err error, content map[string]interface{}, meta interface{}) error {
	var template *templates.ResponseTemplate

	switch httpStatus {
	case http.StatusOK:
		template = templates.Ok(content, err, meta)
	case http.StatusCreated:
		template = templates.Created(content, meta)
	case http.StatusBadRequest:
		template = templates.BadRequest(content, err.Error())
	case http.StatusInternalServerError:
		template = templates.InternalServerError(content, err.Error())
	case http.StatusNotFound:
		template = templates.NotFound(content, err.Error())
	case http.StatusUnprocessableEntity:
		template = templates.UnprocessableEntity(content, err.Error())
	case http.StatusMethodNotAllowed:
		template = templates.MethodNotAllowed(content, err.Error())
	case http.StatusUnauthorized:
		template = templates.Unauthorized(content, err.Error())
	case http.StatusForbidden:
		template = templates.Forbidden(content, err.Error())
	case http.StatusGatewayTimeout:
		template = templates.GatewayTimeOut(content, err.Error())
	case http.StatusLocked:
		template = templates.Locked(content, err.Error())
	case http.StatusNotAcceptable:
		template = templates.NotAcceptable(content, err.Error())
	default:
		template = templates.InternalServerError(content, errors.New("invalid reply request"))
	}

	return ctx.JSON(httpStatus, template)
}

// ExportRoutes ...
func ExportRoutes(e *echo.Echo, prefix string) ([]*echo.Route, error) {
	var apiRoutes []*echo.Route
	routes := e.Routes()
	for _, route := range routes {
		if strings.Index(route.Path, prefix) == 0 {
			apiRoutes = append(apiRoutes, route)
		}
	}

	// data, err := json.Marshal(apiRoutes)
	// if err != nil {
	// 	return nil, err
	// }
	// //ioutil.WriteFile("routes.json", data, 0644)
	// return data, err
	return apiRoutes, nil
}

func DurationToHumanity(d time.Duration) (days int, hours int, minutes int) {
	minutes = int(d.Minutes()) % 60
	hours = int(d.Hours()) % 24
	days = int(d.Hours()) / 24
	return days, hours, minutes
}

// ParseFilterQuery ...
func ParseFilterQuery(ctx echo.Context, key string) []uint {

	_url := ctx.QueryString()
	qs, _ := url.ParseQuery(_url)
	filterArr := qs[key]

	filters := make([]uint, len(filterArr))
	for i := 0; i < len(filterArr); i++ {
		v, _ := strconv.Atoi(filterArr[i])
		filters[i] = uint(v)
	}

	return filters
}

// ParseIDsQuery ...
func ParseIDsQuery(ctx echo.Context, key string) database.PIDs {
	_url := ctx.QueryString()
	qs, _ := url.ParseQuery(_url)
	filterArr := qs[key]
	ids := strings.Split(filterArr[0], ",")

	filters := make(database.PIDs, len(ids))
	for i := 0; i < len(ids); i++ {
		if pid, err := database.ParsePID(ids[i]); err == nil {
			filters[i] = pid
		}
	}

	return filters
}

func ArrayElementExists(a []string, v string) bool {
	for _, s := range a {
		if s == v {
			return true
		}
	}

	return false
}

func ErrorToHttpStatusCode(err error) (status int) {
	switch err {
	case ErrNotFound, ErrRecordNotFound:
		status = http.StatusNotFound
	case ErrInvalidRequest:
		status = http.StatusBadRequest
	case ErrAlreadyExist:
		status = http.StatusNotAcceptable

	default:
		status = http.StatusNotImplemented
	}

	return
}
