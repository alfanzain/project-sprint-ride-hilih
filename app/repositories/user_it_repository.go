package repositories

import (
	repositoryContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/repositories"
	"github.com/alfanzain/project-sprint-halo-suster/app/databases"
	"github.com/alfanzain/project-sprint-halo-suster/app/entities"

	"database/sql"
	"errors"
	"log"
)

type UserITRepository struct {
	DB *sql.DB
}

func NewUserITRepository() repositoryContracts.IUserITRepository {
	return &UserITRepository{DB: databases.PostgreSQLInstance}
}

func (r *UserITRepository) FindByNIP(NIP string) (*entities.User, error) {
	var user entities.User
	err := r.DB.QueryRow(`SELECT id, nip, name, password, role_id, gender_id FROM users WHERE nip = $1`, NIP).Scan(&user.ID, &user.NIP, &user.Name, &user.Password, &user.RoleID, &user.GenderID)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Fatalln(err)
		return nil, err
	}

	return &user, nil
}

func (r *UserITRepository) DoesNIPExist(NIP string) (bool, error) {
	var scannedNIP string
	err := r.DB.QueryRow(`SELECT nip FROM users WHERE nip = $1`, NIP).Scan(&scannedNIP)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Fatalln(err)
		return false, err
	}

	if len(scannedNIP) == 0 {
		return false, nil
	}

	return true, nil
}

func (r *UserITRepository) Store(p *entities.UserITStorePayload) (*entities.User, error) {
	var id string
	err := r.DB.QueryRow(`
			INSERT INTO users (
				name,
				nip,
				password,
				role_id,
				gender_id
			) VALUES (
				$1, $2, $3, $4, $5
			) RETURNING id
		`,
		p.Name,
		p.NIP,
		p.Password,
		p.RoleID,
		p.GenderID,
	).Scan(&id)
	if err != nil {
		log.Printf("Error inserting user IT: %s", err)
		return nil, err
	}

	user := &entities.User{
		ID:       id,
		Name:     p.Name,
		NIP:      p.NIP,
		RoleID:   p.RoleID,
		GenderID: p.GenderID,
	}

	return user, nil
}
