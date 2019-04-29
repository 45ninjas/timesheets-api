package timesheets

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Shift stores details about a shift.
type Shift struct {
	ID     uint    `json:"id"`
	Start  uint    `json:"start"`
	Finish uint    `json:"finish"`
	Total  float32 `json:"total"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{shiftID}", GetShift)
	router.Delete("/{shiftID}", DeleteShift)
	router.Post("/", CreateShift)
	router.Put("/{shiftID}", UpdateShift)
	router.Get("/", GetAllShifts)

	return router
}

// GetShift will get a shift from the database.
func GetShift(w http.ResponseWriter, r *http.Request) {
	// Get the {shiftID} of the url.
	shiftID := chi.URLParam(r, "shiftID")

	render.JSON(w, r, shiftID)
}

// DeleteShift removes a shift from the database.
func DeleteShift(w http.ResponseWriter, r *http.Request) {
	// Get the {shiftID} of the url.
	shiftID := chi.URLParam(r, "shiftID")

	render.JSON(w, r, shiftID)
}

// CreateShift will insert an new shift into the database and returns the new index.
func CreateShift(w http.ResponseWriter, r *http.Request) {
	// Get the {shiftID} of the url.
	shiftID := chi.URLParam(r, "shiftID")

	render.JSON(w, r, shiftID)
}

// UpdateShift will update an existing shift.
func UpdateShift(w http.ResponseWriter, r *http.Request) {
	// Get the {shiftID} of the url.
	shiftID := chi.URLParam(r, "shiftID")

	render.JSON(w, r, shiftID)
}

// GetAllShifts will get all the shifts in the database for the current user.
func GetAllShifts(w http.ResponseWriter, r *http.Request) {

}
