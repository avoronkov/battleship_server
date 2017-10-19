#include "Field.hpp"
#include <stdexcept>
#include <sstream>

const size_t FIELD_SIZE = 10;

Field::Field() {
	field.resize(FIELD_SIZE);
	for (auto & line : field) {
		line.resize(FIELD_SIZE);
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
	for (size_t i = 0; i < FIELD_SIZE; i++) {
		std::string line;
		std::getline(in, line);

		if (line.size() < FIELD_SIZE) {
			std::stringstream err("Line is too short: ");
			err << line;
			throw std::runtime_error(err.str());
		}
		std::stringstream str(line);
		for (size_t j = 0; j < FIELD_SIZE; j++) {
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
