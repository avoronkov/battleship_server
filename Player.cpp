#include "Player.hpp"
#include <stdexcept>
#include <sstream>

namespace bp = boost::process;

Player::Player(const std::string & prog) {
	ch.reset(new bp::child(prog, bp::std_out > out, bp::std_in < in));
}

Field Player::initField() {
	in << "Arrange!" << std::endl;
	Field f;
	out >> f;
	return f;
}

void Player::shootCmd() {
	in << "Shoot!" << std::endl;
}

std::pair<int, int> Player::getShot() {
	std::string sx, sy;
	out >> sx >> sy;
	int x = sx.at(0) - 'A';
	if (sx.size() != 1 || x < 0 || x >= Field::SIZE) {
		std::stringstream err;
		err << "Invalid x-coordinate: '" << sx << "'\n";
		throw std::invalid_argument(err.str());
	}
	int y = sy.at(0) - '0';
	if (sy.size() != 1 || y < 0 || y >= Field::SIZE) {
		std::stringstream err;
		err << "Invalid y-coordinate: '" << sy << "'\n";
		throw std::invalid_argument(err.str());
	}
	return std::pair<int, int>(x, y);
}
