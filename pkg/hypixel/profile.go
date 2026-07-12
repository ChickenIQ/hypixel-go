package hypixel

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	sc "github.com/DuckySoLucky/SkyCrypt-Types"
)

type SkyBlockProfile struct {
	ID                string
	Name              string
	GameMode          string
	Active            bool
	Data              *sc.Member
	Banking           *sc.Banking
	CommunityUpgrades *sc.CommunityUpgrades
}

func (c *Client) GetProfileData(ctx context.Context, uuid string) ([]sc.Profile, error) {
	body, err := c.getWithUUID(ctx, "/skyblock/profiles", parseUUID(uuid))
	if err != nil {
		return nil, err
	}

	var data struct {
		Profiles []sc.Profile `json:"profiles"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("decode Hypixel profiles response: %w", err)
	}

	return data.Profiles, nil
}

func (c *Client) GetSkyBlockProfiles(ctx context.Context, uuid string) ([]SkyBlockProfile, error) {
	uuid = parseUUID(uuid)
	data, err := c.GetProfileData(ctx, uuid)
	if err != nil {
		return nil, err
	}

	profiles := make([]SkyBlockProfile, 0, len(data))
	for i := range data {
		var memberData *sc.Member
		p := &data[i]

		member, ok := p.Members[uuid]
		if ok {
			memberData = &member
		}

		profiles = append(profiles, SkyBlockProfile{
			ID:                p.ProfileID,
			Name:              p.CuteName,
			GameMode:          p.GameMode,
			Active:            p.Selected,
			Data:              memberData,
			Banking:           p.Banking,
			CommunityUpgrades: p.CommunityUpgrades,
		})
	}

	return profiles, nil
}

func (c *Client) GetSkyBlockProfile(ctx context.Context, uuid, profileName string) (*SkyBlockProfile, error) {
	profiles, err := c.GetSkyBlockProfiles(ctx, parseUUID(uuid))
	if err != nil {
		return nil, err
	}

	if profileName == "" {
		for _, profile := range profiles {
			if !profile.Active {
				continue
			}

			if profile.Data == nil {
				return nil, fmt.Errorf("member %s not found in selected profile", uuid)
			}

			return &profile, nil
		}

		return nil, fmt.Errorf("selected profile not found")
	}

	for _, profile := range profiles {
		if strings.EqualFold(profile.Name, profileName) {
			if profile.Data == nil {
				return nil, fmt.Errorf("member %s not found in profile %s", uuid, profileName)
			}
			return &profile, nil
		}
	}

	return nil, fmt.Errorf("profile %s not found", profileName)
}
