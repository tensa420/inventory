package main

import (
	"context"
	in "inventory/pkg/pb"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type InventoryServer struct {
	in.UnimplementedInventoryServiceServer
	mu      sync.Mutex
	details map[string]*in.Part
}

func (p *InventoryServer) GetPart(_ context.Context, req *in.GetPartRequest) (*in.GetPartResponse, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	id := req.Uuid

	part, ok := p.details[id]
	if !ok {
		return nil, status.Error(codes.NotFound, "ID not found")
	}

	return &in.GetPartResponse{
		Part: part,
	}, nil
}

func (p *InventoryServer) ListParts(_ context.Context, req *in.ListPartsRequest) (*in.ListPartsResponse, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	parts := make([]*in.Part, 0)

	uuids := req.Filter.Uuids
	names := req.Filter.Names
	categories := req.Filter.Categories
	manufacturers := req.Filter.ManufacturerCountries
	tags := req.Filter.Tags

	if len(uuids) > 0 {
		for _, v := range p.details {
			for uid := 0; uid < len(uuids); uid++ {
				if v.UUID == uuids[uid] {
					parts = append(parts, v)
				}
			}
		}
	}
	if len(names) > 0 {
		for _, v := range p.details {
			for name := 0; name < len(names); name++ {
				if v.Name == names[name] {
					parts = append(parts, v)
				}
			}
		}
	}
	if len(categories) > 0 {
		for _, v := range p.details {
			for cat := 0; cat < len(categories); cat++ {
				if v.Category == categories[cat] {
					parts = append(parts, v)
				}
			}
		}
	}
	if len(manufacturers) > 0 {
		for _, v := range p.details {
			for countr := 0; countr < len(manufacturers); countr++ {
				if v.Manufacturer.Country == manufacturers[countr] {
					parts = append(parts, v)
				}
			}
		}
	}
	if len(tags) > 0 {
		for _, v := range p.details {
			for tagsInStruct := 0; tagsInStruct < len(v.Tags); tagsInStruct++ {
				for tag := 0; tag < len(tags); tag++ {
					if v.Tags[tagsInStruct] == tags[tag] {
						parts = append(parts, v)
					}
				}
			}
		}
	}

	return &in.ListPartsResponse{
		Parts: parts,
	}, nil
}
func main() {

	lis, err := net.Listen("TCP", ":50052")
	if err != nil {
		log.Printf("Failed to listen: %v", err)
	}
	defer func() {
		err = lis.Close()
		if err != nil {
			log.Printf("Failed to close listener: %v", err)
		}
	}()
	s := grpc.NewServer()
	in.RegisterInventoryServiceServer(s, &InventoryServer{})

	reflection.Register(s)

	go func() {
		err := s.Serve(lis)
		if err != nil {
			log.Printf("Failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 0)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	s.GracefulStop()
	log.Println("Server was stopped")
}
