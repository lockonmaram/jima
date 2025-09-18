package helper

import "gorm.io/gorm/clause"

func FormatUpdatePayloadToClauseColumns(updatePayload map[string]any) (clauseColumns []clause.Column) {
	for column := range updatePayload {
		clauseColumns = append(clauseColumns, clause.Column{Name: column})
	}
	return
}
