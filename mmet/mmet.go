package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

// NodeSessions holds the session counts for each node.
type NodeSessions struct {
	mu     sync.Mutex // Ensure thread-safe access
	Counts map[string]int
}

func (ns *NodeSessions) String() string {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	cities := make([]string, 0, len(ns.Counts))
	for city := range ns.Counts {
		cities = append(cities, city)
	}

	parts := make([]string, 0, len(cities))
	for _, city := range cities {
		parts = append(parts, fmt.Sprintf("%s: %d", city, ns.Counts[city]))
	}

	return "{" + strings.Join(parts, ", ") + "}"
}

// MoveCommand represents a request to move sessions between nodes.
type MoveCommand struct {
	Count int    `json:"count"`
	From  string `json:"from"`
	To    string `json:"to"`
}

func (ns *NodeSessions) ApplyRandomChange() {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	for city := range ns.Counts {
		change := rand.Intn(5) - 2 // Random number between -2 and 2
		newCount := ns.Counts[city] + change
		if newCount < 0 {
			newCount = 0 // Ensure the count never goes below zero
		}
		ns.Counts[city] = newCount
	}
}

func (ns *NodeSessions) MoveHandler(w http.ResponseWriter, r *http.Request) {
	var cmd MoveCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ns.mu.Lock()
	defer ns.mu.Unlock()

	// Apply the move command if valid
	if _, ok := ns.Counts[cmd.From]; ok && ns.Counts[cmd.From] >= cmd.Count {
		ns.Counts[cmd.From] -= cmd.Count
		ns.Counts[cmd.To] += cmd.Count
		log.Printf("Got move command: %+v", cmd)
	} else {
		http.Error(w, "Invalid move command", http.StatusBadRequest)
	}
}

func main() {
	time.Now().UnixNano()

	sessions := NodeSessions{
		Counts: map[string]int{"Gdansk": 10, "Poznan": 12, "Warsaw": 18, "Krakow": 6},
	}

	// HTTP Server for move commands
	http.HandleFunc("/api/move", sessions.MoveHandler)
	go func() {
		if err := http.ListenAndServe(":4545", nil); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Main loop
	ticker := time.NewTicker(1 * time.Second * 5)
	round := 1
	for ; true; <-ticker.C {
		sessions.ApplyRandomChange()
		// Here, you'd also send the JSON to the configured endpoint using an HTTP client.
		log.Printf("Round %d: %+v", round, sessions.String())
		round++
	}
}
