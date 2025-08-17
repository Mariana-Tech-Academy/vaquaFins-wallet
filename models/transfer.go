package models

type Transfer struct { // a transfer is always from a senders point of view
	RecipientAccountNumber uint
	Amount                 uint
	Description            string
}
