package interfaces

type SqlScan interface {
	Scan(...interface{}) error
}