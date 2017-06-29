package goeco

import "encoding/xml"

// SubscriptionData represents a subscription
type SubscriptionData struct {
	Name        string `xml:"Name"`
	Description string `xml:"Description"`
}

type subscriptionHandle struct {
	ID int `xml:"Id"`
}

type subscriptionHandles struct {
	SubscriptionHandles []subscriptionHandle `xml:"SubscriptionHandle"`
}

type subscriptionGetDataArrayResponse struct {
	XMLName                        xml.Name                       `xml:"http://e-conomic.com Subscription_GetDataArrayResponse"`
	SubscriptionGetDataArrayResult subscriptionGetDataArrayResult `xml:"Subscription_GetDataArrayResult"`
}

type subscriptionGetDataArrayResult struct {
	SubscriptionDatas []SubscriptionData `xml:"SubscriptionData"`
}

type subscriptionGetDataArray struct {
	XMLName       xml.Name            `xml:"http://e-conomic.com Subscription_GetDataArray"`
	EntityHandles subscriptionHandles `xml:"entityHandles"`
}

// SubscriptionGetDataArray takes a list of subscription handles and returns the subscriptions
func (c Client) SubscriptionGetDataArray(handles []int) ([]SubscriptionData, error) {
	res := &subscriptionGetDataArrayResponse{}
	shs := make([]subscriptionHandle, len(handles))
	for i, h := range handles {
		shs[i] = subscriptionHandle{ID: h}
	}
	err := c.call(subscriptionGetDataArray{EntityHandles: subscriptionHandles{SubscriptionHandles: shs}}, res)
	if err != nil {
		return []SubscriptionData{}, err
	}
	return res.SubscriptionGetDataArrayResult.SubscriptionDatas, nil
}
