package main

import (
	"context"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/server"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/fatih/color"
	"github.com/pkg/profile"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

const profFlag = "--prof"
const pprofServerFlag = "--pprof-server"

const cpuProf = "cpu"
const memProf = "mem"
const blockingProf = "blocking"
const traceProf = "trace"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg := config.InitConfig()
	ctx := context.Background()
	l := logger.NewConsole()

	checkParams(l)

	srv := server.NewServer(cfg, l)
	srv.ServeHTTP(ctx)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	return nil
}

func checkParams(l *logger.Logger) {
	args := os.Args[1:]

	if len(args) > 0 {
		var doneDebugFlags bool
		for !doneDebugFlags && len(args) > 0 {
			switch args[0] {
			case profFlag:
				switch args[1] {
				case cpuProf:
					l.Println("cpu profiling enabled.\n")
					defer profile.Start(profile.CPUProfile, profile.NoShutdownHook).Stop()
				case memProf:
					l.Println("mem profiling enabled.\n")
					defer profile.Start(profile.MemProfile, profile.NoShutdownHook).Stop()
				case blockingProf:
					l.Println("block profiling enabled\n")
					defer profile.Start(profile.BlockProfile, profile.NoShutdownHook).Stop()
				case traceProf:
					l.Println("trace profiling enabled\n")
					defer profile.Start(profile.TraceProfile, profile.NoShutdownHook).Stop()
				default:
					panic("Unexpected prof flag: " + args[1])
				}
				args = args[2:]

			case pprofServerFlag:
				// serve the pprof endpoints setup in the init function run when "net/http/pprof" is imported
				go func() {
					cyanStar := color.CyanString("*")
					l.Print(cyanStar, "Starting pprof server on port 6060.")
					l.Print("Go to", "http://localhost:6060/debug/pprof in a browser to see supported endpoints.")

					err := http.ListenAndServe("0.0.0.0:6060", nil)

					if err != nil {
						l.Error().Err(err).Msg("pprof server exited with error:")
					}
				}()
				args = args[1:]
			default:
				doneDebugFlags = true
			}
		}
	}
}
