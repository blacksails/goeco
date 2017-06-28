package goeco

import (
	"encoding/xml"
)

type soapEnvelope struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`

	Body soapBody `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
}

// SOAPFault is the error return on failing SOAP api calls
type SOAPFault struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Fault"`

	Reason soapReason `xml:"http://www.w3.org/2003/05/soap-envelope Reason"`
	Code   soapCode   `xml:"http://www.w3.org/2003/05/soap-envelope Code"`
	Detail soapDetail `xml:"http://www.w3.org/2003/05/soap-envelope Detail"`
}

func (f *SOAPFault) Error() string {
	return f.Reason.Text
}

type soapReason struct {
	Text string `xml:"http://www.w3.org/2003/05/soap-envelope Text"`
}

type soapCode struct {
	Value string `xml:"http://www.w3.org/2003/05/soap-envelope Value"`
}

type soapDetail struct{}

type soapBody struct {
	XMLName xml.Name    `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

func (b *soapBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be pointer to struct")
	}

	var (
		token    xml.Token
		consumed bool
		err      error
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}
		if token == nil {
			break
		}
		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			}
			if se.Name.Space == "http://www.w3.org/2003/05/soap-envelope" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil
				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}
				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}
				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}
	return nil
}
