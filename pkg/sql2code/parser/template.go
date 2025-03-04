package parser

import (
	"sync"
	"text/template"

	"github.com/pkg/errors"
)

var (
	modelStructTmpl    *template.Template
	modelStructTmplRaw = `
{{- if .Comment -}}
// {{.TableName}} {{.Comment}}
{{end -}}
type {{.TableName}} struct {
{{- range .Fields}}
	{{.Name}} {{.GoType}} {{if .Tag}}` + "`{{.Tag}}`" + `{{end}}{{if .Comment}} // {{.Comment}}{{end}}
{{- end}}
}
{{if .NameFunc}}
// TableName table name
func (m *{{.TableName}}) TableName() string {
	return "{{.RawTableName}}"
}
{{end}}
`

	modelTmpl    *template.Template
	modelTmplRaw = `package {{.Package}}
{{if .ImportPath}}
import (
	{{- range .ImportPath}}
	"{{.}}"
	{{- end}}
)
{{- end}}
{{range .StructCode}}
{{.}}
{{end}}`

	updateFieldTmpl    *template.Template
	updateFieldTmplRaw = `
{{- range .Fields}}
	if table.{{.Name}} {{.ConditionZero}} {
		update["{{.ColName}}"] = table.{{.Name}}
	}
{{- end}}`

	handlerCreateStructTmpl    *template.Template
	handlerCreateStructTmplRaw = `
// Create{{.TableName}}Request request params
type Create{{.TableName}}Request struct {
{{- range .Fields}}
	{{.Name}}  {{.GoType}} ` + "`" + `json:"{{.JSONName}}" binding:""` + "`" + `{{if .Comment}} // {{.Comment}}{{end}}
{{- end}}
}
`

	handlerUpdateStructTmpl    *template.Template
	handlerUpdateStructTmplRaw = `
// Update{{.TableName}}ByIDRequest request params
type Update{{.TableName}}ByIDRequest struct {
{{- range .Fields}}
	{{.Name}}  {{.GoType}} ` + "`" + `json:"{{.JSONName}}" binding:""` + "`" + `{{if .Comment}} // {{.Comment}}{{end}}
{{- end}}
}
`

	handlerDetailStructTmpl    *template.Template
	handlerDetailStructTmplRaw = `
// {{.TableName}}ObjDetail detail
type {{.TableName}}ObjDetail struct {
{{- range .Fields}}
	{{.Name}}  {{.GoType}} ` + "`" + `json:"{{.JSONName}}"` + "`" + `{{if .Comment}} // {{.Comment}}{{end}}
{{- end}}
}`

	modelJSONTmpl    *template.Template
	modelJSONTmplRaw = `{
{{- range .Fields}}
	"{{.ColName}}" {{.GoZero}}
{{- end}}
}
`

	protoFileTmpl    *template.Template
	protoFileTmplRaw = `syntax = "proto3";

package api.serverNameExample.v1;

import "api/types/types.proto";
import "validate/validate.proto";

option go_package = "github.com/zhufuyi/sponge/api/serverNameExample/v1;v1";

service {{.TName}} {
  // create {{.TName}}
  rpc Create(Create{{.TableName}}Request) returns (Create{{.TableName}}Reply) {}

  // delete {{.TName}} by id
  rpc DeleteByID(Delete{{.TableName}}ByIDRequest) returns (Delete{{.TableName}}ByIDReply) {}

  // delete {{.TName}} by batch id
  rpc DeleteByIDs(Delete{{.TableName}}ByIDsRequest) returns (Delete{{.TableName}}ByIDsReply) {}

  // update {{.TName}} by id
  rpc UpdateByID(Update{{.TableName}}ByIDRequest) returns (Update{{.TableName}}ByIDReply) {}

  // get {{.TName}} by id
  rpc GetByID(Get{{.TableName}}ByIDRequest) returns (Get{{.TableName}}ByIDReply) {}

  // get {{.TName}} by condition
  rpc GetByCondition(Get{{.TableName}}ByConditionRequest) returns (Get{{.TableName}}ByConditionReply) {}

  // list of {{.TName}} by batch id
  rpc ListByIDs(List{{.TableName}}ByIDsRequest) returns (List{{.TableName}}ByIDsReply) {}

  // list {{.TName}} by last id
  rpc ListByLastID(List{{.TableName}}ByLastIDRequest) returns (List{{.TableName}}ByLastIDReply) {}

  // list of {{.TName}} by query parameters
  rpc List(List{{.TableName}}Request) returns (List{{.TableName}}Reply) {}
}

// Some notes on defining fields under message:
// Fill in the validate rules https://github.com/envoyproxy/protoc-gen-validate#constraint-rules

// protoMessageCreateCode

message Create{{.TableName}}Reply {
  uint64   id = 1;
}

message Delete{{.TableName}}ByIDRequest {
  uint64   id = 1 [(validate.rules).uint64.gt  = 0];
}

message Delete{{.TableName}}ByIDReply {

}

message Delete{{.TableName}}ByIDsRequest {
  repeated uint64 ids = 1 [(validate.rules).repeated.min_items = 1];
}

message Delete{{.TableName}}ByIDsReply {

}

// protoMessageUpdateCode

message Update{{.TableName}}ByIDReply {

}

// protoMessageDetailCode

message Get{{.TableName}}ByIDRequest {
  uint64   id = 1 [(validate.rules).uint64.gt  = 0];
}

message Get{{.TableName}}ByIDReply {
  {{.TableName}} {{.TName}} = 1;
}

message Get{{.TableName}}ByConditionRequest {
  types.Conditions conditions = 1;
}

message Get{{.TableName}}ByConditionReply {
  {{.TableName}} {{.TName}} = 1;
}

message List{{.TableName}}ByIDsRequest {
  repeated uint64 ids = 1 [(validate.rules).repeated.min_items = 1];
}

message List{{.TableName}}ByIDsReply {
  repeated {{.TableName}} {{.TName}}s = 1;
}

message List{{.TableName}}ByLastIDRequest {
  uint64 lastID = 1; // last id
  uint32 limit = 2; // page size
  string sort = 3; // sort by column name of table, default is -id, the - sign indicates descending order.
}

message List{{.TableName}}ByLastIDReply {
  repeated {{.TableName}} {{.TName}}s = 1;
}

message List{{.TableName}}Request {
  types.Params params = 1;
}

message List{{.TableName}}Reply {
  int64 total =1;
  repeated {{.TableName}} {{.TName}}s = 2;
}
`

	protoFileForWebTmpl    *template.Template
	protoFileForWebTmplRaw = `syntax = "proto3";

package api.serverNameExample.v1;

import "api/types/types.proto";
import "google/api/annotations.proto";
import "protoc-gen-openapiv2/options/annotations.proto";
import "tagger/tagger.proto";
import "validate/validate.proto";

option go_package = "github.com/zhufuyi/sponge/api/serverNameExample/v1;v1";

// Default settings for generating swagger documents
// NOTE: because json does not support 64 bits, the int64 and uint64 types under *.swagger.json are automatically converted to string types
// Reference https://github.com/grpc-ecosystem/grpc-gateway/blob/db7fbefff7c04877cdb32e16d4a248a024428207/examples/internal/proto/examplepb/a_bit_of_everything.proto  
option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
  host: "localhost:8080"
  base_path: ""
  info: {
    title: "serverNameExample api docs";
    version: "2.0";
  }
  schemes: HTTP;
  schemes: HTTPS;
  consumes: "application/json";
  produces: "application/json";
  security_definitions: {
    security: {
      key: "BearerAuth";
      value: {
        type: TYPE_API_KEY;
        in: IN_HEADER;
        name: "Authorization";
        description: "Input a \"Bearer your-jwt-token\" to Value";
      }
    }
  }
};

service {{.TName}} {
  // create {{.TName}}
  rpc Create(Create{{.TableName}}Request) returns (Create{{.TableName}}Reply) {
    option (google.api.http) = {
      post: "/api/v1/{{.TName}}"
      body: "*"
    };
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      summary: "create {{.TName}}",
      description: "submit information to create {{.TName}}",
    };
  }

  // delete {{.TName}} by id
  rpc DeleteByID(Delete{{.TableName}}ByIDRequest) returns (Delete{{.TableName}}ByIDReply) {
    option (google.api.http) = {
      delete: "/api/v1/{{.TName}}/{id}"
    };
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      summary: "delete {{.TName}}",
      description: "delete {{.TName}} by id",
      //security: {
      //  security_requirement: {
      //    key: "BearerAuth";
      //    value: {}
      //  }
      //}
    };
  }

  // delete {{.TName}} by batch id
  rpc DeleteByIDs(Delete{{.TableName}}ByIDsRequest) returns (Delete{{.TableName}}ByIDsReply) {
    option (google.api.http) = {
      post: "/api/v1/{{.TName}}/delete/ids"
      body: "*"
    };
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      summary: "delete {{.TName}}s by batch id",
      description: "delete {{.TName}}s by batch id",
      //security: {
      //  security_requirement: {
      //    key: "BearerAuth";
      //    value: {}
      //  }
      //}
    };
  }

  // update {{.TName}} by id
  rpc UpdateByID(Update{{.TableName}}ByIDRequest) returns (Update{{.TableName}}ByIDReply) {
    option (google.api.http) = {
      put: "/api/v1/{{.TName}}/{id}"
      body: "*"
    };
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      summary: "update {{.TName}}",
      description: "update {{.TName}} by id",
      //security: {
      //  security_requirement: {
      //    key: "BearerAuth";
      //    value: {}
      //  }
      //}
    };
  }

  // get {{.TName}} by id
  rpc GetByID(Get{{.TableName}}ByIDRequest) returns (Get{{.TableName}}ByIDReply) {
    option (google.api.http) = {
      get: "/api/v1/{{.TName}}/{id}"
    };
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      summary: "get {{.TName}} detail",
      description: "get {{.TName}} detail by id",
      //security: {
      //  security_requirement: {
      //    key: "BearerAuth";
      //    value: {}
      //  }
      //}
    };
  }

  // get {{.TName}} by condition
  rpc GetByCondition(Get{{.TableName}}ByConditionRequest) returns (Get{{.TableName}}ByConditionReply) {
    option (google.api.http) = {
      post: "/api/v1/{{.TName}}/condition"
      body: "*"
    };
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      summary: "get {{.TName}} detail by condition",
      description: "get {{.TName}} detail by condition",
      //security: {
      //  security_requirement: {
      //    key: "BearerAuth";
      //    value: {}
      //  }
      //}
    };
  }

  // list of {{.TName}} by batch id
  rpc ListByIDs(List{{.TableName}}ByIDsRequest) returns (List{{.TableName}}ByIDsReply) {
    option (google.api.http) = {
      post: "/api/v1/{{.TName}}/list/ids"
      body: "*"
    };
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      summary: "list of {{.TName}}s by batch id",
      description: "list of {{.TName}}s by batch id",
      //security: {
      //  security_requirement: {
      //    key: "BearerAuth";
      //    value: {}
      //  }
      //}
    };
  }

  // list {{.TName}} by last id
  rpc ListByLastID(List{{.TableName}}ByLastIDRequest) returns (List{{.TableName}}ByLastIDReply) {
    option (google.api.http) = {
      get: "/api/v1/{{.TName}}/list"
    };
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      summary: "list of {{.TName}} by last id",
      description: "list of {{.TName}} by last id",
      //security: {
      //  security_requirement: {
      //    key: "BearerAuth";
      //    value: {}
      //  }
      //}
    };
  }

  // list of {{.TName}} by query parameters
  rpc List(List{{.TableName}}Request) returns (List{{.TableName}}Reply) {
    option (google.api.http) = {
      post: "/api/v1/{{.TName}}/list"
      body: "*"
    };
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      summary: "list of {{.TName}}s by parameters",
      description: "list of {{.TName}}s by paging and conditions",
      //security: {
      //  security_requirement: {
      //    key: "BearerAuth";
      //    value: {}
      //  }
      //}
    };
  }
}

// Some notes on defining fields under message:
// (1) Fill in the validate rules https://github.com/envoyproxy/protoc-gen-validate#constraint-rules
// (2) When using the protoc-gen-openapiv2 plugin, if the defined fields are snake case,
//      you must add annotations for snake case names, such as string foo_bar = 1 [json_name = "foo_bar"],
//      to ensure that the front end and back end JSON naming is consistent.
// (3) If the route contains the path parameter, such as /api/v1/userExample/{id}, the defined
//      message  must contain the name of the path parameter and the name should be
//      added with a new tag, such as int64 id = 1 [(tagger.tags) = "uri:\"id\""];
// (4) If the request url is followed by a query parameter, such as /api/v1/getUserExample?name=Tom,
//      a form tag must be added when defining the query parameter in the message,
//      such as string name = 1 [(tagger.tags) = "form:\"name\""];


// protoMessageCreateCode

message Create{{.TableName}}Reply {
  uint64   id = 1;
}

message Delete{{.TableName}}ByIDRequest {
  uint64   id =1 [(validate.rules).uint64.gte  = 1, (tagger.tags) = "uri:\"id\"" ];
}

message Delete{{.TableName}}ByIDReply {

}

message Delete{{.TableName}}ByIDsRequest {
  repeated uint64 ids = 1 [(validate.rules).repeated.min_items = 1];
}

message Delete{{.TableName}}ByIDsReply {

}

// protoMessageUpdateCode

message Update{{.TableName}}ByIDReply {

}

// protoMessageDetailCode

message Get{{.TableName}}ByIDRequest {
  uint64   id =1 [(validate.rules).uint64.gte  = 1, (tagger.tags) = "uri:\"id\"" ];
}

message Get{{.TableName}}ByIDReply {
  {{.TableName}} {{.TName}} = 1;
}

message Get{{.TableName}}ByConditionRequest {
  types.Conditions conditions = 1;
}

message Get{{.TableName}}ByConditionReply {
  {{.TableName}} {{.TName}} = 1;
}

message List{{.TableName}}ByIDsRequest {
  repeated uint64 ids = 1 [(validate.rules).repeated.min_items = 1];
}

message List{{.TableName}}ByIDsReply {
  repeated {{.TableName}} {{.TName}}s = 1;
}

message List{{.TableName}}ByLastIDRequest {
  uint64 lastID = 1 [(tagger.tags) = "form:\"lastID\""]; // last id
  uint32 limit = 2 [(tagger.tags) = "form:\"limit\""]; // page size
  string sort = 3 [(tagger.tags) = "form:\"sort\""]; // sort by column name of table, default is -id, the - sign indicates descending order.
}

message List{{.TableName}}ByLastIDReply {
  repeated {{.TableName}} {{.TName}}s = 1;
}

message List{{.TableName}}Request {
  types.Params params = 1;
}

message List{{.TableName}}Reply {
  int64 total =1;
  repeated {{.TableName}} {{.TName}}s = 2;
}
`

	protoMessageCreateTmpl    *template.Template
	protoMessageCreateTmplRaw = `message Create{{.TableName}}Request {
{{- range $i, $v := .Fields}}
	{{$v.GoType}} {{$v.JSONName}} = {{$v.AddOne $i}}; {{if $v.Comment}} // {{$v.Comment}}{{end}}
{{- end}}
}`

	protoMessageUpdateTmpl    *template.Template
	protoMessageUpdateTmplRaw = `message Update{{.TableName}}ByIDRequest {
{{- range $i, $v := .Fields}}
	{{$v.GoType}} {{$v.JSONName}} = {{$v.AddOneWithTag $i}}; {{if $v.Comment}} // {{$v.Comment}}{{end}}
{{- end}}
}`

	protoMessageDetailTmpl    *template.Template
	protoMessageDetailTmplRaw = `message {{.TableName}} {
{{- range $i, $v := .Fields}}
	{{$v.GoType}} {{$v.JSONName}} = {{$v.AddOne $i}}; {{if $v.Comment}} // {{$v.Comment}}{{end}}
{{- end}}
}`

	serviceStructTmpl    *template.Template
	serviceStructTmplRaw = `
		{
			name: "Create",
			fn: func() (interface{}, error) {
				// todo enter parameters before testing
// serviceCreateStructCode
			},
			wantErr: false,
		},

		{
			name: "UpdateByID",
			fn: func() (interface{}, error) {
				// todo enter parameters before testing
// serviceUpdateStructCode
			},
			wantErr: false,
		},
`

	serviceCreateStructTmpl    *template.Template
	serviceCreateStructTmplRaw = `				req := &serverNameExampleV1.Create{{.TableName}}Request{
					{{- range .Fields}}
						{{.Name}}:  {{.GoTypeZero}}, {{if .Comment}} // {{.Comment}}{{end}}
					{{- end}}
				}
				return cli.Create(ctx, req)`

	serviceUpdateStructTmpl    *template.Template
	serviceUpdateStructTmplRaw = `				req := &serverNameExampleV1.Update{{.TableName}}ByIDRequest{
					{{- range .Fields}}
						{{.Name}}:  {{.GoTypeZero}}, {{if .Comment}} // {{.Comment}}{{end}}
					{{- end}}
				}
				return cli.UpdateByID(ctx, req)`

	tmplParseOnce sync.Once
)

func initTemplate() {
	tmplParseOnce.Do(func() {
		var err, errSum error

		modelStructTmpl, err = template.New("goStruct").Parse(modelStructTmplRaw)
		if err != nil {
			errSum = errors.Wrap(err, "modelStructTmplRaw")
		}
		modelTmpl, err = template.New("goFile").Parse(modelTmplRaw)
		if err != nil {
			errSum = errors.Wrap(errSum, "modelTmplRaw:"+err.Error())
		}
		updateFieldTmpl, err = template.New("goUpdateField").Parse(updateFieldTmplRaw)
		if err != nil {
			errSum = errors.Wrap(errSum, "updateFieldTmplRaw:"+err.Error())
		}
		handlerCreateStructTmpl, err = template.New("goPostStruct").Parse(handlerCreateStructTmplRaw)
		if err != nil {
			errSum = errors.Wrap(errSum, "handlerCreateStructTmplRaw:"+err.Error())
		}
		handlerUpdateStructTmpl, err = template.New("goPutStruct").Parse(handlerUpdateStructTmplRaw)
		if err != nil {
			errSum = errors.Wrap(errSum, "handlerUpdateStructTmplRaw:"+err.Error())
		}
		handlerDetailStructTmpl, err = template.New("goGetStruct").Parse(handlerDetailStructTmplRaw)
		if err != nil {
			errSum = errors.Wrap(errSum, "handlerDetailStructTmplRaw:"+err.Error())
		}
		modelJSONTmpl, err = template.New("modelJSON").Parse(modelJSONTmplRaw)
		if err != nil {
			errSum = errors.Wrap(errSum, "modelJSONTmplRaw:"+err.Error())
		}
		protoFileTmpl, err = template.New("protoFile").Parse(protoFileTmplRaw)
		if err != nil {
			errSum = errors.Wrap(errSum, "protoFileTmplRaw:"+err.Error())
		}
		protoFileForWebTmpl, err = template.New("protoFileForWeb").Parse(protoFileForWebTmplRaw)
		if err != nil {
			errSum = errors.Wrap(errSum, "protoFileForWebTmplRaw:"+err.Error())
		}
		protoMessageCreateTmpl, err = template.New("protoMessageCreate").Parse(protoMessageCreateTmplRaw)
		if err != nil {
			errSum = errors.Wrap(errSum, "protoMessageCreateTmplRaw:"+err.Error())
		}
		protoMessageUpdateTmpl, err = template.New("protoMessageUpdate").Parse(protoMessageUpdateTmplRaw)
		if err != nil {
			errSum = errors.Wrap(errSum, "protoMessageUpdateTmplRaw:"+err.Error())
		}
		protoMessageDetailTmpl, err = template.New("protoMessageDetail").Parse(protoMessageDetailTmplRaw)
		if err != nil {
			errSum = errors.Wrap(errSum, "protoMessageDetailTmplRaw:"+err.Error())
		}
		serviceCreateStructTmpl, err = template.New("serviceCreateStruct").Parse(serviceCreateStructTmplRaw)
		if err != nil {
			errSum = errors.Wrap(errSum, "serviceCreateStructTmplRaw:"+err.Error())
		}
		serviceUpdateStructTmpl, err = template.New("serviceUpdateStruct").Parse(serviceUpdateStructTmplRaw)
		if err != nil {
			errSum = errors.Wrap(errSum, "serviceUpdateStructTmplRaw:"+err.Error())
		}
		serviceStructTmpl, err = template.New("serviceStruct").Parse(serviceStructTmplRaw)
		if err != nil {
			errSum = errors.Wrap(errSum, "serviceStructTmplRaw:"+err.Error())
		}

		if errSum != nil {
			panic(errSum)
		}
	})
}
