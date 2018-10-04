package frontend

import (
	"fmt"

	"github.com/davitron/go-micro-services/dialer"
	"github.com/davitron/go-micro-services/registry"
	"github.com/davitron/go-micro-services/services/frontend"
	profile "github.com/davitron/go-micro-services/services/profile/proto"
	search "github.com/davitron/go-micro-services/services/search/proto"
	"github.com/davitron/go-micro-services/tracing"
)

func runFrontend(port int, consul *registry.Client, jaegeraddr string) error {
	tracer, err := tracing.Init("frontend", jaegeraddr)
	if err != nil {
		return fmt.Errorf("tracing init error: %v", err)
	}

	// dial search srv
	sc, err := dialer.Dial(
		searchSrvName,
		dialer.WithTracer(tracer),
		dialer.WithBalancer(consul.Client),
	)
	if err != nil {
		return fmt.Errorf("dialer error: %v", err)
	}

	// dial profile srv
	pc, err := dialer.Dial(
		profileSrvName,
		dialer.WithTracer(tracer),
		dialer.WithBalancer(consul.Client),
	)
	if err != nil {
		return fmt.Errorf("dialer error: %v", err)
	}

	srv := frontend.NewServer(
		search.NewSearchClient(sc),
		profile.NewProfileClient(pc),
		tracer,
	)
	return srv.Run(port)
}
