package tool

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func ResultQuery(rows *sqlx.Rows) (result interface{}, err error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	// we"ll want to end up with a list of name->value maps, a la JSON
	// surely we know how many rows we got but can"t find it now
	allgeneric := make([]map[string]interface{}, 0)
	// we"ll need to pass an interface to sql.Row.Scan
	colvals := make([]interface{}, len(cols))
	for rows.Next() {
		colassoc := make(map[string]interface{}, len(cols))
		// values we"ll be passing will be pointers, themselves to interfaces
		for i, key := range colvals {
			_ = key
			colvals[i] = new(interface{})
		}
		if err := rows.Scan(colvals...); err != nil {
			return nil, err
		}
		for i, col := range cols {

			vallue := *colvals[i].(*interface{})

			colassoc[col] = vallue
			sVal := fmt.Sprintf("%s", colassoc[col])
			dataType := fmt.Sprintf("%T", colassoc[col])
			if dataType == "[]uint8" {
				colassoc[col] = sVal
			}
			fmt.Printf("%s: %T %s\n", col, colassoc[col], colassoc[col])
		}
		fmt.Println(colassoc)
		allgeneric = append(allgeneric, colassoc)
	}
	return allgeneric, nil
}
