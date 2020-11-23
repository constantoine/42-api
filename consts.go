package api

import "errors"

var (
	ErrBadMonth          = errors.New("Wrong month number. Month can only be between 1 and 12 inclusive")
	ErrInvalidSearchKind = errors.New("Wrong search kind number. SearchKind can only be between 0 and 5 inclusive")
	ErrInvalidFieldType  = errors.New("Wrong field type number. FieldType can only be between 0 and 8 inclusive")
)

type FieldType uint

const (
	ID FieldType = iota
	Login
	Email
	PoolYear
	PoolMonth
	Kind
	PrimaryCampusID
	FirstName
	LastName
)

func (FieldType FieldType) String() (string, error) {
	var FieldTypes = [...]string{
		"id",
		"login",
		"email",
		"pool_year",
		"pool_month",
		"kind",
		"primari_campus_id",
		"first_name",
		"last_name",
	}
	if FieldType > 8 {
		return "", ErrInvalidFieldType
	}
	return FieldTypes[FieldType], nil
}

type SearchKind uint

const (
	Sort SearchKind = iota
	Filter
	Range
	PageNumber
	PageSize
)

func (Search SearchKind) String() (string, error) {
	var SearchKinds = [...]string{
		"sort",
		"filter",
		"range",
		"page[number]",
		"page[size]",
	}
	if Search > 4 {
		return "", ErrInvalidSearchKind
	}
	return SearchKinds[Search], nil
}

type Month uint

const (
	January Month = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func (month Month) String() (string, error) {
	var months = [...]string{
		"january",
		"february",
		"march",
		"april",
		"may",
		"june",
		"july",
		"august",
		"september",
		"october",
		"november",
		"december",
	}
	if month < 1 || month > 12 {
		return "", ErrBadMonth
	}
	return months[month], nil
}
