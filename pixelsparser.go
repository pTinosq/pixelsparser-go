package pixelsparser

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Tag struct {
	Type    string   `json:"type"`
	Entries []string `json:"entries"`
}

type Pixel struct {
	Date   time.Time
	Type   string
	Scores []int
	Score  int
	Mood   int
	Notes  string
	Tags   []Tag
}

type PixelJSON struct {
	Date   string `json:"date"`
	Type   string `json:"type"`
	Scores []int  `json:"scores"`
	Notes  string `json:"notes"`
	Tags   []Tag
}

func Load(path string) ([]Pixel, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	var pixelsJSON []PixelJSON
	err = json.Unmarshal(data, &pixelsJSON)
	if err != nil {
		return nil, err
	}

	pixels := make([]Pixel, len(pixelsJSON))
	for i, pj := range pixelsJSON {
		parsedDate, err := time.Parse("2006-1-2", pj.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to parse date '%s': %w", pj.Date, err)
		}
		pixels[i] = Pixel{
			Date:   parsedDate,
			Type:   pj.Type,
			Scores: pj.Scores,
			Score:  pj.Scores[0],
			Mood:   pj.Scores[0],
			Notes:  pj.Notes,
			Tags:   pj.Tags,
		}
	}

	return pixels, nil
}
