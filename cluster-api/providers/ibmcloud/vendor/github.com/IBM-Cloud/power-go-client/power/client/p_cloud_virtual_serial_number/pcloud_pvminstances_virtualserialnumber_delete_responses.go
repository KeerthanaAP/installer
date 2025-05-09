// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_virtual_serial_number

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesVirtualserialnumberDeleteReader is a Reader for the PcloudPvminstancesVirtualserialnumberDelete structure.
type PcloudPvminstancesVirtualserialnumberDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesVirtualserialnumberDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudPvminstancesVirtualserialnumberDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesVirtualserialnumberDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesVirtualserialnumberDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesVirtualserialnumberDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesVirtualserialnumberDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudPvminstancesVirtualserialnumberDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesVirtualserialnumberDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/virtual-serial-number] pcloud.pvminstances.virtualserialnumber.delete", response, response.Code())
	}
}

// NewPcloudPvminstancesVirtualserialnumberDeleteAccepted creates a PcloudPvminstancesVirtualserialnumberDeleteAccepted with default headers values
func NewPcloudPvminstancesVirtualserialnumberDeleteAccepted() *PcloudPvminstancesVirtualserialnumberDeleteAccepted {
	return &PcloudPvminstancesVirtualserialnumberDeleteAccepted{}
}

/*
PcloudPvminstancesVirtualserialnumberDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudPvminstancesVirtualserialnumberDeleteAccepted struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud pvminstances virtualserialnumber delete accepted response has a 2xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances virtualserialnumber delete accepted response has a 3xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances virtualserialnumber delete accepted response has a 4xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances virtualserialnumber delete accepted response has a 5xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances virtualserialnumber delete accepted response a status code equal to that given
func (o *PcloudPvminstancesVirtualserialnumberDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud pvminstances virtualserialnumber delete accepted response
func (o *PcloudPvminstancesVirtualserialnumberDeleteAccepted) Code() int {
	return 202
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/virtual-serial-number][%d] pcloudPvminstancesVirtualserialnumberDeleteAccepted %s", 202, payload)
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/virtual-serial-number][%d] pcloudPvminstancesVirtualserialnumberDeleteAccepted %s", 202, payload)
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteAccepted) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVirtualserialnumberDeleteBadRequest creates a PcloudPvminstancesVirtualserialnumberDeleteBadRequest with default headers values
func NewPcloudPvminstancesVirtualserialnumberDeleteBadRequest() *PcloudPvminstancesVirtualserialnumberDeleteBadRequest {
	return &PcloudPvminstancesVirtualserialnumberDeleteBadRequest{}
}

/*
PcloudPvminstancesVirtualserialnumberDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesVirtualserialnumberDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances virtualserialnumber delete bad request response has a 2xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances virtualserialnumber delete bad request response has a 3xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances virtualserialnumber delete bad request response has a 4xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances virtualserialnumber delete bad request response has a 5xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances virtualserialnumber delete bad request response a status code equal to that given
func (o *PcloudPvminstancesVirtualserialnumberDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud pvminstances virtualserialnumber delete bad request response
func (o *PcloudPvminstancesVirtualserialnumberDeleteBadRequest) Code() int {
	return 400
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/virtual-serial-number][%d] pcloudPvminstancesVirtualserialnumberDeleteBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/virtual-serial-number][%d] pcloudPvminstancesVirtualserialnumberDeleteBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVirtualserialnumberDeleteUnauthorized creates a PcloudPvminstancesVirtualserialnumberDeleteUnauthorized with default headers values
func NewPcloudPvminstancesVirtualserialnumberDeleteUnauthorized() *PcloudPvminstancesVirtualserialnumberDeleteUnauthorized {
	return &PcloudPvminstancesVirtualserialnumberDeleteUnauthorized{}
}

/*
PcloudPvminstancesVirtualserialnumberDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesVirtualserialnumberDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances virtualserialnumber delete unauthorized response has a 2xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances virtualserialnumber delete unauthorized response has a 3xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances virtualserialnumber delete unauthorized response has a 4xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances virtualserialnumber delete unauthorized response has a 5xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances virtualserialnumber delete unauthorized response a status code equal to that given
func (o *PcloudPvminstancesVirtualserialnumberDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud pvminstances virtualserialnumber delete unauthorized response
func (o *PcloudPvminstancesVirtualserialnumberDeleteUnauthorized) Code() int {
	return 401
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/virtual-serial-number][%d] pcloudPvminstancesVirtualserialnumberDeleteUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/virtual-serial-number][%d] pcloudPvminstancesVirtualserialnumberDeleteUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVirtualserialnumberDeleteForbidden creates a PcloudPvminstancesVirtualserialnumberDeleteForbidden with default headers values
func NewPcloudPvminstancesVirtualserialnumberDeleteForbidden() *PcloudPvminstancesVirtualserialnumberDeleteForbidden {
	return &PcloudPvminstancesVirtualserialnumberDeleteForbidden{}
}

/*
PcloudPvminstancesVirtualserialnumberDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesVirtualserialnumberDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances virtualserialnumber delete forbidden response has a 2xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances virtualserialnumber delete forbidden response has a 3xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances virtualserialnumber delete forbidden response has a 4xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances virtualserialnumber delete forbidden response has a 5xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances virtualserialnumber delete forbidden response a status code equal to that given
func (o *PcloudPvminstancesVirtualserialnumberDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud pvminstances virtualserialnumber delete forbidden response
func (o *PcloudPvminstancesVirtualserialnumberDeleteForbidden) Code() int {
	return 403
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/virtual-serial-number][%d] pcloudPvminstancesVirtualserialnumberDeleteForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/virtual-serial-number][%d] pcloudPvminstancesVirtualserialnumberDeleteForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVirtualserialnumberDeleteNotFound creates a PcloudPvminstancesVirtualserialnumberDeleteNotFound with default headers values
func NewPcloudPvminstancesVirtualserialnumberDeleteNotFound() *PcloudPvminstancesVirtualserialnumberDeleteNotFound {
	return &PcloudPvminstancesVirtualserialnumberDeleteNotFound{}
}

/*
PcloudPvminstancesVirtualserialnumberDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesVirtualserialnumberDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances virtualserialnumber delete not found response has a 2xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances virtualserialnumber delete not found response has a 3xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances virtualserialnumber delete not found response has a 4xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances virtualserialnumber delete not found response has a 5xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances virtualserialnumber delete not found response a status code equal to that given
func (o *PcloudPvminstancesVirtualserialnumberDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud pvminstances virtualserialnumber delete not found response
func (o *PcloudPvminstancesVirtualserialnumberDeleteNotFound) Code() int {
	return 404
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/virtual-serial-number][%d] pcloudPvminstancesVirtualserialnumberDeleteNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/virtual-serial-number][%d] pcloudPvminstancesVirtualserialnumberDeleteNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVirtualserialnumberDeleteGone creates a PcloudPvminstancesVirtualserialnumberDeleteGone with default headers values
func NewPcloudPvminstancesVirtualserialnumberDeleteGone() *PcloudPvminstancesVirtualserialnumberDeleteGone {
	return &PcloudPvminstancesVirtualserialnumberDeleteGone{}
}

/*
PcloudPvminstancesVirtualserialnumberDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudPvminstancesVirtualserialnumberDeleteGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances virtualserialnumber delete gone response has a 2xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances virtualserialnumber delete gone response has a 3xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances virtualserialnumber delete gone response has a 4xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances virtualserialnumber delete gone response has a 5xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteGone) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances virtualserialnumber delete gone response a status code equal to that given
func (o *PcloudPvminstancesVirtualserialnumberDeleteGone) IsCode(code int) bool {
	return code == 410
}

// Code gets the status code for the pcloud pvminstances virtualserialnumber delete gone response
func (o *PcloudPvminstancesVirtualserialnumberDeleteGone) Code() int {
	return 410
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteGone) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/virtual-serial-number][%d] pcloudPvminstancesVirtualserialnumberDeleteGone %s", 410, payload)
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteGone) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/virtual-serial-number][%d] pcloudPvminstancesVirtualserialnumberDeleteGone %s", 410, payload)
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVirtualserialnumberDeleteInternalServerError creates a PcloudPvminstancesVirtualserialnumberDeleteInternalServerError with default headers values
func NewPcloudPvminstancesVirtualserialnumberDeleteInternalServerError() *PcloudPvminstancesVirtualserialnumberDeleteInternalServerError {
	return &PcloudPvminstancesVirtualserialnumberDeleteInternalServerError{}
}

/*
PcloudPvminstancesVirtualserialnumberDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesVirtualserialnumberDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances virtualserialnumber delete internal server error response has a 2xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances virtualserialnumber delete internal server error response has a 3xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances virtualserialnumber delete internal server error response has a 4xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances virtualserialnumber delete internal server error response has a 5xx status code
func (o *PcloudPvminstancesVirtualserialnumberDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances virtualserialnumber delete internal server error response a status code equal to that given
func (o *PcloudPvminstancesVirtualserialnumberDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud pvminstances virtualserialnumber delete internal server error response
func (o *PcloudPvminstancesVirtualserialnumberDeleteInternalServerError) Code() int {
	return 500
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/virtual-serial-number][%d] pcloudPvminstancesVirtualserialnumberDeleteInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/virtual-serial-number][%d] pcloudPvminstancesVirtualserialnumberDeleteInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVirtualserialnumberDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
