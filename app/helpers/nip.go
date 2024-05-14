package helpers

import (
	"errors"
	"strconv"
	"time"
)

const (
	CODE_ROLE_IT       int = 615
	CODE_ROLE_NURSE    int = 303
	CODE_GENDER_MALE   int = 1
	CODE_GENDER_FEMALE int = 2
)

type DecodedNIP struct {
	Number     string
	RoleID     int
	GenderID   int
	Year       string
	Month      string
	UniqueCode string
}

func DecodeNIP(nip string) (nipNumber *DecodedNIP, err error) {
	if len(nip) != 13 {
		return nil, errors.New("nip is invalid")
	}

	roleID, err := strconv.Atoi(nip[:3])
	if err != nil {
		return nil, err
	}

	genderID, err := strconv.Atoi(nip[3:4])
	if err != nil {
		return nil, err
	}

	nipNumber = &DecodedNIP{
		Number:     nip,
		RoleID:     roleID,
		GenderID:   genderID,
		Year:       nip[4:8],
		Month:      nip[8:10],
		UniqueCode: nip[10:13],
	}

	roleCodes := Slice[int]{
		CODE_ROLE_IT,
		CODE_ROLE_NURSE,
	}
	if roleCodes.Includes(nipNumber.RoleID) {
		return nil, errors.New("nip is invalid")
	}

	genderCodes := Slice[int]{
		CODE_GENDER_MALE,
		CODE_GENDER_FEMALE,
	}
	if genderCodes.Includes(nipNumber.GenderID) {
		return nil, errors.New("nip is invalid")
	}

	year, err := strconv.Atoi(nipNumber.Year)
	if err != nil {
		return nil, errors.New("nip is invalid")
	}

	currentYear := time.Now().Year()
	if year < 2000 || year > currentYear {
		return nil, errors.New("nip is invalid")
	}

	month, err := strconv.Atoi(nipNumber.Month)
	if err != nil {
		return nil, errors.New("nip is invalid")
	}

	if month < 1 || month > 12 {
		return nil, errors.New("nip is invalid")
	}

	return nipNumber, nil
}
