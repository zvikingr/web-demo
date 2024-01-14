package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"web-demo/config"
	"web-demo/controller/router"
	"web-demo/dao"
	"web-demo/utils/log"
)

func main() {
	var configFile string

	flag.StringVar(&configFile, "f", "service.conf", "config file")
	flag.Parse()

	cfg, err := config.DecodeConfig(configFile)
	if err != nil {
		fmt.Printf("config decode error, configFile=%v, err=%v", configFile, err)
		return
	}

	log.InitLog(cfg.LogConfig)

	if err := dao.InitDatabase(cfg.DBConfig); err != nil {
		log.Errorf("db init error:%v", err)
		return
	}

	r, err := router.InitRouter(cfg.Env)
	if err != nil {
		log.Errorf("router init error:%v", err)
		return
	}

	if err := startService(r, cfg); err != nil {
		log.Errorf("service error:%v", err)
		return
	}

	log.Infof("service quit.")
}

func startService(r http.Handler, cfg *config.Config) error {
	srv := &http.Server{
		Addr:    cfg.ServiceAddr,
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Infof("service start, listen on:%s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Errorf("listenAndServe error, err:%v", err)
			quit <- syscall.SIGINT
		}
	}()

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	return srv.Shutdown(ctx)
}
