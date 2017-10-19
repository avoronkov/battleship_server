#include "Cell.hpp"

std::ostream & operator<<(std::ostream& out, Cell c) {
	switch (c) {
		case Cell::EMPTY:
			out << ".";
			break;
		case Cell::SHIP:
			out << "#";
			break;
		case Cell::MISS:
			out << "-";
			break;
		case Cell::HIT:
			out << "X";
			break;
		default:
			std::cerr << "Unknown state" << std::endl;
	}
	return out;
}

std::istream& operator>>(std::istream& in, Cell& c) {
	char x;
	in >> x;
	switch (x) {
		case '.':
		case '0':
			c = Cell::EMPTY;
			break;
		case '#':
		case '1':
			c = Cell::SHIP;
			break;
		case '-':
			c = Cell::MISS;
			break;
		case 'X':
			c = Cell::HIT;
			break;
		default:
			std::cerr << "Unknown Cell value: " << x << std::endl;
	}
	return in;
}

