# Double Up Poker Game

## Rules
In Double Up Poker, the player tries to guess whether the single card in front of them is small (numbers from 1 to 6) or large (numbers from 8 to 13). If the card is 7, the player always loses.

## Basic technical info
- The game was tested on Windows and Linux (Ubuntu) platforms.
- Backend is written in Go (v1.22.2) and Java (v22) languages, but the solution made in Go is more mature as it contains automated tests.
- Database is made using PostgreSQL (v16.2).

## Setting up
Note: this guide is for running Go side of the project along with database.

### Checking requirements
- `go` version 1.22.2
- `docker` version 26.0.0

### Installing Go dependencies
- Open command prompt or terminal
- Navigate to the directory containing Go project
```bash
cd <project root>/go
```
- Install required dependencies  
Note: Optional step, will be downloaded on the first backend start.
```bash
go get -u github.com/lib/pq
go get -u github.com/google/uuid
go get -u -t double_up/service
```

### Starting DB up
Note: database uses port 5433 so it should be free

- Navigate to the root directory of the project
```bash
cd <project root>
```
- Start docker DB container
```bash
docker-compose up -d
```

### Starting backend
This part is for backend written in Go. Note: uses port 8080 so it should be free.

- Navigate to the directory containing Go project
```bash
cd <project root>/go
```
- Run the project 
```bash
go run main.go
```

### Shutting DB down
- Navigate to the root directory of the project
```bash
cd <project root>
```
- Shut down docker DB container 
```bash
docker-compose down
```

## Playing the game
This part describes the game flow. The game engine basically has 2 states.

### Initial state
This is the game state when there is no money in play for the current player. This is the original state of the game and in addition to that the game ends here every time when the player loses or withdraws money that are currently in play. In this state, only start game -operation is possible, this can be tested e.g.:
- Windows / powershell
```bash
Invoke-RestMethod -Method Post -Uri "http://localhost:8080/play/start" -Body (ConvertTo-Json @{
	playerId = "01162f1f-0bd9-43fe-8032-fa9590ee0e7e"
	betSize = 10
	choice = "SMALL"
}) -ContentType "application/json"
```
- Linux / bash
```bash
curl -X POST -H "Content-Type: application/json" -d '{"playerId":"01162f1f-0bd9-43fe-8032-fa9590ee0e7e", "betSize":10, "choice":"SMALL"}' http://localhost:8080/play/start
```

### Game-in-progress state
The game moves on to this state when the game has already been started. In this state, player can either continue doubling up or withdraw money that are currently in-play, but is not able to start new game until the current one is played to the end (which would be either loss or money withdrawal).

#### Continuing the game
Can be tested e.g.:
- Windows / powershell
```bash
Invoke-RestMethod -Method Post -Uri "http://localhost:8080/play/continue" -Body (ConvertTo-Json @{
	playerId = "01162f1f-0bd9-43fe-8032-fa9590ee0e7e"
	choice = "LARGE"
}) -ContentType "application/json"
```
- Linux / bash
```bash
curl -X POST -H "Content-Type: application/json" -d '{"playerId":"01162f1f-0bd9-43fe-8032-fa9590ee0e7e", "choice":"LARGE"}' http://localhost:8080/play/continue
```

#### In-play money withdrawal
Can be tested e.g.:
- Windows / powershell
```bash
Invoke-RestMethod -Method Post -Uri "http://localhost:8080/withdraw/withdrawmoney" -Body (ConvertTo-Json @{
	playerId = "01162f1f-0bd9-43fe-8032-fa9590ee0e7e"
}) -ContentType "application/json"
```
- Linux / bash
```bash
curl -X POST -H "Content-Type: application/json" -d '{"playerId":"01162f1f-0bd9-43fe-8032-fa9590ee0e7e"}' http://localhost:8080/withdraw/withdrawmoney
```

## Documentation on error responses
This part describes all the error responses that the game engine might return during the game process. Typically backend handler functions return JSON with error -string field of the following format:
```bash
{"error":"<handler name>: <description of error>"}
```
### Error description
#### Serialization / dto format errors
- Only POST method is allowed
```bash
{"error":"<handler name>: Method not allowed"}
```
- Player ID is of incorrect format, built-in deserialization error
```bash
{"error":"<handler name>: invalid UUID length: 10"}
```
- Bet size is of incorrect format, should be number
```bash
{"error":"<handler name>: json: cannot unmarshal string into Go struct field PlayStartRequestDto.BetSize of type int64"}
```
- Player choice is of incorrect format, should be either "SMALL" or "LARGE" in string format
```bash
{"error":"<handler name>: <validator name>: choice is invalid"}
```
#### Serialization / dto format errors
- Player ID is missing from DB
```bash
{"error":"<handler name>: Player not found, id: 9ff66fec-17c4-4594-aa03-d053fc036bad"}
```
- Starting the game is possible only when the game has not been yet
```bash
{"error":"<handler name>: <validator name>: there should be no money in play in order to start!"}
```
- Continuing the game and withdrawing are possible only after the game has been started
```bash
{"error":"<handler name>: <validator name>: money should be in play already!"}
```
- Bet size is zero or negative
```bash
{"error":"<handler name>: <validator name>: bet is too small"}
```
- Bet size exceeds available funds
```bash
{"error":"<handler name>: <validator name>: bet is too large, insufficient funds"}
```

## Running unit and integration tests
There are both unit and integration tests existing for the project.
### Unit tests
Unit tests reside in same folders along with production files. Currently there are unit tests for both service and validation layer of the software. The can be run using e.g.
```bash
cd <project root>/go
go test ./service
go test ./validation
```
Expected output (example):
```bash
ok      double_up/service    0.064s
ok      double_up/validation    0.112s
```
### Integration tests
Integration tests reside in own folder:
```bash
<project root>/integration/tests
```
The DB connection is currently the same as for the "production", there ia a hard-coded test user named `heikki` with hard-coded UUID of `01162f1f-0bd9-43fe-8032-fa9590ee0e7e`. Currently there are 6 integration tests that can be divided into 2 categories:

- basic API flow tests (3 pieces) which test handlers and possible error messages
- process-related tests with db checking (3 pieces) which test

Integration tests can be run using e.g.:
```bash
cd <project root>/go
go test ./integration/tests
```
Expected output (example):
```bash
ok      double_up/integration/tests     0.839s
```

## Technical choices made during implementation

### Money in play -concept (`money_in_play` field inside player DB table)
- The game remembers the state even if backend goes down or some other glitch happens, so the player is able to continue playing at any time
- player is not able to hack bet size after win (modifying the bet size manually), which was the reason why /play/continue -API route was added

### Db-in-docker with in-built test user seeding
- does not require manual local DB installation, new users and roles as everything works straight out of the box
- DB can be accessed e.g (cmd, password: admin_password) using psql
```bash
psql -h localhost -p 5433 -U postgres -d doubleupdb
```
- docker-compose.yml file added for smooth container control, even though there is only one container at the moment
- `postgres:16.2` is put inside Dockerfile as `postgres:latest` label might change at some point and break the software

### Float-point arithmetic is avoided at the moment
- easier to operate and store values when they are of int type compared to float, which is more error-prone
- money -type values can be seen as cents (note: this approach works until we use divisions/percentages where we want very precise calculations)
- can be changed in the future if we need more precision

### Audit logging
- audit logging goes straight into db, which makes it easier to manage and investigate, but can become problematic as db size grows

### Rate limiting
- Fixed Window Rate Limiting was chosen (for simplicity reasons) even though Sliding Window Rate Limiting might be better for this case, but the implementation might be more complex

## Technical trade offs
Trade offs made during development and potential improvements for future iterations:
- Current repository solution (while is it somewhat generic) is still not perfectly generic, copy-paste still could be reduced. This can be done by impleneting generic interface and different implementations for each model (a bit like in Java)
- DB (db *sql.DB) is passed all the way from main.go file down to repository level, which is a bit overheadish. Some kind of wrapper should be considered here
- The same DB (and connection string) is used for integration tests and for "production", which definitely should be changed by implementing own test DB with test users
- All error processing ends up being in handler files (might be the golang way, though, but makes the code a bit messy)
- The program passes English texts to frontend, this can be improved by e.g. introducing translation keys, which in turn can be translated into different languages on frontend. Definitely helps to make the product international
- Deserialization technique might be impoved (e.g. another deserialization library used), as currently there is some manual work required to check all potential corner cases. 
- Golang linter should be considered, especially if the software grows. This insert code consistency across .go files, especially is there are several developers in the project
- Currently there is no hot reload in dev environment, programmer has to restart the program every time the change to source code is made. Not a problem when codebase is small, might be problematic when it grows.
