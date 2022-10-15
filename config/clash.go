package config

import (
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

type ClashConfig struct {
	MixedPort          int    `yaml:"mixed-port,omitempty" mapstructure:"mixed-port"`
	AllowLan           bool   `yaml:"allow-lan,omitempty" mapstructure:"allow-lan"`
	Mode               string `yaml:"mode,omitempty" mapstructure:"mode"`
	LogLevel           string `yaml:"log-level,omitempty" mapstructure:"log-level"`
	ExternalController string `yaml:"external-controller,omitempty" mapstructure:"external-controller"`
	Secret             string `yaml:"secret,omitempty" mapstructure:"secret"`
	DNS                struct {
		Enable         bool     `yaml:"enable,omitempty" mapstructure:"enable"`
		Ipv6           bool     `yaml:"ipv6,omitempty" mapstructure:"ipv6"`
		Listen         string   `yaml:"listen,omitempty" mapstructure:"listen"`
		EnhancedMode   string   `yaml:"enhanced-mode,omitempty" mapstructure:"enhanced-mode"`
		FakeIPFilter   []string `yaml:"fake-ip-filter,omitempty" mapstructure:"fake-ip-filter"`
		Nameserver     []string `yaml:"nameserver,omitempty" mapstructure:"nameserver"`
		Fallback       []string `yaml:"fallback,omitempty" mapstructure:"fallback"`
		FallbackFilter struct {
			Geoip  bool     `yaml:"geoip,omitempty" mapstructure:"geoip"`
			Ipcidr []string `yaml:"ipcidr,omitempty" mapstructure:"ipcidr"`
			Domain []string `yaml:"domain,omitempty" mapstructure:"domain"`
		} `yaml:"fallback-filter,omitempty" mapstructure:"fallback-filter"`
		FakeIPRange       string   `yaml:"fake-ip-range,omitempty" mapstructure:"fake-ip-range"`
		DefaultNameserver []string `yaml:"default-nameserver,omitempty" mapstructure:"default-nameserver"`
	} `yaml:"dns,omitempty" mapstructure:"dns"`
	Proxies []struct {
		Name          string `yaml:"name,omitempty" mapstructure:"name"`
		Type          string `yaml:"type,omitempty" mapstructure:"type"`
		Server        string `yaml:"server,omitempty" mapstructure:"server"`
		Port          int    `yaml:"port,omitempty" mapstructure:"port"`
		Cipher        string `yaml:"cipher,omitempty" mapstructure:"cipher"`
		Password      string `yaml:"password,omitempty" mapstructure:"password"`
		Protocol      string `yaml:"protocol,omitempty" mapstructure:"protocol"`
		ProtocolParam string `yaml:"protocol-param,omitempty" mapstructure:"protocol-param"`
		Obfs          string `yaml:"obfs,omitempty" mapstructure:"obfs"`
		ObfsParam     string `yaml:"obfs-param,omitempty" mapstructure:"obfs-param"`
		UDP           bool   `yaml:"udp,omitempty" mapstructure:"udp"`
	} `yaml:"proxies,omitempty" mapstructure:"proxies"`
	ProxyGroups []struct {
		Name     string   `yaml:"name,omitempty" mapstructure:"name"`
		Type     string   `yaml:"type,omitempty" mapstructure:"type"`
		Proxies  []string `yaml:"proxies,omitempty" mapstructure:"proxies"`
		URL      string   `yaml:"url,omitempty,omitempty" mapstructure:"url"`
		Interval string   `yaml:"interval,omitempty,omitempty" mapstructure:"interval"`
	} `yaml:"proxy-groups,omitempty" mapstructure:"proxy-groups"`
	Rules        []string `yaml:"rules,omitempty" mapstructure:"rules"`
	RedirPort    int      `yaml:"redir-port,omitempty" mapstructure:"redir-port"`
	TproxyPort   int      `yaml:"tproxy-port,omitempty" mapstructure:"tproxy-port"`
	Port         int      `yaml:"port,omitempty" mapstructure:"port"`
	SocksPort    int      `yaml:"socks-port,omitempty" mapstructure:"socks-port"`
	BindAddress  string   `yaml:"bind-address,omitempty" mapstructure:"bind-address"`
	ExternalUI   string   `yaml:"external-ui,omitempty" mapstructure:"external-ui"`
	Ipv6         bool     `yaml:"ipv6,omitempty" mapstructure:"ipv-6"`
	Experimental struct {
		SniffTLSSni bool `yaml:"sniff-tls-sni,omitempty" mapstructure:"sniff-tls-sni"`
	} `yaml:"experimental,omitempty" mapstructure:"experimental"`
	Tun struct {
		Enable              bool     `yaml:"enable,omitempty" mapstructure:"enable"`
		Stack               string   `yaml:"stack,omitempty" mapstructure:"stack"`
		AutoRoute           bool     `yaml:"auto-route,omitempty" mapstructure:"auto-route"`
		AutoDetectInterface bool     `yaml:"auto-detect-interface,omitempty" mapstructure:"auto-detect-interface"`
		DNSHijack           []string `yaml:"dns-hijack,omitempty" mapstructure:"dns-hijack"`
	} `yaml:"tun,omitempty" mapstructure:"tun"`
	Profile struct {
		StoreSelected bool `yaml:"store-selected,omitempty" mapstructure:"store-selected"`
		StoreFakeIP   bool `yaml:"store-fake-ip,omitempty" mapstructure:"store-fake-ip"`
	} `yaml:"profile,omitempty" mapstructure:"profile"`
	Authentication []string `yaml:"authentication,omitempty" mapstructure:"authentication"`
}

func NewClashConfigHandler(f io.Reader) (*ClashConfig, error) {
	var c ClashConfig
	y := yaml.NewDecoder(f)
	err := y.Decode(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil

}

func (s *ClashConfig) Unmarshal(p string) error {
	f, err := os.OpenFile(p, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	err = f.Truncate(0)
	if err != nil {
		return err
	}
	y := yaml.NewEncoder(f)
	defer y.Close()
	err = y.Encode(s)
	if err != nil {
		return err
	}
	return nil
}
