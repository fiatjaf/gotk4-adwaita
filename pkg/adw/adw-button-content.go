// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeButtonContent = coreglib.Type(C.adw_button_content_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeButtonContent, F: marshalButtonContent},
	})
}

// ButtonContentOverrides contains methods that are overridable.
type ButtonContentOverrides struct {
}

func defaultButtonContentOverrides(v *ButtonContent) ButtonContentOverrides {
	return ButtonContentOverrides{}
}

// ButtonContent: helper widget for creating buttons.
//
// <picture> <source srcset="button-content-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="button-content.png"
// alt="button-content"> </picture>
//
// AdwButtonContent is a box-like widget with an icon and a label.
//
// It's intended to be used as a direct child of gtk.Button, gtk.MenuButton or
// splitbutton, when they need to have both an icon and a label, as follows:
//
//    <object class="GtkButton">
//      <property name="child">
//        <object class="AdwButtonContent">
//          <property name="icon-name">document-open-symbolic</property>
//          <property name="label" translatable="yes">_Open</property>
//          <property name="use-underline">True</property>
//        </object>
//      </property>
//    </object>
//
// AdwButtonContent handles style classes and connecting the mnemonic to the
// button automatically.
//
// CSS nodes
//
//    buttoncontent
//    ╰── box
//        ├── image
//        ╰── label
//
// AdwButtonContent's CSS node is called buttoncontent. It contains a box
// subnode that serves as a container for the image and label nodes.
//
// When inside a GtkButton or AdwSplitButton, the button will receive the
// .image-text-button style class. When inside a GtkMenuButton, the internal
// GtkButton will receive it instead.
//
// # Accessibility
//
// AdwButtonContent uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type ButtonContent struct {
	_ [0]func() // equal guard
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*ButtonContent)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ButtonContent, *ButtonContentClass, ButtonContentOverrides](
		GTypeButtonContent,
		initButtonContentClass,
		wrapButtonContent,
		defaultButtonContentOverrides,
	)
}

func initButtonContentClass(gclass unsafe.Pointer, overrides ButtonContentOverrides, classInitFunc func(*ButtonContentClass)) {
	if classInitFunc != nil {
		class := (*ButtonContentClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapButtonContent(obj *coreglib.Object) *ButtonContent {
	return &ButtonContent{
		Widget: gtk.Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: gtk.Accessible{
				Object: obj,
			},
			Buildable: gtk.Buildable{
				Object: obj,
			},
			ConstraintTarget: gtk.ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalButtonContent(p uintptr) (interface{}, error) {
	return wrapButtonContent(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewButtonContent creates a new AdwButtonContent.
//
// The function returns the following values:
//
//   - buttonContent: new created AdwButtonContent.
//
func NewButtonContent() *ButtonContent {
	var _cret *C.GtkWidget // in

	_cret = C.adw_button_content_new()

	var _buttonContent *ButtonContent // out

	_buttonContent = wrapButtonContent(coreglib.Take(unsafe.Pointer(_cret)))

	return _buttonContent
}

// CanShrink gets whether the button can be smaller than the natural size of its
// contents.
//
// The function returns the following values:
//
//   - ok: whether the button can shrink.
//
func (self *ButtonContent) CanShrink() bool {
	var _arg0 *C.AdwButtonContent // out
	var _cret C.gboolean          // in

	_arg0 = (*C.AdwButtonContent)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_button_content_get_can_shrink(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IconName gets the name of the displayed icon.
//
// The function returns the following values:
//
//   - utf8: icon name.
//
func (self *ButtonContent) IconName() string {
	var _arg0 *C.AdwButtonContent // out
	var _cret *C.char             // in

	_arg0 = (*C.AdwButtonContent)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_button_content_get_icon_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Label gets the displayed label.
//
// The function returns the following values:
//
//   - utf8: label.
//
func (self *ButtonContent) Label() string {
	var _arg0 *C.AdwButtonContent // out
	var _cret *C.char             // in

	_arg0 = (*C.AdwButtonContent)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_button_content_get_label(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UseUnderline gets whether an underline in the text indicates a mnemonic.
//
// The function returns the following values:
//
//   - ok: whether an underline in the text indicates a mnemonic.
//
func (self *ButtonContent) UseUnderline() bool {
	var _arg0 *C.AdwButtonContent // out
	var _cret C.gboolean          // in

	_arg0 = (*C.AdwButtonContent)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_button_content_get_use_underline(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCanShrink sets whether the button can be smaller than the natural size of
// its contents.
//
// If set to TRUE, the label will ellipsize.
//
// See gtk.Button.SetCanShrink().
//
// The function takes the following parameters:
//
//   - canShrink: whether the button can shrink.
//
func (self *ButtonContent) SetCanShrink(canShrink bool) {
	var _arg0 *C.AdwButtonContent // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.AdwButtonContent)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if canShrink {
		_arg1 = C.TRUE
	}

	C.adw_button_content_set_can_shrink(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(canShrink)
}

// SetIconName sets the name of the displayed icon.
//
// If empty, the icon is not shown.
//
// The function takes the following parameters:
//
//   - iconName: new icon name.
//
func (self *ButtonContent) SetIconName(iconName string) {
	var _arg0 *C.AdwButtonContent // out
	var _arg1 *C.char             // out

	_arg0 = (*C.AdwButtonContent)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_button_content_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)
}

// SetLabel sets the displayed label.
//
// The function takes the following parameters:
//
//   - label: new label.
//
func (self *ButtonContent) SetLabel(label string) {
	var _arg0 *C.AdwButtonContent // out
	var _arg1 *C.char             // out

	_arg0 = (*C.AdwButtonContent)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_button_content_set_label(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(label)
}

// SetUseUnderline sets whether an underline in the text indicates a mnemonic.
//
// The mnemonic can be used to activate the parent button.
//
// See buttoncontent:label.
//
// The function takes the following parameters:
//
//   - useUnderline: whether an underline in the text indicates a mnemonic.
//
func (self *ButtonContent) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.AdwButtonContent // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.AdwButtonContent)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.adw_button_content_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(useUnderline)
}

// ButtonContentClass: instance of this type is always passed by reference.
type ButtonContentClass struct {
	*buttonContentClass
}

// buttonContentClass is the struct that's finalized.
type buttonContentClass struct {
	native *C.AdwButtonContentClass
}

func (b *ButtonContentClass) ParentClass() *gtk.WidgetClass {
	valptr := &b.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
