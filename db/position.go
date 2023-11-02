package db

import "github.com/Rosya-edwica/statistic-vacancy/models"

func (d *Database) GetPositions() (positions []models.Position) {
	query := "select position_id, city_id from h_vacancy where position_id != 0 and city_id != 0 group by position_id, city_id order by count(*) desc"
	rows, err := d.Connection.Query(query)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var city_id, positiond_id int
		err = rows.Scan(&positiond_id, &city_id)
		checkErr(err)
		positions = append(positions, models.Position{
			CityId:     city_id,
			PositionId: positiond_id,
		})
	}
	return
}

func (d *Database) GetPythonPositions() (positions []models.Position) {
	query := "select position_id, city_id from h_vacancy where position_id IN (323, 675, 676, 677, 678) AND city_id != 0 group by position_id, city_id order by count(*) desc"
	rows, err := d.Connection.Query(query)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var city_id, positiond_id int
		err = rows.Scan(&positiond_id, &city_id)
		checkErr(err)
		positions = append(positions, models.Position{
			CityId:     city_id,
			PositionId: positiond_id,
		})
	}
	return
}
