package db

import (
	"time"

	"syreclabs.com/go/faker"
)

// seed for generating random (test/mock) data into the database
func (r *soccerRepo) seed() error {
	statement, err := r.db.Prepare(`CREATE TABLE IF NOT EXISTS soccer (id INTEGER PRIMARY KEY, league TEXT, team_home TEXT, team_home_manager TEXT, team_away TEXT, team_away_manager TEXT, advertised_start_time DATETIME)`)
	if err == nil {
		_, err = statement.Exec()
	}

	for i := 1; i <= 100; i++ {
		statement, err = r.db.Prepare(`INSERT OR IGNORE INTO soccer(id, league, team_home, team_home_manager, team_away, team_away_manager, advertised_start_time) VALUES (?,?,?,?,?,?,?)`)
		if err == nil {
			_, err = statement.Exec(
				i,
				//TODO: faker library does not have specific soccer name, maybe we can
				// create an extension to faker or use our own name ?
				//this is a test data we are using
				//Bitcoin address - for league
				//Team - for team_home and team_away
				//Company - for team_home_manager and team_away_manager
				faker.Bitcoin().Address(),
				faker.Team().Name(),
				faker.Company().Name(),
				faker.Team().Name(),
				faker.Company().Name(),
				faker.Time().Between(time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 2)).Format(time.RFC3339),
			)
		}
	}

	return err
}
