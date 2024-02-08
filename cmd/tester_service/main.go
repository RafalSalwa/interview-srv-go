package main

import (
	"github.com/RafalSalwa/interview-app-srv/cmd/tester_service/internal/workers"
)

func main() {
	worker := workers.NewWorker("daisy_chain")
	worker.Run()
}
