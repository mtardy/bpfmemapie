package piechart

import (
	"io"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func Render(w io.Writer, serie []opts.PieData) {
	pie := charts.NewPie()

	pie.SetGlobalOptions(
		charts.WithTooltipOpts(opts.Tooltip{
			Show: true,
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "BPF Maps Total Bytes Memlock",
			Subtitle: `The pie chart below represents the total bytes memlock by maps aggregatated by map names.

Refresh to reload data using bpftool. Use ?threshold=<value in %> to group small maps (bytes_memlock < value in % of the total) under the "others" label.
`,
		}),
		charts.WithLegendOpts(opts.Legend{
			Show:   true,
			Type:   "scroll",
			Orient: "vertical",
			Left:   "right",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Height:    "90vh",
			Width:     "90vw",
			PageTitle: "BPF Map Total Bytes Memlock",
		}),
	)

	pie.AddSeries("maps total bytes memlock", serie).SetSeriesOptions(
		charts.WithSeriesAnimation(true),
		charts.WithLabelOpts(opts.Label{
			Show:      true,
			Formatter: "{b}: {c} ({d}%)",
		}),
		charts.WithPieChartOpts(opts.PieChart{
			Radius: "60%",
			Center: []string{"45%", "55%"},
		}),
	)
	pie.Render(w)
}
