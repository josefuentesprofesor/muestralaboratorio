package main

type Translation struct {
	Language string `json:"_language"`
	Text     string `json:"_text"`
}

type DescriptionText struct {
	Translation     []Translation `json:"_translation"`
	ExtensionObject interface{}   `json:"extensionObject"`
}

type HeaderText struct {
	Translation     []Translation `json:"_translation"`
	ExtensionObject interface{}   `json:"extensionObject"`
}

type InformedEntityBus struct {
	AgencyID        string      `json:"_agency_id"`
	RouteID         string      `json:"_route_id"`
	RouteType       int         `json:"_route_type"`
	StopID          string      `json:"_stop_id"`
	Trip            interface{} `json:"_trip"`
	ExtensionObject interface{} `json:"extensionObject"`
}

type AlertBus struct {
	ActivePeriod    []interface{}       `json:"_active_period"`
	Cause           int                 `json:"_cause"`
	DescriptionText DescriptionText     `json:"_description_text"`
	Effect          int                 `json:"_effect"`
	HeaderText      HeaderText          `json:"_header_text"`
	InformedEntity  []InformedEntityBus `json:"_informed_entity"`
	URL             struct {
		Translation     []interface{} `json:"_translation"`
		ExtensionObject interface{}   `json:"extensionObject"`
	} `json:"_url"`
	ExtensionObject interface{} `json:"extensionObject"`
}

type HeaderBus struct {
	GtfsRealtimeVersion string      `json:"_gtfs_realtime_version"`
	Incrementality      int         `json:"_incrementality"`
	Timestamp           int64       `json:"_timestamp"`
	ExtensionObject     interface{} `json:"extensionObject"`
}

type EntityBus struct {
	Alert           AlertBus    `json:"_alert"`
	ID              string      `json:"_id"`
	IsDeleted       bool        `json:"_is_deleted"`
	TripUpdate      interface{} `json:"_trip_update"`
	Vehicle         interface{} `json:"_vehicle"`
	ExtensionObject interface{} `json:"extensionObject"`
}

type BusRealtimeData struct {
	Entity          []EntityBus `json:"_entity"`
	Header          HeaderBus   `json:"_header"`
	ExtensionObject interface{} `json:"extensionObject"`
}
