package invoiceheader

import "time"

type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	updateAt  time.Time
}
