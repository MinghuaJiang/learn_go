package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("Pool has been closed")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Fetch a resource from pool or create a new resource if pool is empty
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire: Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}

		return r, nil

	default:
		log.Println("Acquire: New Resource")
		return p.factory()
	}
}

// Here we have a set of number of go routine which will both be sender and receiver on a buffered channel and there are dropping of workitems
// Add the resource to pool when release the resource or close the resource if pool is full
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	// Just discard the resource is pool has been closed
	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Release: In Queue")
	default:
		log.Println("Release: Closing")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}

	p.closed = true
	// first make sure no write to the channel anymore
	close(p.resources)

	// drain the flush the channel
	for r := range p.resources {
		r.Close()
	}
}
