package store

type Store struct {
	c Client
}

////go:generate mockgen -source=store.go -destination=../../test/mocks/store/source/client.go -package=mock_client
type Client interface {
	DarkMagic(string) string
}

func NewServer(c Client) *Store {
	return &Store{
		c: c,
	}
}

func (s *Store) Verify(something string) string {
	a := s.c.DarkMagic(something)
	switch a {
	case "Owl":
		return "Hello Mr. Porter. Welcome!"
	case "Mouse":
		return "You're in the wrong place."
	default:
		return "Unknown creature. Avada Kedavra!"
	}
}
