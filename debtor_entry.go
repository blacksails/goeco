package goeco

import (
	"encoding/xml"
)

// DebtorEntryData contains data about a debtor entry
type DebtorEntryData struct {
	DueDate                  ecoTime `xml:"DueDate"`
	RemainderDefaultCurrency float64 `xml:"RemainderDefaultCurrency"`
}

type entryHandle struct {
	SerialNumber int `xml:"SerialNumber"`
}

type entityHandleList struct {
	EntityHandles []entryHandle `xml:"DebtorEntryHandle"`
}

type debtorEntryGetDataArrayResponse struct {
	XMLName                       xml.Name                      `xml:"http://e-conomic.com DebtorEntry_GetDataArrayResponse"`
	DebtorEntryGetDataArrayResult debtorEntryGetDataArrayResult `xml:"DebtorEntry_GetDataArrayResult"`
}

type debtorEntryGetDataArray struct {
	XMLName       xml.Name         `xml:"http://e-conomic.com DebtorEntry_GetDataArray"`
	EntityHandles entityHandleList `xml:"entityHandles"`
}

type debtorEntryGetDataArrayResult struct {
	DebtorEntryDatas []DebtorEntryData `xml:"DebtorEntryData"`
}

// DebtorEntryGetDataArray gets a list of debtor entry data from a list of entry handles
func (c Client) DebtorEntryGetDataArray(handles []int) ([]DebtorEntryData, error) {
	res := &debtorEntryGetDataArrayResponse{}
	ehl := &entityHandleList{EntityHandles: make([]entryHandle, len(handles))}
	for i, handle := range handles {
		ehl.EntityHandles[i] = entryHandle{SerialNumber: handle}
	}
	err := c.call(debtorEntryGetDataArray{EntityHandles: *ehl}, res)
	if err != nil {
		return []DebtorEntryData{}, err
	}
	return res.DebtorEntryGetDataArrayResult.DebtorEntryDatas, nil
}
