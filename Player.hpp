#pragma once

#include <iostream>
#include <string>
#include <memory>
#include <utility>
#include "boost/process.hpp"

#include "Field.hpp"

class Player {
private:
	boost::process::opstream in;
	boost::process::ipstream out;

	std::unique_ptr<boost::process::child> ch;
public:
	Player(const std::string & prog);
	Player(const Player &) = delete;
	~Player() = default;

	Player& operator=(const Player&) = delete;

	Field initField();

	void shootCmd();

	std::pair<int, int> getShot();
};
