package main

import (
    "net/http"
    "encoding/xml"
    "strings"
)

const VACATION_MODE = false
const DOORGATE_NUMBER = "+1[INSERT_NUMBER_HERE]"
const TWILIO_PHONE_NUMBER = "+1[INSERT_NUMBER_HERE]"
const SMS_ALERT_PHONE_NUMBER = "+1[INSERT_NUMBER_HERE]"
const OPEN_DOOR_DTMF_TONE = "tones/dtmf9.wav" // Change to different .wav file depending on what's required for your gate

type TwiMLResponse struct {
  XMLName xml.Name `xml:"Response"`
  Say string `xml:"Say"`
  Pause TwiMLPause `xml:"Pause"`
  Play string `xml:"Play"`
  SMS TwiMLSms `xml:"Sms"`
}

type TwiMLPause struct {
  Length int `xml:"length,attr"`
}

type TwiMLSms struct {
  From string `xml:"from,attr"`
  To string `xml:"to,attr"`
  Message string `xml:",chardata"`
}

func init() {
  http.HandleFunc("/", responseHandler)
}

func responseHandler(w http.ResponseWriter, r *http.Request) {

  // Fetch caller ID from GET variables since this is how Twilio sends it
  callerId := r.URL.Query().Get("From")
  var message string
  var smsMessage string
  tone := OPEN_DOOR_DTMF_TONE
  pause := TwiMLPause{Length: 2}
  
  if VACATION_MODE {
  	message = "Not available at this time"
  	tone = "" // No tone is played, door will not be opened
  } else {
    if strings.Contains(callerId, DOORGATE_NUMBER) {
  	  message = "Opening door, please come in"
  	  smsMessage = "Someone at the door."
    } else {
      message = "Please leave a message after the beep"
      smsMessage = "Non-gate number called."
    }
  }

  twimlsms := TwiMLSms{From: TWILIO_PHONE_NUMBER, To: SMS_ALERT_PHONE_NUMBER, Message: smsMessage}
  twiml := TwiMLResponse{Say: message, Pause: pause, Play: tone, SMS: twimlsms}

  x, err := xml.MarshalIndent(twiml, "", "  ")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/xml")
  w.Write(x)
}
