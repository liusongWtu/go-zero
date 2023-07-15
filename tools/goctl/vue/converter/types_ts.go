package converter

var CommonGoDataTypeMapTs = map[string]string{
	// For consistency, all integer types are converted to int64
	// number
	"int64":          "number",
	"float64":        "number",
	"string":         "string",
	"bool":           "boolean",
	"sql.NullString": "string",
	// "byte"
	// "time.Time"
}
