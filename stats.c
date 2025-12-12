#include "stats.h"
#include "linearalg.h"
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <string.h>


double randDouble() {
    // Creates a number from 0 to 1
    return (double)rand() / (double)RAND_MAX;
}

double* createDoubleArray(size_t len) {
    return (double*)malloc(len * sizeof(double)); 
} 

double searchArray(double* arr, size_t len, double search) {
    for (size_t i = 0; i < len; i++) {
        if (search == arr[i])
            return i;
    }

    return -1;
}

int compareDoublesSorting(const void *a, const void *b) {
    double val_a = *(const double *)a;
    double val_b = *(const double *)b;

    if (val_a < val_b) {
        return -1;
    } else if (val_a > val_b) {
        return 1;
    } else {
        return 0;
    }
}

double min(Vector* v) {
    double min = v->data[0];
    for (size_t i = 0; i < v->len; i++) {
        if (v->data[i] < min) 
            min = v->data[i];
    }

    return min;
}

double max(Vector* v) {
    double max = v->data[0];
    for (size_t i = 0; i < v->len; i++) {
        if (v->data[i] > max) 
            max = v->data[i];
    }
    return max; 
}

double maxArray(double* arr, size_t len) {
    double max = arr[0];
    for (size_t i = 0; i < len; i++) {
        if (arr[i] > max) 
            max = arr[i];
    }
    return max; 
}

double minArray(double* arr, size_t len) {
    double min = arr[0];
    for (size_t i = 0; i < len; i++) {
        if (arr[i] < min) 
            min = arr[i];
    }
    return min; 
}

double mean(Vector* v) {
    double sum = 0; 
    for (size_t i = 0; i < v->len; i++) 
        sum += v->data[i];

    return sum / v->len;
}

double _median_even(Vector* v) {
    double* array_sorted = createDoubleArray(v->len);
    memcpy(array_sorted, v->data, (v->len * sizeof(double))); // Create a copy of the array
    qsort(array_sorted, v->len, sizeof(double), compareDoublesSorting); // Sort the array

    // Safeguards
    int hi_midpoint = (int)floor(v->len / 2);
    if (hi_midpoint < 0) hi_midpoint = 0;
    if (hi_midpoint >= v->len) hi_midpoint = v->len - 1;

    double returns = (array_sorted[hi_midpoint - 1] + array_sorted[hi_midpoint]) / 2;
    free(array_sorted);
    return returns; 
}

double _median_odd(Vector* v) {
    double* array_sorted = createDoubleArray(v->len);
    memcpy(array_sorted, v->data, (v->len * sizeof(double))); // Create a copy of the array
    qsort(array_sorted, v->len, sizeof(double), compareDoublesSorting); // Sort the array

    // Safeguards
    int hi_midpoint = (int)floor(v->len / 2);
    if (hi_midpoint < 0) hi_midpoint = 0;
    if (hi_midpoint >= v->len) hi_midpoint = v->len - 1;

    double returns = array_sorted[hi_midpoint];
    free(array_sorted);
    return returns;
}

double median(Vector* v) {
    return v->len % 2 == 0 ? _median_even(v) : _median_odd(v);
}

double quantile(Vector* v, double p) {
    double* array_sorted = createDoubleArray(v->len);
    memcpy(array_sorted, v->data, (v->len * sizeof(double))); // Create a copy of the array
    qsort(array_sorted, v->len, sizeof(double), compareDoublesSorting); // Sort the array

    // Safeguards
    int idx_quantile = (int)floor(p * (v->len - 1));
    if (idx_quantile < 0) idx_quantile = 0;
    if (idx_quantile >= v->len) idx_quantile = v->len - 1;

    double returns = array_sorted[idx_quantile];
    free(array_sorted);
    return returns;
}


Vector mode(Vector* v) {
    double* dictArr = createDoubleArray(v->len);
    double* modes = createDoubleArray(v->len);
    size_t searchIdx = 0;
    size_t lastIdx = 0; 

    // Counter for the number of times a value appears (mode-counter)
    for(size_t i = 0; i < v->len; i++) {
        searchIdx = searchArray(dictArr, lastIdx, v->data[i]); 
        if (searchIdx == -1) {
            dictArr[lastIdx] = v->data[i];
            modes[lastIdx] = 1; 
            lastIdx++;
        } else {
            modes[searchIdx]++; 
        }
    }

    // Get the max from array 
    double max = maxArray(modes, lastIdx);

    // Get the length of the final array
    size_t arrLen = 0;
    for(size_t i = 0; i < lastIdx; i++) { if(modes[i] == max) arrLen++; } 

    // Get the modes (all value that has the max-counter)
    double finalModes[arrLen];
    searchIdx = 0;
    for(size_t i = 0; i < lastIdx; i++) {
        if(modes[i] == max) {
            finalModes[searchIdx] = dictArr[i];
            searchIdx++;
        }
    }
    
    free(dictArr); free(modes);
    return _vector_from_array(finalModes, arrLen);
}

double dataRange(Vector* v) {
    return max(v) - min(v); 
}

Vector deMean(Vector* v) {
    double x_bar = mean(v);
    double* new_data = createDoubleArray(v->len);

    for(size_t i = 0; i < v->len; i++)
        new_data[i] = v->data[i] - x_bar;

    Vector res; 
    res.data = new_data; 
    res.len = v->len;
    return res;
}

double variance(Vector* v) {
    if (v->len < 2) {
        printf("ERROR: Variance requires that the Vector has at least 2 elements in the Vector");
        return NAN;
    }

    Vector deviations = deMean(v);
    double sum = sumOfSquares(&deviations);
    freeVector(&deviations);

    return sum / (v->len - 1);
}

double standardDeviation(Vector* v) {
    return sqrt(variance(v));
}

double interquantileRange(Vector* v) {
    return quantile(v, 0.75) -  quantile(v, 0.25); 
}

double covariance(Vector* xs, Vector* ys) {
    if (xs->len != ys->len) {
        printf("ERROR: Covariance requires that both Vectors have the same length. Got %zu and %zu", xs->len, ys->len);
        return NAN;
    }

    Vector demean_xs = deMean(xs), demean_ys = deMean(ys); 
    return dotProduct(&demean_xs, &demean_ys) / (xs->len - 1);  
}

double correlation(Vector* xs, Vector* ys) {
    double stdev_x = standardDeviation(xs);
    double stdev_y = standardDeviation(ys);

    if (stdev_x > 0 && stdev_y > 0) {
        return covariance(xs, ys) / stdev_x / stdev_y;
    }
    return 0; // If no variation, correlation is zero
}

// ========================== PROBABILITY ==========================
double normalPdf(double x, double mu, double sigma) {
    double deviation = x - mu;
    double variance = sigma * sigma;
    double exponent = - (deviation * deviation) / (2 * variance);
    
    return exp(exponent) / (SQRT_TWO_PI * sigma);
}

double uniformPdf(double x) {
    return x >= 0 && x < 1 ? 1 : 0; 
}

double uniformCdf(double x) {
    if (x < 0) {
        return 0; 
    } else if (x < 1) {
        return x;
    } else {
        return 1; 
    }
}

double normalCdf(double x, double mu, double sigma) {
    return (1 + erf((x - mu) / sqrt(2) / sigma)) / 2;
}
    
double inverseNormalCdf(double p, double mu, double sigma, double tolerance) {
    // If not standard, compute standard and rescale
    if (mu != 0 || sigma != 1) {
        return mu + sigma * inverseNormalCdf(p, 0, 1, tolerance);
    }

    double low_z = -10.0, hi_z = 10, mid_z = 0, mid_p = 0; 
    while (hi_z - low_z > tolerance) {
        mid_z = (low_z + hi_z) / 2; 
        mid_p = normalCdf(mid_z, 0, 1); 
        if (mid_p < p) {
            low_z = mid_z;
        } else {
            hi_z = mid_z;
        }
    }
    return mid_z;
}

int bernoulliTrial(double p) {
    return randDouble() < p ? 1 : 0; 
}

int binomial(int n, double p) {
    double sum = 0; 
    for (size_t i = 0; i < n; i++)
        sum += bernoulliTrial(p);
        
    return sum;
}


int main() {
    double a_data[] = { 1.0, 2.0, 3.0, 3.0, 4.0, 2.0, 4.0, 7.0 };
    Vector a = createVector(a_data);

    Vector mod = mode(&a);
    print_vector(mod);

    print_vector(deMean(&a));

    Vector dmena = deMean(&a);
    printf("%.1f\n", mean(&dmena)); // Holy, it is really zero!
    printf("%.1f\n", mean(&a));
    printf("%.1f\n", standardDeviation(&a));
    printf("%.1f\n", quantile(&a, 0.25));
    printf("%.1f\n", median(&a));
    printf("%f\n", normalPdf(3.5, 0, 1));
    printf("%f\n", normalCdf(3.5, 0, 1));

    // double b_data[] = { 1.0, 2.0, 3.0 };
    // Vector b = createVector(b_data);


    // printf("--- Teste 1 ---\n");
    // Vector vectors1[] = { a, b };
    // Matrix m1 = createMatrixVector(vectors1);
    // Vector v = vectorSum(&m1);

    // double c_data[] = { 1.0, 2.0, 3.0 };
    // Vector c = createVector(c_data);
    
    // double d_data[] = { 1.0, 2.0, 3.0, 4.0, 5.0 };
    // Vector d = createVector(d_data);

    // printf("\n--- Teste 2 ---\n");
    // Vector vectors2[] = { c, d };
    // Matrix m2 = createMatrixVector(vectors2);
    // vectorSum(&m2);

    // print_vector(v);


    // double dot = dotProduct(&a, &b);
    // printf("\n%.1f\n", dot);

    // dot = sumOfSquares(&a);
    // printf("%.1f\n", dot);


    // Vector vm = getColumn(&m1, 1);
    // print_vector(vm);

    return 0;

}