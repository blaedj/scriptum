// Code generated by protoc-gen-twirp v5.2.0, DO NOT EDIT.
// source: rpc/service.proto

/*
Package scriptumproto is a generated twirp stub package.
This code was generated with github.com/twitchtv/twirp/protoc-gen-twirp v5.2.0.

It is generated from these files:
	rpc/service.proto
*/
package scriptumproto

import bytes "bytes"
import strings "strings"
import context "context"
import fmt "fmt"
import ioutil "io/ioutil"
import log "log"
import http "net/http"

import jsonpb "github.com/golang/protobuf/jsonpb"
import proto "github.com/golang/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// Imports only used by utility functions:
import io "io"
import strconv "strconv"
import json "encoding/json"
import url "net/url"

// =========================
// ScriptumService Interface
// =========================

type ScriptumService interface {
	NewDocument(context.Context, *Document) (*NewDocumentResponse, error)

	SaveDocument(context.Context, *Document) (*SaveDocumentResponse, error)

	DeleteDocument(context.Context, *DeleteDocumentRequest) (*DeleteDocumentResponse, error)
}

// ===============================
// ScriptumService Protobuf Client
// ===============================

type scriptumServiceProtobufClient struct {
	client HTTPClient
	urls   [3]string
}

// NewScriptumServiceProtobufClient creates a Protobuf client that implements the ScriptumService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewScriptumServiceProtobufClient(addr string, client HTTPClient) ScriptumService {
	prefix := urlBase(addr) + ScriptumServicePathPrefix
	urls := [3]string{
		prefix + "NewDocument",
		prefix + "SaveDocument",
		prefix + "DeleteDocument",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &scriptumServiceProtobufClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &scriptumServiceProtobufClient{
		client: client,
		urls:   urls,
	}
}

func (c *scriptumServiceProtobufClient) NewDocument(ctx context.Context, in *Document) (*NewDocumentResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "blaedj.scriptum.proto")
	ctx = ctxsetters.WithServiceName(ctx, "ScriptumService")
	ctx = ctxsetters.WithMethodName(ctx, "NewDocument")
	out := new(NewDocumentResponse)
	err := doProtobufRequest(ctx, c.client, c.urls[0], in, out)
	return out, err
}

func (c *scriptumServiceProtobufClient) SaveDocument(ctx context.Context, in *Document) (*SaveDocumentResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "blaedj.scriptum.proto")
	ctx = ctxsetters.WithServiceName(ctx, "ScriptumService")
	ctx = ctxsetters.WithMethodName(ctx, "SaveDocument")
	out := new(SaveDocumentResponse)
	err := doProtobufRequest(ctx, c.client, c.urls[1], in, out)
	return out, err
}

func (c *scriptumServiceProtobufClient) DeleteDocument(ctx context.Context, in *DeleteDocumentRequest) (*DeleteDocumentResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "blaedj.scriptum.proto")
	ctx = ctxsetters.WithServiceName(ctx, "ScriptumService")
	ctx = ctxsetters.WithMethodName(ctx, "DeleteDocument")
	out := new(DeleteDocumentResponse)
	err := doProtobufRequest(ctx, c.client, c.urls[2], in, out)
	return out, err
}

// ===========================
// ScriptumService JSON Client
// ===========================

type scriptumServiceJSONClient struct {
	client HTTPClient
	urls   [3]string
}

// NewScriptumServiceJSONClient creates a JSON client that implements the ScriptumService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewScriptumServiceJSONClient(addr string, client HTTPClient) ScriptumService {
	prefix := urlBase(addr) + ScriptumServicePathPrefix
	urls := [3]string{
		prefix + "NewDocument",
		prefix + "SaveDocument",
		prefix + "DeleteDocument",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &scriptumServiceJSONClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &scriptumServiceJSONClient{
		client: client,
		urls:   urls,
	}
}

func (c *scriptumServiceJSONClient) NewDocument(ctx context.Context, in *Document) (*NewDocumentResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "blaedj.scriptum.proto")
	ctx = ctxsetters.WithServiceName(ctx, "ScriptumService")
	ctx = ctxsetters.WithMethodName(ctx, "NewDocument")
	out := new(NewDocumentResponse)
	err := doJSONRequest(ctx, c.client, c.urls[0], in, out)
	return out, err
}

func (c *scriptumServiceJSONClient) SaveDocument(ctx context.Context, in *Document) (*SaveDocumentResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "blaedj.scriptum.proto")
	ctx = ctxsetters.WithServiceName(ctx, "ScriptumService")
	ctx = ctxsetters.WithMethodName(ctx, "SaveDocument")
	out := new(SaveDocumentResponse)
	err := doJSONRequest(ctx, c.client, c.urls[1], in, out)
	return out, err
}

func (c *scriptumServiceJSONClient) DeleteDocument(ctx context.Context, in *DeleteDocumentRequest) (*DeleteDocumentResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "blaedj.scriptum.proto")
	ctx = ctxsetters.WithServiceName(ctx, "ScriptumService")
	ctx = ctxsetters.WithMethodName(ctx, "DeleteDocument")
	out := new(DeleteDocumentResponse)
	err := doJSONRequest(ctx, c.client, c.urls[2], in, out)
	return out, err
}

// ==============================
// ScriptumService Server Handler
// ==============================

type scriptumServiceServer struct {
	ScriptumService
	hooks *twirp.ServerHooks
}

func NewScriptumServiceServer(svc ScriptumService, hooks *twirp.ServerHooks) TwirpServer {
	return &scriptumServiceServer{
		ScriptumService: svc,
		hooks:           hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *scriptumServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// ScriptumServicePathPrefix is used for all URL paths on a twirp ScriptumService server.
// Requests are always: POST ScriptumServicePathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const ScriptumServicePathPrefix = "/twirp/blaedj.scriptum.proto.ScriptumService/"

func (s *scriptumServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "blaedj.scriptum.proto")
	ctx = ctxsetters.WithServiceName(ctx, "ScriptumService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}

	switch req.URL.Path {
	case "/twirp/blaedj.scriptum.proto.ScriptumService/NewDocument":
		s.serveNewDocument(ctx, resp, req)
		return
	case "/twirp/blaedj.scriptum.proto.ScriptumService/SaveDocument":
		s.serveSaveDocument(ctx, resp, req)
		return
	case "/twirp/blaedj.scriptum.proto.ScriptumService/DeleteDocument":
		s.serveDeleteDocument(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *scriptumServiceServer) serveNewDocument(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveNewDocumentJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveNewDocumentProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *scriptumServiceServer) serveNewDocumentJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "NewDocument")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	defer closebody(req.Body)
	reqContent := new(Document)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request json")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *NewDocumentResponse
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.NewDocument(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *NewDocumentResponse and nil error while calling NewDocument. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		err = wrapErr(err, "failed to marshal json response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	if _, err = resp.Write(buf.Bytes()); err != nil {
		log.Printf("errored while writing response to client, but already sent response status code to 200: %s", err)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *scriptumServiceServer) serveNewDocumentProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "NewDocument")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	defer closebody(req.Body)
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		err = wrapErr(err, "failed to read request body")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}
	reqContent := new(Document)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request proto")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *NewDocumentResponse
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.NewDocument(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *NewDocumentResponse and nil error while calling NewDocument. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		err = wrapErr(err, "failed to marshal proto response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.WriteHeader(http.StatusOK)
	if _, err = resp.Write(respBytes); err != nil {
		log.Printf("errored while writing response to client, but already sent response status code to 200: %s", err)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *scriptumServiceServer) serveSaveDocument(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveSaveDocumentJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveSaveDocumentProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *scriptumServiceServer) serveSaveDocumentJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "SaveDocument")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	defer closebody(req.Body)
	reqContent := new(Document)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request json")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *SaveDocumentResponse
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.SaveDocument(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *SaveDocumentResponse and nil error while calling SaveDocument. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		err = wrapErr(err, "failed to marshal json response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	if _, err = resp.Write(buf.Bytes()); err != nil {
		log.Printf("errored while writing response to client, but already sent response status code to 200: %s", err)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *scriptumServiceServer) serveSaveDocumentProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "SaveDocument")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	defer closebody(req.Body)
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		err = wrapErr(err, "failed to read request body")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}
	reqContent := new(Document)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request proto")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *SaveDocumentResponse
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.SaveDocument(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *SaveDocumentResponse and nil error while calling SaveDocument. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		err = wrapErr(err, "failed to marshal proto response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.WriteHeader(http.StatusOK)
	if _, err = resp.Write(respBytes); err != nil {
		log.Printf("errored while writing response to client, but already sent response status code to 200: %s", err)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *scriptumServiceServer) serveDeleteDocument(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveDeleteDocumentJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveDeleteDocumentProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *scriptumServiceServer) serveDeleteDocumentJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "DeleteDocument")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	defer closebody(req.Body)
	reqContent := new(DeleteDocumentRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request json")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *DeleteDocumentResponse
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.DeleteDocument(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *DeleteDocumentResponse and nil error while calling DeleteDocument. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		err = wrapErr(err, "failed to marshal json response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	if _, err = resp.Write(buf.Bytes()); err != nil {
		log.Printf("errored while writing response to client, but already sent response status code to 200: %s", err)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *scriptumServiceServer) serveDeleteDocumentProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "DeleteDocument")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	defer closebody(req.Body)
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		err = wrapErr(err, "failed to read request body")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}
	reqContent := new(DeleteDocumentRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request proto")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *DeleteDocumentResponse
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.DeleteDocument(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *DeleteDocumentResponse and nil error while calling DeleteDocument. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		err = wrapErr(err, "failed to marshal proto response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.WriteHeader(http.StatusOK)
	if _, err = resp.Write(respBytes); err != nil {
		log.Printf("errored while writing response to client, but already sent response status code to 200: %s", err)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *scriptumServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor0, 0
}

func (s *scriptumServiceServer) ProtocGenTwirpVersion() string {
	return "v5.2.0"
}

// =====
// Utils
// =====

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
//
// HTTPClient implementations should not follow redirects. Redirects are
// automatically disabled if *(net/http).Client is passed to client
// constructors. See the withoutRedirects function in this file for more
// details.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// TwirpServer is the interface generated server structs will support: they're
// HTTP handlers with additional methods for accessing metadata about the
// service. Those accessors are a low-level API for building reflection tools.
// Most people can think of TwirpServers as just http.Handlers.
type TwirpServer interface {
	http.Handler
	// ServiceDescriptor returns gzipped bytes describing the .proto file that
	// this service was generated from. Once unzipped, the bytes can be
	// unmarshalled as a
	// github.com/golang/protobuf/protoc-gen-go/descriptor.FileDescriptorProto.
	//
	// The returned integer is the index of this particular service within that
	// FileDescriptorProto's 'Service' slice of ServiceDescriptorProtos. This is a
	// low-level field, expected to be used for reflection.
	ServiceDescriptor() ([]byte, int)
	// ProtocGenTwirpVersion is the semantic version string of the version of
	// twirp used to generate this file.
	ProtocGenTwirpVersion() string
}

// WriteError writes an HTTP response with a valid Twirp error format.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func WriteError(resp http.ResponseWriter, err error) {
	writeError(context.Background(), resp, err, nil)
}

// writeError writes Twirp errors in the response and triggers hooks.
func writeError(ctx context.Context, resp http.ResponseWriter, err error, hooks *twirp.ServerHooks) {
	// Non-twirp errors are wrapped as Internal (default)
	twerr, ok := err.(twirp.Error)
	if !ok {
		twerr = twirp.InternalErrorWith(err)
	}

	statusCode := twirp.ServerHTTPStatusFromErrorCode(twerr.Code())
	ctx = ctxsetters.WithStatusCode(ctx, statusCode)
	ctx = callError(ctx, hooks, twerr)

	resp.Header().Set("Content-Type", "application/json") // Error responses are always JSON (instead of protobuf)
	resp.WriteHeader(statusCode)                          // HTTP response status code

	respBody := marshalErrorToJSON(twerr)
	_, err2 := resp.Write(respBody)
	if err2 != nil {
		log.Printf("unable to send error message %q: %s", twerr, err2)
	}

	callResponseSent(ctx, hooks)
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// getCustomHTTPReqHeaders retrieves a copy of any headers that are set in
// a context through the twirp.WithHTTPRequestHeaders function.
// If there are no headers set, or if they have the wrong type, nil is returned.
func getCustomHTTPReqHeaders(ctx context.Context) http.Header {
	header, ok := twirp.HTTPRequestHeaders(ctx)
	if !ok || header == nil {
		return nil
	}
	copied := make(http.Header)
	for k, vv := range header {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}
	return copied
}

// closebody closes a response or request body and just logs
// any error encountered while closing, since errors are
// considered very unusual.
func closebody(body io.Closer) {
	if err := body.Close(); err != nil {
		log.Printf("error closing body: %q", err)
	}
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if customHeader := getCustomHTTPReqHeaders(ctx); customHeader != nil {
		req.Header = customHeader
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Twirp-Version", "v5.2.0")
	return req, nil
}

// JSON serialization for errors
type twerrJSON struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Meta map[string]string `json:"meta,omitempty"`
}

// marshalErrorToJSON returns JSON from a twirp.Error, that can be used as HTTP error response body.
// If serialization fails, it will use a descriptive Internal error instead.
func marshalErrorToJSON(twerr twirp.Error) []byte {
	// make sure that msg is not too large
	msg := twerr.Msg()
	if len(msg) > 1e6 {
		msg = msg[:1e6]
	}

	tj := twerrJSON{
		Code: string(twerr.Code()),
		Msg:  msg,
		Meta: twerr.MetaMap(),
	}

	buf, err := json.Marshal(&tj)
	if err != nil {
		buf = []byte("{\"type\": \"" + twirp.Internal + "\", \"msg\": \"There was an error but it could not be serialized into JSON\"}") // fallback
	}

	return buf
}

// errorFromResponse builds a twirp.Error from a non-200 HTTP response.
// If the response has a valid serialized Twirp error, then it's returned.
// If not, the response status code is used to generate a similar twirp
// error. See twirpErrorFromIntermediary for more info on intermediary errors.
func errorFromResponse(resp *http.Response) twirp.Error {
	statusCode := resp.StatusCode
	statusText := http.StatusText(statusCode)

	if isHTTPRedirect(statusCode) {
		// Unexpected redirect: it must be an error from an intermediary.
		// Twirp clients don't follow redirects automatically, Twirp only handles
		// POST requests, redirects should only happen on GET and HEAD requests.
		location := resp.Header.Get("Location")
		msg := fmt.Sprintf("unexpected HTTP status code %d %q received, Location=%q", statusCode, statusText, location)
		return twirpErrorFromIntermediary(statusCode, msg, location)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clientError("failed to read server error response body", err)
	}
	var tj twerrJSON
	if err := json.Unmarshal(respBodyBytes, &tj); err != nil {
		// Invalid JSON response; it must be an error from an intermediary.
		msg := fmt.Sprintf("Error from intermediary with HTTP status code %d %q", statusCode, statusText)
		return twirpErrorFromIntermediary(statusCode, msg, string(respBodyBytes))
	}

	errorCode := twirp.ErrorCode(tj.Code)
	if !twirp.IsValidErrorCode(errorCode) {
		msg := "invalid type returned from server error response: " + tj.Code
		return twirp.InternalError(msg)
	}

	twerr := twirp.NewError(errorCode, tj.Msg)
	for k, v := range tj.Meta {
		twerr = twerr.WithMeta(k, v)
	}
	return twerr
}

// twirpErrorFromIntermediary maps HTTP errors from non-twirp sources to twirp errors.
// The mapping is similar to gRPC: https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md.
// Returned twirp Errors have some additional metadata for inspection.
func twirpErrorFromIntermediary(status int, msg string, bodyOrLocation string) twirp.Error {
	var code twirp.ErrorCode
	if isHTTPRedirect(status) { // 3xx
		code = twirp.Internal
	} else {
		switch status {
		case 400: // Bad Request
			code = twirp.Internal
		case 401: // Unauthorized
			code = twirp.Unauthenticated
		case 403: // Forbidden
			code = twirp.PermissionDenied
		case 404: // Not Found
			code = twirp.BadRoute
		case 429, 502, 503, 504: // Too Many Requests, Bad Gateway, Service Unavailable, Gateway Timeout
			code = twirp.Unavailable
		default: // All other codes
			code = twirp.Unknown
		}
	}

	twerr := twirp.NewError(code, msg)
	twerr = twerr.WithMeta("http_error_from_intermediary", "true") // to easily know if this error was from intermediary
	twerr = twerr.WithMeta("status_code", strconv.Itoa(status))
	if isHTTPRedirect(status) {
		twerr = twerr.WithMeta("location", bodyOrLocation)
	} else {
		twerr = twerr.WithMeta("body", bodyOrLocation)
	}
	return twerr
}
func isHTTPRedirect(status int) bool {
	return status >= 300 && status <= 399
}

// wrappedError implements the github.com/pkg/errors.Causer interface, allowing errors to be
// examined for their root cause.
type wrappedError struct {
	msg   string
	cause error
}

func wrapErr(err error, msg string) error { return &wrappedError{msg: msg, cause: err} }
func (e *wrappedError) Cause() error      { return e.cause }
func (e *wrappedError) Error() string     { return e.msg + ": " + e.cause.Error() }

// clientError adds consistency to errors generated in the client
func clientError(desc string, err error) twirp.Error {
	return twirp.InternalErrorWith(wrapErr(err, desc))
}

// badRouteError is used when the twirp server cannot route a request
func badRouteError(msg string, method, url string) twirp.Error {
	err := twirp.NewError(twirp.BadRoute, msg)
	err = err.WithMeta("twirp_invalid_route", method+" "+url)
	return err
}

// The standard library will, by default, redirect requests (including POSTs) if it gets a 302 or
// 303 response, and also 301s in go1.8. It redirects by making a second request, changing the
// method to GET and removing the body. This produces very confusing error messages, so instead we
// set a redirect policy that always errors. This stops Go from executing the redirect.
//
// We have to be a little careful in case the user-provided http.Client has its own CheckRedirect
// policy - if so, we'll run through that policy first.
//
// Because this requires modifying the http.Client, we make a new copy of the client and return it.
func withoutRedirects(in *http.Client) *http.Client {
	copy := *in
	copy.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if in.CheckRedirect != nil {
			// Run the input's redirect if it exists, in case it has side effects, but ignore any error it
			// returns, since we want to use ErrUseLastResponse.
			err := in.CheckRedirect(req, via)
			_ = err // Silly, but this makes sure generated code passes errcheck -blank, which some people use.
		}
		return http.ErrUseLastResponse
	}
	return &copy
}

// doProtobufRequest is common code to make a request to the remote twirp service.
func doProtobufRequest(ctx context.Context, client HTTPClient, url string, in, out proto.Message) (err error) {
	reqBodyBytes, err := proto.Marshal(in)
	if err != nil {
		return clientError("failed to marshal proto request", err)
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	req, err := newRequest(ctx, url, reqBody, "application/protobuf")
	if err != nil {
		return clientError("could not build request", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return clientError("failed to do request", err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = clientError("failed to close response body", cerr)
		}
	}()

	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clientError("failed to read response body", err)
	}
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	if err = proto.Unmarshal(respBodyBytes, out); err != nil {
		return clientError("failed to unmarshal proto response", err)
	}
	return nil
}

// doJSONRequest is common code to make a request to the remote twirp service.
func doJSONRequest(ctx context.Context, client HTTPClient, url string, in, out proto.Message) (err error) {
	reqBody := bytes.NewBuffer(nil)
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(reqBody, in); err != nil {
		return clientError("failed to marshal json request", err)
	}
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	req, err := newRequest(ctx, url, reqBody, "application/json")
	if err != nil {
		return clientError("could not build request", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return clientError("failed to do request", err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = clientError("failed to close response body", cerr)
		}
	}()

	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(resp.Body, out); err != nil {
		return clientError("failed to unmarshal json response", err)
	}
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}
	return nil
}

// Call twirp.ServerHooks.RequestReceived if the hook is available
func callRequestReceived(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestReceived == nil {
		return ctx, nil
	}
	return h.RequestReceived(ctx)
}

// Call twirp.ServerHooks.RequestRouted if the hook is available
func callRequestRouted(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestRouted == nil {
		return ctx, nil
	}
	return h.RequestRouted(ctx)
}

// Call twirp.ServerHooks.ResponsePrepared if the hook is available
func callResponsePrepared(ctx context.Context, h *twirp.ServerHooks) context.Context {
	if h == nil || h.ResponsePrepared == nil {
		return ctx
	}
	return h.ResponsePrepared(ctx)
}

// Call twirp.ServerHooks.ResponseSent if the hook is available
func callResponseSent(ctx context.Context, h *twirp.ServerHooks) {
	if h == nil || h.ResponseSent == nil {
		return
	}
	h.ResponseSent(ctx)
}

// Call twirp.ServerHooks.Error if the hook is available
func callError(ctx context.Context, h *twirp.ServerHooks, err twirp.Error) context.Context {
	if h == nil || h.Error == nil {
		return ctx
	}
	return h.Error(ctx, err)
}

var twirpFileDescriptor0 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4f, 0xea, 0x40,
	0x10, 0xc7, 0xd3, 0xf2, 0x78, 0xf4, 0x0d, 0x4f, 0xd0, 0x15, 0x4c, 0xed, 0x05, 0xd2, 0x13, 0x41,
	0x2d, 0x09, 0xc6, 0xc4, 0xab, 0x86, 0x83, 0x5c, 0x3c, 0x80, 0x27, 0x62, 0x42, 0x4a, 0x77, 0x24,
	0x35, 0x6d, 0xb7, 0xee, 0x6e, 0xe1, 0xdb, 0xf8, 0x41, 0xfc, 0x74, 0x86, 0x5d, 0x6a, 0x10, 0xdb,
	0xe0, 0xa9, 0x9d, 0xd9, 0xdf, 0xfe, 0xe7, 0x3f, 0x33, 0x0b, 0x27, 0x3c, 0x0d, 0x06, 0x02, 0xf9,
	0x2a, 0x0c, 0xd0, 0x4b, 0x39, 0x93, 0x8c, 0xb4, 0x17, 0x91, 0x8f, 0xf4, 0xd5, 0x13, 0x01, 0x0f,
	0x53, 0x99, 0xc5, 0x3a, 0xed, 0x74, 0x96, 0x8c, 0x2d, 0x23, 0x1c, 0xa8, 0x68, 0x91, 0xbd, 0x0c,
	0x64, 0x18, 0xa3, 0x90, 0x7e, 0x9c, 0x6a, 0xc0, 0x7d, 0x37, 0xc0, 0x1a, 0xb1, 0x20, 0x8b, 0x31,
	0x91, 0xa4, 0x05, 0x55, 0x19, 0xca, 0x08, 0x6d, 0xa3, 0x6b, 0xf4, 0xfe, 0x4d, 0x74, 0x40, 0x1c,
	0xb0, 0x02, 0x96, 0x48, 0x4c, 0xa4, 0xb0, 0x4d, 0x75, 0xf0, 0x15, 0x93, 0x06, 0x98, 0x21, 0xb5,
	0x2b, 0x2a, 0x6b, 0x86, 0x94, 0xdc, 0x80, 0x25, 0xfc, 0x15, 0xd2, 0xb9, 0x2f, 0xed, 0x3f, 0x5d,
	0xa3, 0x57, 0x1f, 0x3a, 0x9e, 0xb6, 0xe0, 0xe5, 0x16, 0xbc, 0xa7, 0xdc, 0xc2, 0xa4, 0xa6, 0xd8,
	0x3b, 0x49, 0xce, 0xc1, 0x62, 0xeb, 0x04, 0xf9, 0x3c, 0xa4, 0x76, 0x55, 0x89, 0xd5, 0x54, 0x3c,
	0xa6, 0xee, 0x03, 0x9c, 0x3e, 0xe2, 0x3a, 0xb7, 0x38, 0x41, 0x91, 0xb2, 0x44, 0x20, 0x39, 0x86,
	0x0a, 0x72, 0xbe, 0x35, 0xba, 0xf9, 0x25, 0x1d, 0xa8, 0xd3, 0x2d, 0xb5, 0x91, 0xd1, 0x4e, 0x21,
	0x4f, 0x8d, 0xa9, 0xdb, 0x83, 0xd6, 0xd4, 0x5f, 0xe1, 0x61, 0x29, 0xf7, 0x16, 0xda, 0x23, 0x8c,
	0x50, 0xee, 0xb0, 0x6f, 0x19, 0x0a, 0xb9, 0x5f, 0xc3, 0xf8, 0x51, 0xa3, 0x0f, 0x67, 0xfb, 0x37,
	0xcb, 0xaa, 0x0c, 0x3f, 0x4c, 0x68, 0x4e, 0xb7, 0xeb, 0x9a, 0xea, 0x65, 0x92, 0x19, 0xd4, 0x77,
	0xba, 0x25, 0x1d, 0xaf, 0x70, 0xad, 0x5e, 0x0e, 0x38, 0xfd, 0x12, 0xa0, 0x68, 0x64, 0xcf, 0xf0,
	0x7f, 0xb7, 0xff, 0xc3, 0xe2, 0x17, 0x25, 0x40, 0xe1, 0x14, 0x63, 0x68, 0x7c, 0xef, 0x9c, 0x5c,
	0x96, 0xe9, 0x17, 0x8d, 0xd6, 0xb9, 0xfa, 0x25, 0xad, 0xcb, 0xdd, 0x37, 0x67, 0x47, 0x39, 0xa8,
	0x1f, 0xd6, 0x5f, 0xf5, 0xb9, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xf0, 0x33, 0x2c, 0x1c,
	0x03, 0x00, 0x00,
}
