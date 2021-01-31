package endpoints

import (
	"testing"
	"time"
)

func TestComputePriceRide(t *testing.T) {
	startTime, _ := time.Parse(time.RFC3339,"2020-06-19T14:01:17.031Z")
	ride := &Ride{
		Id:        1,
		Distance:  5,
		StartTime: startTime,
		Duration:  4000,
	}
	computePriceRide(ride)
	expectedPriceRide := 13.5 // 1 + (5 / (1 / 5)) * 0.5
	got := ride.Price
	if got != expectedPriceRide {
		t.Errorf("Computed price ride = %f; want %f", got, expectedPriceRide)
	}
}

func TestComputePriceRideBetween4Pm7Pm(t *testing.T) {
	startTime, _ := time.Parse(time.RFC3339,"2020-06-19T17:01:17.031Z")
	ride := &Ride{
		Id:        1,
		Distance:  5,
		StartTime: startTime,
		Duration:  4000,
	}
	computePriceRide(ride)
	expectedPriceRide := 14.5 // 1 + (5 / (1 / 5)) * 0.5 + 1
	got := ride.Price
	if got != expectedPriceRide {
		t.Errorf("Computed price ride = %f; want %f", got, expectedPriceRide)
	}
}

func TestComputePriceRideBetween8Pm6Am(t *testing.T) {
	startTime, _ := time.Parse(time.RFC3339,"2020-06-19T04:01:17.031Z")
	ride := &Ride{
		Id:        1,
		Distance:  5,
		StartTime: startTime,
		Duration:  4000,
	}
	computePriceRide(ride)
	expectedPriceRide := 14. // 1 + (5 / (1 / 5)) * 0.5 + 0.5
	got := ride.Price
	if got != expectedPriceRide {
		t.Errorf("Computed price ride = %f; want %f", got, expectedPriceRide)
	}
}