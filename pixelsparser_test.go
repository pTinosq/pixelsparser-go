package pixelsparser

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestPixelCreation(t *testing.T) {
	date := "2023-06-29"
	p := Pixel{
		Date:   mustParse(date),
		Type:   "Mood",
		Scores: []int{4},
		Score:  4,
		Mood:   4,
		Notes:  "Woke up and had breakfast, went to work, had a good day",
		Tags:   []Tag{{Type: "Emotions", Entries: []string{"sad", "tired"}}},
	}

	if p.Date != mustParse(date) {
		t.Errorf("expected date %v, got %v", date, p.Date)
	}
	if p.Type != "Mood" || p.Score != 4 || p.Mood != 4 {
		t.Errorf("unexpected field values")
	}
}

func TestLoadFromJSON(t *testing.T) {
	path := filepath.Join("test_data", "test_diary.json")
	pixels, err := Load(path)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if len(pixels) != 4 {
		t.Fatalf("expected 4 pixels, got %d", len(pixels))
	}

	p := pixels[0]
	if p.Date != mustParse("2022-09-20") {
		t.Errorf("unexpected date: %v", p.Date)
	}
	if p.Type != "Mood" || p.Score != 2 || p.Mood != 2 {
		t.Errorf("unexpected score or type")
	}
	expectedNote := "I accidentally ate a whole bag of chips today. It slipped into my mouth. I swear!"
	if p.Notes != expectedNote {
		t.Errorf("unexpected notes: %s", p.Notes)
	}
	if len(p.Tags) != 1 || p.Tags[0].Type != "Emotions" {
		t.Errorf("unexpected tags: %+v", p.Tags)
	}

	p2 := pixels[1]
	if p2.Date != mustParse("2022-09-21") || p2.Score != 3 {
		t.Errorf("unexpected pixel 2 data")
	}
}

func mustParse(date string) time.Time {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	return t
}

func TestLoad_FileNotFound(t *testing.T) {
	_, err := Load("test_data/nonexistent.json")
	if err == nil {
		t.Fatal("expected error for non-existent file, got nil")
	}
}

func TestLoad_InvalidJSON(t *testing.T) {
	path := filepath.Join("test_data", "invalid.json")
	err := os.WriteFile(path, []byte(`{this is not valid JSON}`), 0644)
	if err != nil {
		t.Fatalf("failed to write invalid.json: %v", err)
	}
	defer os.Remove(path)

	_, err = Load(path)
	if err == nil {
		t.Fatal("expected JSON unmarshal error, got nil")
	}
}

func TestLoad_InvalidDateFormat(t *testing.T) {
	path := filepath.Join("test_data", "bad_date.json")
	err := os.WriteFile(path, []byte(`[
		{
			"date": "2022/09/20",
			"type": "Mood",
			"scores": [2],
			"notes": "bad date",
			"tags": []
		}
	]`), 0644)
	if err != nil {
		t.Fatalf("failed to write bad_date.json: %v", err)
	}
	defer os.Remove(path)

	_, err = Load(path)
	if err == nil {
		t.Fatal("expected date parse error, got nil")
	}
}
