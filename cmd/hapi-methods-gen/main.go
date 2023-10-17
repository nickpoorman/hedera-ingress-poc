package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

const codeTemplate = `package {{.PackageName}}

import (
	"context"

	"github.com/nickpoorman/hedera-ingress/proto"
	"google.golang.org/grpc"
)

func (h *HapiServer) RegisterServices(gsrv *grpc.Server, backendGrpcConn *grpc.ClientConn) error {
	{{- range $receiver, $methods := .MethodsByReceiver}}
	{{$receiver | firstCharToLower}} := &{{$receiver}}{
		backendClient: {{$.ProtoPackage}}.New{{$receiver | replaceServerWithClient}}(backendGrpcConn),
		hapiServer:    h,
	}
	{{$.ProtoPackage}}.Register{{$receiver}}(gsrv, {{$receiver | firstCharToLower}})
	{{end}}
	return nil
}
{{range $receiver, $methods := .MethodsByReceiver}}
type {{$receiver}} struct {
	{{$.ProtoPackage}}.Unimplemented{{$receiver}}
	backendClient {{$.ProtoPackage}}.{{$receiver | replaceServerWithClient}}
	hapiServer    *HapiServer
}

// validateTransaction takes a transaction and validates it. If the transaction
// is invalid, it returns an error.
func (s *{{$receiver}}) validateTransaction(ctx context.Context, tx *{{$.ProtoPackage}}.Transaction) error {
	// For now delegates to the global validateTransaction function.
	return s.hapiServer.validateTransaction(ctx, tx)
}

// handleValidationError takes a transaction and an error and returns a response
// with same error message that would have been returned by the server.
func (s *{{$receiver}}) handleTransactionValidationError(ctx context.Context, tx *{{$.ProtoPackage}}.Transaction, err error) (*{{$.ProtoPackage}}.TransactionResponse, error) {
	// For now delegates to the global handleTransactionValidationError function.
	return s.hapiServer.handleTransactionValidationError(ctx, tx, err)
}

// validateQuery takes a query and validates it. If the query
// is invalid, it returns an error.
func (s *{{$receiver}}) validateQuery(ctx context.Context, query *{{$.ProtoPackage}}.Query) error {
	// For now delegates to the global validateQuery function.
	return s.hapiServer.validateQuery(ctx, query)
}

// handleValidationError takes a query and an error and returns a response
// with same error message that would have been returned by the server.
func (s *{{$receiver}}) handleQueryValidationError(ctx context.Context, query *{{$.ProtoPackage}}.Query, err error) (*{{$.ProtoPackage}}.Response, error) {
	// For now delegates to the global handleQueryValidationError function.
	return s.hapiServer.handleQueryValidationError(ctx, query, err)
}
{{range $methods}}
func (s *{{$receiver}}) {{.MethodName}}({{.Params}}) ({{.Results}}) {
	{{- if .Params | isTransaction}}
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
	{{- end}}
	{{- if .Params | isQuery}}
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
	{{- end}}
    return s.backendClient.{{.MethodName}}(ctx, {{.Params | varFromParams}})
}
{{end}}
{{- end}}`

type MethodsByReceiver map[string][]Method

type Method struct {
	StructName string
	MethodName string
	Params     string
	Results    string
}

func main() {
	protoPackage := flag.String("proto-package", "proto", "Proto package name")

	directoryPath := flag.String("dir", "", "Path to the proto directory")
	flag.Parse()

	if *protoPackage == "" {
		log.Fatal("Please provide a valid proto package name using the -proto-package flag")
	}
	if *directoryPath == "" {
		log.Fatal("Please provide a valid directory path using the -dir flag")
	}

	files, err := os.ReadDir(*directoryPath)
	if err != nil {
		log.Fatal(err)
	}

	methodsByReceiver := MethodsByReceiver{}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), "_service_grpc.pb.go") {
			parseFile(methodsByReceiver, path.Join(*directoryPath, file.Name()), *protoPackage)
		}
	}

	// Generate Go code using templates
	tmpl := template.Must(template.New("code").Funcs(template.FuncMap{
		"firstCharToLower":        firstCharToLower,
		"replaceServerWithClient": replaceServerWithClient,
		"varFromParams":           varFromParams,
		"isTransaction":           isTransaction,
		"isQuery":                 isQuery,
	}).Parse(codeTemplate))
	err = tmpl.Execute(os.Stdout, map[string]interface{}{
		"PackageName":       "hapi", // adjust as needed
		"MethodsByReceiver": methodsByReceiver,
		"ProtoPackage":      *protoPackage,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func parseFile(methodsByReceiver MethodsByReceiver, filePath string, protoPackage string) {
	// methods := []Method{}
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	// Walk the AST and extract methods
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if x.Recv != nil && x.Name.IsExported() && len(x.Recv.List) > 0 {
				var receiverName string

				// We really only care about non-pointer receiver types
				switch typeExpr := x.Recv.List[0].Type.(type) {
				case *ast.StarExpr: // For pointer receiver types
					ident, ok := typeExpr.X.(*ast.Ident)
					if ok {
						receiverName = ident.Name
					}
				case *ast.Ident:
					receiverName = typeExpr.Name
				}

				if !strings.HasSuffix(receiverName, "Server") {
					return true
				}

				// Extracting Params
				params := []string{}
				for _, field := range x.Type.Params.List {
					typeExpr := field.Type
					paramType := exprToString(typeExpr)
					// fmt.Println("paramType", paramType)
					if paramType != "context.Context" {
						paramType = addPackageToType(protoPackage, paramType)
					}
					params = append(params, paramType)
				}

				// Extracting Results
				results := []string{}
				for _, field := range x.Type.Results.List {
					typeExpr := field.Type
					resultType := exprToString(typeExpr)
					if resultType != "error" {
						resultType = addPackageToType(protoPackage, resultType)
					}
					results = append(results, resultType)
				}

				// This is a method. Extract the info.
				method := Method{
					StructName: strings.Replace(receiverName, "Unimplemented", "", 1),
					MethodName: x.Name.Name,
					Params:     strings.Join(paramsWithVariables(params), ", "),
					Results:    strings.Join(results, ", "),
				}
				// methods = append(methods, method)
				methodsByReceiver[method.StructName] = append(methodsByReceiver[method.StructName], method)
			}
		}
		return true
	})
}

func addPackageToType(p string, t string) string {
	if strings.HasPrefix(t, "*") {
		return "*" + addPackageToType(p, t[1:])
	}

	return fmt.Sprintf("%s.%s", p, t)
}

func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + exprToString(t.X)
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		return "[]" + exprToString(t.Elt)
	default:
		return "" // handle other cases as needed
	}
}

func firstCharToLower(s string) string {
	if len(s) == 0 {
		return ""
	}

	uni := []rune(s)
	firstUni := string(uni[0:1]) // handle unicode
	firstChar := strings.ToLower(firstUni)
	return firstChar + string(uni[1:])
}

func paramsWithVariables(params []string) []string {
	var paramsVars []string
	for i, p := range params {
		paramsVars = append(paramsVars, paramTypeToVariable(i, p))
	}
	return paramsVars
}

func paramTypeToVariable(idx int, param string) string {
	if param == "context.Context" {
		return "ctx context.Context"
	}
	if param == "*proto.Transaction" {
		return "tx *proto.Transaction"
	}
	if param == "*proto.Query" {
		return "query *proto.Query"
	}
	// Split on the .
	tokens := strings.Split(param, ".")
	lastToken := tokens[len(tokens)-1]
	lower := strings.ToLower(lastToken)
	character := string([]rune(lower)[0:1]) // handle unicode
	return fmt.Sprintf("%s%d %s", character, idx, param)
}

func replaceServerWithClient(receiver string) string {
	return strings.Replace(receiver, "Server", "Client", 1)
}

func varFromParams(params string) string {
	if isTransaction(params) {
		return "tx"
	} else if isQuery(params) {
		return "query"
	} else {
		panic(fmt.Sprintf("unknown param: %s", params))
	}
}

func isTransaction(params string) bool {
	return strings.Contains(params, "Transaction")
}

func isQuery(params string) bool {
	return strings.Contains(params, "Query")
}
