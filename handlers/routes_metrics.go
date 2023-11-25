package handlers

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

type GraphDataset struct {
	Label       string `json:"label"`
	Data        []int  `json:"data"`
	BorderWidth int    `json:"borderWidth"`
}

type GraphData struct {
	Labels   []string       `json:"labels"`
	Datasets []GraphDataset `json:"datasets"`
}

func HandleMetrics(c *fiber.Ctx) error {
	graphData := GraphData{
		Labels: []string{"Red", "Blue", "Yellow", "Green", "Purple", "Orange"},
		Datasets: []GraphDataset{
			{
				Label:       "# of Votes",
				Data:        []int{12, 19, 3, 5, 2, 3},
				BorderWidth: 1,
			},
		},
	}

	graphDataBuf := &bytes.Buffer{}
	if err := json.NewEncoder(graphDataBuf).Encode(graphData); err != nil {
		return err
	}

	log.Println(graphDataBuf.String())

	return c.Render("routes/metrics", fiber.Map{
		"GraphData": graphDataBuf.String(),
	})
}
