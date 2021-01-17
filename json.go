package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)
// JSON structure from API
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
//JSON structure for output
type jsonWrite struct {
	Gender string `json:"Gender"`
	First_name string `json:"First_name"`
	Last_name string `json:"Last_name"`
	Created_at string `json:"Created_at"`
	Postcode int`json:"Postcode"`
}

func readJson() (jsonWrite) {
	const urlSite = "https://randomuser.me/api/"
	var read jsonRead
	var write jsonWrite

	req, err := http.Get(urlSite)
	if err != nil {
		log.Printf("Error http.Get = %v\n",err)
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		log.Printf("API Error: Unable to get API request because HTTP status %d\n", req.StatusCode)
	}

	// read HTTP JSON body
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("API Error: Unable to get API request because err: %v with HTTP status %d\n", err, req.StatusCode)
	}

	// convert JSON to struct
	erj := json.Unmarshal(body, &read)
	if erj != nil {
		log.Printf("JSON parse err: %v\n", erj)
	}

	if err == nil	{
		write = jsonWrite{
			Gender:     read.Results[0].Gender,
			First_name: read.Results[0].Name.First,
			Last_name:  read.Results[0].Name.Last,
			Created_at: read.Results[0].Registered.Date,
			Postcode:   read.Results[0].Loc.Postcode,
		}
	}
	return write
}

func getListUsers(dateFrom,dateTo time.Time,count int) string {
	const layout = "2006-01-02T15:04:05.999Z"
	var c = 0

	var users = make(map[int]jsonWrite)

	for i:=0; i<count * 10;i++  {
		if c>count-1 {
			break
		}
		users[c] = readJson()
		fmt.Printf( showLog(users[c]) )

		dateRead, errData := time.Parse(layout,users[c].Created_at)
		if dateFrom.Before(dateRead) && dateTo.After(dateRead) {
			c++
		}
		if errData != nil {
			log.Printf("Error Parse=%v",errData)
		}
	}
	jitem, _ := json.Marshal(users)
	return string(jitem)
}

func showLog(r jsonWrite) string {
	return  fmt.Sprintf("First_name: %s Last_name:%s (%s) Created_at %s Postcode = %d \n",
		r.Last_name, r.First_name, r.Gender,r.Created_at, r.Postcode )
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

