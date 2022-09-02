package db

const (
	soccerList = "list"
)

// getSoccerQueries to return the query to be use for soccer sport
func getSoccerQueries() map[string]string {
	return map[string]string{
		soccerList: `
			SELECT 
				id, 
				league, 
				team_home, 
				team_home_manager, 
				team_away, 
				team_away_manager,
				advertised_start_time
			FROM soccer
		`,
	}
}
