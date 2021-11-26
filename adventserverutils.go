package main

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/net/context"

	fcpb "github.com/brotherlogic/filecopier/proto"
)

func (s *Server) loadFile(ctx context.Context, path string) (string, error) {
	err := s.replicate(ctx, path)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (s *Server) replicate(ctx context.Context, path string) error {
	servers, err := s.FFind(ctx, "filecopier")
	if err != nil {
		return err
	}

	for _, se := range servers {
		conn, err := s.FDial(se)
		if err != nil {
			return err
		}
		defer conn.Close()

		client := fcpb.NewFileCopierServiceClient(conn)
		res, err := client.Exists(ctx, &fcpb.ExistsRequest{Path: path})
		if err != nil {
			return err
		}
		if res.GetExists() {
			_, err := client.Replicate(ctx, &fcpb.ReplicateRequest{Path: path})
			return err
		}
	}

	return fmt.Errorf("Cannot locate %v", path)
}
