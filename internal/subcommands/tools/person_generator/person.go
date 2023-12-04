package person_generator

import (
	"strconv"
	"strings"
	"time"
)

type Address struct {
	Room     string
	Building string
	Street   string
	District string
}

type Person struct {
	Name     string
	Birthday time.Time
	Address
}

func (p *Person) Age() int {
	now := time.Now()
	yearDiff := now.Year() - p.Birthday.Year()
	if now.Month() >= p.Birthday.Month() && now.Day() >= p.Birthday.Day() {
		return yearDiff
	} else {
		return yearDiff - 1
	}
}

func (p *Person) AddressView() string {
	b := strings.Builder{}
	addToBuilder := func(attribute string) {
		if attribute != "" {
			if b.Len() != 0 {
				b.WriteString(", ")
			}
			b.WriteString(attribute)
		}
	}

	addToBuilder(p.Address.Room)
	addToBuilder(p.Address.Building)
	addToBuilder(p.Address.Street)
	addToBuilder(p.Address.District)
	return b.String()
}

func (p *Person) ToStringSlice() []string {
	return []string{
		p.Name,
		strconv.Itoa(p.Age()),
		p.AddressView(),
	}
}
