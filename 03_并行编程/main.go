package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	parentCtx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(parentCtx)

	// 端口监听1
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	server1 := http.Server{
		Addr:    ":8090",
		Handler: mux1,
	}
	group.Go(func() error {
		<-ctx.Done()
		return server1.Shutdown(context.Background())
	})
	group.Go(func() error {
		return server1.ListenAndServe()
	})


	// 端口监听2
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, Debug!")
	})
	server2 := http.Server{
		Addr:    ":8091",
		Handler: mux2,
	}
	group.Go(func() error {
		<-ctx.Done()
		return server2.Shutdown(context.Background())
	})
	group.Go(func() error {
		return server2.ListenAndServe()
	})

	// 中断监听
	c := make(chan os.Signal, 1)
	signal.Notify(c, []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}...)
	group.Go(func() error {
		for  {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				cancel()
			}
		}
	})
	if err := group.Wait();err != nil && !errors.Is(err, context.Canceled) {
		log.Fatal(err)
	}
}
