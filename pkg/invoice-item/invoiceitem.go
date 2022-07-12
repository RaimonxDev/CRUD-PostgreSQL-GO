package invoice_item

import "time"

type Model struct {
	ID              uint
	InvoiceHeardeID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
