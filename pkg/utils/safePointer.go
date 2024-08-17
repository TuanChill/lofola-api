package utils

import "time"

func SafeDereferenceString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func SafeDereferenceTime(t *time.Time) time.Time {
	if t != nil {
		return *t
	}
	return time.Time{} // Return zero value of time.Time if nil
}
