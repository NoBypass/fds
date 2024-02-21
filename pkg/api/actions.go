package api

import (
	"net/http"
	"strconv"
)

// Verify is used to link a Discord account to a Hypixel account.
// The backend will store a snapshot of the player's Hypixel stats
// and Mojang profile as well as store the Discord user.
func (c *Client) Verify(input *DiscordVerifyRequest) (*DiscordVerifyResponse, error) {
	req, err := c.newJsonRequest(http.MethodPost, "/discord/verify", input)
	if err != nil {
		return nil, err
	}

	return do[DiscordVerifyResponse](req)
}

// Daily is used to claim the daily reward for a Discord user.
// The backend will return the user's updated stats.
// TODO: Add error docs
func (c *Client) Daily(id string) (*DiscordMemberResponse, error) {
	req, err := c.newJsonRequest(http.MethodPost, "/discord/daily/"+id, nil)
	if err != nil {
		return nil, err
	}

	return do[DiscordMemberResponse](req)
}

// BotLogin is used to login the bot to the Discord API.
// No token is required for this endpoint.
func (c *Client) BotLogin(input *DiscordBotLoginRequest) (*DiscordBotLoginResponse, error) {
	req, err := c.newJsonRequest(http.MethodPost, "/discord/bot-login", input)
	if err != nil {
		return nil, err
	}

	return do[DiscordBotLoginResponse](req)
}

// Leaderboard is used to get the leaderboard for all verified Discord users.
// NOTE: The pagination uses zero-based indexing.
func (c *Client) Leaderboard(page int) (*DiscordLeaderboardResponse, error) {
	req, err := c.newJsonRequest(http.MethodGet, "/discord/leaderboard/"+strconv.Itoa(page), nil)
	if err != nil {
		return nil, err
	}

	return do[DiscordLeaderboardResponse](req)
}

// Member is used to get the stats for a specific Discord user.
func (c *Client) Member(id string) (*DiscordMemberResponse, error) {
	req, err := c.newJsonRequest(http.MethodGet, "/discord/member/"+id, nil)
	if err != nil {
		return nil, err
	}

	return do[DiscordMemberResponse](req)
}

// Revoke is used to unlink a Discord account from a Hypixel account.
func (c *Client) Revoke(id string) error {
	req, err := c.newJsonRequest(http.MethodDelete, "/discord/revoke/"+id, nil)
	if err != nil {
		return err
	}

	_, err = do[DiscordMemberResponse](req)
	return err
}
