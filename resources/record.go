package resources

import "fmt"

type Record struct {
	Issuer string `json:"issuer"`
	Data   string `json:"data"`
}

func (r Record) String() string {
	return fmt.Sprintf("%s, %v", r.Issuer, r.Data)
}
