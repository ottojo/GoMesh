package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type MAC int

func (m MAC) String() string {
	return strconv.Itoa(int(m))
}

func (m MAC) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.Itoa(int(m)))
}

func (m *MAC) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i, err := strconv.Atoi(s)
	*m = MAC(i)
	return err
}

func sprintMACs(MACs []MAC) string {
	s := ""
	for _, m := range MACs {
		s += fmt.Sprint(m) + " "
	}
	return s
}
