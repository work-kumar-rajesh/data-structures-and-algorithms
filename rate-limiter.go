package main

import (
	"container/list"
)

type RateLimiter struct {
	freq   map[int]*list.List // UserID -> List of request times (as integers)
	limit  int                // Max requests allowed
	window int                // Time window (integers instead of timestamps)
}

// Constructor
func NewRateLimiter(limit int, window int) *RateLimiter {
	return &RateLimiter{
		freq:   make(map[int]*list.List),
		limit:  limit,
		window: window,
	}
}

// AllowRequest checks if a request can be made for the given userID at current time `t`
func (rl *RateLimiter) AllowRequest(userID int, t int) bool {
	// If user is not in the map, initialize their list
	if rl.freq[userID] == nil {
		rl.freq[userID] = list.New()
	}

	reqList := rl.freq[userID]

	// Remove expired timestamps (requests older than `window`)
	for reqList.Len() > 0 {
		front := reqList.Front()
		if t-front.Value.(int) > rl.window {
			reqList.Remove(front) // Remove outdated requests
		} else {
			break
		}
	}

	// Check if the user has exceeded request limit
	if reqList.Len() >= rl.limit {
		return false
	}

	// Add new request time
	reqList.PushBack(t)
	return true
}
