#ifndef UTILS_H
#define UTILS_H

#include "linearalg.h"

typedef struct Dataset {
    Matrix features; 
    Vector targets; 
} Dataset;

typedef struct CSVFile {
    char* filename; 
    size_t lineCount;
    size_t rows; 
    size_t columns; 
    char separator;
    int hasHeader;
} CSVFile; 


size_t st_mult(size_t a, size_t b);

// #define ALLOC_ARRAY(x, alloc) (x) = malloc(st_mult(sizeof(*(x)), (alloc)))
#define ALLOC_ARRAY(x, alloc) malloc(st_mult(sizeof(*(x)), alloc))

#define max2(x, y) (x > y ? x : y)

#endif