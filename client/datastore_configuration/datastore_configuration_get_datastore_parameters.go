// Code generated by go-swagger; DO NOT EDIT.

package datastore_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// GetByIDParams creates a new DatastoreConfigurationGetDatastoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func GetByIDParams() *DatastoreConfigurationGetDatastoreParams {
	return &DatastoreConfigurationGetDatastoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDatastoreConfigurationGetDatastoreParamsWithTimeout creates a new DatastoreConfigurationGetDatastoreParams object
// with the ability to set a timeout on a request.
func NewDatastoreConfigurationGetDatastoreParamsWithTimeout(timeout time.Duration) *DatastoreConfigurationGetDatastoreParams {
	return &DatastoreConfigurationGetDatastoreParams{
		timeout: timeout,
	}
}

// NewDatastoreConfigurationGetDatastoreParamsWithContext creates a new DatastoreConfigurationGetDatastoreParams object
// with the ability to set a context for a request.
func NewDatastoreConfigurationGetDatastoreParamsWithContext(ctx context.Context) *DatastoreConfigurationGetDatastoreParams {
	return &DatastoreConfigurationGetDatastoreParams{
		Context: ctx,
	}
}

// NewDatastoreConfigurationGetDatastoreParamsWithHTTPClient creates a new DatastoreConfigurationGetDatastoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewDatastoreConfigurationGetDatastoreParamsWithHTTPClient(client *http.Client) *DatastoreConfigurationGetDatastoreParams {
	return &DatastoreConfigurationGetDatastoreParams{
		HTTPClient: client,
	}
}

/*
DatastoreConfigurationGetDatastoreParams contains all the parameters to send to the API endpoint

	for the datastore configuration get datastore operation.

	Typically these are written to a http.Request.
*/
type DatastoreConfigurationGetDatastoreParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the datastore
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the datastore configuration get datastore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DatastoreConfigurationGetDatastoreParams) WithDefaults() *DatastoreConfigurationGetDatastoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the datastore configuration get datastore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DatastoreConfigurationGetDatastoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the datastore configuration get datastore params
func (o *DatastoreConfigurationGetDatastoreParams) WithTimeout(timeout time.Duration) *DatastoreConfigurationGetDatastoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the datastore configuration get datastore params
func (o *DatastoreConfigurationGetDatastoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the datastore configuration get datastore params
func (o *DatastoreConfigurationGetDatastoreParams) WithContext(ctx context.Context) *DatastoreConfigurationGetDatastoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the datastore configuration get datastore params
func (o *DatastoreConfigurationGetDatastoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the datastore configuration get datastore params
func (o *DatastoreConfigurationGetDatastoreParams) WithHTTPClient(client *http.Client) *DatastoreConfigurationGetDatastoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the datastore configuration get datastore params
func (o *DatastoreConfigurationGetDatastoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the datastore configuration get datastore params
func (o *DatastoreConfigurationGetDatastoreParams) WithXRequestID(xRequestID *string) *DatastoreConfigurationGetDatastoreParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the datastore configuration get datastore params
func (o *DatastoreConfigurationGetDatastoreParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the datastore configuration get datastore params
func (o *DatastoreConfigurationGetDatastoreParams) WithID(id string) *DatastoreConfigurationGetDatastoreParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the datastore configuration get datastore params
func (o *DatastoreConfigurationGetDatastoreParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DatastoreConfigurationGetDatastoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
