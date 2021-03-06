package httpinfrastructuregroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
// 
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
)

// HTTPSuccessClient is the test Infrastructure for AutoRest
type HTTPSuccessClient struct {
    ManagementClient
}
// NewHTTPSuccessClient creates an instance of the HTTPSuccessClient client.
func NewHTTPSuccessClient() HTTPSuccessClient {
    return NewHTTPSuccessClientWithBaseURI(DefaultBaseURI, )
}

// NewHTTPSuccessClientWithBaseURI creates an instance of the
// HTTPSuccessClient client.
func NewHTTPSuccessClientWithBaseURI(baseURI string, ) HTTPSuccessClient {
    return HTTPSuccessClient{NewWithBaseURI(baseURI, )}
}

// Delete200 delete simple boolean value true returns 200
//
// booleanValue is simple boolean value true
func (client HTTPSuccessClient) Delete200(booleanValue *bool) (result autorest.Response, err error) {
    req, err := client.Delete200Preparer(booleanValue)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Delete200", nil , "Failure preparing request")
    }

    resp, err := client.Delete200Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Delete200", resp, "Failure sending request")
    }

    result, err = client.Delete200Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Delete200", resp, "Failure responding to request")
    }

    return
}

// Delete200Preparer prepares the Delete200 request.
func (client HTTPSuccessClient) Delete200Preparer(booleanValue *bool) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsDelete(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/200"))
    if booleanValue != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithJSON(booleanValue))
    }
    return preparer.Prepare(&http.Request{})
}

// Delete200Sender sends the Delete200 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Delete200Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Delete200Responder handles the response to the Delete200 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Delete200Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Delete202 delete true Boolean value in request returns 202 (accepted)
//
// booleanValue is simple boolean value true
func (client HTTPSuccessClient) Delete202(booleanValue *bool) (result autorest.Response, err error) {
    req, err := client.Delete202Preparer(booleanValue)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Delete202", nil , "Failure preparing request")
    }

    resp, err := client.Delete202Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Delete202", resp, "Failure sending request")
    }

    result, err = client.Delete202Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Delete202", resp, "Failure responding to request")
    }

    return
}

// Delete202Preparer prepares the Delete202 request.
func (client HTTPSuccessClient) Delete202Preparer(booleanValue *bool) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsDelete(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/202"))
    if booleanValue != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithJSON(booleanValue))
    }
    return preparer.Prepare(&http.Request{})
}

// Delete202Sender sends the Delete202 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Delete202Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Delete202Responder handles the response to the Delete202 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Delete202Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Delete204 delete true Boolean value in request returns 204 (no content)
//
// booleanValue is simple boolean value true
func (client HTTPSuccessClient) Delete204(booleanValue *bool) (result autorest.Response, err error) {
    req, err := client.Delete204Preparer(booleanValue)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Delete204", nil , "Failure preparing request")
    }

    resp, err := client.Delete204Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Delete204", resp, "Failure sending request")
    }

    result, err = client.Delete204Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Delete204", resp, "Failure responding to request")
    }

    return
}

// Delete204Preparer prepares the Delete204 request.
func (client HTTPSuccessClient) Delete204Preparer(booleanValue *bool) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsDelete(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/204"))
    if booleanValue != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithJSON(booleanValue))
    }
    return preparer.Prepare(&http.Request{})
}

// Delete204Sender sends the Delete204 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Delete204Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Delete204Responder handles the response to the Delete204 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Delete204Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Get200 get 200 success
func (client HTTPSuccessClient) Get200() (result Bool, err error) {
    req, err := client.Get200Preparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Get200", nil , "Failure preparing request")
    }

    resp, err := client.Get200Sender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Get200", resp, "Failure sending request")
    }

    result, err = client.Get200Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Get200", resp, "Failure responding to request")
    }

    return
}

// Get200Preparer prepares the Get200 request.
func (client HTTPSuccessClient) Get200Preparer() (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/200"))
    return preparer.Prepare(&http.Request{})
}

// Get200Sender sends the Get200 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Get200Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Get200Responder handles the response to the Get200 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Get200Responder(resp *http.Response) (result Bool, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result.Value),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// Head200 return 200 status code if successful
func (client HTTPSuccessClient) Head200() (result autorest.Response, err error) {
    req, err := client.Head200Preparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Head200", nil , "Failure preparing request")
    }

    resp, err := client.Head200Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Head200", resp, "Failure sending request")
    }

    result, err = client.Head200Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Head200", resp, "Failure responding to request")
    }

    return
}

// Head200Preparer prepares the Head200 request.
func (client HTTPSuccessClient) Head200Preparer() (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsHead(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/200"))
    return preparer.Prepare(&http.Request{})
}

// Head200Sender sends the Head200 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Head200Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Head200Responder handles the response to the Head200 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Head200Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Head204 return 204 status code if successful
func (client HTTPSuccessClient) Head204() (result autorest.Response, err error) {
    req, err := client.Head204Preparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Head204", nil , "Failure preparing request")
    }

    resp, err := client.Head204Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Head204", resp, "Failure sending request")
    }

    result, err = client.Head204Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Head204", resp, "Failure responding to request")
    }

    return
}

// Head204Preparer prepares the Head204 request.
func (client HTTPSuccessClient) Head204Preparer() (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsHead(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/204"))
    return preparer.Prepare(&http.Request{})
}

// Head204Sender sends the Head204 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Head204Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Head204Responder handles the response to the Head204 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Head204Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Head404 return 404 status code
func (client HTTPSuccessClient) Head404() (result autorest.Response, err error) {
    req, err := client.Head404Preparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Head404", nil , "Failure preparing request")
    }

    resp, err := client.Head404Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Head404", resp, "Failure sending request")
    }

    result, err = client.Head404Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Head404", resp, "Failure responding to request")
    }

    return
}

// Head404Preparer prepares the Head404 request.
func (client HTTPSuccessClient) Head404Preparer() (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsHead(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/404"))
    return preparer.Prepare(&http.Request{})
}

// Head404Sender sends the Head404 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Head404Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Head404Responder handles the response to the Head404 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Head404Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent,http.StatusNotFound),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Patch200 patch true Boolean value in request returning 200
//
// booleanValue is simple boolean value true
func (client HTTPSuccessClient) Patch200(booleanValue *bool) (result autorest.Response, err error) {
    req, err := client.Patch200Preparer(booleanValue)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Patch200", nil , "Failure preparing request")
    }

    resp, err := client.Patch200Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Patch200", resp, "Failure sending request")
    }

    result, err = client.Patch200Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Patch200", resp, "Failure responding to request")
    }

    return
}

// Patch200Preparer prepares the Patch200 request.
func (client HTTPSuccessClient) Patch200Preparer(booleanValue *bool) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPatch(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/200"))
    if booleanValue != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithJSON(booleanValue))
    }
    return preparer.Prepare(&http.Request{})
}

// Patch200Sender sends the Patch200 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Patch200Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Patch200Responder handles the response to the Patch200 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Patch200Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Patch202 patch true Boolean value in request returns 202
//
// booleanValue is simple boolean value true
func (client HTTPSuccessClient) Patch202(booleanValue *bool) (result autorest.Response, err error) {
    req, err := client.Patch202Preparer(booleanValue)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Patch202", nil , "Failure preparing request")
    }

    resp, err := client.Patch202Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Patch202", resp, "Failure sending request")
    }

    result, err = client.Patch202Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Patch202", resp, "Failure responding to request")
    }

    return
}

// Patch202Preparer prepares the Patch202 request.
func (client HTTPSuccessClient) Patch202Preparer(booleanValue *bool) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPatch(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/202"))
    if booleanValue != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithJSON(booleanValue))
    }
    return preparer.Prepare(&http.Request{})
}

// Patch202Sender sends the Patch202 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Patch202Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Patch202Responder handles the response to the Patch202 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Patch202Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Patch204 patch true Boolean value in request returns 204 (no content)
//
// booleanValue is simple boolean value true
func (client HTTPSuccessClient) Patch204(booleanValue *bool) (result autorest.Response, err error) {
    req, err := client.Patch204Preparer(booleanValue)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Patch204", nil , "Failure preparing request")
    }

    resp, err := client.Patch204Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Patch204", resp, "Failure sending request")
    }

    result, err = client.Patch204Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Patch204", resp, "Failure responding to request")
    }

    return
}

// Patch204Preparer prepares the Patch204 request.
func (client HTTPSuccessClient) Patch204Preparer(booleanValue *bool) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPatch(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/204"))
    if booleanValue != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithJSON(booleanValue))
    }
    return preparer.Prepare(&http.Request{})
}

// Patch204Sender sends the Patch204 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Patch204Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Patch204Responder handles the response to the Patch204 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Patch204Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Post200 post bollean value true in request that returns a 200
//
// booleanValue is simple boolean value true
func (client HTTPSuccessClient) Post200(booleanValue *bool) (result autorest.Response, err error) {
    req, err := client.Post200Preparer(booleanValue)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Post200", nil , "Failure preparing request")
    }

    resp, err := client.Post200Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Post200", resp, "Failure sending request")
    }

    result, err = client.Post200Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Post200", resp, "Failure responding to request")
    }

    return
}

// Post200Preparer prepares the Post200 request.
func (client HTTPSuccessClient) Post200Preparer(booleanValue *bool) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPost(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/200"))
    if booleanValue != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithJSON(booleanValue))
    }
    return preparer.Prepare(&http.Request{})
}

// Post200Sender sends the Post200 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Post200Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Post200Responder handles the response to the Post200 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Post200Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Post201 post true Boolean value in request returns 201 (Created)
//
// booleanValue is simple boolean value true
func (client HTTPSuccessClient) Post201(booleanValue *bool) (result autorest.Response, err error) {
    req, err := client.Post201Preparer(booleanValue)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Post201", nil , "Failure preparing request")
    }

    resp, err := client.Post201Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Post201", resp, "Failure sending request")
    }

    result, err = client.Post201Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Post201", resp, "Failure responding to request")
    }

    return
}

// Post201Preparer prepares the Post201 request.
func (client HTTPSuccessClient) Post201Preparer(booleanValue *bool) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPost(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/201"))
    if booleanValue != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithJSON(booleanValue))
    }
    return preparer.Prepare(&http.Request{})
}

// Post201Sender sends the Post201 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Post201Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Post201Responder handles the response to the Post201 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Post201Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Post202 post true Boolean value in request returns 202 (Accepted)
//
// booleanValue is simple boolean value true
func (client HTTPSuccessClient) Post202(booleanValue *bool) (result autorest.Response, err error) {
    req, err := client.Post202Preparer(booleanValue)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Post202", nil , "Failure preparing request")
    }

    resp, err := client.Post202Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Post202", resp, "Failure sending request")
    }

    result, err = client.Post202Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Post202", resp, "Failure responding to request")
    }

    return
}

// Post202Preparer prepares the Post202 request.
func (client HTTPSuccessClient) Post202Preparer(booleanValue *bool) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPost(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/202"))
    if booleanValue != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithJSON(booleanValue))
    }
    return preparer.Prepare(&http.Request{})
}

// Post202Sender sends the Post202 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Post202Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Post202Responder handles the response to the Post202 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Post202Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Post204 post true Boolean value in request returns 204 (no content)
//
// booleanValue is simple boolean value true
func (client HTTPSuccessClient) Post204(booleanValue *bool) (result autorest.Response, err error) {
    req, err := client.Post204Preparer(booleanValue)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Post204", nil , "Failure preparing request")
    }

    resp, err := client.Post204Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Post204", resp, "Failure sending request")
    }

    result, err = client.Post204Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Post204", resp, "Failure responding to request")
    }

    return
}

// Post204Preparer prepares the Post204 request.
func (client HTTPSuccessClient) Post204Preparer(booleanValue *bool) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPost(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/204"))
    if booleanValue != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithJSON(booleanValue))
    }
    return preparer.Prepare(&http.Request{})
}

// Post204Sender sends the Post204 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Post204Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Post204Responder handles the response to the Post204 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Post204Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Put200 put boolean value true returning 200 success
//
// booleanValue is simple boolean value true
func (client HTTPSuccessClient) Put200(booleanValue *bool) (result autorest.Response, err error) {
    req, err := client.Put200Preparer(booleanValue)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Put200", nil , "Failure preparing request")
    }

    resp, err := client.Put200Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Put200", resp, "Failure sending request")
    }

    result, err = client.Put200Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Put200", resp, "Failure responding to request")
    }

    return
}

// Put200Preparer prepares the Put200 request.
func (client HTTPSuccessClient) Put200Preparer(booleanValue *bool) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPut(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/200"))
    if booleanValue != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithJSON(booleanValue))
    }
    return preparer.Prepare(&http.Request{})
}

// Put200Sender sends the Put200 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Put200Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Put200Responder handles the response to the Put200 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Put200Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Put201 put true Boolean value in request returns 201
//
// booleanValue is simple boolean value true
func (client HTTPSuccessClient) Put201(booleanValue *bool) (result autorest.Response, err error) {
    req, err := client.Put201Preparer(booleanValue)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Put201", nil , "Failure preparing request")
    }

    resp, err := client.Put201Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Put201", resp, "Failure sending request")
    }

    result, err = client.Put201Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Put201", resp, "Failure responding to request")
    }

    return
}

// Put201Preparer prepares the Put201 request.
func (client HTTPSuccessClient) Put201Preparer(booleanValue *bool) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPut(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/201"))
    if booleanValue != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithJSON(booleanValue))
    }
    return preparer.Prepare(&http.Request{})
}

// Put201Sender sends the Put201 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Put201Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Put201Responder handles the response to the Put201 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Put201Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Put202 put true Boolean value in request returns 202 (Accepted)
//
// booleanValue is simple boolean value true
func (client HTTPSuccessClient) Put202(booleanValue *bool) (result autorest.Response, err error) {
    req, err := client.Put202Preparer(booleanValue)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Put202", nil , "Failure preparing request")
    }

    resp, err := client.Put202Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Put202", resp, "Failure sending request")
    }

    result, err = client.Put202Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Put202", resp, "Failure responding to request")
    }

    return
}

// Put202Preparer prepares the Put202 request.
func (client HTTPSuccessClient) Put202Preparer(booleanValue *bool) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPut(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/202"))
    if booleanValue != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithJSON(booleanValue))
    }
    return preparer.Prepare(&http.Request{})
}

// Put202Sender sends the Put202 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Put202Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Put202Responder handles the response to the Put202 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Put202Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Put204 put true Boolean value in request returns 204 (no content)
//
// booleanValue is simple boolean value true
func (client HTTPSuccessClient) Put204(booleanValue *bool) (result autorest.Response, err error) {
    req, err := client.Put204Preparer(booleanValue)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Put204", nil , "Failure preparing request")
    }

    resp, err := client.Put204Sender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Put204", resp, "Failure sending request")
    }

    result, err = client.Put204Responder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPSuccessClient", "Put204", resp, "Failure responding to request")
    }

    return
}

// Put204Preparer prepares the Put204 request.
func (client HTTPSuccessClient) Put204Preparer(booleanValue *bool) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPut(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/success/204"))
    if booleanValue != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithJSON(booleanValue))
    }
    return preparer.Prepare(&http.Request{})
}

// Put204Sender sends the Put204 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPSuccessClient) Put204Sender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// Put204Responder handles the response to the Put204 request. The method always
// closes the http.Response Body.
func (client HTTPSuccessClient) Put204Responder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
            autorest.ByClosing())
    result.Response = resp
    return
}

