package main

import (
	"testing"
)

func TestPluralityLearner(t *testing.T) {
	iris := ReadData("iris.csv")

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

	simple := []Item{
		Item{features: []float64{3.1, 3.3}, class: "A"},
		Item{features: []float64{3.3, 3.0}, class: "A"},
		Item{features: []float64{3.5, 3.5}, class: "A"},
		Item{features: []float64{0.7, 0.3}, class: "B"},
		Item{features: []float64{0.5, 0.9}, class: "B"},
	}

	pl = PluralityLearner(simple)
	if pl([]float64{3.5, 3.0}) != "A" {
		t.Error("Expected A")
	}
}

func TestNearestNeighbors(t *testing.T) {
	iris := ReadData("iris.csv")
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
