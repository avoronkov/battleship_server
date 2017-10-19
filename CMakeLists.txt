cmake_minimum_required(VERSION 3.0)
project(BattleshipServer)

set(CMAKE_CXX_STANDARD 14)

find_package(Boost REQUIRED COMPONENTS system)

include_directories(".")

set(SOURCE_FILES main.cpp Player.cpp Field.cpp Cell.cpp)
add_executable(battleship_server ${SOURCE_FILES})

target_link_libraries(battleship_server ${Boost_LIBRARIES})
