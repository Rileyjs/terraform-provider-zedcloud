package image

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

	"github.com/zededa/terraform-provider/models"
)

// NewImageConfigurationMarkEveImageLatestParams creates a new ImageConfigurationMarkEveImageLatestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageConfigurationMarkEveImageLatestParams() *ImageConfigurationMarkEveImageLatestParams {
	return &ImageConfigurationMarkEveImageLatestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageConfigurationMarkEveImageLatestParamsWithTimeout creates a new ImageConfigurationMarkEveImageLatestParams object
// with the ability to set a timeout on a request.
func NewImageConfigurationMarkEveImageLatestParamsWithTimeout(timeout time.Duration) *ImageConfigurationMarkEveImageLatestParams {
	return &ImageConfigurationMarkEveImageLatestParams{
		timeout: timeout,
	}
}

// NewImageConfigurationMarkEveImageLatestParamsWithContext creates a new ImageConfigurationMarkEveImageLatestParams object
// with the ability to set a context for a request.
func NewImageConfigurationMarkEveImageLatestParamsWithContext(ctx context.Context) *ImageConfigurationMarkEveImageLatestParams {
	return &ImageConfigurationMarkEveImageLatestParams{
		Context: ctx,
	}
}

// NewImageConfigurationMarkEveImageLatestParamsWithHTTPClient creates a new ImageConfigurationMarkEveImageLatestParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageConfigurationMarkEveImageLatestParamsWithHTTPClient(client *http.Client) *ImageConfigurationMarkEveImageLatestParams {
	return &ImageConfigurationMarkEveImageLatestParams{
		HTTPClient: client,
	}
}

/*
ImageConfigurationMarkEveImageLatestParams contains all the parameters to send to the API endpoint

	for the image configuration mark eve image latest operation.

	Typically these are written to a http.Request.
*/
type ImageConfigurationMarkEveImageLatestParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.Image

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image configuration mark eve image latest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationMarkEveImageLatestParams) WithDefaults() *ImageConfigurationMarkEveImageLatestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image configuration mark eve image latest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationMarkEveImageLatestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image configuration mark eve image latest params
func (o *ImageConfigurationMarkEveImageLatestParams) WithTimeout(timeout time.Duration) *ImageConfigurationMarkEveImageLatestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image configuration mark eve image latest params
func (o *ImageConfigurationMarkEveImageLatestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image configuration mark eve image latest params
func (o *ImageConfigurationMarkEveImageLatestParams) WithContext(ctx context.Context) *ImageConfigurationMarkEveImageLatestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image configuration mark eve image latest params
func (o *ImageConfigurationMarkEveImageLatestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image configuration mark eve image latest params
func (o *ImageConfigurationMarkEveImageLatestParams) WithHTTPClient(client *http.Client) *ImageConfigurationMarkEveImageLatestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image configuration mark eve image latest params
func (o *ImageConfigurationMarkEveImageLatestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the image configuration mark eve image latest params
func (o *ImageConfigurationMarkEveImageLatestParams) WithXRequestID(xRequestID *string) *ImageConfigurationMarkEveImageLatestParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the image configuration mark eve image latest params
func (o *ImageConfigurationMarkEveImageLatestParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the image configuration mark eve image latest params
func (o *ImageConfigurationMarkEveImageLatestParams) WithBody(body *models.Image) *ImageConfigurationMarkEveImageLatestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the image configuration mark eve image latest params
func (o *ImageConfigurationMarkEveImageLatestParams) SetBody(body *models.Image) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ImageConfigurationMarkEveImageLatestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}