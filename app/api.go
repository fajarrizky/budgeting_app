package app

import (
	"budgetapp/module"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartApi() {

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	apiMod := module.NewApiModule(serverCtx)

	mux := apiMod.GetMux()

	configService := apiMod.GetConfigService()

	server := &http.Server{Addr: ":" + configService.GetServerPort(), Handler: mux}

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx,
			time.Duration(configService.GetServerShutdownGracePeriod()*int(time.Second)))
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}

		apiMod.CleanUp()
		serverStopCtx()
	}()

	// Run the server
	fmt.Printf("server starting on port: %v\n", configService.GetServerPort())
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()

}
