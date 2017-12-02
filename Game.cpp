#include <iostream>
#include "Game.hpp"

Game::Game(const std::string & _p1, const std::string & _p2):
	p1{_p1},
	p2{_p2}
{
}

void Game::start() {
	f1 = p1.initField();
	std::cerr << "Player1 field:\n" << f1 << std::endl;
	if (!checkField(f1)) {
		std::cout << "Player1 loose: incorrect field\n";
		return;
	}
	f2 = p2.initField();
	std::cerr << "Player2 field:\n" << f2 << std::endl;
	if (!checkField(f2)) {
		std::cout << "Player2 loose: incorrect field\n";
		return;
	}

	while (true) {
		p1.shootCmd();
		auto shoot = p1.getShot();
	}

}

bool Game::checkField(const Field & f) const {
	// TODO
	return true;
}
