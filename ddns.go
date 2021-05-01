package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cloudflare/cloudflare-go"
)

type Config struct {
	Token        string
	Zones        map[string][]string
	UpdatePeriod *int64 `json:update_period,omitempty`
}

func loadConfig(path string) *Config {
	configFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open config file %v", err)
	}
	defer func() {
		if err := configFile.Close(); err != nil {
			log.Fatalf("Failed to close config file: %v", err)
		}
	}()
	configContents, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	var c Config
	if err = json.Unmarshal(configContents, &c); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}
	return &c
}

func getMyIP() (string, error) {
	response, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	ip, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(ip), nil
}

// List zones
// Map name -> ID
// For each zone
//   List DNS records
//   Map name -> ID
//   Update DNS record
func main() {
	// TODO: add -c flag for config path
	configPath := "cloudflare-ddns.json"
	log.Printf("Using config file at %v", configPath)
	config := loadConfig(configPath)

	api, err := cloudflare.NewWithAPIToken(config.Token)
	if err != nil {
		log.Fatal(err)
	}

	var updatePeriod = 5 * time.Minute
	if config.UpdatePeriod != nil {
		// Update period is specified in seconds
		updatePeriod = time.Duration(*config.UpdatePeriod) * time.Second
	}
	log.Printf("Using update period = %v", updatePeriod)

	for {
		myIP, err := getMyIP()
		if err != nil {
			log.Printf("Failed to get IP: %v", err)
		}

		for zoneName, names := range config.Zones {
			zoneID, err := api.ZoneIDByName(zoneName)
			if err != nil {
				log.Printf("Failed to get zone ID for zone %v: %v", zoneName, err)
				continue
			}
			for _, name := range names {
				records, err := api.DNSRecords(context.Background(), zoneID, cloudflare.DNSRecord{Name: name})
				if err != nil {
					log.Printf("Failed to get DNS record with name %v: %v", name, err)
					continue
				}
				for _, record := range records {

					if record.Content == myIP {
						log.Printf("In zone %v: DNS record %v already up-to-date", zoneName, name)
					} else {
						record.Content = myIP
						if err = api.UpdateDNSRecord(context.Background(), zoneID, record.ID, record); err != nil {
							log.Printf("Failed to update DNS record %v: %v", name, err)
						} else {
							log.Printf("In zone %v: updated DNS record %v to %v", zoneName, name, myIP)
						}
					}
				}
			}
		}
		time.Sleep(updatePeriod)
	}
}
