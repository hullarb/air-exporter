package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.einride.tech/can/pkg/candevice"
	"go.einride.tech/can/pkg/socketcan"
)

func main() {
	log.SetFlags(log.Lmicroseconds | log.LstdFlags)
	setup := flag.Bool("setup", false, "configure and set can device up - needs permission")
	td := flag.Bool("setdwon", false, "set can device down - needs permission")
	dev := flag.String("d", "can0", "name of the can device")
	addr := flag.String("l", ":2112", "address to listen on")
	flag.Parse()

	if *setup {
		if err := start(*dev); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
	if *td {
		if err := stop(*dev); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	metrics := registerMetrics()

	http.Handle("/metrics", promhttp.Handler())
	go func() { log.Fatal(http.ListenAndServe(*addr, nil)) }()

	log.Printf("exporter started")
	conn, err := socketcan.DialContext(context.Background(), "can", *dev)
	if err != nil {
		log.Fatal(err)
	}
	recv := socketcan.NewReceiver(conn)
	for recv.Receive() {
		frame := recv.Frame()
		if f, ok := frameIDMap[frame.ID]; ok {
			metrics[frame.ID].Set(f.dataConverter(frame.Data))
		} else {
			log.Printf("unknown frame: %s", frame.String())
		}
	}

}

func registerMetrics() map[uint32]prometheus.Gauge {
	metrics := map[uint32]prometheus.Gauge{}
	for id, m := range frameIDMap {
		if strings.HasPrefix(m.name, "0") || strings.HasPrefix(m.name, "1") {
			m.name = "unkn_" + m.name
		}
		metrics[id] = promauto.NewGauge(prometheus.GaugeOpts{
			Name: m.name,
			Help: fmt.Sprintf("%s unit: %s", m.name, m.unit),
		})
	}
	return metrics
}

func start(dev string) error {
	d, err := candevice.New(dev)
	if err != nil {
		return fmt.Errorf("failed to create new device: %v", err)
	}
	err = d.SetBitrate(50000)
	if err != nil {
		return fmt.Errorf("failed to set bitrate: %v", err)
	}
	err = d.SetUp()
	if err != nil {
		return fmt.Errorf("failed to start the can device: %v", err)
	}
	return nil
}

func stop(dev string) error {
	d, err := candevice.New(dev)
	if err != nil {
		return fmt.Errorf("failed to create new device: %v", err)
	}
	err = d.SetDown()
	if err != nil {
		return fmt.Errorf("failed to stop the can device: %v", err)
	}
	return nil
}
