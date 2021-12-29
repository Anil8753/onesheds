package interfaces

type HandlerDependency interface {
	GetDB() Database
}
