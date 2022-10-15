package installer

import (
	"log"
	"os"
	"path/filepath"
)

func (s *InstallArgs) Remove() error {
	err := Stop()
	if err != nil {
		return err
	}
	err = os.Remove(filepath.Join(s.Path, "clash"))
	if err != nil {
		return err
	}
	err = os.Remove("/etc/systemd/system/clash.service")
	if err != nil {
		return err
	}
	err = daemonReload()
	if err != nil {
		return err
	}
	log.Printf("Remove Successfully")
	return nil
}
