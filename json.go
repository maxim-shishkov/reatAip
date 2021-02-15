package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type jsonRead struct {
	Results []struct {
		Gender string `json:"gender"`
		Name   struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Loc struct {
			Postcode int `json:"postcode"`
		} `json:"location"`

		Email string `json:"email"`
		Registered struct{
			Date string `json:"date"`
		} `json:"registered"`

	} `json:"results"`
}

type jsonWrite struct {
	Gender string `json:"Gender"`
	First_name string `json:"First_name"`
	Last_name string `json:"Last_name"`
	Created_at string `json:"Created_at"`
	Postcode int `json:"Postcode"`
}


func checkDate(from, to, readDate time.Time) bool {
	return from.Before(readDate) && to.After(readDate)
}

func checkInputDate(from,to string) (error, time.Time,time.Time) {
	const layout  = "2006-01-02T15:04"
	var err error

	dateFrom, errFrom := time.Parse(layout,from)
	dateTo, errTo := time.Parse(layout,to)
	if errFrom != nil || errTo != nil {
		err = errors.New("Invalid format date")
	}
	return err,dateFrom,dateTo
}

func readData(dateFrom,dateTo time.Time) jsonWrite {
	const urlSite = "https://randomuser.me/api/"
	const LAYOUT = "2006-01-02T15:04:05.999Z"
	var read jsonRead
	var write jsonWrite
	for {
		req, err := http.Get(urlSite)
		if err != nil{
			log.Printf("Error http.Get: %v\n",err)
			continue
		}
		defer req.Body.Close()

		if req.StatusCode != http.StatusOK {
			log.Printf("API Error: Unable to get API request because HTTP status %d\n", req.StatusCode)
		}

		// read HTTP JSON body
		body, err := ioutil.ReadAll(req.Body)
		if err != nil{
			log.Printf("API Error: Unable to get API request because err: %v\n",err)
			continue
		}

		// convert JSON to struct
		err = json.Unmarshal(body, &read)
		if err != nil{
			log.Printf("JSON parse err: %v\n",err)
			continue
		}
		c := &count

		dateRead , _ := time.Parse(LAYOUT,read.Results[0].Registered.Date)

		if checkDate(dateFrom,dateTo,dateRead) {
			write = jsonWrite{
				Gender:     read.Results[0].Gender,
				First_name: read.Results[0].Name.First,
				Last_name:  read.Results[0].Name.Last,
				Created_at: read.Results[0].Registered.Date,
				Postcode:   read.Results[0].Loc.Postcode,
			}
			*c++
			break
		}
	}
	return write
}

func writePost(from,to time.Time,site string) string {
	requestBody, err := json.Marshal( map[string]string{
		"status": "Success",
		"from": from.String(),
		"to": to.String(),
	})
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(site,"application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	boby, err := ioutil.ReadAll( resp.Body )
	if err != nil {
		log.Printf("Invalid format date %v\n",err)
	}
	return string(boby)
}
