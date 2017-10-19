#include <iostream>

#include "Player.hpp"

int main(int argc, char **argv) {
	if (argc < 2) {
		std::cerr << "Usage: <prog1> ..." << std::endl;
		return 2;
	}

	Player p1(argv[1]);
	Field f = p1.initField();
	std::cout << "Player 1 filed:\n" << f << std::endl;
}
