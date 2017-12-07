#include <iostream>
#include <stdexcept>

#include "Game.hpp"

int main(int argc, char **argv) {
	if (argc < 3) {
		std::cerr << "Usage: <prog1> <prog2> ..." << std::endl;
		return 2;
	}

	try {
		Game game(argv[1], argv[2]);
		game.start();
	}
	catch (const std::exception & e) {
		std::cerr << e.what() << std::endl;
		return 1;
	}
}
