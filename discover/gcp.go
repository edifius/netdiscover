package discover

import (
	"net"
	"log"
)

const (
	gcpPrivateIPv4URL = "http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/ip"
	gcpPublicIPv4URL  = "http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/access-configs/0/external-ip"
	gcpHostnameURL    = "http://metadata.google.internal/computeMetadata/v1/instance/hostname"
)

// NewGCPDiscoverer returns a new Google Cloud Platform network discoverer
func NewGCPDiscoverer() Discoverer {
	return NewDiscoverer(
		PrivateIPv4DiscovererOption(gcpPrivateIPv4),
		PublicIPv4DiscovererOption(gcpPublicIPv4),
		PublicHostnameDiscovererOption(gcpHostname),
	)
}

func gcpPrivateIPv4() (net.IP, error) {
	log.Print(StandardIPFromHTTP(gcpPrivateIPv4URL, map[string]string{"Metadata-Flavor": "Google"}))
	return StandardIPFromHTTP(gcpPrivateIPv4URL, map[string]string{"Metadata-Flavor": "Google"})
}

func gcpPublicIPv4() (net.IP, error) {
	log.Print(StandardIPFromHTTP(gcpPublicIPv4URL, map[string]string{"Metadata-Flavor": "Google"}))
	return StandardIPFromHTTP(gcpPublicIPv4URL, map[string]string{"Metadata-Flavor": "Google"})
}

func gcpHostname() (string, error) {
	return StandardHostnameFromHTTP(gcpHostnameURL, map[string]string{"Metadata-Flavor": "Google"})
}
