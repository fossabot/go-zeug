package sqlite

import (
	"database/sql/driver"
	"time"
)

// SQLTime will represent an time.Time for sqlite with milliseconds support.
type SQLTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// SQLTimeFromTime will create a SQLTime from time.Time.
// It converts the time to time.UTC().
func SQLTimeFromTime(t time.Time) SQLTime {
	return SQLTime{
		Time:  t.UTC(),
		Valid: true,
	}
}

// Scan implements the sql.Scanner interface.
func (t *SQLTime) Scan(src interface{}) error {
	if src == nil {
		t.Time, t.Valid = time.Time{}, false
		return nil
	}

	switch v := src.(type) {
	case time.Time:
		t.Time, t.Valid = v.UTC(), true
		return nil
	case int:
		t.Time = time.Unix(0, int64(v)*int64(time.Millisecond))
		t.Valid = true
		return nil
	case int64:
		t.Time = time.Unix(0, v*int64(time.Millisecond))
		t.Valid = true
		return nil
	default:
		return ErrSQLTimeUnmarshallType
	}
}

// Value implements the driver driver.Valuer interface.
func (t SQLTime) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}

	return t.Time.UnixNano() / 1000000, nil
}

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in RFC 3339 format, with sub-second precision added if present.
func (t SQLTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time.MarshalJSON()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (t *SQLTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}

	tmp := &time.Time{}
	err := tmp.UnmarshalJSON(data)
	t.Valid = err == nil
	if err != nil {
		return err
	}
	t.Time = *tmp

	return nil
}
