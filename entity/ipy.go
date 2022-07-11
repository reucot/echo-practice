package entity

import (
	"fmt"
	"strconv"
	"strings"
)

type IncomePerYear struct {
	Icy int64
}

func (icy IncomePerYear) MarshalJSON() ([]byte, error) {
	icyStr := fmt.Sprintf(`"%d.%d"`, icy.Icy/100, icy.Icy%100)
	return []byte(icyStr), nil
}

//Можно было через регулярное выражение, но как тогда описывать ошибки?
func (icy *IncomePerYear) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	if strings.Contains(s, ",") {
		return fmt.Errorf("replace  \",\" to \".\"")
	}

	ns := strings.Split(s, ".")

	if len(ns) > 2 {
		return fmt.Errorf("income_per_year should be contains one \".\"")
	}

	n1, err := strconv.ParseInt(ns[0], 0, 64)
	if err != nil {
		return fmt.Errorf("income_per_year isn't number")
	}

	n2 := int64(0)
	if len(ns) > 1 {
		if len(ns[1]) > 2 {
			return fmt.Errorf("income_per_year, should be contains maximum two digits after \".\"")
		}

		var b strings.Builder
		fmt.Fprint(&b, ns[1])
		//Добавляем нули
		for i := len(ns[1]); i < 2; i++ {
			b.WriteString("0")
		}

		n2, err = strconv.ParseInt(b.String(), 0, 64)
		if err != nil {
			return fmt.Errorf("income_per_year isn't number")
		}
	}

	icy.Icy = n1*100 + n2
	return nil
}
