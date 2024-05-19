package repositories

import (
	"strconv"
	"strings"

	"github.com/alfanzain/project-sprint-halo-suster/app/consts"
	repositoryContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/repositories"
	"github.com/alfanzain/project-sprint-halo-suster/app/databases"
	"github.com/alfanzain/project-sprint-halo-suster/app/entities"

	"database/sql"
	"errors"
	"log"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository() repositoryContracts.IUserRepository {
	return &UserRepository{DB: databases.PostgreSQLInstance}
}

func (r *UserRepository) FindByID(userID string) (*entities.User, error) {
	var user entities.User
	err := r.DB.QueryRow(`SELECT id, nip, name, password, role_id, gender_id FROM users WHERE id = $1`, userID).Scan(&user.ID, &user.NIP, &user.Name, &user.Password, &user.RoleID, &user.GenderID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Fatalln(err)
		return nil, err
	}

	return &user, err
}

func (r *UserRepository) FindByNIP(NIP string) (*entities.User, error) {
	var user entities.User
	err := r.DB.QueryRow(`SELECT id, nip, name, password, role_id, gender_id FROM users WHERE nip = $1`, NIP).Scan(&user.ID, &user.NIP, &user.Name, &user.Password, &user.RoleID, &user.GenderID)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Fatalln(err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindITByNIP(NIP string) (*entities.User, error) {
	var user entities.User
	err := r.DB.QueryRow(`SELECT id, nip, name, password, role_id, gender_id FROM users WHERE nip = $1 AND role_id = $2`, NIP, consts.NIP_CODE_ROLE_IT).Scan(&user.ID, &user.NIP, &user.Name, &user.Password, &user.RoleID, &user.GenderID)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Fatalln(err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindNurseByNIP(NIP string) (*entities.User, error) {
	var user entities.User
	err := r.DB.QueryRow(`SELECT id, nip, name, password, role_id, gender_id FROM users WHERE nip = $1 AND role_id = $2`, NIP, consts.NIP_CODE_ROLE_NURSE).Scan(&user.ID, &user.NIP, &user.Name, &user.Password, &user.RoleID, &user.GenderID)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Fatalln(err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) DoesNIPExist(NIP string) (bool, error) {
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

func (r *UserRepository) Store(p *entities.UserStorePayload) (*entities.User, error) {
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

	NIP, err := strconv.Atoi(p.NIP)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		ID:       id,
		Name:     p.Name,
		NIP:      NIP,
		RoleID:   p.RoleID,
		GenderID: p.GenderID,
	}

	return user, nil
}

func (r *UserRepository) GetUsers(filters *entities.UserGetFilterParams) ([]*entities.User, error) {
	query := "SELECT id, name, nip, created_at FROM users WHERE 1=1 "
	params := []interface{}{}
	conditions := []string{}

	if filters.ID != "" {
		conditions = append(conditions, "id = $"+strconv.Itoa(len(params)+1))
		params = append(params, filters.ID)
	}
	if filters.Name != "" {
		query += "AND name ILIKE '%' || $" + strconv.Itoa(len(params)+1) + " || '%'"
		params = append(params, filters.Name)
	}
	if filters.NIP != "" {
		query += "AND nip = $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.NIP)
	}
	if filters.Role != "" && (filters.Role != "it") {
		query += "AND role_id = $" + strconv.Itoa(len(params)+1)
		params = append(params, consts.NIP_CODE_ROLE_IT)
	}
	if filters.Role != "" && (filters.Role != "nurse") {
		query += "AND role_id = $" + strconv.Itoa(len(params)+1)
		params = append(params, consts.NIP_CODE_ROLE_NURSE)
	}
	if len(conditions) > 0 {
		query += " AND "
	}

	query += strings.Join(conditions, " AND ")
	if filters.Limit == 0 {
		filters.Limit = 5
	}

	query += " ORDER BY created_at DESC"

	query += " LIMIT $" + strconv.Itoa(len(params)+1)
	params = append(params, filters.Limit)

	if filters.Offset == 0 {
		filters.Offset = 0
	} else {
		query += " OFFSET $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Offset)
	}

	rows, err := r.DB.Query(query, params...)
	if err != nil {
		log.Printf("Error finding cat: %s", err)
		return nil, err
	}
	defer rows.Close()

	users := make([]*entities.User, 0)
	for rows.Next() {
		u := new(entities.User)
		err := rows.Scan(&u.ID, &u.Name, &u.NIP, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Update(p *entities.UserUpdatePayload) (*entities.UserUpdateResponse, error) {
	var NIP int
	err := r.DB.QueryRow(`SELECT nip FROM users WHERE id = $1`, p.ID).Scan(&NIP)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	NIPPayload, err := strconv.Atoi(p.NIP)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	if NIP == NIPPayload {
		_, err := r.DB.Exec("UPDATE users SET name = $2 WHERE id = $1",
			p.ID,
			p.Name,
		)

		if err != nil {
			log.Printf("Error updating user: %s", err)
			return nil, err
		}
	} else {
		_, err := r.DB.Exec("UPDATE users SET name = $2, nip = $3 WHERE id = $1",
			p.ID,
			p.Name,
			p.NIP,
		)

		if err != nil {
			log.Printf("Error updating user: %s", err)
			return nil, err
		}
	}

	user := &entities.UserUpdateResponse{
		ID:   p.ID,
		Name: p.Name,
		NIP:  NIP,
	}

	return user, nil
}

func (r *UserRepository) UpdatePassword(p *entities.UserNurseGrantAccessPayload) (*entities.UserUpdatePasswordResponse, error) {
	_, err := r.DB.Exec("UPDATE users SET password = $2 WHERE id = $1",
		p.ID,
		p.Password,
	)
	if err != nil {
		log.Printf("Error updating user: %s", err)
		return nil, err
	}

	user := &entities.UserUpdatePasswordResponse{
		ID: p.ID,
	}

	return user, err
}

func (r *UserRepository) Destroy(userID string) (bool, error) {
	_, err := r.DB.Exec("DELETE FROM users WHERE id = $1;", userID)

	if err != nil {
		log.Printf("Error deleting user: %s", err)
		return false, err
	}

	return true, nil
}
