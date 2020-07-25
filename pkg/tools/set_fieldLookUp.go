package tool

import (
	"fmt"
	"strings"
)

func SetFieldListLookup(SourceField string, DisplayField string) string {
	var (
		result       = ""
		FieldLookUps []string
		Display      []string
	)
	FieldLookUps = strings.Split(SourceField, ",")
	Display = strings.Split(DisplayField, ",")
	DisplayLookUp := "concat("
	fieldLookUp := ""
	for i, key := range Display {
		_ = key
		fieldLookUp += fmt.Sprintf(",'|',%s::varchar", Display[i])
	}
	fieldLookUp = fieldLookUp[strings.Index(fieldLookUp, ",'|',")+5:]
	fmt.Printf(fieldLookUp)
	DisplayLookUp += fieldLookUp + ")"

	for x, key := range FieldLookUps {
		_ = key
		if x == 0 {
			result += fmt.Sprintf(",%s as value", FieldLookUps[x])
		} else if x == 1 {
			result += fmt.Sprintf(",%s as label ", DisplayLookUp)
			result += fmt.Sprintf(",%s", FieldLookUps[x])
		} else {
			result += fmt.Sprintf(",%s", FieldLookUps[x])
		}
	}
	result = result[strings.Index(result, ",")+1:]

	return result

}
