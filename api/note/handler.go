package note

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	HTTPHandler struct {
		NoteUsecase Usecase
	}
)

func NewHTTPHandler(e *echo.Echo, noteUcase Usecase) {
	handler := &HTTPHandler{
		NoteUsecase: noteUcase,
	}

	g := e.Group("/note")
	g.GET("", handler.GetAllNoteHandler)
	g.POST("", handler.CreateNoteHandler)
}

// CreateNoteHandler handling request to create new note
func (handler *HTTPHandler) CreateNoteHandler(c echo.Context) (err error) {
	n := new(Note)
	if err = c.Bind(n); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err = n.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = handler.NoteUsecase.Store(n); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, n)
}

// GetAllNoteHandler handling request to get all stored notes
func (handler *HTTPHandler) GetAllNoteHandler(c echo.Context) (err error) {
	notes, err := handler.NoteUsecase.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, notes)
}
