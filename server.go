package main

import (
	"apiserver/pkg/databases"
	"apiserver/pkg/loggers"
	"apiserver/pkg/queues"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	Router http.Handler
	Conn   *Connections
}

type Connections struct {
	DB    databases.DBInterface
	Queue queues.QueueInterface
}

func NewConnections(db databases.DBInterface, queue queues.QueueInterface) *Connections {
	return &Connections{DB: db, Queue: queue}
}

func NewGracefulServer(db databases.DBInterface, queue queues.QueueInterface) *Server {
	conn := NewConnections(db, queue)
	s := Server{Conn: conn}
	s.Router = NewRoutes(conn)
	return &s
}

func (s *Server) RunGracefulServer() {
	server := &http.Server{
		Addr:         os.Getenv("APP_ADDRESS"),
		Handler:      s.Router,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Run workers
	var workerWg sync.WaitGroup
	if err := s.Workers(serverCtx, &workerWg); err != nil {
		loggers.Logger.Fatal(err)
	}

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()
		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				loggers.Logger.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()
		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			loggers.Logger.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	fmt.Printf("Server started at port %s %s\n", server.Addr, time.Now().String())
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		loggers.Logger.Fatal(err)
	}

	workerWg.Wait()
	fmt.Println("All Workers closed.")

	// Wait for server context to be stopped
	<-serverCtx.Done()
	fmt.Printf("Server closed at %s\n", time.Now().String())

}

func (s *Server) Workers(ctx context.Context, wg *sync.WaitGroup) error {
	return nil
}
