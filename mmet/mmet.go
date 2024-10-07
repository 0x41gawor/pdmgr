package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sort"
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

	// Collect cities into a slice
	cities := make([]string, 0, len(ns.Counts))
	for city := range ns.Counts {
		cities = append(cities, city)
	}

	// Sort the cities alphabetically
	sort.Strings(cities)

	// Create parts of the final string
	parts := make([]string, 0, len(cities))
	for _, city := range cities {
		parts = append(parts, fmt.Sprintf("%s: %2d", city, ns.Counts[city]))
	}

	return "{" + strings.Join(parts, ", ") + "}"
}

// MoveCommand represents a request to move sessions between nodes.
type MoveCommand struct {
	Count int    `json:"count"`
	From  string `json:"from"`
	To    string `json:"to"`
}

func (mc *MoveCommand) String() string {
	return "{From: " + mc.From + ", To: " + mc.To + ", Count: " + fmt.Sprintf("%d", mc.Count) + "}"
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

	// Apply the move command if valid
	if _, ok := ns.Counts[cmd.From]; ok && ns.Counts[cmd.From] >= cmd.Count {
		ns.mu.Lock()
		ns.Counts[cmd.From] -= cmd.Count
		ns.Counts[cmd.To] += cmd.Count
		ns.mu.Unlock()
		log.Printf("Got move command: %+v", cmd.String())
		// log.Printf("Round %d:", round)
		log.Printf("Round %d: %+v", round, ns.String())
	} else {
		http.Error(w, "Invalid move command", http.StatusBadRequest)
	}
}

var round = 0

func main() {
	// Command-line flag to get the interval value
	interval := flag.Int("interval", 5, "Interval in seconds for applying random changes")
	flag.Parse()
	log.Printf("Round interval time set as: %d seconds", *interval)

	time.Now().UnixNano()

	sessions := NodeSessions{
		Counts: map[string]int{"Gdansk": 10, "Poznan": 12, "Warsaw": 18, "Krakow": 6},
	}

	// HTTP Server for move commands
	http.HandleFunc("/api/move", sessions.MoveHandler)
	go func() {
		if err := http.ListenAndServe(":4040", nil); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Main loop
	ticker := time.NewTicker(1 * time.Second * time.Duration(*interval))
	for ; true; <-ticker.C {
		sessions.ApplyRandomChange()
		// Here, you'd also send the JSON to the configured endpoint using an HTTP client.
		round++
		log.Printf("%-10s %s", fmt.Sprintf("Round %d:", round), sessions.String())
	}
}
