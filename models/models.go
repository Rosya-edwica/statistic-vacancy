package models


type Position struct {
	PositionId int
	CityId int
}

type Statistic struct {
	PositionId int
	CityId int
	ListVacancyId []string
	VacanciesCount int
	Areas []string
	Specs []string
	AverageExperience string
	AverageSalary float64
}

type Vacancy struct {
	PositionId int
	CityId int
	Id string
	Experience string
	SalaryFrom float64
	SalaryTo float64
	Spec string
	Area string
}