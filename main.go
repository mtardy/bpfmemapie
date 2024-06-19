package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/mtardy/bpfmemapie/internal/mapsdata"
	"github.com/mtardy/bpfmemapie/internal/piechart"

	"github.com/go-echarts/go-echarts/v2/opts"
)

const (
	defaultThreshold     = 0.1
	defaultWebserverPort = 8080
)

var webserverPort string

func convertToPieItems(data mapsdata.MapsData) []opts.PieData {
	items := make([]opts.PieData, 0)
	for name, d := range data {
		items = append(items, opts.PieData{Name: name, Value: d.TotalBytesMemlock})
	}
	return items
}

func main() {
	flag.StringVar(&webserverPort, "port", fmt.Sprint(defaultWebserverPort), "Port to listen to for the webserver")
	flag.Parse()

	// browser will try to retrieve the favicon on root, ignore
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		maps, err := mapsdata.BPFtoolFetchMapsData()

		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		aggregatedMaps := mapsdata.AggregateMapsPerName(maps)

		threshold := defaultThreshold
		rawThreshold := strings.TrimSpace(req.URL.Query().Get("threshold"))
		if rawThreshold != "" {
			threshold, err = strconv.ParseFloat(rawThreshold, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, fmt.Errorf("failed to parse threshold, using default value %f: %w", defaultThreshold, err))
			}
		}

		serie := convertToPieItems(mapsdata.AggregateUnderThreshold(aggregatedMaps, threshold))

		piechart.Render(w, serie)
	})

	fmt.Printf("Listening on http://localhost:%s/\n", webserverPort)

	err := http.ListenAndServe(fmt.Sprintf("localhost:%s", webserverPort), nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
