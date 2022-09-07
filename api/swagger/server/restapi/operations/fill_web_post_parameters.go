// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// FillWebPostMaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var FillWebPostMaxParseMemory int64 = 32 << 20

// NewFillWebPostParams creates a new FillWebPostParams object
//
// There are no default values defined in the spec.
func NewFillWebPostParams() FillWebPostParams {

	return FillWebPostParams{}
}

// FillWebPostParams contains all the bound params for the fill web post operation
// typically these are obtained from a http.Request
//
// swagger:parameters fillWebPost
type FillWebPostParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Website's short name.
	  Required: true
	  In: path
	*/
	Website string
	/*Contents of the ZIP file.
	  Required: true
	  In: formData
	*/
	Zipfile io.ReadCloser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFillWebPostParams() beforehand.
func (o *FillWebPostParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(FillWebPostMaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}

	rWebsite, rhkWebsite, _ := route.Params.GetOK("website")
	if err := o.bindWebsite(rWebsite, rhkWebsite, route.Formats); err != nil {
		res = append(res, err)
	}

	zipfile, zipfileHeader, err := r.FormFile("zipfile")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "zipfile", err))
	} else if err := o.bindZipfile(zipfile, zipfileHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.Zipfile = &runtime.File{Data: zipfile, Header: zipfileHeader}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindWebsite binds and validates parameter Website from path.
func (o *FillWebPostParams) bindWebsite(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Website = raw

	return nil
}

// bindZipfile binds file parameter Zipfile.
//
// The only supported validations on files are MinLength and MaxLength
func (o *FillWebPostParams) bindZipfile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}