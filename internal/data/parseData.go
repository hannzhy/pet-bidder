package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"pet-bidder/internal/types"
)

func GetInitialData() (campaigns []types.AdCampaign, err error) {
	path, err := filepath.Abs("./internal/data/data.json")
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %v", err)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed on open file: %v\n", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&campaigns); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %v\n", err)
	}

	return campaigns, nil
}
