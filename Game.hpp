#pragma once

#include <string>
#include <memory>
#include "Player.hpp"

class Game {
	Player p1;
	Player p2;

	Field f1;
	Field f2;
public:
	Game(const std::string &, const std::string &);

	void start();
private:
	bool checkField(const Field & f) const;
};
