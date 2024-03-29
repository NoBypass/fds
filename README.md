# FDS Core

## API
### `/`
##### Methods

- `NewFDSClient(url string) *Client` NewFDSClient creates a new client for the FDS API. For most cases, it is
  recommended to use the SetToken method soon after to set the token.
- `SetToken(token string)` SetToken (on Client) sets the token for the client. This token is used for authentication.

### `/discord`
<details>
 <summary><code>POST</code> <code><b>/verify</b></code> <code>Meant for linking Discord to Minecraft accounts</code></summary>

##### Request Body (JSON)

  ``` go
  type DiscordVerifyRequest struct {
	ID   string `json:"id"`
	Nick string `json:"nick"`
	Name string `json:"name"`
  }
  ```

##### Response Body (JSON)

  ``` go
  type DiscordVerifyResponse struct {
	Actual string `json:"actual"`
  }
  ```

##### Method (on Client)
`Verify(input *DiscordVerifyRequest) (*DiscordVerifyResponse, error)` Verify is used to link a Discord
account to a Hypixel account. The backend will store a snapshot of the player's Hypixel stats and Mojang profile as
well as store the Discord user.
---
</details>
<details>
 <summary><code>POST</code> <code><b>/daily/:id</b></code> <code>Claim a daily reward for a Discord user by id</code></summary>

##### Request Parameters

- `id` the Discord id of the user whose daily should be claimed

##### Response Body (JSON)

  ``` go
  type DiscordDailyResponse struct {
    Actual string `json:"actual"`
  }
  ```

##### Method (on Client)
`ClaimDaily(id string) (*DiscordDailyResponse, error)` Daily is used to claim the daily reward for a Discord user.
The backend will return the user's updated stats.
---
</details>
<details>
  <summary><code>POST</code> <code><b>/bot-login</b></code> <code>Log in port for a Discord bot instance</code> <code>no auth</code></summary>

##### Request Body (JSON)

  ``` go
  type DiscordBotLoginRequest struct {
	Pwd string `json:"pwd" query:"pwd"`
  }
  ```

##### Response Body (JSON)

  ``` go
  type DiscordBotLoginResponse struct {
    Actual string `json:"actual"`
  }
  ```

##### Method (on Client)
`BotLogin(input *DiscordBotLoginRequest) (*DiscordBotLoginResponse, error)` BotLogin is used to login the bot to the
Discord API. No token is required for this endpoint.
---
</details>
<details>
 <summary><code>GET</code> <code><b>/leaderboard/:page</b></code> <code>Get the leaderboard by page</code></summary>

##### Request Parameters

- `page` the page of the leaderboard to get (zero-based indexing)

##### Response Body (JSON)

  ``` go
  type DiscordLeaderboardResponse []struct{
	DiscordID string  `json:"discord_id"`
	Level     int     `json:"level"`
	XP        float64 `json:"xp"`
  }
  ```

##### Method (on Client)
`Leaderboard(page int) (*DiscordLeaderboardResponse, error)` Leaderboard is used to get the leaderboard for all verified 
Discord users. NOTE: The pagination uses zero-based indexing.
---
</details>
<details>
 <summary><code>GET</code> <code><b>/member/:id</b></code> <code>Get the stats of a Discord member</code></summary>

##### Request Parameters

- `id` the Discord id of the user whose stats should be retrieved

##### Response Body (JSON)

  ``` go
  type DiscordMemberResponse struct {
	DiscordID   string  `json:"discord_id"`
	Name        string  `json:"name"`
	Nick        string  `json:"nick"`
	XP          float64 `json:"xp"`
	LastDailyAt string  `json:"last_daily_at"`
	Level       int     `json:"level"`
	Streak      int     `json:"streak"`
  }
  ```

##### Method (on Client)
`Member(id string) (*DiscordMemberResponse, error)` Member is used to get the stats for a specific Discord user.
---
</details>

## Environment Variables
- `port` - The port to listen on
- `db_host` - The host of the database
- `db_port` - The port of the database
- `db_user` - The user of the database
- `db_pwd` - The password of the database
- `db_name` - The name of the database
- `db_namespace` - The namespace of the database
- `jwt_secret` - The secret used to sign JWTs
- `bot_pwd` - The password of the Discord bot
- `hypixel_api_key` - The Hypixel API key

## ToDo
- [x] Migrate from GraphQL to REST
- [x] Migrate from Neo4j to SurrealDB
- [ ] Add response tables to API definition
- [ ] Create REST documentation
- [ ] Create HTTP bindings
- [ ] Document HTTP bindings
- [ ] Create WebSocket bindings
- [ ] Document WebSocket bindings
- [ ] Create SSE bindings
- [ ] Document SSE bindings
- [x] Migrate from Mux to Echo
- [x] Implement authentication
