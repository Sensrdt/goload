package main

import (
	"log"
	"time"
	"github.com/go-co-op/gocron"
)
func startHealthCheck() {
	s := gocron.NewScheduler(time.Local)
	for _, host := range serverList {
		log.Printf("%s", host)
		_, err := s.Every(2).Seconds().Do(func(s *server) {
			healthy := s.checkHealth()
			if healthy {
				log.Printf("'%s' server is alive", s.URL)
			} else {
				log.Printf("'%s' server is dead", s.URL)
			}
		}, host)
		if err != nil {
			log.Fatalln("%s", err)
		}
	}
	s.StartAsync()
}
