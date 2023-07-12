package db

import (
	"fmt"
	"strings"

	"github.com/Rosya-edwica/statistic-vacancy/models"
)

func (d *Database) SaveStatistic(statistic models.Statistic) {
	columns := buildPatternInsertValues(8)
	smt := fmt.Sprintf(`INSERT INTO h_vacancy_statistic(position_id, city_id, vacancy_id, count, average_salary, average_experience, prof_areas, specs) VALUES %s ON DUPLICATE KEY UPDATE average_salary = %f`, columns, statistic.AverageSalary)
	tx, _ := d.Connection.Begin()
	_, err := d.Connection.Exec(smt, statistic.PositionId, statistic.CityId, strings.Join(statistic.ListVacancyId, "|"), statistic.VacanciesCount, statistic.AverageSalary, statistic.AverageExperience, strings.Join(statistic.Areas, "|"), strings.Join(statistic.Specs, "|"))
	if err != nil {
		if err.Error() == "Error 1054 (42S22): Unknown column 'NaN' in 'field list'" {
			fmt.Println("Error: Unknown column 'NaN' in 'field list'\n", smt, "\n", statistic)
		} else {
			checkErr(err)
			return	
		}
	}
	err = tx.Commit()
	checkErr(err)
	fmt.Printf("Успешно сохранили статистику по %d вакансиям\n", len(statistic.ListVacancyId))
}