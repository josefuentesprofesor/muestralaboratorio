package main

type Header struct {
	GTFSRealtimeVersion string `json:"gtfs_realtime_version"`
	Incrementality      int    `json:"incrementality"`
	Timestamp           int64  `json:"timestamp"`
}

type InformedEntity struct {
	AgencyID  string      `json:"agency_id"`
	RouteID   string      `json:"route_id"`
	RouteType int         `json:"route_type"`
	Trip      interface{} `json:"trip"`
	StopID    string      `json:"stop_id"`
}

type Alert struct {
	ActivePeriod   []interface{}       `json:"active_period"`
	InformedEntity []InformedEntityBus `json:"informed_entity"`
	Cause          int                 `json:"cause"`
	Effect         int                 `json:"effect"`
	URL            interface{}         `json:"url"`
	HeaderText     struct {
		Translation []struct {
			Text     string `json:"text"`
			Language string `json:"language"`
		} `json:"translation"`
	} `json:"header_text"`
	DescriptionText struct {
		Translation []struct {
			Text     string `json:"text"`
			Language string `json:"language"`
		} `json:"translation"`
	} `json:"description_text"`
}

type Entity struct {
	ID         string      `json:"id"`
	IsDeleted  bool        `json:"is_deleted"`
	TripUpdate interface{} `json:"trip_update"`
	Vehicle    interface{} `json:"vehicle"`
	Alert      AlertBus    `json:"alert"`
}

type GTFSRealtimeData struct {
	Header Header   `json:"header"`
	Entity []Entity `json:"entity"`
}
