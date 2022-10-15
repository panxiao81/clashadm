BINDIR=bin
VERSION=$(shell git describe --tags || echo "unknown")
COMMITID=$(shell git describe --abbrev=4 HEAD)
BUILDTIME=$(shell date -u)
GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags '-X "github.com/panxiao81/clashadm/cmd.Version=$(VERSION)" \
	-X "github.com/panxiao81/clashadm/cmd.Commit=$(COMMITID)" \
	-X "github.com/panxiao81/clashadm/cmd.BuildTime=$(BUILDTIME)" \
	-w -s -buildid='

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/clashadm-linux-amd64

clean:
	rm $(BINDIR)/*