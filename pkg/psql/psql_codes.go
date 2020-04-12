package psql

import (
	"log"

	"github.com/jackc/pgx"
)

const (
	// NotNullViolation - not_null_violation
	NotNullViolation = "23502"
	// ForeignKeyViolation - foreign_key_violation
	ForeignKeyViolation = "23503"
	// UniqueViolation - unique_violation
	UniqueViolation = "23505"
	// CheckViolation - check_violation
	CheckViolation = "23514"
)

// Print - print psql error
func Print(e pgx.PgError) {
	log.Println(e.Code, e.ColumnName, e.ConstraintName,
		e.DataTypeName, e.Detail, e.File, e.Hint,
		e.InternalPosition, e.InternalQuery, e.Line,
		e.Message, e.Position, e.Routine, e.SchemaName,
		e.Severity, e.TableName, e.Where)

	// log.Print(e.Code)
	// log.Print(e.ColumnName)
	// log.Print(e.ConstraintName)
	// log.Print(e.DataTypeName)
	// log.Print(e.Detail)
	// log.Print(e.File)
	// log.Print(e.Hint)
	// log.Print(e.InternalPosition)
	// log.Print(e.InternalQuery)
	// log.Print(e.Line)
	// log.Print(e.Message)
	// log.Print(e.Position)
	// log.Print(e.Routine)
	// log.Print(e.SchemaName)
	// log.Print(e.Severity)
	// log.Print(e.TableName)
	// log.Print(e.Where)
}
