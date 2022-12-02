/*
 * Ory Keto API
 *
 * Documentation for all of Ory Keto's REST APIs. gRPC is documented separately.
 *
 * API version: 1.0.0
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

type RelationshipApi interface {

	/*
	 * CheckOplSyntax Performs a syntax check request.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return RelationshipApiApiCheckOplSyntaxRequest
	 */
	CheckOplSyntax(ctx context.Context) RelationshipApiApiCheckOplSyntaxRequest

	/*
	 * CheckOplSyntaxExecute executes the request
	 * @return CheckOplSyntaxResult
	 */
	CheckOplSyntaxExecute(r RelationshipApiApiCheckOplSyntaxRequest) (*CheckOplSyntaxResult, *http.Response, error)

	/*
	 * CreateRelationship Creates a relationship
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return RelationshipApiApiCreateRelationshipRequest
	 */
	CreateRelationship(ctx context.Context) RelationshipApiApiCreateRelationshipRequest

	/*
	 * CreateRelationshipExecute executes the request
	 * @return Relationship
	 */
	CreateRelationshipExecute(r RelationshipApiApiCreateRelationshipRequest) (*Relationship, *http.Response, error)

	/*
	 * DeleteRelationships Deletes relationships based on relation query
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return RelationshipApiApiDeleteRelationshipsRequest
	 */
	DeleteRelationships(ctx context.Context) RelationshipApiApiDeleteRelationshipsRequest

	/*
	 * DeleteRelationshipsExecute executes the request
	 */
	DeleteRelationshipsExecute(r RelationshipApiApiDeleteRelationshipsRequest) (*http.Response, error)

	/*
	 * GetRelationships Lists ACL relationships.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return RelationshipApiApiGetRelationshipsRequest
	 */
	GetRelationships(ctx context.Context) RelationshipApiApiGetRelationshipsRequest

	/*
	 * GetRelationshipsExecute executes the request
	 * @return Relationships
	 */
	GetRelationshipsExecute(r RelationshipApiApiGetRelationshipsRequest) (*Relationships, *http.Response, error)

	/*
	 * ListRelationshipNamespaces Lists Namespaces
	 * Get all namespaces.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return RelationshipApiApiListRelationshipNamespacesRequest
	 */
	ListRelationshipNamespaces(ctx context.Context) RelationshipApiApiListRelationshipNamespacesRequest

	/*
	 * ListRelationshipNamespacesExecute executes the request
	 * @return RelationshipNamespaces
	 */
	ListRelationshipNamespacesExecute(r RelationshipApiApiListRelationshipNamespacesRequest) (*RelationshipNamespaces, *http.Response, error)

	/*
	 * PatchRelationships Writes one or more relationships in a single transaction.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return RelationshipApiApiPatchRelationshipsRequest
	 */
	PatchRelationships(ctx context.Context) RelationshipApiApiPatchRelationshipsRequest

	/*
	 * PatchRelationshipsExecute executes the request
	 */
	PatchRelationshipsExecute(r RelationshipApiApiPatchRelationshipsRequest) (*http.Response, error)
}

// RelationshipApiService RelationshipApi service
type RelationshipApiService service

type RelationshipApiApiCheckOplSyntaxRequest struct {
	ctx        context.Context
	ApiService RelationshipApi
	body       *string
}

func (r RelationshipApiApiCheckOplSyntaxRequest) Body(body string) RelationshipApiApiCheckOplSyntaxRequest {
	r.body = &body
	return r
}

func (r RelationshipApiApiCheckOplSyntaxRequest) Execute() (*CheckOplSyntaxResult, *http.Response, error) {
	return r.ApiService.CheckOplSyntaxExecute(r)
}

/*
 * CheckOplSyntax Performs a syntax check request.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return RelationshipApiApiCheckOplSyntaxRequest
 */
func (a *RelationshipApiService) CheckOplSyntax(ctx context.Context) RelationshipApiApiCheckOplSyntaxRequest {
	return RelationshipApiApiCheckOplSyntaxRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return CheckOplSyntaxResult
 */
func (a *RelationshipApiService) CheckOplSyntaxExecute(r RelationshipApiApiCheckOplSyntaxRequest) (*CheckOplSyntaxResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *CheckOplSyntaxResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipApiService.CheckOplSyntax")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/opl/syntax/check"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ErrorGeneric
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RelationshipApiApiCreateRelationshipRequest struct {
	ctx                    context.Context
	ApiService             RelationshipApi
	createRelationshipBody *CreateRelationshipBody
}

func (r RelationshipApiApiCreateRelationshipRequest) CreateRelationshipBody(createRelationshipBody CreateRelationshipBody) RelationshipApiApiCreateRelationshipRequest {
	r.createRelationshipBody = &createRelationshipBody
	return r
}

func (r RelationshipApiApiCreateRelationshipRequest) Execute() (*Relationship, *http.Response, error) {
	return r.ApiService.CreateRelationshipExecute(r)
}

/*
 * CreateRelationship Creates a relationship
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return RelationshipApiApiCreateRelationshipRequest
 */
func (a *RelationshipApiService) CreateRelationship(ctx context.Context) RelationshipApiApiCreateRelationshipRequest {
	return RelationshipApiApiCreateRelationshipRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Relationship
 */
func (a *RelationshipApiService) CreateRelationshipExecute(r RelationshipApiApiCreateRelationshipRequest) (*Relationship, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *Relationship
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipApiService.CreateRelationship")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/relation-tuples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createRelationshipBody == nil {
		return localVarReturnValue, nil, reportError("createRelationshipBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createRelationshipBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ErrorGeneric
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RelationshipApiApiDeleteRelationshipsRequest struct {
	ctx                 context.Context
	ApiService          RelationshipApi
	namespace           *string
	object              *string
	relation            *string
	subjectId           *string
	subjectSetNamespace *string
	subjectSetObject    *string
	subjectSetRelation  *string
}

func (r RelationshipApiApiDeleteRelationshipsRequest) Namespace(namespace string) RelationshipApiApiDeleteRelationshipsRequest {
	r.namespace = &namespace
	return r
}
func (r RelationshipApiApiDeleteRelationshipsRequest) Object(object string) RelationshipApiApiDeleteRelationshipsRequest {
	r.object = &object
	return r
}
func (r RelationshipApiApiDeleteRelationshipsRequest) Relation(relation string) RelationshipApiApiDeleteRelationshipsRequest {
	r.relation = &relation
	return r
}
func (r RelationshipApiApiDeleteRelationshipsRequest) SubjectId(subjectId string) RelationshipApiApiDeleteRelationshipsRequest {
	r.subjectId = &subjectId
	return r
}
func (r RelationshipApiApiDeleteRelationshipsRequest) SubjectSetNamespace(subjectSetNamespace string) RelationshipApiApiDeleteRelationshipsRequest {
	r.subjectSetNamespace = &subjectSetNamespace
	return r
}
func (r RelationshipApiApiDeleteRelationshipsRequest) SubjectSetObject(subjectSetObject string) RelationshipApiApiDeleteRelationshipsRequest {
	r.subjectSetObject = &subjectSetObject
	return r
}
func (r RelationshipApiApiDeleteRelationshipsRequest) SubjectSetRelation(subjectSetRelation string) RelationshipApiApiDeleteRelationshipsRequest {
	r.subjectSetRelation = &subjectSetRelation
	return r
}

func (r RelationshipApiApiDeleteRelationshipsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRelationshipsExecute(r)
}

/*
 * DeleteRelationships Deletes relationships based on relation query
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return RelationshipApiApiDeleteRelationshipsRequest
 */
func (a *RelationshipApiService) DeleteRelationships(ctx context.Context) RelationshipApiApiDeleteRelationshipsRequest {
	return RelationshipApiApiDeleteRelationshipsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *RelationshipApiService) DeleteRelationshipsExecute(r RelationshipApiApiDeleteRelationshipsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipApiService.DeleteRelationships")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/relation-tuples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.object != nil {
		localVarQueryParams.Add("object", parameterToString(*r.object, ""))
	}
	if r.relation != nil {
		localVarQueryParams.Add("relation", parameterToString(*r.relation, ""))
	}
	if r.subjectId != nil {
		localVarQueryParams.Add("subject_id", parameterToString(*r.subjectId, ""))
	}
	if r.subjectSetNamespace != nil {
		localVarQueryParams.Add("subject_set.namespace", parameterToString(*r.subjectSetNamespace, ""))
	}
	if r.subjectSetObject != nil {
		localVarQueryParams.Add("subject_set.object", parameterToString(*r.subjectSetObject, ""))
	}
	if r.subjectSetRelation != nil {
		localVarQueryParams.Add("subject_set.relation", parameterToString(*r.subjectSetRelation, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ErrorGeneric
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type RelationshipApiApiGetRelationshipsRequest struct {
	ctx                 context.Context
	ApiService          RelationshipApi
	pageSize            *int32
	pageToken           *string
	namespace           *string
	object              *string
	relation            *string
	subjectId           *string
	subjectSetNamespace *string
	subjectSetObject    *string
	subjectSetRelation  *string
}

func (r RelationshipApiApiGetRelationshipsRequest) PageSize(pageSize int32) RelationshipApiApiGetRelationshipsRequest {
	r.pageSize = &pageSize
	return r
}
func (r RelationshipApiApiGetRelationshipsRequest) PageToken(pageToken string) RelationshipApiApiGetRelationshipsRequest {
	r.pageToken = &pageToken
	return r
}
func (r RelationshipApiApiGetRelationshipsRequest) Namespace(namespace string) RelationshipApiApiGetRelationshipsRequest {
	r.namespace = &namespace
	return r
}
func (r RelationshipApiApiGetRelationshipsRequest) Object(object string) RelationshipApiApiGetRelationshipsRequest {
	r.object = &object
	return r
}
func (r RelationshipApiApiGetRelationshipsRequest) Relation(relation string) RelationshipApiApiGetRelationshipsRequest {
	r.relation = &relation
	return r
}
func (r RelationshipApiApiGetRelationshipsRequest) SubjectId(subjectId string) RelationshipApiApiGetRelationshipsRequest {
	r.subjectId = &subjectId
	return r
}
func (r RelationshipApiApiGetRelationshipsRequest) SubjectSetNamespace(subjectSetNamespace string) RelationshipApiApiGetRelationshipsRequest {
	r.subjectSetNamespace = &subjectSetNamespace
	return r
}
func (r RelationshipApiApiGetRelationshipsRequest) SubjectSetObject(subjectSetObject string) RelationshipApiApiGetRelationshipsRequest {
	r.subjectSetObject = &subjectSetObject
	return r
}
func (r RelationshipApiApiGetRelationshipsRequest) SubjectSetRelation(subjectSetRelation string) RelationshipApiApiGetRelationshipsRequest {
	r.subjectSetRelation = &subjectSetRelation
	return r
}

func (r RelationshipApiApiGetRelationshipsRequest) Execute() (*Relationships, *http.Response, error) {
	return r.ApiService.GetRelationshipsExecute(r)
}

/*
 * GetRelationships Lists ACL relationships.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return RelationshipApiApiGetRelationshipsRequest
 */
func (a *RelationshipApiService) GetRelationships(ctx context.Context) RelationshipApiApiGetRelationshipsRequest {
	return RelationshipApiApiGetRelationshipsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Relationships
 */
func (a *RelationshipApiService) GetRelationshipsExecute(r RelationshipApiApiGetRelationshipsRequest) (*Relationships, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *Relationships
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipApiService.GetRelationships")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.pageToken != nil {
		localVarQueryParams.Add("page_token", parameterToString(*r.pageToken, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.object != nil {
		localVarQueryParams.Add("object", parameterToString(*r.object, ""))
	}
	if r.relation != nil {
		localVarQueryParams.Add("relation", parameterToString(*r.relation, ""))
	}
	if r.subjectId != nil {
		localVarQueryParams.Add("subject_id", parameterToString(*r.subjectId, ""))
	}
	if r.subjectSetNamespace != nil {
		localVarQueryParams.Add("subject_set.namespace", parameterToString(*r.subjectSetNamespace, ""))
	}
	if r.subjectSetObject != nil {
		localVarQueryParams.Add("subject_set.object", parameterToString(*r.subjectSetObject, ""))
	}
	if r.subjectSetRelation != nil {
		localVarQueryParams.Add("subject_set.relation", parameterToString(*r.subjectSetRelation, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ErrorGeneric
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RelationshipApiApiListRelationshipNamespacesRequest struct {
	ctx        context.Context
	ApiService RelationshipApi
}

func (r RelationshipApiApiListRelationshipNamespacesRequest) Execute() (*RelationshipNamespaces, *http.Response, error) {
	return r.ApiService.ListRelationshipNamespacesExecute(r)
}

/*
 * ListRelationshipNamespaces Lists Namespaces
 * Get all namespaces.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return RelationshipApiApiListRelationshipNamespacesRequest
 */
func (a *RelationshipApiService) ListRelationshipNamespaces(ctx context.Context) RelationshipApiApiListRelationshipNamespacesRequest {
	return RelationshipApiApiListRelationshipNamespacesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return RelationshipNamespaces
 */
func (a *RelationshipApiService) ListRelationshipNamespacesExecute(r RelationshipApiApiListRelationshipNamespacesRequest) (*RelationshipNamespaces, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *RelationshipNamespaces
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipApiService.ListRelationshipNamespaces")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/namespaces"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ErrorGeneric
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RelationshipApiApiPatchRelationshipsRequest struct {
	ctx               context.Context
	ApiService        RelationshipApi
	relationshipDelta *[]RelationshipDelta
}

func (r RelationshipApiApiPatchRelationshipsRequest) RelationshipDelta(relationshipDelta []RelationshipDelta) RelationshipApiApiPatchRelationshipsRequest {
	r.relationshipDelta = &relationshipDelta
	return r
}

func (r RelationshipApiApiPatchRelationshipsRequest) Execute() (*http.Response, error) {
	return r.ApiService.PatchRelationshipsExecute(r)
}

/*
 * PatchRelationships Writes one or more relationships in a single transaction.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return RelationshipApiApiPatchRelationshipsRequest
 */
func (a *RelationshipApiService) PatchRelationships(ctx context.Context) RelationshipApiApiPatchRelationshipsRequest {
	return RelationshipApiApiPatchRelationshipsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *RelationshipApiService) PatchRelationshipsExecute(r RelationshipApiApiPatchRelationshipsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipApiService.PatchRelationships")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/relation-tuples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.relationshipDelta == nil {
		return nil, reportError("relationshipDelta is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.relationshipDelta
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ErrorGeneric
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
