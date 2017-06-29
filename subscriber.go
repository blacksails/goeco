package goeco

import "encoding/xml"

// SubscriberData represents data of a subscriber
type SubscriberData struct {
	SubscriptionHandle  subscriptionHandle `xml:"SubscriptionHandle"`
	StartDate           ecoTime            `xml:"StartDate"`
	RegisteredDate      ecoTime            `xml:"RegisteredDate"`
	EndDate             ecoTime            `xml:"EndDate"`
	ExpiryDate          ecoTime            `xml:"ExpiryDate"`
	ExtraTextForInvoice string             `xml:"ExtraTextForInvoice"`
	SpecialPrice        float64            `xml:"SpecialPrice"`
	QuantityFactor      float64            `xml:"QuantityFactor"`
}

type subscriberHandle struct {
	SubscriberID       int                `xml:"SubscriberId"`
	SubscriptionHandle subscriptionHandle `xml:"SubscriptionHandle"`
}

type subscriberGetDataArrayResponse struct {
	XMLName                      xml.Name                     `xml:"http://e-conomic.com Subscriber_GetDataArrayResponse"`
	SubscriberGetDataArrayResult subscriberGetDataArrayResult `xml:"Subscriber_GetDataArrayResult"`
}

type subscriberGetDataArrayResult struct {
	SubscriberDatas []SubscriberData `xml:"SubscriberData"`
}

type subscriberGetDataArray struct {
	EntityHandles subscriberHandles `xml:"entityHandles"`
}

type subscriberHandles struct {
	SubscriberHandles []subscriberHandle `xml:"SubscriberHandle"`
}

// SubscriberGetDataArray takes a list of subscriber handles and returns the subscriptions
func (c Client) SubscriberGetDataArray(handles []int) ([]SubscriberData, error) {
	res := &subscriberGetDataArrayResponse{}
	shs := make([]subscriberHandle, len(handles))
	for i, h := range handles {
		shs[i] = subscriberHandle{SubscriberID: h}
	}
	err := c.call(subscriberGetDataArray{EntityHandles: subscriberHandles{SubscriberHandles: shs}}, res)
	if err != nil {
		return []SubscriberData{}, err
	}
	return res.SubscriberGetDataArrayResult.SubscriberDatas, nil
}
