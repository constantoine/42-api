package api

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
func main() {
	var s Search
	str, err := s.PageNumber(2).PageSize(100).Filter(PrimaryCampusID, "1").Range(PoolYear, 2017, 2020).Filter(ID, "1", "5").Filter(PoolMonth, "january").SortAsc(PoolMonth).SortDesc(PoolYear).QueryString()
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
*/

type Search struct {
	Params []Param
}

type Param struct {
	Type  SearchKind
	Field FieldType
	Value string
}

func (s *Search) querySort() (string, []Param, error) {
	var params []string
	var parmasan []Param
	for i := range s.Params {
		if s.Params[i].Type != Sort {
			parmasan = append(parmasan, s.Params[i])
			continue
		}
		pField, err := s.Params[i].Field.String()
		if err != nil {
			return "", nil, err
		}
		params = append(params, strings.Join([]string{s.Params[i].Value, pField}, ""))
	}
	var query string
	if len(params) != 0 {
		pName, err := Sort.String()
		if err != nil {
			return "", nil, err
		}
		query = fmt.Sprintf("?%s=%s", pName, strings.Join(params, ","))
	}
	return query, parmasan, nil
}

func (s *Search) QueryString() (string, error) {
	var params []string
	sortQuery, parmasan, err := s.querySort()
	if err != nil {
		return "", err
	}
	if sortQuery != "" {
		params = append(params, sortQuery)
	}
	sort.Slice(parmasan, func(i, j int) bool {
		return parmasan[i].Type < parmasan[j].Type
	})
	for i := range parmasan {
		var (
			pName  string
			pField string
			err    error
		)
		pName, err = parmasan[i].Type.String()
		if err != nil {
			return "", err
		}
		if parmasan[i].Type != PageSize && parmasan[i].Type != PageNumber {
			pField, err = parmasan[i].Field.String()
			if err != nil {
				return "", err
			}
		}
		switch parmasan[i].Type {
		case PageSize, PageNumber:
			params = append(params, fmt.Sprintf("%s=%s", pName, parmasan[i].Value))
		case Filter, Range:
			params = append(params, fmt.Sprintf("%s[%s]=%s", pName, pField, parmasan[i].Value))
		}
		if i == 0 && sortQuery == "" {
			params[0] = fmt.Sprintf("?%s", params[0])
		}
	}
	var query string
	query = strings.Join(params, "&")
	return query, nil
}

func (s *Search) PageSize(size uint) *Search {
	for i := range s.Params {
		if s.Params[i].Type == PageSize {
			s.Params[i].Value = strconv.FormatUint(uint64(size), 10)
			return s
		}
	}
	s.Params = append(s.Params, Param{
		Type:  PageSize,
		Value: strconv.FormatUint(uint64(size), 10),
	})
	return s
}

func (s *Search) PageNumber(number uint) *Search {
	for i := range s.Params {
		if s.Params[i].Type == PageNumber {
			s.Params[i].Value = strconv.FormatUint(uint64(number), 10)
			return s
		}
	}
	s.Params = append(s.Params, Param{
		Type:  PageNumber,
		Value: strconv.FormatUint(uint64(number), 10),
	})
	return s
}

func (s *Search) SortAsc(on FieldType) *Search {
	for i := range s.Params {
		if s.Params[i].Type == Sort && s.Params[i].Field == on {
			s.Params[i].Value = ""
			return s
		}
	}
	s.Params = append(s.Params, Param{
		Type:  Sort,
		Field: on,
		Value: "",
	})
	return s
}

func (s *Search) SortDesc(on FieldType) *Search {
	for i := range s.Params {
		if s.Params[i].Type == Sort && s.Params[i].Field == on {
			s.Params[i].Value = "-"
			return s
		}
	}
	s.Params = append(s.Params, Param{
		Type:  Sort,
		Field: on,
		Value: "-",
	})
	return s
}

func (s *Search) Filter(on FieldType, value ...string) *Search {
	if len(value) == 0 {
		return s
	}
	params := strings.Join(value, ",")
	for i := range s.Params {
		if s.Params[i].Type == Filter && s.Params[i].Field == on {
			if s.Params[i].Value != "" {
				s.Params[i].Value = strings.Join([]string{s.Params[i].Value, params}, ",")
			} else {
				s.Params[i].Value = params
			}
			return s
		}
	}
	s.Params = append(s.Params, Param{
		Type:  Filter,
		Field: on,
		Value: params,
	})
	return s
}

func (s *Search) Range(on FieldType, min uint, max uint) *Search {
	minimum := strconv.FormatUint(uint64(min), 10)
	maximum := strconv.FormatUint(uint64(max), 10)
	params := strings.Join([]string{minimum, maximum}, ",")
	for i := range s.Params {
		if s.Params[i].Type == Range && s.Params[i].Field == on {
			s.Params[i].Value = params
			return s
		}
	}
	s.Params = append(s.Params, Param{
		Type:  Range,
		Field: on,
		Value: params,
	})
	return s
}
