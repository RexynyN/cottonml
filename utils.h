#ifndef UTILS_H
#define UTILS_H


// #define ALLOC_ARRAY(x, alloc) (x) = malloc(st_mult(sizeof(*(x)), (alloc)))
#define ALLOC_ARRAY(x, alloc) malloc(sizeof(*(x)) *(alloc))

#define max2(x, y) (x > y ? x : y)

#endif