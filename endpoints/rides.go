package endpoints

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Rides struct {
	Rides []*Ride `json:"rides"`
}

type Ride struct {
	Id int `json:"id"`
	Distance float64  `json:"distance"`
	StartTime time.Time `json:"startTime"`
	Duration int  `json:"duration"`
	Price float64 `json:"price,omitempty"`
}

func NewRidesEndpoint(ridesFilePath string) (httprouter.Handle, error) {
	return func(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
		jsonFile, err := os.Open(ridesFilePath)
		if err != nil {
			glog.Errorf("Failed to open file: %v", err)
			return
		}
		defer func() {
			if err := jsonFile.Close(); err != nil {
				glog.Errorf("Failed to close file: %v", err)
			}
		}()
		var rides Rides
		byteValue, _ := ioutil.ReadAll(jsonFile)
		err = json.Unmarshal(byteValue, &rides)
		if err != nil {
			glog.Errorf("Error decoding stored json: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		computePrices(&rides)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(rides)
		if err != nil {
			glog.Errorf("Error encoding json: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}, nil
}

func computePrices(rides *Rides) {
	for _, ride := range rides.Rides {
		computePriceRide(ride)
	}
}

func computePriceRide(ride *Ride) {
	hours, _, _ := ride.StartTime.Clock()
	var between8Pm6Am float64
	var between4Pm7Pm float64
	if hours >= 20 || hours < 6 {
		between8Pm6Am = 1
		between4Pm7Pm = 0
	} else if hours >= 16 && hours < 19 {
		between8Pm6Am = 0
		between4Pm7Pm = 1
	} else {
		between8Pm6Am = 0
		between4Pm7Pm = 0
	}
	ride.Price = 1 + 0.5 * 5 * ride.Distance + 0.5 * between8Pm6Am + 1 * between4Pm7Pm
}
