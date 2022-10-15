package installer

import (
	"context"
	"github.com/coreos/go-systemd/v22/dbus"
	"time"
)

func daemonReload() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()
	conn, err := dbus.NewSystemdConnectionContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	err = conn.ReloadContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func Start() error {

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()
	conn, err := dbus.NewSystemdConnectionContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.StartUnitContext(ctx, "clash.service", "replace", nil)
	if err != nil {
		return err
	}
	return nil
}

func Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()
	conn, err := dbus.NewSystemdConnectionContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.StopUnitContext(ctx, "clash.service", "replace", nil)
	if err != nil {
		return err
	}
	return nil
}
