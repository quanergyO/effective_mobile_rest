package postgres

import (
	"effective_mobile_rest/types"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"log/slog"
	"strings"
)

type Cars struct {
	db *sqlx.DB
}

func NewCars(db *sqlx.DB) *Cars {
	return &Cars{
		db: db,
	}
}

func (r *Cars) CreateCars(cars []types.Cars) error {
	tx, err := r.db.Beginx()
	values := ""
	for _, car := range cars {
		values += fmt.Sprintf("('%s', '%s', '%s'),", car.People.Name, car.People.Surname, car.People.Patronymic)
	}
	values = values[:len(values)-1]
	query := fmt.Sprintf("INSERT INTO %s (name, surname, patronymic) VALUES %s RETURNING ID", peopleTable, values)
	slog.Debug(query)
	_, err = r.db.Exec(query)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	values = ""
	for _, car := range cars {
		values += fmt.Sprintf("('%s', '%s', '%s', '%d', (SELECT id FROM %s WHERE name='%s' AND surname='%s' LIMIT 1)),",
			car.RegNum, car.Mark, car.Model, car.Year, peopleTable, car.People.Name, car.People.Surname)
	}
	values = values[:len(values)-1]
	rowsName := "regNum, mark, model, year, owner_id"
	query = fmt.Sprintf("INSERT INTO %s (%s) VALUES %s", carsTable, rowsName, values)
	slog.Debug(query)
	_, err = r.db.Exec(query)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *Cars) GetAllCars(limit, offset int) ([]types.Cars, error) {
	carsDAO := make([]types.CarsDAO, 0)
	rowsName := "regNum, mark, model, year, name, surname, patronymic"
	query := fmt.Sprintf("SELECT %s FROM %s JOIN %s ON %s.id=%s.owner_id LIMIT $1 OFFSET $2", rowsName, carsTable, peopleTable, peopleTable, carsTable)
	slog.Info("GetAllCars query:", query)
	if err := r.db.Select(&carsDAO, query, limit, offset); err != nil {
		slog.Info("GetAllCars error:", err.Error())
		return []types.Cars{}, err
	}

	cars := make([]types.Cars, len(carsDAO))
	for i, car := range carsDAO {
		cars[i].RegNum = car.RegNum
		cars[i].Mark = car.Mark
		cars[i].Model = car.Model
		cars[i].Year = car.Year
		cars[i].People = types.People{
			Name:       car.Name,
			Surname:    car.Surname,
			Patronymic: car.Patronymic,
		}
	}

	return cars, nil
}

func (r *Cars) DeleteCar(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", carsTable)
	slog.Debug("DeleteCar query:", query)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *Cars) UpdateCarById(id int, newData types.UpdateCar) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if newData.RegNum != nil {
		setValue = append(setValue, fmt.Sprintf("regNum=$%d", argId))
		args = append(args, *newData.RegNum)
		argId++
	}

	if newData.Year != nil {
		setValue = append(setValue, fmt.Sprintf("year=$%d", argId))
		args = append(args, *newData.Year)
		argId++
	}

	if newData.Model != nil {
		setValue = append(setValue, fmt.Sprintf("model=$%d", argId))
		args = append(args, *newData.Model)
		argId++
	}

	if newData.Mark != nil {
		setValue = append(setValue, fmt.Sprintf("mark=$%d", argId))
		args = append(args, *newData.Mark)
		argId++
	}

	setQuery := strings.Join(setValue, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s.id=%d", carsTable, setQuery, carsTable, id)
	log.Println(query)
	log.Println(args)
	slog.Debug("UpdateCarById query:", query)
	_, err := r.db.Exec(query, args...)

	return err
}
