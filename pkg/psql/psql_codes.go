/*
 * Copyright (C) 2020. Leonid Moguchev
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
