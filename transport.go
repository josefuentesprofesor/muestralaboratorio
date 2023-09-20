package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"google.golang.org/protobuf/proto"
)

func Transport(clientId string, clientSecret string) string {
	url := "https://apitransporte.buenosaires.gob.ar/ecobici/gbfs/stationInformation?client_id=" + clientId + "&client_secret=" + clientSecret

	fname := "response.json"

	fmt.Println(url)

	f, err := os.Create(fname)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	res, err := http.Get(url)

	var stationData StationData
	// Deserializar los datos JSON en la estructura Person
	err = json.NewDecoder(res.Body).Decode(&stationData)
	if err != nil {
		fmt.Println("Error al decodificar los datos JSON:", err)
	}

	fmt.Println(stationData.Data.Stations[0].Address)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	var nbytes int64
	nbytes, err = io.Copy(f, res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fname, " downloaded, ", nbytes, " bytes")
	return "nada"

}

func StationInfo(clientId string, clientSecret string, stationId int) string {
	var stationData StationData
	var stationStatus StationStatus
	var result string

	serviceUrl := "https://apitransporte.buenosaires.gob.ar/ecobici/gbfs/stationInformation?client_id=" + clientId + "&client_secret=" + clientSecret
	res, err := http.Get(serviceUrl)

	// Deserializar los datos JSON en la estructura Person
	err = json.NewDecoder(res.Body).Decode(&stationData)
	if err != nil {
		fmt.Println("Error al decodificar los datos JSON:", err)
	}

	for _, station := range stationData.Data.Stations {
		if station.StationID == strconv.Itoa(stationId) {
			result = "Nombre: " + station.Name + "\n"
			result = result + "Direccion: " + station.Address + "\n"
			result = result + "Capacidad: " + strconv.Itoa(station.Capacity) + "\n"
			result = result + "https://maps.google.com/?q=" + strconv.FormatFloat(station.Lat, 'f', 14, 64) + "," + strconv.FormatFloat(station.Lon, 'f', 14, 64) + "\n"
		}

	}

	serviceUrl = "https://apitransporte.buenosaires.gob.ar/ecobici/gbfs/stationStatus?client_id=" + clientId + "&client_secret=" + clientSecret
	res, err = http.Get(serviceUrl)

	// Deserializar los datos JSON en la estructura Person
	err = json.NewDecoder(res.Body).Decode(&stationStatus)
	if err != nil {
		fmt.Println("Error al decodificar los datos JSON:", err)
	}

	for _, station := range stationStatus.Data.Stations {
		if station.StationID == strconv.Itoa(stationId) {
			t := time.Unix(station.LastReported, 0)
			result = result + "Ultimo reporte: " + t.Format("2006-01-02 15:04:05") + "\n"
			result = result + "Estado: " + station.Status + "\n"
			result = result + "Bicicletas disponibles: " + strconv.Itoa(station.NumBikesAvailable) + "\n"
			result = result + "Bicicletas averiadas: " + strconv.Itoa(station.NumBikesDisabled) + "\n"
			result = result + "Docks disponibles: " + strconv.Itoa(station.NumDocksAvailable) + "\n"
			result = result + "Docks inhabilitados: " + strconv.Itoa(station.NumDocksDisabled) + "\n"
		}

	}
	return result

}

func SubwayAlerts(clientId string, clientSecret string) string {
	var gtfsData GTFSRealtimeData

	var result string

	serviceUrl := "https://apitransporte.buenosaires.gob.ar/subtes/serviceAlerts?json=1&client_id=" + clientId + "&client_secret=" + clientSecret

	fmt.Println(serviceUrl)
	res, err := http.Get(serviceUrl)

	// Deserializar los datos JSON en la estructura Person
	err = json.NewDecoder(res.Body).Decode(&gtfsData)
	if err != nil {
		fmt.Println("Error al decodificar los datos JSON:", err)
	}

	fmt.Println(gtfsData)

	t := time.Unix(gtfsData.Header.Timestamp, 0)
	result = "Fecha: " + t.Format("2006-01-02 15:04:05") + "\n"

	for _, entity := range gtfsData.Entity {
		for _, tr := range entity.Alert.HeaderText.Translation {
			result = result + "Alerta: " + tr.Text + "\n"
		}
	}

	return result
}

func BusAlerts(clientId string, clientSecret string) string {
	//var busData BusRealtimeData

	var result string

	serviceUrl := "https://apitransporte.buenosaires.gob.ar/colectivos/serviceAlerts/?client_id=" + clientId + "&client_secret=" + clientSecret

	/*	fmt.Println(serviceUrl)
		resp, err := http.Get(serviceUrl)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	*/
	client := &http.Client{}
	req, err := http.NewRequest("GET", serviceUrl, nil)
	//req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	feed := gtfs.FeedMessage{}
	err = proto.Unmarshal(body, &feed)
	if err != nil {
		log.Fatal(err)
	}

	for _, entity := range feed.Entity {
		tripUpdate := entity.GetTripUpdate()
		trip := tripUpdate.GetTrip()
		tripId := trip.GetTripId()
		fmt.Printf("Trip ID: %s\n", tripId)
	}

	/*	for _, entity := range feed.Entity {
			tripUpdate := entity.GetTripUpdate()
			trip := tripUpdate.GetTrip()
			tripId := trip.GetTripId()
			fmt.Printf("Trip ID: %s\n", tripId)
			result = result + tripId + "\n"
		}
	*/

	/*
		// Deserializar los datos JSON en la estructura Person
		err = json.NewDecoder(res.Body).Decode(&busData)
		if err != nil {
			fmt.Println("Error al decodificar los datos JSON:", err)
		}

		//	fmt.Println(busData)

		t := time.Unix(busData.Header.Timestamp, 0)
		result = "Fecha: " + t.Format("2006-01-02 15:04:05") + "\n"

		for _, entity := range busData.Entity {
			for _, tr := range entity.Alert.DescriptionText.Translation {
				result = result + "Alerta: " + tr.Text + "\n"
			}
		}
	*/
	return result
}
