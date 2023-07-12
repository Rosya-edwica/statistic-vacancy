package main

import (
	"os"
	"strings"

	"github.com/Rosya-edwica/statistic-vacancy/db"
	"github.com/Rosya-edwica/statistic-vacancy/models"
	"github.com/joho/godotenv"
)

type Experience struct {
	Name string
	Level int
}

var ExperienceList = []Experience{
	Experience{Name: "Нет опыта", Level: 1},
	Experience{Name: "От 1 года до 3 лет", Level: 2},
	Experience{Name: "От 3 до 6 лет", Level: 3},
	Experience{Name: "Более 6 лет", Level: 4},
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Настройте переменные окружения!")
	}
}

func main() {
	database := db.Database{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		User: os.Getenv("USER"),
		Password: os.Getenv("PASS"),
		Name: os.Getenv("NAME"),
	}
	database.Connect()
	positions := database.GetPositions()[10000:]
	for _, pos := range positions {
		vacancies := database.GetVacancies(pos)
		statistic := buildStatistic(vacancies)
		database.SaveStatistic(statistic)
	}
}

func buildStatistic(vacancies []models.Vacancy) (statistic models.Statistic) {
	if len(vacancies) == 0 {
		return
	}

	statistic.VacanciesCount = len(vacancies)
	statistic.ListVacancyId = getListVacancyId(vacancies)
	statistic.CityId = vacancies[0].CityId
	statistic.PositionId = vacancies[0].PositionId
	statistic.Areas = getAreas(vacancies, "areas")
	statistic.Specs = getAreas(vacancies, "specs")
	statistic.AverageExperience = getAverageExperience(vacancies)
	statistic.AverageSalary = getAverageSalary(vacancies)
	return
}

func getListVacancyId(vacancies []models.Vacancy) (ids []string) {
	for _, item := range vacancies {
		ids = append(ids, item.Id)
	}
	return
}

func getAreas(vacancies []models.Vacancy, typeArea string) (items []string) {
	for _, vac := range vacancies {
		switch typeArea {
			case "areas": items = append(items, strings.Split(vac.Area, "|")...)
			case "specs": items = append(items, strings.Split(vac.Spec, "|")...)
		} 
	}
	return uniqueStringsInList(items)
}

func uniqueStringsInList(items []string) (uniqueItems []string){
	inResult := make(map[string]bool)
    for _, str := range items {
        if _, ok := inResult[str]; !ok {
            inResult[str] = true
            uniqueItems = append(uniqueItems, str)
        }
    }
    return uniqueItems

}

func getAverageExperience(vacancies []models.Vacancy) (experience string) {
	var sum int
	for _, item := range vacancies {
		switch item.Experience {
		case "Нет опыта": sum += 1
		case "От 1 года до 3 лет": sum += 2
		case "От 3 до 6 лет": sum += 3
		case "Более 6 лет": sum += 4
		}
	}
	average := sum / len(vacancies)
	for _, i := range ExperienceList {
		if i.Level == average {
			return i.Name
		}
	}
	return ""
}

func getAverageSalary(vacancies []models.Vacancy) (average float64) {
	var sum float64
	var count int
	for _, vac := range vacancies {
		if vac.SalaryFrom != 0 && vac.SalaryTo != 0 {
			vacAverage := (vac.SalaryFrom + vac.SalaryTo) / 2
			sum += vacAverage
			count++
		} else if vac.SalaryFrom != 0 {
			sum += vac.SalaryFrom
			count++
		} else if vac.SalaryTo != 0 {
			sum += vac.SalaryTo
			count++
		}
	}
	average = sum / float64(count)
	return
}