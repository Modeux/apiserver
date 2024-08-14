package databases

import (
	"database/sql"
	"github.com/goccy/go-json"
	"time"
)

type JsonNullString struct {
	sql.NullString
}

func (v JsonNullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	}
	return json.Marshal(nil)
}

type JsonNullTime struct {
	sql.NullTime
}

func (v JsonNullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time.Format(time.DateTime))
	}
	return json.Marshal(nil)
}

type JsonNullInt64 struct {
	sql.NullInt64
}

func (v JsonNullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}
	return json.Marshal(nil)
}

type JSONDate struct {
	sql.NullTime
}

func (t JSONDate) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time.Format(time.DateOnly))
	}
	return json.Marshal(nil)
}
