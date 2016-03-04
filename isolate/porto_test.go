// +build linux

package isolate

import (
	"testing"

	log "github.com/Sirupsen/logrus"
	logformatter "github.com/noxiouz/stout/pkg/formatter"

	"golang.org/x/net/context"
)

func init() {
	log.SetFormatter(&logformatter.CombaineFormatter{})
	log.SetLevel(log.DebugLevel)
}

func TestMain(t *testing.T) {
	config := &PortoIsolationConfig{
		RootNamespace: "cocs",
		Layers:        "/tmp/isolate",
		Volumes:       "/cocaine-porto",
	}

	p, err := NewPortoIsolation(config)
	t.Log(p, err)
	ctx := context.Background()
	if err := p.Spool(ctx, "registry.ape.yandex.net/echo", "latest"); err != nil {
		t.Fatal(err)
	}

	profile := Profile{
		WorkingDir:  "/",
		NetworkMode: "host",
		Image:       "registry.ape.yandex.net/echo",
		Command:     "ls -la /",
		Bind:        "",
	}
	container, err := p.Create(ctx, profile)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(container)
	if err := p.Start(ctx, container); err != nil {
		t.Fatal(err)
	}

	output, err := p.Output(ctx, container)
	if err != nil {
		t.Fatal(err)
	}
	defer output.Close()
	t.Log(output)
	stderr := make([]byte, 10)
	n, err := output.Read(stderr)
	t.Logf("STDERR: %d %v %s", n, err, stderr)

	if err := p.Terminate(ctx, container); err != nil {
		t.Fatal(err)
	}
}
