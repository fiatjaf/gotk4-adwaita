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
	GTypePreferencesGroup = coreglib.Type(C.adw_preferences_group_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePreferencesGroup, F: marshalPreferencesGroup},
	})
}

// PreferencesGroupOverrides contains methods that are overridable.
type PreferencesGroupOverrides struct {
}

func defaultPreferencesGroupOverrides(v *PreferencesGroup) PreferencesGroupOverrides {
	return PreferencesGroupOverrides{}
}

// PreferencesGroup: group of preference rows.
//
// <picture> <source srcset="preferences-group-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="preferences-group.png"
// alt="preferences-group"> </picture>
//
// An AdwPreferencesGroup represents a group or tightly related preferences,
// which in turn are represented by preferencesrow.
//
// To summarize the role of the preferences it gathers, a group can have both a
// title and a description. The title will be used by preferencesdialog to let
// the user look for a preference.
//
// # AdwPreferencesGroup as GtkBuildable
//
// The AdwPreferencesGroup implementation of the gtk.Buildable interface
// supports adding preferencesrows to the list by omitting "type". If "type"
// is omitted and the widget isn't a preferencesrow the child is added to a box
// below the list.
//
// When the "type" attribute of a child is header-suffix, the child is set as
// the suffix on the end of the title and description.
//
// # CSS nodes
//
// AdwPreferencesGroup has a single CSS node with name preferencesgroup.
//
// # Accessibility
//
// AdwPreferencesGroup uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type PreferencesGroup struct {
	_ [0]func() // equal guard
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*PreferencesGroup)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*PreferencesGroup, *PreferencesGroupClass, PreferencesGroupOverrides](
		GTypePreferencesGroup,
		initPreferencesGroupClass,
		wrapPreferencesGroup,
		defaultPreferencesGroupOverrides,
	)
}

func initPreferencesGroupClass(gclass unsafe.Pointer, overrides PreferencesGroupOverrides, classInitFunc func(*PreferencesGroupClass)) {
	if classInitFunc != nil {
		class := (*PreferencesGroupClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapPreferencesGroup(obj *coreglib.Object) *PreferencesGroup {
	return &PreferencesGroup{
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

func marshalPreferencesGroup(p uintptr) (interface{}, error) {
	return wrapPreferencesGroup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPreferencesGroup creates a new AdwPreferencesGroup.
//
// The function returns the following values:
//
//   - preferencesGroup: newly created AdwPreferencesGroup.
//
func NewPreferencesGroup() *PreferencesGroup {
	var _cret *C.GtkWidget // in

	_cret = C.adw_preferences_group_new()

	var _preferencesGroup *PreferencesGroup // out

	_preferencesGroup = wrapPreferencesGroup(coreglib.Take(unsafe.Pointer(_cret)))

	return _preferencesGroup
}

// Add adds a child to self.
//
// The function takes the following parameters:
//
//   - child: widget to add.
//
func (self *PreferencesGroup) Add(child gtk.Widgetter) {
	var _arg0 *C.AdwPreferencesGroup // out
	var _arg1 *C.GtkWidget           // out

	_arg0 = (*C.AdwPreferencesGroup)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.adw_preferences_group_add(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// Description gets the description of self.
//
// The function returns the following values:
//
//   - utf8 (optional): description of self.
//
func (self *PreferencesGroup) Description() string {
	var _arg0 *C.AdwPreferencesGroup // out
	var _cret *C.char                // in

	_arg0 = (*C.AdwPreferencesGroup)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_preferences_group_get_description(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// HeaderSuffix gets the suffix for self's header.
//
// The function returns the following values:
//
//   - widget (optional): suffix for self's header.
//
func (self *PreferencesGroup) HeaderSuffix() gtk.Widgetter {
	var _arg0 *C.AdwPreferencesGroup // out
	var _cret *C.GtkWidget           // in

	_arg0 = (*C.AdwPreferencesGroup)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_preferences_group_get_header_suffix(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gtk.Widgetter)
				return ok
			})
			rv, ok := casted.(gtk.Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Title gets the title of self.
//
// The function returns the following values:
//
//   - utf8: title of self.
//
func (self *PreferencesGroup) Title() string {
	var _arg0 *C.AdwPreferencesGroup // out
	var _cret *C.char                // in

	_arg0 = (*C.AdwPreferencesGroup)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_preferences_group_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Remove removes a child from self.
//
// The function takes the following parameters:
//
//   - child to remove.
//
func (self *PreferencesGroup) Remove(child gtk.Widgetter) {
	var _arg0 *C.AdwPreferencesGroup // out
	var _arg1 *C.GtkWidget           // out

	_arg0 = (*C.AdwPreferencesGroup)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.adw_preferences_group_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetDescription sets the description for self.
//
// The function takes the following parameters:
//
//   - description (optional): description.
//
func (self *PreferencesGroup) SetDescription(description string) {
	var _arg0 *C.AdwPreferencesGroup // out
	var _arg1 *C.char                // out

	_arg0 = (*C.AdwPreferencesGroup)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if description != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(description)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_preferences_group_set_description(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(description)
}

// SetHeaderSuffix sets the suffix for self's header.
//
// Displayed above the list, next to the title and description.
//
// Suffixes are commonly used to show a button or a spinner for the whole group.
//
// The function takes the following parameters:
//
//   - suffix (optional) to set.
//
func (self *PreferencesGroup) SetHeaderSuffix(suffix gtk.Widgetter) {
	var _arg0 *C.AdwPreferencesGroup // out
	var _arg1 *C.GtkWidget           // out

	_arg0 = (*C.AdwPreferencesGroup)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if suffix != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(suffix).Native()))
	}

	C.adw_preferences_group_set_header_suffix(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(suffix)
}

// SetTitle sets the title for self.
//
// The function takes the following parameters:
//
//   - title: title.
//
func (self *PreferencesGroup) SetTitle(title string) {
	var _arg0 *C.AdwPreferencesGroup // out
	var _arg1 *C.char                // out

	_arg0 = (*C.AdwPreferencesGroup)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_preferences_group_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}

// PreferencesGroupClass: instance of this type is always passed by reference.
type PreferencesGroupClass struct {
	*preferencesGroupClass
}

// preferencesGroupClass is the struct that's finalized.
type preferencesGroupClass struct {
	native *C.AdwPreferencesGroupClass
}

// ParentClass: parent class.
func (p *PreferencesGroupClass) ParentClass() *gtk.WidgetClass {
	valptr := &p.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
