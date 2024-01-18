# FDS Core

## API
### `/discord`
<details>
 <summary><code>POST</code> <code><b>/signup</b></code> <code>Meant for linking Discord to Minecraft accounts</code></summary>

  ##### Request Body (JSON)

  ``` go
  type DiscordSignupInput struct {
    ID   string `json:"id"`
    Nick string `json:"nick"`
  }
  ```  

  ##### Responses
  
  | http code | content-type       | response                  |
  |-----------|--------------------|---------------------------|
  | `200`     | `application/text` | `"success"`               |
  | `400`     | `application/text` | `"invalid request query"` |
  | `500`     | `application/text` | `"internal server error"` |
</details>
<details>
 <summary><code>POST</code> <code><b>/:id/daily</b></code> <code>Claim a daily reward for a Discord user by id</code></summary>

  ##### Request Parameters
  - `id` the Discord id of the user whos daily should be claimed

  ##### Responses
  
  | http code | content-type       | response                  |
  |-----------|--------------------|---------------------------|
  | `200`     | `application/json` | `"success"`               |
  | `400`     | `application/json` | `"invalid request query"` |
  | `404`     | `application/json` | `"invalid request query"` |
  | `500`     | `application/json` | `"internal server error"` |
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

## ToDo
- [x] Migrate from GraphQL to REST
- [x] Migrate from Neo4j to SurrealDB
- [ ] Create REST documentation
- [ ] Create HTTP bindings
- [ ] Document HTTP bindings
- [ ] Create WebSocket bindings
- [ ] Document WebSocket bindings
- [ ] Create SSE bindings
- [ ] Document SSE bindings
- [x] Migrate from Mux to Echo
- [x] Implement authentication
