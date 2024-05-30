package service

import (
	"github.com/JLavrin/mapp.git/server/internal/util"
	"time"
)

var VehicleUpdates chan Res

type Vehicle struct {
	Generated              string  `json:"generated"`
	RouteShortName         string  `json:"routeShortName"`
	RouteId                uint    `json:"routeId"`
	VehicleCode            string  `json:"vehicleCode"`
	VehicleService         string  `json:"vehicleService"`
	VehicleId              uint    `json:"vehicleId"`
	Speed                  uint    `json:"speed"`
	Delay                  int     `json:"delay"`
	Lat                    float32 `json:"lat"`
	Lon                    float32 `json:"lon"`
	GpsQuality             int     `json:"gpsQuality"`
	HeadSign               string  `json:"headSign"`
	Direction              uint    `json:"direction"`
	ScheduledTripStartTime string  `json:"scheduledTripStartTime"`
}

type Res struct {
	LastUpdate string    `json:"lastUpdate"`
	Vehicles   []Vehicle `json:"vehicles"`
}

func Position() {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	var lastUpdate Res

	for {
		select {
		case <-ticker.C:
			res := util.Request[Res](util.Req{
				Url:    "https://ckan2.multimediagdansk.pl/gpsPositions?v=2",
				Method: "GET",
			})

			if res.LastUpdate != lastUpdate.LastUpdate {
				lastUpdate = res
				VehicleUpdates <- res
			}
		}
	}
}

func StartUpdating() {
	VehicleUpdates = make(chan Res)
	go Position()
}
