package store

////go:generate mockgen -source=villain.go -destination=../../test/mocks/store/source/villain.go -package=mock_client
type Villain interface {
	Lose() bool
}
