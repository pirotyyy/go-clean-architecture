package main

import (
	adaptorHTTP "ca-tech/adaptor/http"
	"ca-tech/config"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()

	addr := config.LoadConfig().HTTPInfo.Addr
	router := adaptorHTTP.InitRouter()
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		log.Println("start server at :", addr)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("server exiting")
}
