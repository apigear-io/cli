package simjs

type Service struct {
	conn *Conn
	univ *Universe
}

func NewService(conn *Conn, univ *Universe) *Service {
	s := &Service{
		conn: conn,
		univ: univ,
	}
	conn.SetHandler(s.HandleMessage)
	return s
}

func (s *Service) HandleMessage(msg *SimuMessage) *SimuMessage {
	if msg.WorldID == "" {
		return s.univ.HandleMessage(msg)
	}
	w := s.univ.GetWorld(msg.WorldID)
	if w == nil {
		return nil
	}
	if msg.ActorID == "" {
		return w.HandleMessage(msg)
	}
	a := w.GetActor(msg.ActorID)
	if a == nil {
		return nil
	}
	return a.HandleMessage(msg)
}
