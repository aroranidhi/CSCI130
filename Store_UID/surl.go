package mem

import (
        "net/http"
        "errors"
        )

// called if no cookie was found, returns the uuid from the URL path. Retruens an error if not found.
func idFromURL(res http.ResponseWriter, req *http.Request) (string, error) {
    var err error = nil
    id := req.FormValue("id")
    if id == "" {
        errors.New("No id value found.")
    }
    return id, err
}