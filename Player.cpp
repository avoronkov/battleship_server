#include "Player.hpp"
#include <stdexcept>
#include <sstream>
#include <iostream>

namespace bp = boost::process;
using std::cerr;

Player::Player(const std::string & _name, const std::string & prog):
	name{_name}
{
	ch.reset(new bp::child(prog, bp::std_out > out, bp::std_in < in));
}

Field Player::initField() {
	cerr << name << ": initField()\n";
	in << "Arrange!" << std::endl;
	cerr << "1in is good = " << in.good() << std::endl;
	in.flush();
	cerr << "2in is fail = '" << in.bad() << "'" << std::endl;
	Field f;
	out >> f;
	cerr << name << ": field:\n" << f;
	return f;
}

void Player::shootCmd() {
	cerr << name << ": shootCmd()\n";
	cerr << "in is good = " << in.good() << std::endl;
	in << "Shoot!" << std::endl;
	in.flush();
}

std::pair<int, int> Player::getShot() {
	cerr << name << ": getShot()\n";
	std::string sx, sy;
	out >> sx >> sy;
	cerr << name << ": getShot: " << sx << ", " << sy << std::endl;
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
