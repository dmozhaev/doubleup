# Double Up Poker Game

## Rules
In Double Up Poker, the player tries to guess whether the single card in front of them is small (numbers from 1 to 6) or large (numbers from 8 to 13). If the card is 7, the player always loses.

## Basic technical info
- The game was tested on Windows and Linux (Ubuntu) platforms.
- Backend is written in Go (v1.22.2) and Java (v22) languages, but the solution made in Go is more mature as it contains automated tests.
- Database is made using PostgreSQL (v16.2).

## Setting up
Note: This guide is for running the Go side of the project along with the database.

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
This part is for the backend written in Go. Note: It uses port 8080, so it should be free.

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
This is the game state when there is no money in play for the current player. This is the original state of the game, and in addition to that, the game ends here every time the player loses or withdraws money that is currently in play. In this state, only the "start game" operation is possible, and this can be tested, for example:
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

The game moves to this state when the game has already been started. In this state, the player can either continue doubling up or withdraw money that is currently in play, but they are not able to start a new game until the current one is played to the end (which would either result in a loss or a withdrawal of money).

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
This part describes all the error responses that the game engine might return during the game process. Typically, backend handler functions return JSON with an error string field in the following format:
```bash
{"error":"<handler name>: <description of error>"}
```
### Error description
#### Serialization / dto format errors
- Only POST method is allowed
```bash
{"error":"<handler name>: Method not allowed"}
```
- Player ID is in the incorrect format, produces a built-in deserialization error
```bash
{"error":"<handler name>: invalid UUID length: 10"}
```
- Bet size is in the incorrect format; it should be a number
```bash
{"error":"<handler name>: json: cannot unmarshal string into Go struct field PlayStartRequestDto.BetSize of type int64"}
```
- Player choice is in the incorrect format; it should be either `SMALL` or `LARGE` in string format
```bash
{"error":"<handler name>: <validator name>: choice is invalid"}
```
#### Game engine -related errors
- Player ID is missing from DB
```bash
{"error":"<handler name>: Player not found, id: 9ff66fec-17c4-4594-aa03-d053fc036bad"}
```
- Starting the game is only possible when the game has not yet begun
```bash
{"error":"<handler name>: <validator name>: there should be no money in play in order to start!"}
```
- Continuing the game and withdrawing are only possible after the game has begun
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
Unit tests reside in the same folders along with production files. Currently, there are unit tests for both the service and validation layers of the software. They can be run using, for example:
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
The test DB connection is currently the same as for the "production", there ia a hard-coded test user named `heikki` with hard-coded UUID of `01162f1f-0bd9-43fe-8032-fa9590ee0e7e`. Currently there are 6 integration tests that can be divided into 2 categories:

- basic API flow tests (3 pieces) which test handlers and possible error messages
- process-related tests with db checking (3 pieces) which test the game flow

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

### Money in play concept (`money_in_play` field inside player DB table)
- The game remembers the state even if the backend goes down or experiences another glitch, so the player is able to continue playing at any time
- The player is not able to manipulate the bet size after a win (by modifying the bet size manually), which was the main reason why the `/play/continue` API route was added

### Database in docker with in-built test user seeding
- Does not require manual local database installation; new users and roles are set up automatically, making everything work straight out of the box.
- The database can be accessed, for example, using `psql` (command line), with the password: `admin_password`:
```bash
psql -h localhost -p 5433 -U postgres -d doubleupdb
```
- The docker-compose.yml file has been added for smooth container control, even though there is only one container at the moment.
- `postgres:16.2` is specified inside the Dockerfile as using `postgres:latest` label might change at some point and potentially break the software

### Float-point arithmetic is avoided at the moment
- It's easier to operate and store values when they are of the int type compared to float, which is more error-prone.
- Money-type values can be seen as cents (note: this approach works until we use divisions/percentages where we want very precise calculations).
- This approach can be changed in the future if we need more precision.

### Audit logging
- Audit logging goes straight into the database, which makes it easier to manage and investigate, but it can become problematic as the database size grows.

### Rate limiting
- Fixed Window Rate Limiting was chosen (for simplicity reasons), even though Sliding Window Rate Limiting might be better for this case. However, the implementation might become more complex.

## Technical trade offs
Trade offs made during development and potential improvements for future iterations:
- The current repository/dao solution, while somewhat generic, is still not perfectly generic. Copy-paste could still be reduced. This can be done by implementing a generic interface and different implementations for each model, similar to the approach used in Java.
- The database (`db *sql.DB`) is passed all the way from the `main.go` file down to the repository level, which is a bit overheadish. Some kind of wrapper should be considered here.
- The same database (and connection string) is used for integration tests and for "production," which definitely should be changed by implementing our own test database with test users.
- All error processing ends up being in handler files. While this might be the Go way, it does make the code a bit messy.
- The program passes English texts to the frontend. This can be improved by introducing translation keys, which can then be translated into different languages on the frontend. This approach definitely helps to make the product international.
- The deserialization technique might be improved (e.g., by using another deserialization library), as currently, there is some manual work required to check all potential corner cases.
- A Golang linter should be considered, especially if the software grows. This ensures code consistency across .go files, especially if there are several developers in the project.
- Currently, there is no hot reload in the development environment. Programmers have to restart the program every time a change to the source code is made. While not a problem when the codebase is small, it might become problematic as it grows.
