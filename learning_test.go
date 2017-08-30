package main

import (
	"testing"
)

func TestPluralityLearner(t *testing.T) {
	iris := ReadData("iris.txt")

	pl := PluralityLearner(iris)
	if pl([]float64{4.9, 3.0, 1.4, 0.2}) != "" {
		t.Error("Expected empty string")
	}

	if pl([]float64{5.0, 2.0, 3.5, 1.0}) != "" {
		t.Error("Expected empty string")
	}

	if pl([]float64{6.7, 3.3, 5.7, 2.1}) != "" {
		t.Error("Expected empty string")
	}
}

func TestNearestNeighbors(t *testing.T) {
	iris := ReadData("iris.txt")
	kNN := NearestNeighborLearner(iris, 3)

	if kNN([]float64{4.9, 3.0, 1.4, 0.2}) != "setosa" {
		t.Error("Expected setosa")
	}

	if kNN([]float64{6.0, 4.8, 3.1, 1.5}) != "versicolor" {
		t.Error("Expected versicolor")
	}

	if kNN([]float64{7.5, 4.1, 6.1, 2.0}) != "virginica" {
		t.Error("Expected virginica")
	}
}
