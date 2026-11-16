// Package spentenergy calculates distances, speeds and calories spent during activity.
package spentenergy

import (
	"fmt"
	"math"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

// Distance returns the distance in kilometres for a height in metres.
func Distance(steps int, height float64) float64 {
	return float64(steps) * (height * stepLengthCoefficient) / mInKm
}

// MeanSpeed returns the average speed in kilometres per hour.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// RunningSpentCalories validates the activity and calculates running calories.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}
	if weight <= 0 || math.IsNaN(weight) || math.IsInf(weight, 0) {
		return 0, fmt.Errorf("вес должен быть положительным конечным числом")
	}
	if height <= 0 || math.IsNaN(height) || math.IsInf(height, 0) {
		return 0, fmt.Errorf("рост должен быть положительным конечным числом")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("длительность должна быть больше нуля")
	}
	speed := MeanSpeed(steps, height, duration)
	return weight * speed * duration.Minutes() / minInH, nil
}

// WalkingSpentCalories calculates walking calories using the correction factor.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calories * walkingCaloriesCoefficient, nil
}
