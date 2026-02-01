#ifndef COTTONSTRING_H
#define COTTONSTRING_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unicode/ustring.h>
#include <unicode/uchar.h>
#include <unicode/ustdio.h>


typedef UChar* string;
typedef UChar character; 

typedef struct SplittedString {
    string* strings; 
    int num; 
} SplittedString;

#define stringLen(src) u_strlen(src)  

#endif