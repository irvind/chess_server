package dao

import (
	"database/sql"
	"encoding/json"
)

type JsonNullInt64 struct {
	sql.NullInt64
}

type JsonNullBool struct {
	sql.NullBool
}

func (v JsonNullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullInt64) UnmarshalJSON(data []byte) error {
	var val *int64
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	if val != nil {
		v.Valid = true
		v.Int64 = *val
	} else {
		v.Valid = false
	}

	return nil
}

func (v JsonNullBool) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Bool)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullBool) UnmarshalJSON(data []byte) error {
	var val *bool
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	if val != nil {
		v.Valid = true
		v.Bool = *val
	} else {
		v.Valid = false
	}

	return nil
}
