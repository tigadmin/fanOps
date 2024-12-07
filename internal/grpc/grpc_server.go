package grpc

import (
	"context"
	"github.com/your-username/fanOps/internal/api"
	"github.com/your-username/fanOps/proto/sportsdata"
	"log"
	"net"
)

type Server struct {
	sportsdata.UnimplementedSportsServiceServer
}

func (s *Server) FetchTeams(ctx context.Context, req *sportsdata.Empty) (*sportsdata.TeamList, error) {
	// Fetch teams from the API and store in the database
	teams, err := api.FetchTeams()
	if err != nil {
		return nil, err
	}

	// Prepare the list of teams for the gRPC response
	var teamList sportsdata.TeamList
	for _, team := range teams {
		teamList.Teams = append(teamList.Teams, &sportsdata.Team{
			Id:           int32(team.ID),
			Name:         team.Name,
			City:         team.City,
			Abbreviation: team.Abbreviation,
			Conference:   team.Conference,
			Division:     team.Division,
		})
	}

	return &teamList, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	sportsdata.RegisterSportsServiceServer(grpcServer, &Server{})
	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
