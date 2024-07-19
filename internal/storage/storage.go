package storage

type Storage interface {
	CreateUser(tgID int64) error
	SetPaylaod(tgID int64, payload string) error
	UpdateMessageStatus(tgID int64, status string) error
	GetMessageStatus(tgID int64) (string, error)
	GetPayload(tgID int64) (string, error)
	SetOperations(tgID int64, operations int) error
	GetOperations(tgID int64) (int, error)
	SetFileName(tgID int64, fileName string) error
	GetFileName(tgID int64) (string, error)
}
