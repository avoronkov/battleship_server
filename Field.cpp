#include "Field.hpp"
#include <stdexcept>
#include <sstream>

Field::Field() {
	field.resize(Field::SIZE);
	for (auto & line : field) {
		line.resize(Field::SIZE);
		for (auto & cell : line) {
			cell = Cell::EMPTY;
		}
	}
}

void Field::setCell(int x, int y, Cell c) {
	field.at(y).at(x) = c;
}

Cell Field::getCell(int x, int y) const {
	return field.at(y).at(x);
}

void Field::put(std::ostream & out) const {
	for (const auto & line : field) {
		for (const auto & cell : line) {
			out << cell;
		}
		out << std::endl;
	}
}

void Field::read(std::istream & in) {
	for (size_t i = 0; i < Field::SIZE; i++) {
		std::string line;
		std::getline(in, line);

		if (line.size() < Field::SIZE) {
			std::stringstream err("Line is too short: ");
			err << line;
			throw std::runtime_error(err.str());
		}
		std::stringstream str(line);
		for (size_t j = 0; j < Field::SIZE; j++) {
			Cell c;
			str >> c;
			field.at(i).at(j) = c;
		}
	}
}

std::ostream& operator<<(std::ostream& out, const Field & field) {
	field.put(out);
	return out;
}

std::istream& operator>>(std::istream& in, Field & field) {
	field.read(in);
	return in;
}
