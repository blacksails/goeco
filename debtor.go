package goeco

import "encoding/xml"

type debtorGetDataResponse struct {
	XMLName             xml.Name   `xml:"http://e-conomic.com Debtor_GetDataResponse"`
	DebtorGetDataResult DebtorData `xml:"Debtor_GetDataResult"`
}

// DebtorData is an economic type representing data of a debtor
type DebtorData struct {
	Number     int     `xml:"Number"`
	Name       string  `xml:"Name"`
	Email      string  `xml:"Email"`
	Address    string  `xml:"Address"`
	City       string  `xml:"City"`
	PostalCode string  `xml:"PostalCode"`
	Balance    float64 `xml:"Balance"`
}

type debtorHandle struct {
	Number int `xml:"Number"`
}

type debtorGetData struct {
	XMLName      xml.Name     `xml:"http://e-conomic.com Debtor_GetData"`
	EntityHandle debtorHandle `xml:"entityHandle"`
}

// DebtorGetData gets a DebtorData from the debtor with the given handle
func (c Client) DebtorGetData(handle int) (DebtorData, error) {
	res := new(debtorGetDataResponse)
	err := c.call(debtorGetData{EntityHandle: debtorHandle{Number: handle}}, res)
	if err != nil {
		return DebtorData{}, err
	}
	return res.DebtorGetDataResult, nil
}

type debtorGetNameResponse struct {
	XMLName             xml.Name `xml:"http://e-conomic.com Debtor_GetNameResponse"`
	DebtorGetNameResult string   `xml:"Debtor_GetNameResult"`
}

type debtorGetName struct {
	XMLName      xml.Name     `xml:"http://e-conomic.com Debtor_GetName"`
	DebtorHandle debtorHandle `xml:"debtorHandle"`
}

// DebtorGetName gets the name of the debtor with the given handle
func (c Client) DebtorGetName(handle int) (string, error) {
	res := new(debtorGetNameResponse)
	err := c.call(debtorGetName{DebtorHandle: debtorHandle{Number: handle}}, res)
	if err != nil {
		return "", err
	}
	return res.DebtorGetNameResult, nil
}
