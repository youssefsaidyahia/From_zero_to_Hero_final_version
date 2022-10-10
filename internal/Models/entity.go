package Models

import "github.com/uptrace/bun"
import "github.com/google/uuid"

type Tranactions struct {
	bun.BaseModel `bun:"table:transactions"`
	Id            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Amount        int       `bun:"amount,notnull" validate:"required"`
	Currency      string    `bun:"currency,notnull" validate:"required"`
	Createdat     string    `bun:"createdat,notnull" validate:"required"`
	Status        bool      `bun:"status",default:"false"`
}
