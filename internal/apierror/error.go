package apierror

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

var (
	db *sql.DB
)

// APIError stores details about an api error. This is returned to the end user in json
type APIError struct {
	Title     string
	Message   string
	Reference string
}

// IfErr will send an error to the end user with the message set to err
func IfErr(err error, apiErr APIError, w http.ResponseWriter, r *http.Request) {

	// Make sure that the error exists.
	if err == nil {
		return
	}

	// Set the message to the error
	apiErr.Message = err.Error()

	// Create the json and send it to the user.
	render.JSON(w, r, apiErr)

	// Log something.
	log.Printf("API Error: %s: %s\n%s", apiErr.Title, apiErr.Message, err.Error())
}

// IfErrMsg only shows the error message if err is not nil.
func IfErrMsg(err error, message string, apiErr APIError, w http.ResponseWriter, r *http.Request) {

	// Make sure that the error exists.
	if err == nil {
		return
	}

	// Set the message to the one provided, this is used to hide the underlying error from the end-user.
	apiErr.Message = message

	// Create the json and send it to the user.
	render.JSON(w, r, apiErr)

	// Log something.
	log.Printf("API Error: %s: %s\n%s", apiErr.Title, apiErr.Message, err.Error())
}
