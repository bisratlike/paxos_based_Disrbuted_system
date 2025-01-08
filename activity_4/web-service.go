package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "sync"
    "activity_4/paxos"
)

var (
    acceptors = []*paxos.Acceptor{{}, {}, {}}
    mu        sync.Mutex
)

func proposeHandler(w http.ResponseWriter, r *http.Request) {
    var body struct {
        ProposalNumber int
        Value          string
    }
    json.NewDecoder(r.Body).Decode(&body)

    proposer := paxos.Proposer{ProposalNumber: body.ProposalNumber, Value: body.Value}
    mu.Lock()
    value := proposer.Propose(body.Value, acceptors)
    mu.Unlock()

    if value != nil {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "Consensus reached: %s\n", value)
    } else {
        w.WriteHeader(http.StatusConflict)
        fmt.Fprintf(w, "Consensus not reached\n")
    }
}

func main() {
    http.HandleFunc("/propose", proposeHandler)
	fmt.Println("Web service running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
