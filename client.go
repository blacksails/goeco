package goeco

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// Client is the main interaction point for the economic client
type Client struct {
	token    string
	appToken string
	cookies  []*http.Cookie
}

// New instantiates a new Client with the given token and app token
func New(token, appToken string) *Client {
	return &Client{
		token:    token,
		appToken: appToken,
	}
}

type connectWithToken struct {
	XMLName  xml.Name `xml:"http://e-conomic.com ConnectWithToken"`
	Token    string   `xml:"token"`
	AppToken string   `xml:"appToken"`
}

func (c Client) call(req, res interface{}) error {
	{
		// login
		res, err := c.buildAndSendRequest(connectWithToken{Token: c.token, AppToken: c.appToken})
		if err != nil {
			return err
		}
		c.cookies = res.Cookies()
	}

	httpRes, err := c.buildAndSendRequest(req)
	if err != nil {
		return err
	}

	httpBody, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	//log.Println(string(httpBody))

	// Decode response
	soapEnvel := new(soapEnvelope)
	soapEnvel.Body = soapBody{Content: res}
	err = xml.Unmarshal(httpBody, soapEnvel)
	if err != nil {
		return err
	}
	if fault := soapEnvel.Body.Fault; fault != nil {
		return fault
	}
	return nil
}

func (c Client) buildAndSendRequest(req interface{}) (*http.Response, error) {
	var b bytes.Buffer
	b.WriteString("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n")
	soapEnvel := soapEnvelope{
		Body: soapBody{Content: req},
	}
	xml.NewEncoder(&b).Encode(soapEnvel)

	//log.Println(b.String())

	httpReq, err := http.NewRequest("POST", "https://api.e-conomic.com/secure/api1/EconomicWebService.asmx", &b)
	if err != nil {
		return &http.Response{}, err
	}
	httpReq.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
	for _, cookie := range c.cookies {
		httpReq.AddCookie(cookie)
	}
	httpClient := http.Client{}
	return httpClient.Do(httpReq)
}
