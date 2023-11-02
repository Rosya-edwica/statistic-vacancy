package db

import (
	"fmt"
	"math"
	"strings"

	"github.com/Rosya-edwica/statistic-vacancy/logger"
	"github.com/Rosya-edwica/statistic-vacancy/models"
)

func (d *Database) SaveStatistic(stat models.Statistic) {
	statId := d.CheckStatisticExist(stat.PositionId, stat.CityId)
	if statId == 0 {
		logger.LogInfo.Printf(
			"Создали стату\tPositionId: %d\tCityId:%d\tCount:%d\tExp:%s\tSalary:%f\tSpecsCount:%d\tAreasCount:%d\t",
			stat.PositionId, stat.CityId, stat.VacanciesCount, stat.AverageExperience, stat.AverageSalary,
			len(stat.Specs), len(stat.Areas),
		)
		d.SaveNewStatistic(stat)

	} else {
		logger.LogInfo.Printf(
			"Обновляем стату: %d.\tPositionId: %d\tCityId:%d\tCount:%d\tExp:%s\tSalary:%f\tSpecsCount:%d\tAreasCount:%d\t",
			statId, stat.PositionId, stat.CityId, stat.VacanciesCount, stat.AverageExperience, stat.AverageSalary,
			len(stat.Specs), len(stat.Areas),
		)
		d.UpdateStatistic(statId, stat)

	}
}

func (d *Database) SaveNewStatistic(statistic models.Statistic) {
	columns := buildPatternInsertValues(8)
	smt := fmt.Sprintf(`INSERT INTO h_vacancy_statistic(position_id, city_id, vacancy_id, count, average_salary, average_experience, prof_areas, specs) VALUES %s`, columns)
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
	// fmt.Printf("Успешно сохранили статистику по %d вакансиям\n", len(statistic.ListVacancyId))
}

func (d *Database) UpdateStatistic(statId int64, stat models.Statistic) {
	var salary float64
	if !math.IsNaN(stat.AverageSalary) {
		salary = stat.AverageSalary
	}

	query := fmt.Sprintf(`
		UPDATE h_vacancy_statistic
		SET
			vacancy_id = '%s',
			count = %d,
			average_salary = %.3f,
			average_experience = '%s',
			prof_areas = '%s',
			specs = '%s'
		
		WHERE id = %d
	`, strings.Join(stat.ListVacancyId, "|"), stat.VacanciesCount, salary, stat.AverageExperience,
		strings.Join(stat.Areas, "|"), strings.Join(stat.Specs, "|"), statId)
	tx, _ := d.Connection.Begin()
	_, err := d.Connection.Exec(query)
	checkErr(err)
	tx.Commit()
}

func (d *Database) CheckStatisticExist(posId int, cityId int) (statId int64) {
	query := fmt.Sprintf("SELECT id FROM h_vacancy_statistic WHERE position_id = %d AND city_id = %d LIMIT 1", posId, cityId)
	err := d.Connection.QueryRow(query).Scan(&statId)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return 0
		}
		return 0
	}
	return
}
