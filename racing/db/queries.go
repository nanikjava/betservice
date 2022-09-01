package db

const (
	racesList = "list"
	raceGet   = "get"
)

func getRaceQueries() map[string]string {
	return map[string]string{
		racesList: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time,
				CASE
           			WHEN advertised_start_time < datetime() THEN
               			'CLOSED'
           			ELSE
               			'OPEN'
           		END AS status
			FROM races
		`,
		raceGet: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time,
				CASE
           			WHEN advertised_start_time < datetime() THEN
               			'CLOSED'
           			ELSE
               			'OPEN'
           		END AS status
			FROM races WHERE id=?
		`,
	}
}
