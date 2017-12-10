package field

import (
	"battleship_server/cell"
	"testing"
)

func TestKilled1(t *testing.T) {
	// "X"
	f := Field{}
	f.field[1][1] = cell.HIT
	if !f.killed(1, 1) {
		t.Errorf("Ship at 1,1 is not kiled:\n%s", f)
	}
}

func TestKilled2(t *testing.T) {
	// "X#"
	f := Field{}
	f.field[1][1] = cell.HIT
	f.field[1][2] = cell.SHIP
	if f.killed(1, 1) {
		t.Errorf("Ship at 1,1 is kiled:\n%s", f)
	}
}

func TestKilled3(t *testing.T) {
	// "#X#"
	f := Field{}
	f.field[1][1] = cell.SHIP
	f.field[1][2] = cell.HIT
	f.field[1][3] = cell.SHIP
	if f.killed(2, 1) {
		t.Errorf("Ship at 2,1 is kiled:\n%s", f)
	}
}

func TestKilled4(t *testing.T) {
	// "Xx#"
	f := Field{}
	f.field[1][1] = cell.HIT
	f.field[1][2] = cell.HIT
	f.field[1][3] = cell.SHIP
	if f.killed(1, 1) {
		t.Errorf("Ship at 1,1 is kiled:\n%s", f)
	}
}

func TestKilled5(t *testing.T) {
	// "Xxx#"
	f := Field{}
	f.field[1][1] = cell.HIT
	f.field[1][2] = cell.HIT
	f.field[1][3] = cell.HIT
	f.field[1][4] = cell.SHIP
	if f.killed(1, 1) {
		t.Errorf("Ship at 1,1 is kiled:\n%s", f)
	}
}

func TestKilled6(t *testing.T) {
	// "xXx#"
	f := Field{}
	f.field[1][1] = cell.HIT
	f.field[1][2] = cell.HIT
	f.field[1][3] = cell.HIT
	f.field[1][4] = cell.SHIP
	if f.killed(2, 1) {
		t.Errorf("Ship at 2,1 is kiled:\n%s", f)
	}
}

func TestKilledBorder1(t *testing.T) {
	// "|x"
	f := Field{}
	f.field[0][0] = cell.HIT
	if !f.killed(0, 0) {
		t.Errorf("Ship at 0,0 is not kiled:\n%s", f)
	}
}

func TestKilledBorder2(t *testing.T) {
	// "x|"
	f := Field{}
	f.field[FIELD_SIZE-1][FIELD_SIZE-1] = cell.HIT
	if !f.killed(FIELD_SIZE-1, FIELD_SIZE-1) {
		t.Errorf("Ship at FIELD_SIZE-1,FIELD_SIZE-1 is not kiled:\n%s", f)
	}
}

func TestKilledBorder3(t *testing.T) {
	// "x#|"
	f := Field{}
	f.field[FIELD_SIZE-1][FIELD_SIZE-2] = cell.HIT
	f.field[FIELD_SIZE-1][FIELD_SIZE-1] = cell.SHIP
	if f.killed(FIELD_SIZE-2, FIELD_SIZE-1) {
		t.Errorf("Ship at 8,9 is not kiled:\n%s", f)
	}
}

func TestKilledBorder4(t *testing.T) {
	// "x#|"
	f := Field{}
	f.field[FIELD_SIZE-2][FIELD_SIZE-1] = cell.HIT
	f.field[FIELD_SIZE-1][FIELD_SIZE-1] = cell.SHIP
	if f.killed(FIELD_SIZE-1, FIELD_SIZE-2) {
		t.Errorf("Ship at 9,8 is not kiled:\n%s", f)
	}
}
