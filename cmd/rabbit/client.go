package main

import (
	"context"
	"fmt"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/RafalSalwa/interview-app-srv/config"
	amqpHandlers "github.com/RafalSalwa/interview-app-srv/internal/amqp/handlers"
	intrvamqp "github.com/RafalSalwa/interview-app-srv/pkg/amqp"
)

func main() {
	c := config.New()
	ctx, rejectContext := context.WithCancel(NewContextCancellableByOsSignals(context.Background()))

	connection := intrvamqp.NewConnectionFromCredentials(c.AMQP)
	if err := connection.Connect(); err != nil {
		fmt.Println(ctx, err)
		os.Exit(2)
	}
	defer connection.Close()

	var client intrvamqp.Client = intrvamqp.NewClient(connection)
	client.SetHandler("subscription_create", amqpHandlers.WrapHandleUserSubscription)
	client.SetHandler("password_changed", amqpHandlers.WrapHandleUserPasswordChange)
	client.SetHandler("account_deletion", amqpHandlers.WrapHandleUserAccountDeletion)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		_ = client.HandleChannel(ctx, "interview", false)
		rejectContext()
	}()

	wg.Wait()
}

func NewContextCancellableByOsSignals(parent context.Context) context.Context {
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	newCtx, cancel := context.WithCancel(parent)

	go func() {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			fmt.Println("Received Interrupt signal")
			cancel()
		case syscall.SIGTERM:
			fmt.Println("Received sigterm signal")
			cancel()
		}
	}()

	return newCtx
}
