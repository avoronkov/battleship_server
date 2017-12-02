#include <iostream>

#include "Game.hpp"

int main(int argc, char **argv) {
	if (argc < 3) {
		std::cerr << "Usage: <prog1> <prog2> ..." << std::endl;
		return 2;
	}

	Game game(argv[1], argv[2]);
	game.start();
}
