package main

import (
	"fmt"
	"strings"

	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/desc/protoprint"

	ep "github.com/emicklei/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/pubgo/xerror"
)

const (
	Annotations = "google/api/annotations.proto"
)

func main() {
	md, err := desc.LoadMessageDescriptorForMessage((*empty.Empty)(nil))
	xerror.Exit(err)

	file := builder.NewFile("foo/bar.proto").SetPackageName("foo.bar")
	en := builder.NewEnum("Options").
		AddValue(builder.NewEnumValue("OPTION_1").SetComments(builder.Comments{LeadingComment: " OPTION_1"})).
		AddValue(builder.NewEnumValue("OPTION_2")).
		AddValue(builder.NewEnumValue("OPTION_3"))
	file.AddEnum(en)

	msg := builder.NewMessage("FooRequest").
		AddField(builder.NewField("id", builder.FieldTypeInt64())).
		AddField(builder.NewField("name", builder.FieldTypeString())).
		AddField(builder.NewField("options", builder.FieldTypeEnum(en)).
			SetRepeated())
	file.AddMessage(msg)

	sb := builder.NewService("FooService").
		AddMethod(builder.NewMethod("DoSomething", builder.RpcTypeMessage(msg, false), builder.RpcTypeMessage(msg, false))).
		AddMethod(builder.NewMethod("ReturnThings", builder.RpcTypeImportedMessage(md, false), builder.RpcTypeMessage(msg, true)))
	file.AddService(sb)

	fd, err := file.Build()
	xerror.Exit(err)
	fmt.Println(fd.String())
	fmt.Println(fd.AsProto().String())
	var p protoprint.Printer
	fmt.Println(xerror.ExitErr(p.PrintProtoToString(fd)))
	fmt.Print("\n\n\n\n")

	var data = `
package hello;
syntax = "proto3";
import "google/protobuf/descriptor.proto";

option go_package = "./;user_pb";

message Test {}
message Foo {
  repeated Bar bar = 1;
  message Bar {
    Baz baz = 1;
    string name = 2;
  }
  enum Baz {
	ZERO = 0;
	FROB = 1;
	NITZ = 2;
  }
}
extend google.protobuf.MethodOptions {
  Foo foo = 54321;
}
service TestService {
  rpc Hello (Test) returns (Test) {
    option (foo).bar = { 
		baz:NITZ,
		name:"xyz" 
	};
  }

  rpc Hello1 (Test) returns (Test) {}
  rpc Get (Test) returns (Test) {
 option (google.api.http) = {
    post: "/hello/TestService/Get"
    body: "*"
  };

    option (foo).bar = { baz:FROB name:"abc" };
    option (foo).bar = { baz:NITZ name:"xyz" };
  }
}
`
	files := map[string]string{"test.proto": data}

	pa := &protoparse.Parser{Accessor: protoparse.FileContentsFromMap(files)}
	fds, err := pa.ParseFiles("test.proto")
	//proto.Equal()
	_ = fds

	//fmt.Printf("%#v\n", fds[0].GetServices()[0].GetMethods()[0].GetMethodOptions().GetUninterpretedOption())
	//fds[0].GetServices()[0].GetMethods()[0].GetMethodOptions().ProtoReflect().Set()

	//gp.SetExtension(fds[0].GetServices()[0].GetMethods()[0].GetMethodOptions(), annotations.E_Http, &annotations.HttpRule{
	//	Pattern: &annotations.HttpRule_Post{
	//		Post: "/hello/test",
	//	},
	//	Body: "*",
	//})
	//fmt.Println(gp.GetExtension(fds[0].GetServices()[0].GetMethods()[0].GetMethodOptions(), annotations.E_Http))
	//var p1 protoprint.Printer
	//fmt.Println(xerror.ExitErr(p1.PrintProtoToString(fds[0])))

	parser := ep.NewParser(strings.NewReader(data))
	definition, err := parser.Parse()

	var pkg string
	ep.Walk(definition, ep.WithPackage(func(p *ep.Package) {
		var replacer = strings.NewReplacer(".", "/", "-", "/")
		pkg = replacer.Replace(p.Name)
	}))
	
	ep.WithOption(func(option *ep.Option) {
		fmt.Println(option.Name)
	})

	fmt.Println(pkg)
	var rpcs []*ep.RPC
	ep.Walk(definition, ep.WithService(func(srv *ep.Service) {
		for _, e := range srv.Elements {
			var rpc, ok = e.(*ep.RPC)
			if !ok {
				continue
			}

			rpcs = append(rpcs, rpc)
		}
	}))

	var dataLine = strings.Split(data, "\n")

	var mod bool
	for i := range rpcs {
		rpc := rpcs[i]
		insert := fmt.Sprintf(`
rpc %s (%s) returns (%s) {
  option (google.api.http) = {
    post: "/%s/%s/%s"
    body: "*"
  };`, rpc.Name, rpc.RequestType, rpc.ReturnsType, pkg, rpc.Parent.(*ep.Service).Name, rpc.Name)

		var hasHttp bool
		for i := range rpc.Options {
			if rpc.Options[i].Name == "(google.api.http)" {
				hasHttp = true
			}
		}

		// 如果option为0, 那么可以整体替换, 通过正则表达式
		if len(rpc.Options) == 0 || !hasHttp {
			mod = true
			_ = insert
			var rpcData = strings.Trim(dataLine[rpc.Position.Line-1], ";")
			// 以}结尾
			if rpcData[len(rpcData)-1] == '}' {
				dataLine[rpc.Position.Line-1] = insert + "\n}\n"
			} else {
				dataLine[rpc.Position.Line-1] = insert
			}
			//fmt.Println("line=>", strings.Split(data, "\n")[rpc.Position.Line-1])
			//fmt.Println("no=>", data[rpc.Position.Offset:rpc.Position.Offset+50])
			//} else {
			//	fmt.Println("has=>", data[rpc.Position.Offset:rpc.Options[0].Position.Offset])
		}
	}

	if mod {
		var imports = make(map[string]struct{})
		ep.Walk(definition, ep.WithImport(func(i *ep.Import) {
			imports[i.Filename] = struct{}{}
		}))
		if _, ok := imports[Annotations]; !ok {
			dataLine = append([]string{fmt.Sprintf(`import "%s"`, Annotations)}, dataLine...)
		}
	}

	data = strings.Join(dataLine, "\n")
	fmt.Println(data)
}
