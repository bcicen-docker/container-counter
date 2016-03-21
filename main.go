package main

import (
	"log"

	"github.com/cyberdelia/statsd"
	"github.com/fsouza/go-dockerclient"
	"github.com/kelseyhightower/envconfig"
)

type Options struct {
	Statsd string `short:"s" long:"statsd" description:"Statsd Server url. e.g. 'localhost:8125'"`
}

var (
	opts Options
)

func assert(err error, context string) {
	if err != nil {
		log.Fatal(context+": ", err)
	}
}

func main() {
	err := envconfig.Process("COUNTER", &opts)
	if err != nil {
		log.Fatalf("Error parsing ENV vars %s", err)
	}

	if opts.Statsd == "" {
		log.Fatalf("ENV var STATSD not set")
	}

	//setup docker client
	client, err := docker.NewClient("unix:///var/run/docker.sock")
	assert(err, "docker")

	//setup statsd client
	statsdclient, err := statsd.Dial(opts.Statsd)
	assert(err, "statsd")
	defer statsdclient.Close()

	events := make(chan *docker.APIEvents)
	assert(client.AddEventListener(events), "attacher")
	log.Println("listening for events")
	log.Println("sending events to:", opts.Statsd)
	for msg := range events {
		if msg.Status == "create" {
			log.Println("create:", msg.ID[:12], msg.Status)
			err = statsdclient.IncrementGauge("dockercount.statsd.gauge", 1, 1)
			assert(err, "statsd")
			err = statsdclient.Flush()
			assert(err, "statsd")
		}
	}
}
