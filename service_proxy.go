package koap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"reflect"
)

type serviceProxy struct {
	endpoint       string
	client         *Client
	service        *Service
	serviceVersion *ServiceVersion
}

func (s *serviceProxy) String() string {
	return fmt.Sprintf("%s version=%s endpoint=%s", s.service.Name, s.serviceVersion.Version, s.endpoint)
}

func (s *serviceProxy) CallWithDocument(operation Operation, document interface{}) (interface{}, error) {
	param, err := xml.Marshal(document)
	if err != nil {
		return nil, fmt.Errorf("error marshalling: %w", err)
	}

	envelope := Envelope{
		Body: Body{
			Content: param,
		},
	}

	request, err := newSOAPRequest(s.endpoint, operation.SOAPAction(), &envelope)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	resp, err := s.client.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, parseFaultResponse(resp, operation.FaultType())
	}

	var respEnvelope Envelope
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&respEnvelope)
	if err != nil {
		return nil, fmt.Errorf("decoding SOAP XML response: %v", err)
	}

	output := reflect.New(operation.OutputType()).Interface()
	err = xml.Unmarshal(respEnvelope.Body.Content, output)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling XML response into '%s': %w", operation.OutputType().String(), err)
	}

	return output, nil
}

func parseFaultResponse(resp *http.Response, faultType reflect.Type) error {
	var respEnvelope FaultEnvelope
	decoder := xml.NewDecoder(resp.Body)
	err := decoder.Decode(&respEnvelope)
	if err != nil {
		return fmt.Errorf("decoding fault response: %w", err)
	}
	fault := reflect.New(faultType).Interface()
	err = xml.Unmarshal([]byte(respEnvelope.Body.Fault.Detail.DetailRaw), fault)
	if err != nil {
		return fmt.Errorf("unmarshalling fault response into '%s': %w", faultType.String(), err)
	}

	respEnvelope.Body.Fault.Detail.Detail = fault

	return &respEnvelope.Body.Fault
}

func newSOAPRequest(ep, soapAction string, envelope *Envelope) (*http.Request, error) {
	payload, err := xml.Marshal(envelope)
	if err != nil {
		return nil, fmt.Errorf("marshalling envelope to xml: %w", err)
	}
	req, err := http.NewRequest("POST", ep, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("creating HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", soapAction)
	return req, nil
}
