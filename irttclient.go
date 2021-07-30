// POC: using irtt as a library with min config
package main

import (
	"context"
	"fmt"
	"github.com/heistp/irtt"
	"time"
)

func main() {
	cfg := irtt.NewClientConfig()

	cfg.LocalAddress = ":0"
	cfg.RemoteAddress = "127.0.0.1:2112"
	cfg.OpenTimeouts = []time.Duration{1 * time.Second, 1 * time.Second, 1 * time.Second}
	cfg.Duration = 10 * time.Second
	cfg.Interval = 1 * time.Second
	cfg.Length = 0
	cfg.Clock = 0x03 // Wall | Monotonic
	cfg.IPVersion = 1
	cfg.TTL = 64
	cfg.HMACKey = []byte("wazzup")

	c := irtt.NewClient(cfg)
	ctx := context.Background()
	r, err := c.Run(ctx)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r.RoundTripIPDVStats)

	fmt.Println("all done")
}
