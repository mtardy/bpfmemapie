package mapsdata

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type BPFToolMap struct {
	// ID           int    `json:"id"`
	// Type         string `json:"type"`
	Name string `json:"name"`
	// Flags        int    `json:"flags"`
	// BytesKey     int    `json:"bytes_key"`
	// BytesValue   int    `json:"bytes_value"`
	// MaxEntries   int    `json:"max_entries"`
	BytesMemlock int `json:"bytes_memlock"`
	// Frozen       int    `json:"frozen"`
	// BtfID        int    `json:"btf_id"`
	// Pids []struct {
	// 	Pid  int    `json:"pid"`
	// 	Comm string `json:"comm"`
	// } `json:"pids"`
}

type MapDataValue struct {
	TotalBytesMemlock int
	Maps              int
}

type MapsData map[string]MapDataValue

// AggregateMapsPerName regroups all entries that have the same map name
func AggregateMapsPerName(rawData []BPFToolMap) MapsData {
	m := make(MapsData)

	for _, d := range rawData {
		entry, exist := m[d.Name]
		if exist {
			entry.Maps++
			entry.TotalBytesMemlock += d.BytesMemlock
			m[d.Name] = entry
		} else {
			m[d.Name] = MapDataValue{
				Maps:              1,
				TotalBytesMemlock: d.BytesMemlock,
			}
		}
	}

	return m
}

// AggregateUnderThreshold regroup all entries that value is inferior to
// "threshold / 100 * total" under the entry name "others"
func AggregateUnderThreshold(data MapsData, threshold float64) MapsData {
	m := make(MapsData)

	totalSum := 0
	for _, d := range data {
		totalSum += d.TotalBytesMemlock
	}

	for name, d := range data {
		min := float64(threshold) / 100 * float64(totalSum)
		if float64(d.TotalBytesMemlock) < min {
			others, exist := m["others"]
			if exist {
				others.Maps += d.Maps
				others.TotalBytesMemlock += d.TotalBytesMemlock
				m["others"] = others
			} else {
				m["others"] = MapDataValue{
					Maps:              d.Maps,
					TotalBytesMemlock: d.TotalBytesMemlock,
				}
			}
		} else {
			m[name] = d
		}
	}

	return m
}

// BPFToolFetchMapsData retrieves data using "bpftool map" and unmarshal the
// result into an appropriate Go struct
func BPFtoolFetchMapsData() ([]BPFToolMap, error) {
	cmd := exec.Command("bpftool", "map", "-j")
	stdout, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("failed to run bpftool, make sure it's installed and you run as root: %w", err)
	}

	maps := []BPFToolMap{}
	err = json.Unmarshal(stdout, &maps)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal bpftool output: %w", err)
	}

	return maps, nil
}
