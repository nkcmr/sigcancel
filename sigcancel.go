package sigcancel // import "code.nkcmr.net/sigcancel"

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// CancelOnSignal is a compact, idiomatic way to listen for OS signals in the
// main of a Go program. The intent here is to reduce the tedium of setting up
// proper signal handling and just compact the whole thing into what Go writers
// are used to dealing with: context cancellations. Usage looks like this:
//     func main() {
//         ctx, cancel := context.WithCancel(context.Background())
//         go sigcancel.CancelOnSignal(cancel)
//         go startWorking(ctx)
//         <-ctx.Done()
//         fmt.Println("shutting down...")
//         // clean up services / workload
//     }
func CancelOnSignal(cancel context.CancelFunc) {
	defer cancel()
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("signal: %s\n", <-c)
}
