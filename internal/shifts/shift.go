package shifts

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"timesheets/internal/apierror"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var (
	db     *sql.DB
	apiErr apierror.APIError
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

	apiErr = apierror.APIError{"Shift", "", "docs/endpoints/shifts"}
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

	rows, err := db.Query("SELECT s.id, s.start, s.finish, TIMESTAMPDIFF(MINUTE, s.start, s.finish) / 60 AS total FROM shifts AS s WHERE s.id = ?", shiftID)
	defer rows.Close()

	if err != nil {
		// This is bad. SQL is bad.
		panic(err.Error)
	}

	if rows.Next() {
		// Get the first shift from the rows.
		shift, err := ScanShift(rows)

		if err != nil {
			panic(err)
		}

		// Output the shift.
		render.JSON(w, r, shift)
	} else {
		// Nothing was found.
		http.Error(w, "", http.StatusNotFound)
	}
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

	stmt, err := db.Prepare("UPDATE shifts SET start=?, finish=? WHERE id=?")
	defer stmt.Close()

	if err != nil {
		log.Fatalf("Error updating a shift, %s", err.Error())
	}

	// Decode the json in the request's body.
	decoder := json.NewDecoder(r.Body)
	var shift Shift
	err = decoder.Decode(&shift)

	// TODO: Send this error message back to the end-user.
	// apiErr.Message = "Unable to parse your json"
	apierror.IfErr(err, apiErr, w, r)

	// TODO: Make sure the shiftID exists.

	// Execute the database query.
	stmt.Exec(shift.Start, shift.Finish, shiftID)

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
		shift, err := ScanShift(rows)

		if err != nil {
			panic(err)
		}

		shifts = append(shifts, shift)
	}
	// There where no results, return 204.
	if len(shifts) == 0 {
		http.Error(w, "", http.StatusNoContent)
	} else {
		// Dump that array to the client.
		render.JSON(w, r, shifts)
	}
}

// ScanShift scan a shift from a database row.
func ScanShift(rows *sql.Rows) (Shift, error) {
	shift := Shift{}
	err := rows.Scan(&shift.ID, &shift.Start, &shift.Finish, &shift.Total)

	return shift, err
}
