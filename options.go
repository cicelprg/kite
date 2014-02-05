package kite

import (
	"encoding/json"
	"io/ioutil"
	"koding/kite/protocol"
	"log"
	"net/url"
	"strings"
)

// Options is passed to kite.New when creating new instance.
type Options struct {
	Username              string
	Kitename              string
	PublicIP              string
	Environment           string
	Region                string
	Port                  string
	Path                  string
	Version               string
	KontrolURL            *url.URL
	DisableAuthentication bool
	Dependencies          string
	Visibility            protocol.Visibility
}

// validate is validating the fields of the options struct. It exits if an
// error is occured.
func (o *Options) validate() {
	if o.Kitename == "" {
		log.Fatal("ERROR: options.Kitename field is not set")
	}

	if digits := strings.Split(o.Version, "."); len(digits) != 3 {
		log.Fatal("ERROR: please use 3-digits semantic versioning for options.version")
	}

	if o.Region == "" {
		log.Fatal("ERROR: options.Region field is not set")
	}

	if o.Environment == "" {
		log.Fatal("ERROR: options.Environment field is not set")
	}

	if o.PublicIP == "" {
		o.PublicIP = "127.0.0.1"
	}

	if o.Port == "" {
		o.Port = "0" // OS binds to an automatic port
	}
	if o.Path == "" {
		o.Path = "/kite"
	}

	if o.Path[0] != '/' {
		o.Path = "/" + o.Path
	}

	if o.KontrolURL == nil {
		o.KontrolURL = &url.URL{
			Scheme: "ws",
			Host:   "127.0.0.1:4000", // local fallback address
			Path:   "/dnode",
		}
	}

	if o.Visibility == protocol.Visibility("") {
		o.Visibility = protocol.Private
	}
}

// Read options from a file.
func ReadKiteOptions(configfile string) (*Options, error) {
	file, err := ioutil.ReadFile(configfile)
	if err != nil {
		return nil, err
	}

	options := &Options{}
	err = json.Unmarshal(file, &options)
	if err != nil {
		return nil, err
	}

	return options, nil
}