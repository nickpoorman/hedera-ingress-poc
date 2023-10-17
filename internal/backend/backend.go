package backend

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nickpoorman/hedera-ingress/internal/business/mirror"
	"github.com/nickpoorman/hedera-ingress/internal/ticker"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type BackendID string

// A backend points to a given node.
type Backend struct {
	// backendID is the ID of the specific backend.
	backendID BackendID

	mu sync.RWMutex

	// A backend has many backend connections.
	conns []Connection
}

type Connection struct {
	proxyConn *grpc.ClientConn
}

type Config struct {
	Logger *zap.Logger
}

// Router represents the router of the backends.
// In order to route to a given backend, we need to know the backend ID.
type Router struct {
	Logger *zap.Logger

	ctx    context.Context
	quit   chan struct{}
	doneWg sync.WaitGroup

	mu       sync.RWMutex
	backends map[BackendID]*Backend
}

func NewRouter(config Config) *Router {
	ctx, cancel := context.WithCancel(context.Background())
	ps := &Router{
		Logger:   config.Logger.With(zap.String("backend-router", uuid.New().String())),
		backends: make(map[BackendID]*Backend),
		ctx:      ctx,
	}

	// Automatically close the context when stopping.
	ps.doneWg.Add(1)
	go func() {
		defer ps.doneWg.Done()
		select {
		case <-ctx.Done():
		case <-ps.quit:
			cancel()
		}
	}()

	ps.updateNodesBackground()
	return ps
}

func (s *Router) Close() {
	// Grab the lock so something doesn't try to update.
	s.mu.Lock()
	defer s.mu.Unlock()

	close(s.quit)
	s.doneWg.Wait()
}

func (s *Router) AddNode() {
	panic("not implemented")
}

func (s *Router) RemoveNode() {
	panic("not implemented")
}

// UpdateNodes fetches the address book from the mirror node and updates the nodes map. It starts a background goroutine that periodically updates the nodes map.
func (s *Router) UpdateNodes() error {
	return s.updateNodes()
}

// updateNodes blocks until the nodes have been updated or the context has been cancelled.
func (s *Router) updateNodes() error {
	addressBook, err := mirror.FetchAddressBook(s.ctx)
	if err != nil {
		return fmt.Errorf("update nodes: %w", err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// Update the nodes
	if err := s.upsertNodes(addressBook); err != nil {
		return fmt.Errorf("update nodes: %w", err)
	}

	return nil
}

// updateNodesBackground starts a background goroutine that periodically updates the nodes in the background.
// The background task is stopped when the Router is shut down.
func (s *Router) updateNodesBackground() {
	// Refresh the address book every 60 seconds.
	s.doneWg.Add(1)

	go func() {
		defer s.doneWg.Done()
		t := ticker.New(60 * time.Second)
		go func() {
			<-s.quit
			t.Stop()
		}()
		t.Loop(func() bool {
			addressBook, err := mirror.FetchAddressBook(s.ctx)
			if err != nil {
				// Check if err is a context Cancelled error.
				if errors.Is(err, context.Canceled) {
					return false
				}

				// Log the error
				s.Logger.Error("update nodes (background task)", zap.Error(err))
				return true
			}

			s.mu.Lock()
			defer s.mu.Unlock()

			// Update the nodes
			if err := s.upsertNodes(addressBook); err != nil {
				// Log the error
				s.Logger.Error("update nodes (background task)", zap.Error(err))
				return true
			}

			return true
		})
	}()
}

// upsertNodes takes a new address book and merges it with connections
// to the current nodes. This must be called with the lock held.
func (s *Router) upsertNodes(addressBook *mirror.AddressBook) error {
	panic("not implemented")
}

func (s *Router) Route(backendID BackendID) (Connection, error) {
	panic("not implemented")
}
