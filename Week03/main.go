package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HttpService struct {
	Server http.Server
}

func NewHttpSerivce(addr string) *HttpService {
	return &HttpService{
		Server: http.Server{
			Addr: addr,
		},
	}
}

func (service *HttpService) start() error {
	return service.Server.ListenAndServe()
}

func (service *HttpService) shutdown(ctx context.Context) error {
	return service.Server.Shutdown(ctx)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, _ := errgroup.WithContext(ctx)
	// start http1
	http1 := NewHttpSerivce(":8111")
	g.Go(func() error {
		if err := http1.start(); err != nil {
			cancel()
			return err
		}
		return nil
	})
	// start http2
	http2 := NewHttpSerivce(":8112")
	g.Go(func() error {
		if err := http2.start(); err != nil {
			cancel()
			return err
		}
		return nil
	})


	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		for {
			select {
			case s := <-c:
				switch s {
				case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
					cancel()
				default:
				}
			}
		}
	}()

	
	go func() {
		select {
		case <-ctx.Done():
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			go func() {
				if err := http1.shutdown(ctx); err != nil {
					log.Println("http1 shutdown err: ", err)
				}
			}()
			go func() {
				if err := http2.shutdown(ctx); err != nil {
					log.Println("http2 shutdown err: ", err)
				}
			}()
		}

	}()

	if err := g.Wait(); err != nil {
		log.Println("all server exit: ", err)
	}
}

