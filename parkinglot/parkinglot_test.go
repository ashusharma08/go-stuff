package parkinglot

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func initParkingLot() (*ParkingLot, error) {
	slots := make([]*Slot, 0, 100)
	slotType := func(i int) string {
		if i > 80 {
			return "hmv"
		}
		if i%3 == 0 {
			return "motorcycle"
		}
		return "lmv"
	}
	for i := 0; i < 100; i++ {
		slots = append(slots, &Slot{
			ID:   fmt.Sprintf("SL-%d", i),
			Type: OccupancyType(slotType(i)),
		})
	}
	return NewParkingLot(context.Background(), slots)
}
func Test_Enter(t *testing.T) {
	pk, err := initParkingLot()
	if err != nil {
		t.Fatalf("expected nil got %#v", err)
		return
	}
	tests := []struct {
		vehicleNumber string
		vehicleType   OccupancyType
		exitStrategy  func(string)
	}{
		{
			vehicleNumber: "HP37E4190",
			vehicleType:   "lmv",
		},
		{
			vehicleNumber: "KA03JT1747",
			vehicleType:   "motorcycle",
		},
	}

	for _, tt := range tests {
		tid, err := pk.Enter(tt.vehicleNumber, tt.vehicleType)
		if err != nil {
			t.Fatalf("expected nil got %#v", err)
			return
		}
		if len(tid) == 0 {
			t.Fatalf("expected ticket get empty string")
		}
		fmt.Println("vehicle", tt.vehicleNumber, "ticket:", tid)
		time.Sleep(20 * time.Second)

		err = pk.Exit(tid)
		if err != nil {
			t.Fatalf("expected nil got %#v", err)
			return
		}
	}
}
