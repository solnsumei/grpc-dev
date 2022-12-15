package tutorial

import context "context"

type Server struct {
	UnimplementedTutorialServer
}

func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello, " + in.GetName()}, nil
}

func (s *Server) TreasureChest(ctx context.Context, in *Offering) (*Chest, error) {
	switch {
	case in.Gift == "Penny" && in.Value <= 0:
		return &Chest{Success: false, Amount: 0}, nil
	case in.Gift == "Naira" && in.Value >= 1:
		return &Chest{Success: true, Amount: 1_000}, nil
	case in.Gift == "Dollars" && in.Value >= 10:
		return &Chest{Success: true, Amount: 1_00_000}, nil
	case in.Gift == "Pounds" && in.Value >= 100:
		return &Chest{Success: true, Amount: 1_000_000}, nil
	case in.Gift == "Euro" && in.Value >= 1_000:
		return &Chest{Success: true, Amount: 1_000_000_000}, nil
	default:
		return &Chest{Success: false, Amount: 0}, nil
	}
}
