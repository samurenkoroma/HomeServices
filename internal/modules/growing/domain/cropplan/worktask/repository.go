package worktask

type Repository interface {
	SaveMany(tasks []Task) error
}
