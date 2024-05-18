package helpers

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/alfanzain/project-sprint-halo-suster/app/consts"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/errs"
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
		fmt.Println("nip length is invalid")
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
		consts.NIP_CODE_ROLE_IT,
		consts.NIP_CODE_ROLE_NURSE,
	}
	if !roleCodes.Includes(nipNumber.RoleID) {
		fmt.Println("nip role code is invalid")
		return nil, errors.New("nip is invalid")
	}

	genderCodes := Slice[int]{
		consts.NIP_CODE_GENDER_MALE,
		consts.NIP_CODE_GENDER_FEMALE,
	}
	if !genderCodes.Includes(nipNumber.GenderID) {
		fmt.Println("nip gender code is invalid")
		return nil, errors.New("nip is invalid")
	}

	year, err := strconv.Atoi(nipNumber.Year)
	if err != nil {
		fmt.Println("nip year is invalid")
		return nil, errors.New("nip is invalid")
	}

	currentYear := time.Now().Year()
	if year < 2000 || year > currentYear {
		fmt.Println("nip year is invalid")
		return nil, errors.New("nip is invalid")
	}

	month, err := strconv.Atoi(nipNumber.Month)
	if err != nil {
		fmt.Println("nip month is invalid")
		return nil, errors.New("nip is invalid")
	}

	if month < 1 || month > 12 {
		fmt.Println("nip month is invalid")
		return nil, errors.New("nip is invalid")
	}

	return nipNumber, nil
}

func IsNIPNurseValid(NIProleID int) (bool, error) {
	if NIProleID == consts.NIP_CODE_ROLE_NURSE {
		return true, nil
	}

	return false, errs.ErrInvalidNIP
}

func IsNIPITValid(NIProleID int) (bool, error) {
	if NIProleID == consts.NIP_CODE_ROLE_IT {
		return true, nil
	}

	log.Fatalln(errs.ErrInvalidNIP)
	return false, errs.ErrInvalidNIP
}
