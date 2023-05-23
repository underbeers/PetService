package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/underbeers/PetService/pkg/models"
	"strings"
)

type PetCardPostgres struct {
	db *sqlx.DB
}

func NewPetCardPostgres(db *sqlx.DB) *PetCardPostgres {
	return &PetCardPostgres{db: db}
}

func (r *PetCardPostgres) Create(petCard models.PetCard) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	createPetCardQuery := fmt.Sprintf("INSERT INTO %s (pet_type_id, user_id, pet_name, breed_id, origin_photo, "+
		"thumbnail_photo, birth_date, male, color, care, pet_character, pedigree, sterilization, vaccinations) VALUES "+
		"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)", petCardTable)
	_, err = tx.Exec(createPetCardQuery, petCard.PetTypeId, petCard.UserId, petCard.Name, petCard.BreedId, petCard.Photo,
		petCard.ThumbnailPhoto, petCard.BirthDate, petCard.Male, petCard.Color, petCard.Care, petCard.Character,
		petCard.Pedigree, petCard.Sterilization, petCard.Vaccinations)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

func createPetCardQuery(filter models.PetCardFilter) string {

	query := fmt.Sprintf("SELECT pc.id, pc.pet_type_id, pc.user_id, pc.pet_name, pc.breed_id, pc.origin_photo, "+
		"pc.thumbnail_photo, pc.birth_date, pc.male, CASE pc.male WHEN True THEN 'Мальчик' WHEN False THEN 'Девочка' "+
		"END AS gender, pc.color, pc.care, pc.pet_character, pc.pedigree, pc.sterilization, pc.vaccinations, "+
		"pt.pet_type, br.breed_name FROM %s pc ",
		petCardTable)

	query += "INNER JOIN pet_type pt ON pc.pet_type_id = pt.id INNER JOIN breed br ON pc.breed_id = br.id "

	if filter.PetCardId != 0 || filter.UserId != uuid.Nil || filter.Gender != "" || filter.PetTypeId != 0 ||
		filter.BreedId != 0 {

		query += "WHERE "
		setValues := make([]string, 0)

		if filter.PetCardId != 0 {
			setValues = append(setValues, fmt.Sprintf("pc.id = %d", filter.PetCardId))
		}

		if filter.PetTypeId != 0 {
			setValues = append(setValues, fmt.Sprintf("pc.pet_type_id = %d", filter.PetTypeId))
		}

		if filter.BreedId != 0 {
			setValues = append(setValues, fmt.Sprintf("pc.breed_id = %d", filter.BreedId))
		}

		if filter.UserId != uuid.Nil {
			setValues = append(setValues, fmt.Sprintf("pc.user_id = '%s'", filter.UserId.String()))
		}

		if filter.Gender != "" {
			if filter.Gender == "male" {
				setValues = append(setValues, fmt.Sprintf("pc.male = True"))
			} else if filter.Gender == "female" {
				setValues = append(setValues, fmt.Sprintf("pc.male = False"))
			}
		}

		query += strings.Join(setValues, " AND ")

	}

	return query
}

func (r *PetCardPostgres) GetAll(filter models.PetCardFilter) ([]models.PetCard, error) {
	var lists []models.PetCard

	query := createPetCardQuery(filter)
	err := r.db.Select(&lists, query)

	return lists, err
}

func createMainCardInfoQuery(filter models.PetCardFilter) string {

	query := fmt.Sprintf("SELECT pc.id, pc.pet_name, pc.thumbnail_photo, pc.birth_date, "+
		"CASE pc.male WHEN True THEN 'Мальчик' WHEN False THEN 'Девочка' END AS gender, pt.pet_type, br.breed_name "+
		"FROM %s pc ",
		petCardTable)

	query += "INNER JOIN pet_type pt ON pc.pet_type_id = pt.id INNER JOIN breed br ON pc.breed_id = br.id "

	if filter.PetCardId != 0 || filter.UserId != uuid.Nil || filter.Gender != "" || filter.PetTypeId != 0 ||
		filter.BreedId != 0 {

		query += "WHERE "
		setValues := make([]string, 0)

		if filter.PetCardId != 0 {
			setValues = append(setValues, fmt.Sprintf("pc.id = %d", filter.PetCardId))
		}

		if filter.PetTypeId != 0 {
			setValues = append(setValues, fmt.Sprintf("pc.pet_type_id = %d", filter.PetTypeId))
		}

		if filter.BreedId != 0 {
			setValues = append(setValues, fmt.Sprintf("pc.breed_id = %d", filter.BreedId))
		}

		if filter.UserId != uuid.Nil {
			setValues = append(setValues, fmt.Sprintf("pc.user_id = '%s'", filter.UserId.String()))
		}

		if filter.Gender != "" {
			if filter.Gender == "male" {
				setValues = append(setValues, fmt.Sprintf("pc.male = True"))
			} else if filter.Gender == "female" {
				setValues = append(setValues, fmt.Sprintf("pc.male = False"))
			}
		}

		query += strings.Join(setValues, " AND ")

	}

	return query
}

func (r *PetCardPostgres) GetMain(filter models.PetCardFilter) ([]models.PetCardMainInfo, error) {
	var lists []models.PetCardMainInfo

	query := createMainCardInfoQuery(filter)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *PetCardPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s pc WHERE pc.id = $1",
		petCardTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *PetCardPostgres) Update(id int, input models.UpdateCardInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.PetTypeId != nil {
		setValues = append(setValues, fmt.Sprintf("pet_type_id=$%d", argId))
		args = append(args, *input.PetTypeId)
		argId++
	}

	if input.UserId != nil {
		setValues = append(setValues, fmt.Sprintf("user_id=$%d", argId))
		args = append(args, *input.UserId)
		argId++
	}

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("pet_name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.BreedId != nil {
		setValues = append(setValues, fmt.Sprintf("breed_id=$%d", argId))
		args = append(args, *input.BreedId)
		argId++
	}

	if input.BirthDate != nil {
		setValues = append(setValues, fmt.Sprintf("birth_date=$%d", argId))
		args = append(args, *input.BirthDate)
		argId++
	}

	if input.Male != nil {
		setValues = append(setValues, fmt.Sprintf("male=$%d", argId))
		args = append(args, *input.Male)
		argId++
	}

	if input.Color != nil {
		setValues = append(setValues, fmt.Sprintf("color=$%d", argId))
		args = append(args, *input.Color)
		argId++
	}

	if input.Care != nil {
		setValues = append(setValues, fmt.Sprintf("care=$%d", argId))
		args = append(args, *input.Care)
		argId++
	}

	if input.Character != nil {
		setValues = append(setValues, fmt.Sprintf("pet_character=$%d", argId))
		args = append(args, *input.Character)
		argId++
	}

	if input.Pedigree != nil {
		setValues = append(setValues, fmt.Sprintf("pedigree=$%d", argId))
		args = append(args, *input.Pedigree)
		argId++
	}

	if input.Sterilization != nil {
		setValues = append(setValues, fmt.Sprintf("sterilization=$%d", argId))
		args = append(args, *input.Sterilization)
		argId++
	}

	if input.Vaccinations != nil {
		setValues = append(setValues, fmt.Sprintf("vaccinations=$%d", argId))
		args = append(args, *input.Vaccinations)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s pc SET %s WHERE pc.id = $%d",
		petCardTable, setQuery, argId)
	args = append(args, id)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
