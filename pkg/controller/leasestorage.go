package controller

import (
	"sync"
	"time"
)

// LeaseStorage stores user leases
type LeaseStorage interface {
	// GetLease returns a user's lease time for a specific tag, if present, nil otherwise
	GetLease(tag, user string) *time.Time
	// SetLease sets a user's lease time for a specific tag
	SetLease(tag, user string, expire time.Time)
	// NumActiveUsers returns the number of active users for a certain tag
	NumActiveUsers(tag string, now time.Time) int
}

// NewInMemoryLeaseStorage returns a lease storage that keeps the information in memory
func NewInMemoryLeaseStorage() LeaseStorage {
	return &inMemoryLeaseStorage{
		map[string]tagUsersLease{},
		sync.Mutex{},
	}
}

type inMemoryLeaseStorage struct {
	tags map[string]tagUsersLease
	mut  sync.Mutex
}

func (s *inMemoryLeaseStorage) GetLease(tag, user string) *time.Time {
	s.mut.Lock()
	defer s.mut.Unlock()

	if tagLease, ok := s.tags[tag]; !ok {
		return tagLease.GetLease(user)
	}
	return nil
}

func (s *inMemoryLeaseStorage) SetLease(tag, user string, expire time.Time) {
	s.mut.Lock()
	defer s.mut.Unlock()

	if _, ok := s.tags[tag]; !ok {
		s.tags[tag] = tagUsersLease{}
	}
	s.tags[tag][user] = expire
}

func (s *inMemoryLeaseStorage) NumActiveUsers(tag string, now time.Time) int {
	s.mut.Lock()
	defer s.mut.Unlock()

	if tl, ok := s.tags[tag]; ok {
		return tl.NumActiveUsers(now)
	}
	return 0
}

// tagUsersLease maps users to expiration times
type tagUsersLease map[string]time.Time

func (tl tagUsersLease) GetLease(user string) *time.Time {
	if lease, ok := tl[user]; ok {
		return &lease
	}
	return nil
}

func (tl tagUsersLease) NumActiveUsers(now time.Time) int {
	for user, expTime := range tl {
		if expTime.Before(now) {
			delete(tl, user)
		}
	}
	return len(tl)
}
