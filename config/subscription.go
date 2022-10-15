package config

import (
	"errors"
	"fmt"
	"github.com/imdario/mergo"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"text/tabwriter"
)

type Subscription []SubUrl
type SubUrl struct {
	Name    string
	URL     string
	Enabled bool
}

func (s *Subscription) Add(name string, u string) error {
	if _, err := url.ParseRequestURI(u); err != nil {
		log.Print(err)
	}
	t := SubUrl{
		Name:    name,
		URL:     u,
		Enabled: false,
	}
	*s = append(*s, t)
	viper.Set("subscription_url", *s)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func (s *Subscription) Get(i int) string {
	return (*s)[i].Name
}

func (s *Subscription) GetByUrl(u string) (int, error) {
	for k, v := range *s {
		if v.URL == u {
			return k, nil
		}
	}
	return 0, errors.New("not Found")
}

func (s *Subscription) GetByName(n string) (int, error) {
	for k, v := range *s {
		if v.Name == n {
			return k, nil
		}
	}
	return 0, errors.New("not Found")
}

func (s *Subscription) SetDefault(n string) error {
	i, err := s.GetByName(n)
	if err != nil {
		return err
	}
	if (*s)[i].Enabled {
		return nil
	}
	for _, v := range *s {
		if v.Enabled {

			v.Enabled = false
		}
	}
	(*s)[i].Enabled = true

	viper.Set("subscription_url", *s)
	err = viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func (s *Subscription) List() {
	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	// Title
	fmt.Fprintln(writer, "NAME\tURL\tLAST_UPDATE\tENABLED")
	for _, v := range *s {
		if len(v.URL) > 12 {
			fmt.Fprintf(writer, "%s\t%s...\t%s\t%s\n", v.Name, v.URL[:12], "2006-01-02 15:04:05", strconv.FormatBool(v.Enabled))
		} else {
			fmt.Fprintf(writer, "%s\t%s\t%s\t%s\n", v.Name, v.URL, "2006-01-02 15:04:05", strconv.FormatBool(v.Enabled))
		}
	}
	writer.Flush()
}

func (s *Subscription) Download(name string, path string) (io.ReadCloser, error) {
	i, err := s.GetByName(name)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get((*s)[i].URL)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	return resp.Body, nil
}

func (s *Subscription) Update(name string) error {
	p := viper.GetString("install.config_path")
	u, err := s.Download(name, p)
	if err != nil {
		return err
	}
	defer u.Close()

	data, err := ioutil.ReadAll(u)
	if err != nil {
		return err
	}

	var c ClashConfig
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return err
	}
	c.Port = 0
	c.TproxyPort = 0
	c.SocksPort = 0
	c.MixedPort = 0
	c.RedirPort = 0

	m := viper.Sub("mixed-in")
	var mix ClashConfig
	err = m.Unmarshal(&mix)
	if err != nil {
		return err
	}
	if mix.MixedPort == 0 {
		mix.MixedPort = 7890
	}

	// merge original config and mixed-in in config file
	if err = mergo.Merge(&mix, c); err != nil {
		return err
	}

	o, err := yaml.Marshal(&mix)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(p, name+".yaml"), o, 0644)

	if err != nil {
		return err
	}
	log.Println(string(o))

	return nil
}
