#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
// #include "linearalg.h"

#include "utils.h"

#define FREAD_BUFFER 65536

typedef struct CSVFile {
    char* filename; 
    size_t lineCount;
    char separator;
    int hasHeader;
} CSVFile; 

size_t st_mult(size_t a, size_t b) {
    if (a > ULONG_MAX / b) {
        printf("ERROR: Overflow trying to multiply size_t's: %zu * %zu", a, b);
        exit(1);
    }

    return a * b;
}

// static Vector parseCSVLine() {
//     return (Vector) {};
// }


size_t countFileLines(FILE* file) {
    fseek(file, 0, SEEK_SET); // Set to the start of the file
    char *stringBuffer = ALLOC_ARRAY(stringBuffer, FREAD_BUFFER);
    int counter = 0; 

    for(;;) {
        size_t nBytes = fread(stringBuffer, sizeof(char), FREAD_BUFFER, file);
        if (ferror(file)) 
            return -1; 

        for(size_t i = 0; i < nBytes; i++) {
            if(stringBuffer[i] == '\n')
                counter++; 
        }

        if(feof(file))
            break;
    }

    fseek(file, 0, SEEK_SET); // Set to the start of the file
    return counter;
}

CSVFile createCSV(char* path, char separator, int hasHeader) {
    CSVFile csv; 
    csv.filename = path; 
    csv.separator = separator; 
    csv.hasHeader = hasHeader;

    FILE* csvfile = fopen(path, "r");
    if (csvfile == NULL) {
        printf("ERROR: The CSV path %s could not be opened! Check if the file exists.", csv.filename);
        exit(1);
    }   

    csv.lineCount = countFileLines(csvfile);
    fclose(csvfile);

    return csv; 
}

char* readFileToString(char* path){
    FILE *file = fopen(path, "r");
    if (file == NULL) {
        printf("ERROR: The file path %s could not be opened! Check if the file exists.", path);
        exit(1);
    }

    // Get the file length in bytes
    fseek(file, 0, SEEK_END);
    size_t fileLength = ftell(file);
    fseek(file, 0, SEEK_SET);

    // Create the buffer and the string where the file will sit
    char *entireFuckingFile = malloc(st_mult((fileLength + 1), sizeof(char))); // +1 for the null terminator "\0"
    char *stringBuffer = ALLOC_ARRAY(stringBuffer, FREAD_BUFFER);

    int entireIdx = 0, scanned = 0;
    for(;;) {
        scanned = fread(stringBuffer, sizeof(char), FREAD_BUFFER, file);
        
        // Does the csv string treatment
        for (size_t i = 0; i < scanned; i++) {
            entireFuckingFile[i] = stringBuffer[i];
            entireIdx++; 
        }

        if(feof(file))
            break; 
    }
    fclose(file);
    entireFuckingFile[entireIdx] = '\0';
    return entireFuckingFile;  
}

// // Read all the contents of .csv file and converts then to double-filled Matrix
// void readCsvToMatrix(CSVFile csv){
//     char *singleString, *stopPtr;
//     if (csv.separator == NULL)
//         csv.separator = ',';    

//     FILE *csvFile = fopen(csv.filename, "r");
//     if (csvFile == NULL) {
//         printf("ERROR: The CSV path %s could not be opened! Check if the file exists.", csv.filename);
//         exit(1);
//     }

//     // Get the file length in bytes
//     fseek(csvFile, 0, SEEK_END);
//     size_t fileLength = ftell(csvFile);
//     fseek(csvFile, 0, SEEK_SET);

//     char *entireFuckingFile = malloc(st_mult((fileLength + 1), sizeof(char))); // +1 for the null terminator "\0"
//     char *stringBuffer = ALLOC_ARRAY(stringBuffer, FREAD_BUFFER);
//     char *lineBuilder = ALLOC_ARRAY(stringBuffer, FREAD_BUFFER);

//     int bytesRead = 0, scanned = 0;
//     for(;;) {
//         scanned = fread(stringBuffer, sizeof(char), FREAD_BUFFER, csvFile);
        
//         // Does the csv string treatment
//         for (size_t i = 0; i < scanned; i++) {
//             lineBuilder[i] = stringBuffer[i];
//             if (stringBuffer[i] == '\n') {
//                 // It completes a line, and parses it
//                 parseLine(lineBuilder);
//             }
//         }


//         // TODO: Get the logic behind when: 
//         // - The buffer ends, and there's the end of the line missing
//         //      -- Offset the memory with fseek and seek the rest of the line
//         // - When line is actually longer than lineBuilder

//         bytesRead += scanned;


//         if(feof(csvFile))
//             break; 

//     }


//     double result = strtod(singleString, &stopPtr);

//     fclose(csvFile);
//    // return (Matrix){};
// }



char* stringSplit(char* s) {

}


char* stringTrim(char* s) {
    while (isspace((unsigned char)*s)) {
        s++; // Trim leading whitespace
    }

    // If the string is all whitespace, s will point to the null terminator
    if (*s == '\0')
        return s;

    // Trim trailing whitespace
    char* end = s + strlen(s) - 1;
    while (end >= s && isspace((unsigned char)*end)) {
        *end = '\0'; // Overwrite with null terminator
        end--;
    }

    return s; // String is modified in-place
}

int main() {
    // Must use a mutable character array, not a string literal (which is read-only)
    char str1[] = "   Hello, World!   ";
    char str2[] = "\t  Another string with tabs and newlines\n";
    char str3[] = "     "; // String with only spaces

    printf("Original 1: '%s'\n", str1);
    printf("Trimmed 1 : '%s'\n", trim(str1));
    
    printf("\nOriginal 2: '%s'\n", str2);
    printf("Trimmed 2 : '%s'\n", trim(str2));

    printf("\nOriginal 3: '%s'\n", str3);
    printf("Trimmed 3 : '%s'\n", trim(str3));


    printf(readFileToString("Iris.csv"));
    return 0;
}