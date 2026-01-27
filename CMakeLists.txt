cmake_minimum_required(VERSION 3.25.0)

project(cottonml)

add_executable(cottonml linearalg.c stats.c counter.c utils.c knn.c dataset.c)

target_link_libraries(cottonml m) # Links the <math.h> to the .exe
