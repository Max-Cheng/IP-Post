package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	PORT := os.Getenv("PORT")
	var RWL sync.Mutex
	hostMap := make(map[string]time.Time)
	r := fiber.New()
	r.Get("/", func(c *fiber.Ctx) error {
		RWL.Lock()
		hostSlice := make([]string, 0)
		for i, v := range hostMap {
			if time.Since(v) >= time.Hour {
				delete(hostMap, i)
				continue
			}
			hostSlice = append(hostSlice, i)
		}
		RWL.Unlock()
		return c.JSON(hostSlice)
	})
	r.Post("/host", func(c *fiber.Ctx) error {
		RWL.Lock()
		hostMap[c.IP()] = time.Now()
		RWL.Unlock()
		return nil
	})
	log.Fatal(r.Listen(":" + PORT))
}
