// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// sommelier HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/cellar/design

package client

import (
	"encoding/json"
	"fmt"

	"goa.design/plugins/goakit/examples/cellar/gen/sommelier"
)

// BuildCriteria builds the payload for the sommelier pick endpoint from CLI
// flags.
func BuildCriteria(sommelierPickBody string) (*sommelier.Criteria, error) {
	var body PickRequestBody
	{
		err := json.Unmarshal([]byte(sommelierPickBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"name\": \"Blue\\'s Cuvee\",\n      \"varietal\": [\n         \"pinot noir\",\n         \"merlot\",\n         \"cabernet franc\"\n      ],\n      \"winery\": \"longoria\"\n   }'")
		}
	}
	v := &sommelier.Criteria{
		Name:   body.Name,
		Winery: body.Winery,
	}
	if body.Varietal != nil {
		v.Varietal = make([]string, len(body.Varietal))
		for j, val := range body.Varietal {
			v.Varietal[j] = val
		}
	}

	return v, nil
}
