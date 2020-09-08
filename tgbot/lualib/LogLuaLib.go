package lualib

import (
	"github.com/yuin/gopher-lua"
	"tgbot/tglog"
)

type LuaLogModule struct {
}

var (
	logger = &tglog.TGLog{}
)

func (l *LuaLogModule) Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		// "debug": l.logDebug,
		"info":   l.Info,
		"error":  l.Error,
		"notice": l.Notice,
	})
	L.Push(mod)
	return 1
}

func NewLogModule() *LuaLogModule {
	return &LuaLogModule{}
}

func (l *LuaLogModule) Info(L *lua.LState) int {
	tpl := L.CheckString(1)
	top := L.GetTop()

	args := make([]interface{}, top-1)
	for i := 2; i <= top; i++ {
		args[i-2] = L.Get(i)
	}

	if top == 2 {
		logger.Critical(tpl, L.Get(2))
	}

	return 1
}

func (l *LuaLogModule) Error(L *lua.LState) int {
	tpl := L.CheckString(1)

	top := L.GetTop()

	args := make([]interface{}, top-1)
	for i := 2; i <= top; i++ {
		args[i-2] = L.Get(i)
	}

	if top == 2 {
		logger.Error(tpl, L.Get(2))
	}
	return 1
}

func (l *LuaLogModule) Notice(L *lua.LState) int {
	tpl := L.CheckString(1)

	top := L.GetTop()

	args := make([]interface{}, top-1)
	for i := 2; i <= top; i++ {
		args[i-2] = L.Get(i)
	}

	if top == 2 {
		logger.Notice(tpl, L.Get(2))
	}
	return 1
}
