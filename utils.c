#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
// #include "linearalg.h"

#include <utils.h>

#define FREAD_BUFFER 10000

size_t st_mult(size_t a, size_t b) {
    if (a > ULONG_MAX / b) {
        printf("ERROR: Overflow trying to multiply size_t's: %zu * %zu", a, b);
        exit(1);
    }

    return a * b;
}

// Read all the contents of .csv file and converts then to double-filled Matrix
void readCsvToMatrix(const char *path, char *separator){
    char *singleString, *stopPtr;
    if (separator == NULL)
        separator = ",";

    FILE *csvFile = fopen(path, "r");
    if (csvFile == NULL) {
        printf("ERROR: The CSV path %s could not be opened! Check if the file exists.", path);
        exit(1);
    }

    // Get the file length in bytes
    fseek(csvFile, 0, SEEK_END);
    size_t fileLength = ftell(csvFile);
    fseek(csvFile, 0, SEEK_SET);

    char *entireFuckingFile = malloc(st_mult((fileLength + 1), sizeof(char))); // +1 for the null terminator "\0"
    char *stringBuffer = ALLOC_ARRAY(stringBuffer, FREAD_BUFFER);
    char *lineBuilder = ALLOC_ARRAY(stringBuffer, FREAD_BUFFER);

    int bytesRead = 0, scanned = 0;
    while (!feof(csvFile)) {
        scanned = fread(stringBuffer, sizeof(char), FREAD_BUFFER, csvFile);
        
        // Does the csv string treatment
        for (size_t i = 0; i < scanned; i++) {
            lineBuilder[i] = stringBuffer[i];
            if (stringBuffer[i] == '\n') {
                // It completes a line, and parses it
                parseLine(lineBuilder);
            } 
        }


        // TODO: Get the logic behind when: 
        // - The buffer ends, and there's the end of the line missing
        //      -- Offset the memory with fseek and seek the rest of the line
        // - When line is actually longer than lineBuilder

        bytesRead += scanned;

    }


    double result = strtod(singleString, &stopPtr);

    fclose(csvFile);
   // return (Matrix){};
}

int main(void){
    printf("%zu", sizeof(char));
}
