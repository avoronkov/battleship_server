#include "Player.hpp"

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
