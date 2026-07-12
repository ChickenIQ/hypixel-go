package hypixel

import (
	"context"
	"fmt"

	nw "github.com/SkyCryptWebsite/SkyHelper-Networth-Go"
)

type Networth struct {
	Total       float64
	NonCosmetic float64
}

func (c *Client) GetNetworth(ctx context.Context, uuid, profileName string) (*Networth, error) {
	profile, err := c.GetSkyBlockProfile(ctx, uuid, profileName)
	if err != nil {
		return nil, err
	}

	museumData, err := c.GetMuseum(ctx, profile.ID, uuid)
	if err != nil {
		return nil, err
	}

	var balance float64
	if profile.Banking != nil && profile.Banking.Balance != nil {
		balance = *profile.Banking.Balance
	}

	calculator, err := nw.NewProfileNetworthCalculator(profile.Data, museumData, balance)
	if err != nil {
		return nil, fmt.Errorf("create networth calculator: %w", err)
	}

	return &Networth{
		NonCosmetic: calculator.GetNonCosmeticNetworth().Networth,
		Total:       calculator.GetNetworth().Networth,
	}, nil
}
