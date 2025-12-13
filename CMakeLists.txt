cmake_minimum_required(VERSION 3.28.3)

project(cottonml)

add_executable(cottonml  linearalg.c stats.c)

target_link_libraries(cottonml m) # Links the <math.h> to the .exe