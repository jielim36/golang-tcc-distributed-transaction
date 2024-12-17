package models

type TCCStatus string

const (
	TCC_PENDING TCCStatus = "PENDING"
	TCC_TRYING  TCCStatus = "TRYING"
	TCC_CONFIRM TCCStatus = "CONFIRM"
	TCC_CANCEL  TCCStatus = "CANCEL"
)

func (s TCCStatus) String() string {
	return string(s)
}
