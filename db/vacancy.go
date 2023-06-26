package db

import (
	"fmt"

	"github.com/Rosya-edwica/statistic-vacancy/models"
)

func (d *Database) GetVacancies(pos models.Position) (vacancies []models.Vacancy){
	query := fmt.Sprintf("select id, experience, salary_from, salary_to, prof_areas, specs FROM h_vacancy WHERE position_id=%d AND city_id=%d", pos.PositionId, pos.CityId)
	rows, err := d.Connection.Query(query)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var id, experience, prof_areas, specs string
		var salary_from, salary_to float64
		err = rows.Scan(&id, &experience, &salary_from, &salary_to, &prof_areas, &specs)
		vacancies = append(vacancies, models.Vacancy{
			Id: id,
			Experience: experience,
			SalaryFrom: salary_from,
			SalaryTo: salary_to,
			Area: prof_areas,
			Spec: specs,
			CityId: pos.CityId,
			PositionId: pos.PositionId,
		})
	}
	return
}