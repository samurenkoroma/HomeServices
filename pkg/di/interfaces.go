package di

type IStatRepo interface {
	AddClick(uint)
}

type CRUDRepository[T any] interface {
	Save(*T) error
	Get(uint) (*T, error)
	List(interface{}) ([]*T, error)
	Update(uint, *T) (bool, error)
	Delete(uint) error
}
