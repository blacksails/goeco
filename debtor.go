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
	res := &debtorGetDataResponse{}
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
	res := &debtorGetNameResponse{}
	err := c.call(debtorGetName{DebtorHandle: debtorHandle{Number: handle}}, res)
	if err != nil {
		return "", err
	}
	return res.DebtorGetNameResult, nil
}

type debtorGetOpenEntriesResponse struct {
	XMLName                    xml.Name                   `xml:"http://e-conomic.com Debtor_GetOpenEntriesResponse"`
	DebtorGetOpenEntriesResult debtorGetOpenEntriesResult `xml:"Debtor_GetOpenEntriesResult"`
}

type debtorGetOpenEntriesResult struct {
	DebtorEntryHandles []entryHandle `xml:"DebtorEntryHandle"`
}

type debtorGetOpenEntries struct {
	XMLName      xml.Name     `xml:"http://e-conomic.com Debtor_GetOpenEntries"`
	DebtorHandle debtorHandle `xml:"debtorHandle"`
}

// DebtorGetOpenEntries gets the open entries of the debtor with the given handle
func (c Client) DebtorGetOpenEntries(handle int) ([]int, error) {
	res := &debtorGetOpenEntriesResponse{}
	err := c.call(debtorGetOpenEntries{DebtorHandle: debtorHandle{Number: handle}}, res)
	if err != nil {
		return []int{}, err
	}
	handles := make([]int, len(res.DebtorGetOpenEntriesResult.DebtorEntryHandles))
	for i, entry := range res.DebtorGetOpenEntriesResult.DebtorEntryHandles {
		handles[i] = entry.SerialNumber
	}
	return handles, nil
}

type debtorGetSubscribersResponse struct {
	XMLName                    xml.Name                   `xml:"http://e-conomic.com Debtor_GetSubscribersResponse"`
	DebtorGetSubscribersResult debtorGetSubscribersResult `xml:"Debtor_GetSubscribersResult"`
}

type debtorGetSubscribersResult struct {
	SubscriberHandles []subscriberHandle `xml:"SubscriberHandle"`
}

type debtorGetSubscribers struct {
	XMLName      xml.Name     `xml:"http://e-conomic.com Debtor_GetSubscribers"`
	DebtorHandle debtorHandle `xml:"debtorHandle"`
}

// DebtorGetSubscribers gets the handles of the subscribers tied to the debtor
func (c Client) DebtorGetSubscribers(handle int) ([]int, error) {
	res := &debtorGetSubscribersResponse{}
	err := c.call(debtorGetSubscribers{DebtorHandle: debtorHandle{Number: handle}}, res)
	if err != nil {
		return []int{}, err
	}
	sHandles := make([]int, len(res.DebtorGetSubscribersResult.SubscriberHandles))
	for i, handle := range res.DebtorGetSubscribersResult.SubscriberHandles {
		sHandles[i] = handle.SubscriberID
	}
	return sHandles, nil
}
