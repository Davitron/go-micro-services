package frontend

import (
	"flag"
	"fmt"
	"os"

	"github.com/davitron/go-micro-services/registry"
)

func main() {
	var (
		port       = flag.Int("port", 5000, "The server port")
		jaegeraddr = flag.String("jaeger_addr", "jaeger:6831", "Jaeger address")
		consuladdr = flag.String("consul_addr", "consul:8500", "Consul address")
	)
	flag.Parse()

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	runFrontend(5000, consul*registry.Client, jaegeraddr)

	consul, err := registry.NewClient(*consuladdr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err := run(*port, consul, *jaegeraddr); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
