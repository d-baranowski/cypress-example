package model

import (
	"github.com/uptrace/bun"
	"time"
)

// Todo represents a task in the form
type Todo struct {
	bun.BaseModel `bun:"table:todos"` // Specify the table name for Todos

	ID          int64  `bun:",pk,autoincrement" json:"id"`
	Task        string `bun:"task,notnull" json:"task"`
	FormModelID int64  `bun:"form_model_id,notnull" json:"form_model_id"` // Foreign key to FormModel
}

// FormModel represents the form model structure
type FormModel struct {
	bun.BaseModel `bun:"table:form_models"` // Specify the table name for FormModel

	ID                int64     `bun:",pk,autoincrement" json:"id"`
	Name              string    `bun:"name,notnull" json:"name"`
	Age               int32     `bun:"age,notnull" json:"age"`
	DateTime          time.Time `bun:"date_time,notnull" json:"date_time"` // ISO 8601 formatted DateTime
	Skills            []string  `bun:"skills,array" json:"skills"`         // Array of skills
	City              string    `bun:"city,notnull" json:"city"`
	Range             []int32   `bun:"range,array" json:"range"` // Integer range (min/max)
	TermsAccepted     bool      `bun:"terms_accepted,notnull" json:"terms_accepted"`
	SatisfactionLevel int32     `bun:"satisfaction_level,notnull" json:"satisfaction_level"`
	Gender            string    `bun:"gender,notnull" json:"gender"` // Enum "male", "female", "other"

	// Relationship to Todo (One FormModel can have many Todos)
	Todos []Todo `bun:"rel:has-many,join:id=form_model_id" json:"todos"`
}
