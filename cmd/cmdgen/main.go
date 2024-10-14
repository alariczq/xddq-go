package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"strings"

	"xddq/pkg/game/db"
	_ "xddq/pkg/game/protocol/pb"

	"github.com/samber/lo"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const smTmpl = `
{{- range .}}

// {{.Comment}}
func (c *Client) On{{.ReqName | title}}(f func(ctx Context, msg *pb.{{.RespName | title}})error) {
	c.RegisterHandler({{.MsgId}}, func(ctx Context) error {
		msg := &pb.{{.RespName | title}}{}
		if err := proto.Unmarshal(ctx.Message().Body, msg); err != nil {
			return err
		}
		return f(ctx, msg)
	})
}

{{- end}}
`

type CmdSet struct {
	ReqName  string
	Req      string
	RespName string
	MsgId    int32
	Comment  string
}

func smHandlerGenerate(w io.Writer) {

	cmds := db.ListRecvCmds()
	cmds = lo.Filter(cmds, func(c *db.Cmd, _ int) bool {
		return hasMsg(c.SmMethod) || hasMsg(c.CmMethod)
	})

	seen := make(map[string][]int)

	for _, c := range cmds {
		seen[c.CmMethod+c.SmMethod] = append(seen[c.CmMethod+c.SmMethod], int(c.CmMsgId))
	}

	cmdSets := make([]*CmdSet, 0, len(cmds))

	for _, c := range cmds {
		name := c.CmMethod
		if name == "" {
			name = c.SmMethod
		}

		if len(seen[c.CmMethod+c.SmMethod]) > 1 {
			name = fmt.Sprintf("%s%d", name, c.CmMsgId)
		}

		cmdSets = append(cmdSets, &CmdSet{
			ReqName:  name,
			RespName: c.SmMethod,
			MsgId:    c.SmMsgId,
			Comment:  c.Comment,
		})
	}

	tmpl := template.Must(template.New("cmds").Funcs(template.FuncMap{
		"title": strings.Title,
	}).Parse(smTmpl))

	if err := tmpl.Execute(w, cmdSets); err != nil {
		log.Fatal(err)
	}
}

func hasMsg(name string) bool {
	_, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(fmt.Sprintf("com.yq.msg.CityMsg.%s", name)))
	return err == nil
}

const cmTmpl = `
{{- range .}}

// {{.Comment}}
func (c *Client) Send{{.ReqName | title}}(ctx context.Context, req *pb.{{.Req | title}}) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return c.Send(ctx, {{.MsgId}}, data)
}
{{- end}}
`

func cmHandlerGenerate(w io.Writer) {
	cmds := db.ListCmds()
	cmds = lo.Filter(cmds, func(c *db.Cmd, _ int) bool {
		return hasMsg(c.SmMethod) || hasMsg(c.CmMethod)
	})

	seen := make(map[string][]int)

	for _, c := range cmds {
		seen[c.CmMethod+c.SmMethod] = append(seen[c.CmMethod+c.SmMethod], int(c.CmMsgId))
	}

	cmdSets := make([]*CmdSet, 0, len(cmds))

	for _, c := range cmds {
		name := c.CmMethod

		if len(seen[c.CmMethod+c.SmMethod]) > 1 {
			name = fmt.Sprintf("%s%d", name, c.CmMsgId)
		}

		cmdSets = append(cmdSets, &CmdSet{
			ReqName:  name,
			Req:      c.CmMethod,
			RespName: c.SmMethod,
			MsgId:    c.CmMsgId,
			Comment:  c.Comment,
		})
	}

	tmpl := template.Must(template.New("cmds").Funcs(template.FuncMap{
		"title": strings.Title,
	}).Parse(cmTmpl))

	if err := tmpl.Execute(w, cmdSets); err != nil {
		log.Fatal(err)
	}
}

const header = `package client

import (
	"context"

	"xddq/pkg/game/protocol/pb"

	"google.golang.org/protobuf/proto"
)
`

func main() {
	f, err := os.OpenFile("./pkg/game/client/methods.go", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	io.WriteString(f, header)

	cmHandlerGenerate(f)
	smHandlerGenerate(f)
}
