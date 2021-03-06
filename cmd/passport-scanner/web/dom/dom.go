// +build js,wasm

package dom

import (
	"regexp"
	"syscall/js"
)

var (
	null      = js.Null()
	undefined = js.Undefined()

	Document Doc
	Body     js.Value
	alert    js.Value
)

func init() {
	Document = Doc{NodeBase{js.Global().Get("document")}}
	Body = Document.Get("body")
	alert = js.Global().Get("alert")
}

// EventCallbackFlag is for backward compatibility.
type EventCallbackFlag int

const (
	PreventDefault EventCallbackFlag = 1 << iota
	StopPropagation
	StopImmediatePropagation
)

func Alert(s string) {
	alert.Invoke(s)
}

type Node interface {
	JSValue() js.Value

	AppendChild(c Node)
	RemoveChild(c Node) Node
}

type NodeBase struct{ js.Value }

func (n NodeBase) JSValue() js.Value { return n.Value }

func (n NodeBase) AppendChild(c Node)      { n.Call("appendChild", c.JSValue()) }
func (n NodeBase) RemoveChild(c Node) Node { return NodeBase{n.Call("removeChild", c.JSValue())} }
func (n NodeBase) FirstChild() Node        { return NodeBase{n.Get("firstChild")} }
func (n NodeBase) RemoveAllChildren() {
	for c := n.FirstChild(); validJSValue(c.JSValue()); c = n.FirstChild() {
		n.RemoveChild(c)
	}
}
func (n NodeBase) ParentNode() Node { return NodeBase{n.Get("parentNode")} }
func (n NodeBase) Remove()          { n.ParentNode().RemoveChild(n) }
func (n NodeBase) AsElement() Elt   { return Elt{n} }

func validJSValue(v js.Value) bool {
	return v != js.Value{} && v != null && v != undefined
}

func (n NodeBase) AddEventListener(flags EventCallbackFlag, typ string, fn func(js.Value)) js.Func {
	callBack := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		if flags&PreventDefault != 0 {
			e.Call("preventDefault")
		}
		if flags&StopPropagation != 0 {
			e.Call("stopPropagation")
		}
		if flags&StopImmediatePropagation != 0 {
			e.Call("stopImmediatePropagation")
		}
		go func() {
			fn(e)
		}()
		return nil
	})
	n.Call("addEventListener", typ, callBack)
	return callBack
}

type Doc struct{ NodeBase }

func (d Doc) CreateElement(tag string) Elt { return Elt{NodeBase{Document.Call("createElement", tag)}} }
func (d Doc) GetElementById(id string) Elt { return Elt{NodeBase{Document.Call("getElementById", id)}} }

type Attrs map[string]string

type Elt struct{ NodeBase }

func Element(tag string) Elt {
	return Document.CreateElement(tag)
}

func (e Elt) SetInnerHTML(s string) { e.Set("innerHTML", s) }

func (e Elt) SetAttribute(k, v string) { e.Call("setAttribute", k, v) }

func (e Elt) RemoveAttribute(k string) { e.Call("removeAttribute", k) }

func (e Elt) SetClass(c string) { e.SetAttribute("class", c) }

func (e Elt) GetClass() *string {
	if ca := e.Call("getAttribute", "class"); validJSValue(ca) {
		s := ca.String()
		return &s
	}
	return nil
}

func hasClass(className, classes string) bool {
	m, _ := regexp.MatchString("(\\s|^)"+className+"(\\s|$)", classes)
	return m
}

func removeClass(className, classes string) string {
	r, err := regexp.Compile("(\\s|^)" + className + "(\\s|$)")
	if err != nil {
		return classes
	}
	return r.ReplaceAllString(classes, "")
}

func (e Elt) AddClass(c string) {
	if cl := e.GetClass(); cl != nil {
		if !hasClass(c, *cl) {
			e.SetClass(*cl + " " + c)
		}
	} else {
		e.SetClass(c)
	}
}

func (e Elt) RemoveClass(c string) {
	if cl := e.GetClass(); cl != nil {
		if hasClass(c, *cl) {
			e.SetClass(removeClass(c, *cl))
		}
	}
}

func (e Elt) SetRole(r string) { e.SetAttribute("role", r) }

func (e Elt) WithAttribute(k, v string) Elt {
	e.SetAttribute(k, v)
	return e
}

func (e Elt) WithAttributeRemoved(k string) Elt {
	e.RemoveAttribute(k)
	return e
}

func (e Elt) WithClass(c string) Elt {
	e.SetClass(c)
	return e
}

func (e Elt) WithClassAdded(c string) Elt {
	e.AddClass(c)
	return e
}

func (e Elt) WithClassRemoved(c string) Elt {
	e.RemoveClass(c)
	return e
}

func (e Elt) WithRole(r string) Elt {
	e.SetRole(r)
	return e
}

func (e Elt) WithAttributes(attrs Attrs) Elt {
	for k, v := range attrs {
		e.SetAttribute(k, v)
	}
	return e
}

func (e Elt) WithChildren(ns ...Node) Elt {
	for _, n := range ns {
		e.AppendChild(n)
	}
	return e
}

func (e Elt) AsInput() Inp {
	return Inp{e}
}

func (e Elt) AsButton() Btn {
	return Btn{e}
}

func (e Elt) AsAnchor() A {
	return A{e}
}

func Text(s string) Elt { return Elt{NodeBase{Document.Call("createTextNode", s)}} }

func Span(s string) Elt {
	sp := Element("span")
	sp.SetInnerHTML(s)
	return sp
}

func Label(s string) Elt {
	l := Element("label")
	l.SetInnerHTML(s)
	return l
}

func Div() Elt { return Element("div") }

type A struct{ Elt }

func Anchor(s string) A {
	a := Element("a")
	a.SetInnerHTML(s)
	return A{a}
}

func (a A) WithText(s string) A {
	a.SetInnerHTML(s)
	return a
}

func Form() Elt { return Element("form") }

type Tbl struct {
	Elt
	thead Elt
	tbody Elt
	tfoot Elt
}

func Table() Tbl {
	thead := Element("thead")
	tbody := Element("tbody")
	tfoot := Element("tfoot")

	return Tbl{
		Elt:   Element("table").WithChildren(thead, tbody, tfoot),
		thead: thead,
		tbody: tbody,
		tfoot: tfoot,
	}
}

func (t Tbl) WithClass(c string) Tbl {
	t.SetClass(c)
	return t
}

func (t Tbl) WithHeader(ns ...Node) Tbl {
	newRow := NodeBase{t.thead.Call("insertRow", -1)}
	for _, n := range ns {
		headerCell := Element("th")
		headerCell.AppendChild(n)
		newRow.AppendChild(headerCell)
	}
	return t
}

func (t Tbl) WithHeaderClass(class string) Tbl {
	t.thead.SetClass(class)
	return t
}

func (t Tbl) AppendRow(ns ...Node) Elt {
	newRow := NodeBase{t.tbody.Call("insertRow", -1)}
	for _, n := range ns {
		NodeBase{newRow.Call("insertCell", -1)}.AppendChild(n)
	}
	return Elt{newRow}
}

type Inp struct{ Elt }

func Input(typ string) Inp {
	return Inp{Element("input").WithAttribute("type", typ)}
}

func (i Inp) WithClass(c string) Inp {
	i.SetClass(c)
	return i
}

func (i Inp) WithPlaceholder(p string) Inp {
	i.Set("placeholder", p)
	return i
}

func (i Inp) WithValue(val string) Inp {
	i.Set("value", val)
	return i
}

func (i Inp) Value() string { return i.Get("value").String() }

func TextInput() Inp { return Input("text") }

func (i Inp) OnKeyUp(flags EventCallbackFlag, fn func(js.Value)) js.Func {
	return i.AddEventListener(flags, "keyup", fn)
}

type Btn struct{ Elt }

func Button(s string) Btn {
	btn := Element("button")
	btn.SetInnerHTML(s)
	return Btn{btn}
}

func (b Btn) WithClass(c string) Btn {
	b.SetClass(c)
	return b
}

func (b Btn) OnClick(flags EventCallbackFlag, fn func(js.Value)) js.Func {
	return b.AddEventListener(flags, "click", fn)
}
