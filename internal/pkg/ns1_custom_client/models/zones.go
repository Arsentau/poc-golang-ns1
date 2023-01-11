// Package models contains the models for the NS1 Client.
package models

// Zone is a model for a DNS zone.
type Zone struct {
	DNSServers    []string `json:"dns_servers"`
	Expiry        int      `json:"expiry"`
	PrimaryMaster string   `json:"primary_master"`
	Hostmaster    string   `json:"hostmaster"`
	ID            string   `json:"id"`
	Meta          struct {
		Asn           []string `json:"asn"`
		CaProvince    []string `json:"ca_province"`
		Connections   int      `json:"connections"`
		Country       []string `json:"country"`
		Georegion     []string `json:"georegion"`
		HighWatermark int      `json:"high_watermark"`
		IPPrefixes    []string `json:"ip_prefixes"`
		Latitude      int      `json:"latitude"`
		LoadAvg       int      `json:"loadAvg"`
		Longitude     int      `json:"longitude"`
		LowWatermark  int      `json:"low_watermark"`
		Note          string   `json:"note"`
		Priority      int      `json:"priority"`
		Pulsar        string   `json:"pulsar"`
		Requests      int      `json:"requests"`
		Up            bool     `json:"up"`
		UsState       []string `json:"us_state"`
		Weight        int      `json:"weight"`
	} `json:"meta"`
	NetworkPools []string `json:"network_pools"`
	Networks     []int    `json:"networks"`
	NxTTL        int      `json:"nx_ttl"`
	Primary      struct {
		Enabled     bool `json:"enabled"`
		Secondaries []struct {
			IP       string `json:"ip"`
			Networks []int  `json:"networks"`
			Notify   bool   `json:"notify"`
			Port     int    `json:"port"`
		} `json:"secondaries"`
	} `json:"primary"`
	Records []struct {
		Domain       string   `json:"domain"`
		ID           string   `json:"id"`
		ShortAnswers []string `json:"short_answers"`
		Tier         int      `json:"tier"`
		TTL          int      `json:"ttl"`
		Type         string   `json:"type"`
	} `json:"records"`
	Refresh   int      `json:"refresh"`
	Retry     int      `json:"retry"`
	TTL       int      `json:"ttl"`
	Zone      string   `json:"zone"`
	Views     []string `json:"views"`
	LocalTags []string `json:"local_tags"`
	Tags      struct {
	} `json:"tags"`
}

// Zones is a list of Zone
type Zones []Zone
