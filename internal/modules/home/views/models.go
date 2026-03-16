// Package views contains the views and viewmodels for the home module.
package views

import "example.com/webserver/internal/app/db/store"

// IndexVM is the viewmodel for the index view.
type IndexVM struct {
	Users []store.User
}
