package main

// Estructura de datos para el JSON
type StationData struct {
	LastUpdated int64 `json:"last_updated"`
	TTL         int   `json:"ttl"`
	Data        struct {
		Stations []struct {
			StationID             string   `json:"station_id"`
			Name                  string   `json:"name"`
			PhysicalConfiguration string   `json:"physical_configuration"`
			Lat                   float64  `json:"lat"`
			Lon                   float64  `json:"lon"`
			Altitude              int      `json:"altitude"`
			Address               string   `json:"address"`
			PostCode              string   `json:"post_code"`
			Capacity              int      `json:"capacity"`
			IsChargingStation     bool     `json:"is_charging_station"`
			RentalMethods         []string `json:"rental_methods"`
			Groups                []string `json:"groups"`
			OBCN                  string   `json:"obcn"`
			NearbyDistance        int      `json:"nearby_distance"`
			RideCodeSupport       bool     `json:"_ride_code_support"`
			RentalURIs            struct{} `json:"rental_uris"`
		} `json:"stations"`
	} `json:"data"`
}

type StationStatus struct {
	LastUpdated int64 `json:"last_updated"`
	TTL         int   `json:"ttl"`
	Data        struct {
		Stations []struct {
			StationID              string `json:"station_id"`
			NumBikesAvailable      int    `json:"num_bikes_available"`
			NumBikesAvailableTypes struct {
				Mechanical int `json:"mechanical"`
				Ebike      int `json:"ebike"`
			} `json:"num_bikes_available_types"`
			NumBikesDisabled  int       `json:"num_bikes_disabled"`
			NumDocksAvailable int       `json:"num_docks_available"`
			NumDocksDisabled  int       `json:"num_docks_disabled"`
			LastReported      int64     `json:"last_reported"`
			IsChargingStation bool      `json:"is_charging_station"`
			Status            string    `json:"status"`
			IsInstalled       int       `json:"is_installed"`
			IsRenting         int       `json:"is_renting"`
			IsReturning       int       `json:"is_returning"`
			Traffic           *struct{} `json:"traffic"`
		} `json:"stations"`
	} `json:"data"`
}
