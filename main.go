package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/coreos/go-systemd/daemon"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	prom "github.com/prometheus/client_golang/prometheus"
)

const (
	appname = "typidor_exporter"
)

var (
	version    string
	iteration  string
	commit     string
	env        string
	appversion = appname + " [" + version + "] " + commit + " on " + iteration + " @ " + env

	listen string
	ver    bool
)

func init() {
	flag.StringVar(&listen, "bind", ":9824", "Bind the HTTP server")
	flag.BoolVar(&ver, "v", false, "Print version")
	flag.Parse()
}

func main() {
	if ver {
		fmt.Println(appversion)
		os.Exit(0)
	}

	metric := prom.NewGauge(
		prom.GaugeOpts{
			Namespace: "typidor",
			Name:      "status",
			Help:      "TY PIDOR status",
		})

	prom.MustRegister(metric)

	metric.Set(1)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	daemon.SdNotify(false, "READY=1")

	e.Logger.Fatal(e.Start(listen))
}
