package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

const version = "0.0.2"

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-gin %v\n", version)
		return
	}

	var flags flag.FlagSet

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			// {dirname}_ext.pb.go
			generateExtFile(gen, f)
			// {package_name}_http_server.pb.go
			generateHttpServer(gen, f)
			// {package_name}_http_client.pb.go
			generateHttpClient(gen, f)
			// {package_name}_json.pb.go
			generateJsonFile(gen, f)
		}
		return nil
	})
}

const (
	gocoreApi       = protogen.GoImportPath("github.com/sunmi-OS/gocore/v2/api")
	ecodePackage    = protogen.GoImportPath("github.com/sunmi-OS/gocore/v2/api/ecode")
	utilsPacakge    = protogen.GoImportPath("github.com/sunmi-OS/gocore/v2/utils")
	httpRequest     = protogen.GoImportPath("github.com/sunmi-OS/gocore/v2/utils/http-request")
	ginPackage      = protogen.GoImportPath("github.com/gin-gonic/gin")
	sonicPackage    = protogen.GoImportPath("github.com/bytedance/sonic")
	httpPackage     = protogen.GoImportPath("net/http")
	ctxPackage      = protogen.GoImportPath("context")
	strconvPackage  = protogen.GoImportPath("strconv")
	metadataPackage = protogen.GoImportPath("google.golang.org/grpc/metadata")
)

func generateFileHeader(g *protogen.GeneratedFile, file *protogen.File, gen *protogen.Plugin) {
	g.P("// Code generated by protoc-gen-go-gin. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc-gen-go-gin v", version)
	g.P("// - protoc            ", protocVersion(gen))
	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
}

func generateExtFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	// printErr("%#v", file)
	extFilename := string(file.GoImportPath) + "/" + path.Base(string(file.GoImportPath)) + "_ext.pb.go"
	g := gen.NewGeneratedFile(extFilename, file.GoImportPath)
	generateFileHeader(g, file, gen)
	generateExtContent(file, g)
	return g
}

// generateExtContent generates the http service definitions, excluding the package statement.
func generateExtContent(file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}

	// TResponse
	g.P("type TResponse[T any] struct {\n\tCode int    `json:\"code\"`\n\tData *T     `json:\"data\"`\n\tMsg  string `json:\"msg\"`\n}")
	g.P()

	g.P("var validateErr error = ", gocoreApi.Ident("ErrorBind"))
	g.P(`var releaseShowDetail bool
	var disableValidate bool

	// set you error or use api.ErrorBind(diable:是否启用自动validate, 如果启用则返回 validateErr or 原始错误)
	func SetAutoValidate(disable bool, validatErr error, releaseShowDetail bool) {
		disableValidate = disable
		validateErr = validatErr
		releaseShowDetail = releaseShowDetail
	}
	`)

	g.P(`func checkValidate(err error) error {
		if disableValidate || err == nil {
			return nil
	}`)
	g.P("if ", utilsPacakge.Ident("IsRelease"), "() && !releaseShowDetail {")
	g.P(`return validateErr
		}
		return err
	}`)
	g.P()

	g.P(`const customReturnKey = "x-md-local-customreturn"

	func SetCustomReturn(ctx *api.Context, flag bool) {
		c := ctx.Request.Context()
		md, ok := metadata.FromIncomingContext(c)
		if ok {
			md.Set(customReturnKey, []string{strconv.FormatBool(flag)}...)
		} else {
			md = metadata.Pairs(customReturnKey, strconv.FormatBool(flag))
		}
		c = metadata.NewIncomingContext(c, md)
		ctx.Request = ctx.Request.WithContext(c)
	}`)
	g.P()

	g.P(`func GetCustomReturn(ctx *api.Context) bool {
	c := ctx.Request.Context()`)
	g.P("md, ok := ", metadataPackage.Ident("FromIncomingContext"), "(c)")
	g.P("if ok {")
	g.P("flag, err := ", strconvPackage.Ident("ParseBool"), "(md.Get(customReturnKey)[0])")
	g.P(`if err != nil {
				return false
			}
			return flag
		}
		return false
	}`)
	g.P()

	g.P(`func setRetJSON(ctx *api.Context, resp interface{}, err error) {
	if GetCustomReturn(ctx) {
		return
	}
	ctx.RetJSON(resp, err)
	}`)
	g.P()

}

// generateFile generates a _grpc.pb.go file containing gRPC service definitions.
func generateHttpServer(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	g := gen.NewGeneratedFile(file.GeneratedFilenamePrefix+"_http_server.pb.go", file.GoImportPath)
	generateFileHeader(g, file, gen)
	generateHttpServerContent(file, g)
	return g
}

func generateJsonFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_json.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	generateFileHeader(g, file, gen)
	generateJsonContent(file, g)
	return g
}

func generateHttpClient(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	g := gen.NewGeneratedFile(file.GeneratedFilenamePrefix+"_http_client.pb.go", file.GoImportPath)
	generateFileHeader(g, file, gen)
	generateHttpClientContent(file, g)
	return g
}

// generate mashal/unmarshal method
func generateJsonContent(file *protogen.File, g *protogen.GeneratedFile) {
	for _, msg := range file.Messages {
		name := msg.Desc.FullName().Name()
		g.P()
		g.P(fmt.Sprintf(`func (m *%s) Marshal() ([]byte, error) {
			return %s(m)
		}

		func (m *%s) MarshalString() (string, error) {
			return %s(m)
		}

		func (m *%s)Unmarshal(buf []byte) (error) {
			m = new(%s)
			err := %s(buf, m)
			return err
		}

		func (m *%s)UnmarshalString(str string) (error) {
			m = new(%s)
			err := %s(str, m)
			return err
		}
		`, name, g.QualifiedGoIdent(sonicPackage.Ident("Marshal")),
			name, g.QualifiedGoIdent(sonicPackage.Ident("MarshalString")),
			name, name, g.QualifiedGoIdent(sonicPackage.Ident("Unmarshal")),
			name, name, g.QualifiedGoIdent(sonicPackage.Ident("UnmarshalString")),
		))
		g.P()
	}
}

func protocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}

func printErr(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format, a...)
}

// generateHttpServerContent generates the http service definitions, excluding the package statement.
func generateHttpServerContent(file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}

	for _, service := range file.Services {
		genService(g, service)
	}
}

// generateHttpClientContent generates the http client definitions, excluding the package statement.
func generateHttpClientContent(file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}

	for _, service := range file.Services {
		genClient(g, service)
	}
}

func serverSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	ret := "(*" + g.QualifiedGoIdent(method.Output.GoIdent) + ", error)"
	var reqArgs []string
	reqArgs = append(reqArgs, "*"+g.QualifiedGoIdent(gocoreApi.Ident("Context")))
	reqArgs = append(reqArgs, "*"+g.QualifiedGoIdent(method.Input.GoIdent))
	return method.GoName + "(" + strings.Join(reqArgs, ", ") + ") " + ret
}

func clientSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	ret := "(*TResponse[" + g.QualifiedGoIdent(method.Output.GoIdent) + "], error)"
	var reqArgs []string
	reqArgs = append(reqArgs, g.QualifiedGoIdent(ctxPackage.Ident("Context")))
	reqArgs = append(reqArgs, "*"+g.QualifiedGoIdent(method.Input.GoIdent))
	return method.GoName + "(" + strings.Join(reqArgs, ", ") + ") " + ret
}

func genService(g *protogen.GeneratedFile, service *protogen.Service) {
	// Server interface.
	serverType := service.GoName + "HTTPServer"
	g.P("// ", serverType, " is the server API for ", service.GoName, " service.")

	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
	}
	g.Annotate(serverType, service.Location)
	g.P("type ", serverType, " interface {")
	for _, m := range service.Methods {
		if m.Desc.IsStreamingClient() || m.Desc.IsStreamingServer() {
			continue
		}
		g.Annotate(serverType+"."+m.GoName, m.Location)
		if m.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated() {
		}
		g.P(m.Comments.Leading, serverSignature(g, m))
	}
	g.P("}")
	g.P()

	var methods []*method
	for _, m := range service.Methods {
		rule, ok := proto.GetExtension(m.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
		if rule != nil && ok {
			for _, bind := range rule.AdditionalBindings {
				methods = append(methods, buildHTTPRule(m, bind))
			}
			methods = append(methods, buildHTTPRule(m, rule))
		}
	}

	// Register service HttpServer.
	g.P("func Register", serverType, "(s *", ginPackage.Ident("Engine"), ", srv ", serverType, ") {")
	g.P(`r := s.Group("/")`)
	for _, m := range methods {
		g.P(fmt.Sprintf(`r.POST("%v", %v(srv))`, m.Path, httpHandlerName(service.GoName, m.Name, m.Num)))
	}
	g.P("}")
	g.P()

	// http method func
	for _, m := range methods {
		g.P("func ", httpHandlerName(service.GoName, m.Name, m.Num), "(srv ", serverType, ") func(g *gin.Context) {")
		g.P("return func(g *", ginPackage.Ident("Context"), ") {")
		g.P("req := &", m.Request, "{}")
		g.P(`ctx := api.NewContext(g)
			err := ctx.ShouldBind(req)
			err = checkValidate(err)
			if err != nil {
				setRetJSON(&ctx, nil, err)
				return
		}`)
		g.P("resp, err := srv.", m.Name, "(&ctx, req)")
		g.P(`setRetJSON(&ctx, resp, err)
		}}`)
		g.P()
	}
}

func genClient(g *protogen.GeneratedFile, service *protogen.Service) {
	// Server interface.
	serverType := service.GoName + "HTTPClient"
	g.P("// ", serverType, " is the client API for ", service.GoName, " service.")

	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
	}
	g.Annotate(serverType, service.Location)
	g.P("type ", serverType, " interface {")
	for _, m := range service.Methods {
		if m.Desc.IsStreamingClient() || m.Desc.IsStreamingServer() {
			continue
		}
		g.Annotate(serverType+"."+m.GoName, m.Location)
		g.P(m.Comments.Leading, clientSignature(g, m))
	}
	g.P("}")
	g.P()

	var methods []*method
	for _, m := range service.Methods {
		rule, ok := proto.GetExtension(m.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
		if rule != nil && ok {
			// 跳过additional的client生成，一般只需要请求一个接口
			//for _, bind := range rule.AdditionalBindings {
			//	methods = append(methods, buildHTTPRule(m, bind))
			//}
			methods = append(methods, buildHTTPRule(m, rule))
		}
	}

	// type XXXHttpClientImpl struct
	g.P("type ", serverType, "Impl struct {")
	g.P("hh *", httpRequest.Ident("HttpClient"))
	g.P("}")
	g.P()

	// func NewXXXHttpClient
	g.P("func New", serverType, "(hh *", httpRequest.Ident("HttpClient"), ") ", serverType, " {")
	g.P("return &", serverType, "Impl{hh: hh}")
	g.P("}")
	g.P()

	// http method func
	for _, m := range methods {
		// func (c *XXXHttpClientImpl) XXX(ctx *Context, req *XXXRequest) (*XXXResponse, error)
		g.P("func (c *", serverType, "Impl) ", m.Name, "(ctx ", ctxPackage.Ident("Context"), ", req *", m.Request, ") (*TResponse[", m.Reply, "], error) {")
		g.P("resp := &TResponse[", m.Reply, "]{}")
		g.P("_, err := c.hh.Client.R().SetContext(ctx).SetBody(req).SetResult(resp).Post(\"", m.Path, "\")")
		g.P(`if err != nil {
				return nil, err
			}
			return resp, err
		}`)
		g.P()
	}
}

// _{ServiceName}_{MethodName}_HTTPServer_Handler is the handler invoked by the HTTP transport layer for service
func httpHandlerName(serivceName, methodName string, num int) string {
	return fmt.Sprintf("_%s_%s%d_HTTP_Handler", serivceName, methodName, num)
}

type method struct {
	Name    string // SayHello
	Num     int    // 一个 rpc 方法可以对应多个 http 请求
	Request string // SayHelloReq
	Reply   string // SayHelloResp
	// http_rule
	Path         string // 路由
	Method       string // HTTP Method
	Body         string
	ResponseBody string
}

// HandlerName for gin handler name
func (m *method) HandlerName() string {
	return fmt.Sprintf("%s_%d", m.Name, m.Num)
}

// HasPathParams 是否包含路由参数
func (m *method) HasPathParams() bool {
	paths := strings.Split(m.Path, "/")
	for _, p := range paths {
		if len(p) > 0 && (p[0] == '{' && p[len(p)-1] == '}' || p[0] == ':') {
			return true
		}
	}
	return false
}

// initPathParams 转换参数路由 {xx} --> :xx
func (m *method) initPathParams() {
	paths := strings.Split(m.Path, "/")
	for i, p := range paths {
		if len(p) > 0 && (p[0] == '{' && p[len(p)-1] == '}' || p[0] == ':') {
			paths[i] = ":" + p[1:len(p)-1]
		}
	}
	m.Path = strings.Join(paths, "/")
}

var methodSets = make(map[string]int)

func buildHTTPRule(m *protogen.Method, rule *annotations.HttpRule) *method {
	path, ok := rule.Pattern.(*annotations.HttpRule_Post)
	if !ok {
		panic("method not support")
	}
	md := buildMethodDesc(m, "POST", path.Post)
	return md
}

func buildMethodDesc(m *protogen.Method, httpMethod, path string) *method {
	defer func() { methodSets[m.GoName]++ }()
	md := &method{
		Name:    m.GoName,
		Num:     methodSets[m.GoName],
		Request: m.Input.GoIdent.GoName,
		Reply:   m.Output.GoIdent.GoName,
		Path:    path,
		Method:  httpMethod,
	}
	md.initPathParams()
	return md
}
