// Package layout provides layout models.
package layout

// NavItem represents a navigation item.
type NavItem struct {
	Label  string
	Href   string
	Active bool
}

// UserSummary represents a user summary for the layout.
type UserSummary struct {
	Name  string
	Email string
}
