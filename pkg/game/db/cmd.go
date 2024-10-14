package db

import (
	"fmt"
	"sort"
)

type Cmd struct {
	CmMsgId    int32  `json:"cmMsgId,omitempty"`
	SmMsgId    int32  `json:"smMsgId,omitempty"`
	Comment    string `json:"comment,omitempty"`
	CmMethod   string `json:"cmMethod,omitempty"`
	SmMethod   string `json:"smMethod,omitempty"`
	ByteDecode string `json:"byteDecode,omitempty"`
}

func (c *Cmd) String() string {
	return fmt.Sprintf("CmMsgId: %d, SmMsgId: %d, Comment: %s, CmMethod: %s, SmMethod: %s, ByteDecode: %s", c.CmMsgId, c.SmMsgId, c.Comment, c.CmMethod, c.SmMethod, c.ByteDecode)
}

func initCmds() {
	var (
		cmdList     []*Cmd
		recvCmdList []*Cmd
	)
	load("cmdList.json", &cmdList)
	load("recvCmdList.json", &recvCmdList)

	for _, cmd := range cmdList {
		if cmd == nil {
			continue
		}
		cmds[cmd.CmMsgId] = cmd
	}

	for _, cmd := range recvCmdList {
		if cmd == nil {
			continue
		}
		recvCmds[cmd.SmMsgId] = cmd
	}
}

func ListCmds() []*Cmd {
	var _cmds []*Cmd
	for _, cmd := range cmds {
		_cmds = append(_cmds, cmd)
	}
	sort.Slice(_cmds, func(i, j int) bool {
		return _cmds[i].CmMsgId < _cmds[j].CmMsgId
	})
	return _cmds
}

func ListRecvCmds() []*Cmd {
	var _cmds []*Cmd
	for _, cmd := range recvCmds {
		_cmds = append(_cmds, cmd)
	}
	sort.Slice(_cmds, func(i, j int) bool {
		return _cmds[i].SmMsgId < _cmds[j].SmMsgId
	})
	return _cmds
}

func GetCmd(id int32) *Cmd {
	return cmds[id]
}

func GetRecvCmd(id int32) *Cmd {
	return recvCmds[id]
}
