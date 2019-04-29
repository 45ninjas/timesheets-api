package timesheets

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var (
	db *sql.DB
)

// Shift stores details about a shift.
type Shift struct {
	ID     uint      `json:"id"`
	Start  time.Time `json:"start"`
	Finish time.Time `json:"finish"`
	Total  float32   `json:"total"`
}

// Init initializes.
func Init(dbc *sql.DB) {
	db = dbc
}

// Routes retruns all the routes for this module.
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
	// Query the database.
	rows, err := db.Query("SELECT s.id, s.start, s.finish, TIMESTAMPDIFF(MINUTE, s.start, s.finish) / 60 AS total from shifts s")
	defer rows.Close()

	if err != nil {
		log.Fatalf("Error getting all shifts, %s", err.Error())
	}

	// put all the shifts into an array.
	var shifts []Shift
	for rows.Next() {
		shift := Shift{}

		err = rows.Scan(&shift.ID, &shift.Start, &shift.Finish, &shift.Total)

		if err != nil {
			panic(err)
		}

		shifts = append(shifts, shift)
	}

	// Dump that array to the client.
	render.JSON(w, r, shifts)
}
