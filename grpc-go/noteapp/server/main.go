package main

import (
	"context"

	pb "note/note/proto"

	"github.com/google/uuid"
)

type server struct {
	pb.UnimplementedNoteServiceServer
	notes map[string]*pb.GetNoteResponse
}

func (s *server) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	id := uuid.New().String()
	s.notes[id] = &pb.GetNoteResponse{
		Title:   req.Title,
		Content: req.Content,
	}
	return &pb.CreateNoteResponse{Id: id}, nil
}
