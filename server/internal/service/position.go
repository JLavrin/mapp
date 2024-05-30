package service

import (
	"fmt"
	"github.com/JLavrin/mapp.git/server/internal/util"
)

type Vehicle struct {
	Generated              string `json:"generated"`
	RouteShortName         string `json:"routeShortName"`
	RouteId                uint   `json:"routeId"`
	VehicleCode            string `json:"vehicleCode"`
	VehicleService         string `json:"vehicleService"`
	VehicleId              uint   `json:"vehicleId"`
	Speed                  uint   `json:"speed"`
	Delay                  int    `json:"delay"`
	Lat                    string `json:"lat"`
	Lon                    string `json:"lon"`
	GpsQuality             int    `json:"gpsQuality"`
	HeadSign               string `json:"headSign"`
	Direction              uint   `json:"direction"`
	ScheduledTripStartTime string `json:"scheduledTripStartTime"`
}

type Res struct {
	LastUpdate string    `json:"lastUpdate"`
	Vehicles   []Vehicle `json:"vehicles"`
}

func Position() {
	res := util.Request[Vehicle](util.Req{
		Url:    "https://ckan2.multimediagdansk.pl/gpsPositions?v=2",
		Method: "GET",
	})

	fmt.Println(res)
}
