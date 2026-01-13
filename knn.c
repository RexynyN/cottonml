#include <stdlib.h>
#include <stdio.h> 
#include <float.h>
#include "utils.h"
#include "counter.h"
#include "linearalg.h"


typedef struct KNN {
    Matrix data;
    Vector labels; 
    int kNeighbors; 
} KNN;


static double rawMajorityVote(Vector* labels) {
    Counter counter = counterFromVector(labels);
    CounterResult result = counterMostCommon(&counter, 1);

    return result.data[0].label;
}

static double majorityVote(Vector* labels) {
    Counter counter = counterFromVector(labels);
    int numWinners = counterWinnerCount(&counter); 

    if (numWinners == 1) { // We have a winning label
        CounterResult result = counterMostCommon(&counter, 1);
        return result.data[0].label;
    } 
    
    // Try again, without the furthest
    numWinners -= 1; 
    CounterResult result = counterMostCommon(&counter, numWinners);  
    
    double* data = ALLOC_ARRAY(data, numWinners);
    for (size_t i = 0; i < numWinners; i++){
        data[i] = result.data[i].label;
    }

    Vector v = _vector_from_array(data, numWinners); 
    return majorityVote(&v);
}

double knnClassify(KNN* model, Vector* newPoint) {
    double high = DBL_MAX; // Basically infinity 
    size_t highIdx = 0;
    Vector vec; 

    double* kNearest = ALLOC_ARRAY(kNearest, model->kNeighbors);
    // Initialize the array, for good measure
    for (size_t i = 0; i < model->data.rows; i++)  kNearest[i] = high;

    // Will get the k-neighbors
    for (size_t i = 0; i < model->data.rows; i++){
        vec = matrixGetRow(&(model->data), i);
        double dist = distance(&vec, newPoint);

        if (dist >= high) 
            continue; // Its not lower than any of the K-lowest neighbors

        kNearest[highIdx] = dist; // New value in K-nearest
        
        high = kNearest[0]; 
        for (size_t i = 0; i < model->kNeighbors; i++){ // Gets the new ceiling variable
            if (kNearest[i] > high) {
                high = kNearest[i];
                highIdx = i; 
            }
        }
    }

    Vector kLabels = _vector_from_array(kNearest, model->kNeighbors);
    return majorityVote(&kLabels);
}

