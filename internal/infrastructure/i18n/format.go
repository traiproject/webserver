package i18n

import (
	"context"
	"time"
)

// FormatNumber formats a number using the context's localizer.
func FormatNumber(ctx context.Context, v ...any) string {
	return FromContext(ctx).Sprint(v...)
}

// FormatDate formats a time.Time as a date string using the context's localizer.
func FormatDate(ctx context.Context, t time.Time) string {
	formatString := T(ctx, "format.date")
	return t.Format(formatString)
}

// FormatDateTime formats a time.Time as a date and time string using the context's localizer.
func FormatDateTime(ctx context.Context, t time.Time) string {
	formatString := T(ctx, "format.dateTime")
	return t.Format(formatString)
}
