package base

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Forever run forever
func Forever() {
	TrapSignal(func() {
		fmt.Println("  Bye bye!")
	})
}

// TrapSignal catch os.Interrupt and do [cb] func when exit
func TrapSignal(cb func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			fmt.Printf("captured %v, exiting...\n", sig)
			if cb != nil {
				cb()
			}
			time.Sleep(233 * time.Millisecond)
			os.Exit(1)
		}
	}()
	select {}
}
