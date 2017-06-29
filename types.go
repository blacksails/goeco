package goeco

import (
	"encoding/xml"
	"time"
)

type ecoTime struct {
	T time.Time
}

func (et *ecoTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var timeStr string
	ecoLayout := "2006-01-02T15:04:05"
	d.DecodeElement(&timeStr, &start)
	if timeStr == "" {
		et.T = time.Time{}
		return nil
	}
	t, err := time.Parse(ecoLayout, timeStr)
	if err != nil {
		return err
	}
	et.T = t
	return nil
}
