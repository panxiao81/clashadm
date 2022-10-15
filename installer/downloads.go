package installer

import (
	"compress/gzip"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

// InstallArgs install manager
type InstallArgs struct {
	Path       string `mapstructure:"path"`
	Release    string `mapstructure:"release"`
	ConfigPath string `mapstructure:"config_path"`
}

var ReleaseUrl = map[string]string{
	"premium": "https://api.github.com/repos/Dreamacro/clash/releases/tags/premium",
	"clash":   "https://api.github.com/repos/Dreamacro/clash/releases",
}

func (s *InstallArgs) check() error {
	if _, ok := ReleaseUrl[s.Release]; ok {
		return nil
	} else {
		return errors.New("Not support this release")
	}
}

func (s *InstallArgs) Install() {
	if os.Geteuid() != 0 {
		log.Fatal("You must run clashadm install as root!")
	}
	log.Printf("Installing clash release %s to %s\n", s.Release, s.Path)
	s.Download()
	log.Printf("Generate Systemd Unit File")
	err := s.genericSystemdUnit()
	if err != nil {
		log.Fatal(err)
	}
	err = daemonReload()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("You can start clash now by \"clashadm start\"\n")
}

func (s *InstallArgs) Download() {
	url, err := s.getDownloadUrl()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Download from %s", url)
	os.Remove(filepath.Join(s.Path, "clash"))
	out, err := os.Create(filepath.Join(s.Path, "clash"))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Download Failed", resp.StatusCode)
	}

	zip, err := gzip.NewReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(out, zip)
	if err != nil {
		log.Fatal(err)
	}
	out.Chmod(os.FileMode(0755))
	log.Printf("clash %s Download Successful at %s\n", s.Release, filepath.Join(s.Path, "clash"))
}

func (s *InstallArgs) genericSystemdUnit() error {
	systemdUnitFileTmpl := `
[Unit]
Description=Clash daemon, A rule-based proxy in Go.
After=network.target

[Service]
Type=simple
Restart=always
ExecStart={{ .Path }}/clash -d {{ .ConfigPath }}

[Install]
WantedBy=multi-user.target
`
	tmpl, err := template.New("systemdUnitFileTmpl").Parse(systemdUnitFileTmpl)
	if err != nil {
		return err
	}

	unit, err := os.OpenFile("/etc/systemd/system/clash.service", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer unit.Close()

	err = tmpl.Execute(unit, *s)
	if err != nil {
		return err
	}
	return nil
}
