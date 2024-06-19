package pgx

import "reflect"

func parseColumnsForScan(sub interface{}, addColumns ...interface{}) []interface{} {
	s := reflect.ValueOf(sub).Elem()
	numCols := s.NumField()

	columns := []interface{}{}
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns = append(columns, field.Addr().Interface())
	}
	columns = append(columns, addColumns...)
	return columns
}
