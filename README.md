## Existing Service

The existing racing service has been modified, following are the `curl` commands used for testing:

1. Using `filter` to filter based on `meeting_ids`
```
curl -X "POST" "http://localhost:8000/v1/list-races"      -H 'Content-Type: application/json'     \
  -d '{"filter": {"meeting_ids": [1, 2]}}' | jq .
```

2. Using `visible` filter to filter only races that are visible
```  
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json'     \
  -d '{"filter": {"meeting_ids": [1, 2], "visible":"1"}}' | jq .
```
 
3. Using `order_by` parameter to specify the sorting order. Format is `field_name asc/desc` 
```
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json'     \
  -d '{"filter": {"visible":"1"}, "order_by":"advertised_start_time asc"}' | jq .
```

## New Service

### How to run and test ?

1. The new service is inside `soccer` folder as follows:
```
cd <your_main_folder>/soccer
go run main.go
```

The service will run on port `9001`

2. Make sure the API Server is running, for example in my machine it is run as follows

```
cd <your_main_folder>/api
go run main.go
```

The API server will be running on port `8000`

3. To test the soccer service use the following `curl` command:

```
curl -X "GET" "http://localhost:8000/v1/list-matches" | jq .
```

### Internals

1. The service use it's own table called `soccer.db` having the following structure:

```
create table soccer
(
    id                    INTEGER primary key,
    league                TEXT,
    team_home             TEXT,
    team_home_manager     TEXT,
    team_away             TEXT,
    team_away_manager     TEXT,
    advertised_start_time DATETIME
);
```
