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
	cfg.OpenTimeouts, _ = irtt.ParseDurations("1s")
	cfg.Duration, _ = time.ParseDuration("1s")
	cfg.Interval, _ = time.ParseDuration("20ms")
	cfg.Length = 100
	cfg.Clock = irtt.BothClocks
	cfg.IPVersion = irtt.IPVersionFromBooleans(true, true, irtt.DualStack)
	cfg.TTL = 64
	cfg.HMACKey = []byte("wazzup")

	c := irtt.NewClient(cfg)
	ctx := context.Background()
	r, err := c.Run(ctx)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("RTT min/mean/max: %s/%s/%s\n", r.RTTStats.Min, r.RTTStats.Mean(), r.RTTStats.Max)
	fmt.Printf("jitter min/mean/max: %s/%s/%s\n", r.RoundTripIPDVStats.Min, r.RoundTripIPDVStats.Mean(), r.RoundTripIPDVStats.Max)
	fmt.Printf("late packet count/percent: %d/%f\n", r.LatePackets, r.LatePacketsPercent)
	fmt.Printf("lost packet percent: %f\n", r.PacketLossPercent)

	fmt.Println("all done")
}
