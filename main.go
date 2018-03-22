package main

import (
		"context"
		"crypto/tls"
		"errors"
		"net"
		"net/http"
		"net/url"
		"os"
		"strings"
		"sync/errgroup"
		
		"google.golang.org/grpc"
		"google.golang.org/grpc/keepalive"
		"google.golang.org/grpc/metadata"

		"golang.org/x/crypto/acme/autocert"
		"github.com/cncd/logging"
		"github.com/marjoram/pipelinne/pipeline/rpc/proto"
		"github.com/marjoram/pubsub"
)

func main() {
		ctx := context.Background()
		
		var g errgroup.Group
		
		// start grpc server
		g.Go(func() error {
				lis, err := net.Listen("tcp", ":9000")
				if err != nil {
						return err
				}
				grpc := grpc.NewServer()
				piped := new(pipeline.Server)
				proto.RegisterServer(grpc, piped)
				
				err = grpc.Serve(lis)
				if err != nil {
						return err
				}
				return nil
		})
		
		address, err := c.String("host")
		if err != nil {
				return err
		}
		
		// tmp dir for letsencrypt
		dir := cacheDir()
		os.MkdirAll(dir, 0700)
		
		// manager enables lets encrypt
		manager := &autocert.Manager{
				Prompt:		autocert.AcceptTOS,
				HostPolicy:	autocert.HostWhitelist(address.Host),
				Cache:		autocert.DirCache(dir),
		}
		
		// and now we serve
		g.Go(func() error {
				return http.ListenAndServe(":http", manager.HTTPHandler(http.HandlerFunc(redirect)))
		})
		g.Go(func() error {
				server := &http.Server{
						Addr:		":https",
						Handler:	handler,
						TLSConfig:	&tls.Config{
								GetCertificate:	manager.GetCertificate,
								NextProtos:		[]string{"http/2"},
						},
				}
				return serve.ListenAndServeTLS("", "")
		})
		
		return g.Wait()
}

func redirect(w http.ResponseWriter, req *http.Request) {
	var serverHost string = s.Host
	serverHost = strings.TrimPrefix(serverHost, "http://")
	serverHost = strings.TrimPrefix(serverHost, "https://")
	req.URL.Scheme = "https"
	req.URL.Host = serverHost

	w.Header().Set("Strict-Transport-Security", "max-age=31536000")

	http.Redirect(w, req, req.URL.String(), http.StatusMovedPermanently)
}

func cacheDir() string {
	const base = "golang-autocert"
	if xdg := os.Getenv("XDG_CACHE_HOME"); xdg != "" {
		return filepath.Join(xdg, base)
	}
	return filepath.Join(os.Getenv("HOME"), ".cache", base)
}
		
