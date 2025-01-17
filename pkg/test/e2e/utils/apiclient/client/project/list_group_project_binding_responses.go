// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListGroupProjectBindingReader is a Reader for the ListGroupProjectBinding structure.
type ListGroupProjectBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGroupProjectBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGroupProjectBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListGroupProjectBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListGroupProjectBindingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListGroupProjectBindingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGroupProjectBindingOK creates a ListGroupProjectBindingOK with default headers values
func NewListGroupProjectBindingOK() *ListGroupProjectBindingOK {
	return &ListGroupProjectBindingOK{}
}

/* ListGroupProjectBindingOK describes a response with status code 200, with default header values.

GroupProjectBinding
*/
type ListGroupProjectBindingOK struct {
	Payload []*models.GroupProjectBinding
}

func (o *ListGroupProjectBindingOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/groupbindings][%d] listGroupProjectBindingOK  %+v", 200, o.Payload)
}
func (o *ListGroupProjectBindingOK) GetPayload() []*models.GroupProjectBinding {
	return o.Payload
}

func (o *ListGroupProjectBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGroupProjectBindingUnauthorized creates a ListGroupProjectBindingUnauthorized with default headers values
func NewListGroupProjectBindingUnauthorized() *ListGroupProjectBindingUnauthorized {
	return &ListGroupProjectBindingUnauthorized{}
}

/* ListGroupProjectBindingUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListGroupProjectBindingUnauthorized struct {
}

func (o *ListGroupProjectBindingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/groupbindings][%d] listGroupProjectBindingUnauthorized ", 401)
}

func (o *ListGroupProjectBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListGroupProjectBindingForbidden creates a ListGroupProjectBindingForbidden with default headers values
func NewListGroupProjectBindingForbidden() *ListGroupProjectBindingForbidden {
	return &ListGroupProjectBindingForbidden{}
}

/* ListGroupProjectBindingForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListGroupProjectBindingForbidden struct {
}

func (o *ListGroupProjectBindingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/groupbindings][%d] listGroupProjectBindingForbidden ", 403)
}

func (o *ListGroupProjectBindingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListGroupProjectBindingDefault creates a ListGroupProjectBindingDefault with default headers values
func NewListGroupProjectBindingDefault(code int) *ListGroupProjectBindingDefault {
	return &ListGroupProjectBindingDefault{
		_statusCode: code,
	}
}

/* ListGroupProjectBindingDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListGroupProjectBindingDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list group project binding default response
func (o *ListGroupProjectBindingDefault) Code() int {
	return o._statusCode
}

func (o *ListGroupProjectBindingDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/groupbindings][%d] listGroupProjectBinding default  %+v", o._statusCode, o.Payload)
}
func (o *ListGroupProjectBindingDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGroupProjectBindingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
