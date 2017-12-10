package shot

import (
	"fmt"
)

type Shot struct {
	X int
	Y int
}

func (s Shot) String() string {
	return fmt.Sprintf("%c %d", byte(s.X)+'A', s.Y)
}
