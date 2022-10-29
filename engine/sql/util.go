package sql

import (
	stdSql "database/sql"
	"fmt"
)

func StringToNullableString(v interface{}) stdSql.NullString {

	switch t := v.(type) {
	case string:
		return stdSql.NullString{
			String: fmt.Sprint(t),
			Valid:  true,
		}
	case []byte:
		return stdSql.NullString{
			String: fmt.Sprint(t),
			Valid:  true,
		}
	case nil:
		return stdSql.NullString{
			Valid: false,
		}
	default:
		return stdSql.NullString{
			Valid: false,
		}

	}

}
