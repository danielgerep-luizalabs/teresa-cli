package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/teresa-api/models"
)

// UpdateAppReader is a Reader for the UpdateApp structure.
type UpdateAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UpdateAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateAppDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewUpdateAppOK creates a UpdateAppOK with default headers values
func NewUpdateAppOK() *UpdateAppOK {
	return &UpdateAppOK{}
}

/*UpdateAppOK handles this case with default header values.

Updated version of the app
*/
type UpdateAppOK struct {
	Payload *models.App
}

func (o *UpdateAppOK) Error() string {
	return fmt.Sprintf("[PUT /apps/{app_name}][%d] updateAppOK  %+v", 200, o.Payload)
}

func (o *UpdateAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.App)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppDefault creates a UpdateAppDefault with default headers values
func NewUpdateAppDefault(code int) *UpdateAppDefault {
	return &UpdateAppDefault{
		_statusCode: code,
	}
}

/*UpdateAppDefault handles this case with default header values.

Unexpected error
*/
type UpdateAppDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update app default response
func (o *UpdateAppDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAppDefault) Error() string {
	return fmt.Sprintf("[PUT /apps/{app_name}][%d] updateApp default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAppDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
