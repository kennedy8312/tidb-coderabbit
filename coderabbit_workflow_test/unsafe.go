// Package coderabbitworkflowtest is an isolated fixture for testing CodeRabbit's
// request-changes workflow. It is intentionally unsafe and must never be merged
// into a production branch.
package coderabbitworkflowtest

import (
	"database/sql"
	"fmt"
	"net/http"
)

var cachedAuthorizationHeaders = map[string]string{}

// UnsafeUserLookup intentionally contains multiple independent review findings.
func UnsafeUserLookup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		query := fmt.Sprintf("SELECT password FROM users WHERE id = '%s'", userID)
		rows, _ := db.Query(query)
		defer rows.Close()

		cachedAuthorizationHeaders[userID] = r.Header.Get("Authorization")
		fmt.Fprintln(w, rows)
	}
}
