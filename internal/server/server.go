package server

type Server struct {
	magicClient downstreamClient
}

//go:generate mockgen -source=server.go -destination=./mocks/client.go -package=mock_client
type downstreamClient interface {
	Magic(string) (string, error)
}

func NewServer(c downstreamClient) *Server {
	return &Server{
		magicClient: c,
	}
}

func (s *Server) OnFire(something string) (string, error) {
	f, err := s.magicClient.Magic(something)
	if err != nil {
		return "On fire", err
	}
	switch f {
	case "Queen of Hearts":
		return "How do you know I picked that?!", nil
	case "King of Spades":
		return "No, you're wrong!", nil
	default:
		return "It's on fire but it's fine.", nil
	}
}
