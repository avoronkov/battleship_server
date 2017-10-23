#pragma once

#include <iostream>
#include <vector>

#include "Cell.hpp"

class Field {
private:
	std::vector<std::vector<Cell>> field;
public:
	Field();
	~Field() = default;

	void setCell(int x, int y, Cell c);
	Cell getCell(int x, int y) const;

	void put(std::ostream & out) const;
	void read(std::istream & in);

	static const size_t SIZE = 10;
};

std::ostream& operator<<(std::ostream& out, const Field &);
std::istream& operator>>(std::istream& in, Field &);
