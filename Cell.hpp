#pragma once

#include <iostream>

enum class Cell {
	EMPTY,
	SHIP,
	MISS,
	HIT
};

std::ostream& operator<<(std::ostream& out, Cell c);
std::istream& operator>>(std::istream& in, Cell& c);
