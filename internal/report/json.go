package report

import (
	"encoding/json"
	"os"

	"github.com/M4elstr0m/gophoner/internal/recon"
)

func PrintJSON(phoneNumber string, results []recon.Output) error {
	return json.NewEncoder(os.Stdout).Encode(struct {
		Target  string         `json:"target"`
		Results []recon.Output `json:"results"`
	}{phoneNumber, results})
}
