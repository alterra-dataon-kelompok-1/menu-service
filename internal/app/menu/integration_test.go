package menu

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alterra-kelompok-1/menu-service/database"
	"github.com/alterra-kelompok-1/menu-service/database/migration"
	"github.com/alterra-kelompok-1/menu-service/database/seeder"
	"github.com/alterra-kelompok-1/menu-service/internal/dto"
	"github.com/alterra-kelompok-1/menu-service/internal/factory"
	"github.com/alterra-kelompok-1/menu-service/internal/middleware"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func createTestApp() (*echo.Echo, *gorm.DB, handler) {
	database.CreateConnection()
	migration.Migrate()
	seeder.Seed()

	e := echo.New()
	f := factory.NewFactory()

	// menuRepo := repository.NewMenu(database.GetConnection())
	// menuService := NewService(f)
	menuHandler := NewHandler(f)

	middleware.Init(e)

	// return e, database.GetConnection(), menuHandler
	return e, database.GetConnection(), *menuHandler
}

func TestGetMenu(t *testing.T) {
	e, _, h := createTestApp()
	// defer database.DropTables(db)
	// defer migration.Rollback()

	req := httptest.NewRequest("GET", "/menus", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, h.Get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	}
}

func TestCreateMenu_MissingField(t *testing.T) {
	e, _, h := createTestApp()
	// defer database.DropTables(db)
	// defer migration.Rollback()

	body := dto.CreateMenuRequest{
		Name:  "Nasi Goreng",
		Price: 123,
	}
	reqBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/menus", bytes.NewReader(reqBody))
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, h.Get(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
		assert.Contains(t, rec.Body.String(), "Bad Request")
	}
}

func TestGetOrderByID(t *testing.T) {
	// Setup
	e, _, h := createTestApp()
	// defer database.DropTables(db)
	// defer migration.Rollback()

	testCase := []struct {
		Case             string
		ParamID          string
		WantResponseCode int
		WantBodyContains string
	}{
		{
			Case:             "success",
			ParamID:          "aca1522a-07b6-4c0c-aed6-04a1d123835f",
			WantResponseCode: http.StatusOK,
			WantBodyContains: "aca1522a-07b6-4c0c-aed6-04a1d123835f",
		},
		{
			Case:             "incorrect UUID format",
			ParamID:          "aca12a-07b6-4c0c-aed6-0423835f",
			WantResponseCode: http.StatusBadRequest,
			WantBodyContains: "E_BAD_REQUEST",
		},
		{
			Case:             "ID not found",
			ParamID:          "1c08b996-92bb-4c09-aa3b-989b4c5092ca",
			WantResponseCode: http.StatusNotFound,
			WantBodyContains: "E_NOT_FOUND",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.Case, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest("GET", "/menus", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tc.ParamID)

			// Assertion
			if assert.NoError(t, h.GetByID(c)) {
				assert.Equal(t, tc.WantResponseCode, rec.Code)
				assert.Contains(t, rec.Body.String(), tc.WantBodyContains)
			}
		})
	}
}
