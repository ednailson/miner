package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/ednailson/miner/internal/domain/entity"
	"github.com/ednailson/miner/internal/domain/usecase"
	"go.uber.org/zap"
	"io"
	"net"
)

type Server struct {
	addr     string
	listener net.Listener
	uc       usecase.Miner
}

func NewServer(addr string, uc usecase.Miner) *Server {
	return &Server{addr: addr, uc: uc}
}

func (s *Server) Start() error {
	var err error
	s.listener, err = net.Listen("tcp", fmt.Sprintf("%s", s.addr))
	if err != nil {
		zap.S().Error("server: failed to start connection")
		return err
	}

	zap.S().Info("server: it has been started")

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			zap.S().Errorf("server: failed to accept a connection, error: %s", err.Error())
			continue
		}

		zap.S().Info("server: new connection accepted")

		go func(conn net.Conn) {
			reader := bufio.NewReader(conn)

			for {
				line, _, err := reader.ReadLine()
				if err != nil {
					if err == io.EOF {
						zap.S().Info("server: a connection was lost")
						return
					}
					zap.S().Errorf("server: failed to read line, error: %s", err.Error())
					continue
				}

				var req entity.Request

				err = json.Unmarshal(line, &req)
				if err != nil {
					sendResponse(
						conn,
						entity.NewFail(nil, entity.ErrorParse()),
					)
					continue
				}

				err = req.Validate()
				if err != nil {
					sendResponse(
						conn,
						entity.NewFail(nil, entity.ErrorInvalidRequest()),
					)
					continue
				}

				sendResponse(
					conn,
					s.uc.Save(req),
				)
			}
		}(conn)
	}
}

func (s *Server) Close() {
	if err := s.listener.Close(); err != nil {
		zap.S().Errorf("server: failed to close server connection")
	}
}

func sendResponse(conn net.Conn, resp interface{}) {
	body, _ := json.Marshal(resp)

	_, err := conn.Write(body)
	if err != nil {
		zap.S().Errorf("server: failed to write response, error: %s", err.Error())
	}
}
