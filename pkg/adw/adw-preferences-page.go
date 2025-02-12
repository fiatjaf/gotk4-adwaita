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
	GTypePreferencesPage = coreglib.Type(C.adw_preferences_page_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePreferencesPage, F: marshalPreferencesPage},
	})
}

// PreferencesPageOverrides contains methods that are overridable.
type PreferencesPageOverrides struct {
}

func defaultPreferencesPageOverrides(v *PreferencesPage) PreferencesPageOverrides {
	return PreferencesPageOverrides{}
}

// PreferencesPage: page from preferencesdialog.
//
// <picture> <source srcset="preferences-page-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="preferences-page.png"
// alt="preferences-page"> </picture>
//
// The AdwPreferencesPage widget gathers preferences groups into a single page
// of a preferences window.
//
// # CSS nodes
//
// AdwPreferencesPage has a single CSS node with name preferencespage.
//
// # Accessibility
//
// AdwPreferencesPage uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type PreferencesPage struct {
	_ [0]func() // equal guard
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*PreferencesPage)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*PreferencesPage, *PreferencesPageClass, PreferencesPageOverrides](
		GTypePreferencesPage,
		initPreferencesPageClass,
		wrapPreferencesPage,
		defaultPreferencesPageOverrides,
	)
}

func initPreferencesPageClass(gclass unsafe.Pointer, overrides PreferencesPageOverrides, classInitFunc func(*PreferencesPageClass)) {
	if classInitFunc != nil {
		class := (*PreferencesPageClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapPreferencesPage(obj *coreglib.Object) *PreferencesPage {
	return &PreferencesPage{
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

func marshalPreferencesPage(p uintptr) (interface{}, error) {
	return wrapPreferencesPage(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPreferencesPage creates a new AdwPreferencesPage.
//
// The function returns the following values:
//
//   - preferencesPage: newly created AdwPreferencesPage.
//
func NewPreferencesPage() *PreferencesPage {
	var _cret *C.GtkWidget // in

	_cret = C.adw_preferences_page_new()

	var _preferencesPage *PreferencesPage // out

	_preferencesPage = wrapPreferencesPage(coreglib.Take(unsafe.Pointer(_cret)))

	return _preferencesPage
}

// Add adds a preferences group to self.
//
// The function takes the following parameters:
//
//   - group to add.
//
func (self *PreferencesPage) Add(group *PreferencesGroup) {
	var _arg0 *C.AdwPreferencesPage  // out
	var _arg1 *C.AdwPreferencesGroup // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.AdwPreferencesGroup)(unsafe.Pointer(coreglib.InternObject(group).Native()))

	C.adw_preferences_page_add(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(group)
}

// Description gets the description of self.
//
// The function returns the following values:
//
//   - utf8: description of self.
//
func (self *PreferencesPage) Description() string {
	var _arg0 *C.AdwPreferencesPage // out
	var _cret *C.char               // in

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_preferences_page_get_description(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IconName gets the icon name for self.
//
// The function returns the following values:
//
//   - utf8 (optional): icon name for self.
//
func (self *PreferencesPage) IconName() string {
	var _arg0 *C.AdwPreferencesPage // out
	var _cret *C.char               // in

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_preferences_page_get_icon_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Name gets the name of self.
//
// The function returns the following values:
//
//   - utf8 (optional): name of self.
//
func (self *PreferencesPage) Name() string {
	var _arg0 *C.AdwPreferencesPage // out
	var _cret *C.char               // in

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_preferences_page_get_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Title gets the title of self.
//
// The function returns the following values:
//
//   - utf8: title of self.
//
func (self *PreferencesPage) Title() string {
	var _arg0 *C.AdwPreferencesPage // out
	var _cret *C.char               // in

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_preferences_page_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UseUnderline gets whether an embedded underline in the title indicates a
// mnemonic.
//
// The function returns the following values:
//
//   - ok: whether an embedded underline in the title indicates a mnemonic.
//
func (self *PreferencesPage) UseUnderline() bool {
	var _arg0 *C.AdwPreferencesPage // out
	var _cret C.gboolean            // in

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_preferences_page_get_use_underline(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Remove removes a group from self.
//
// The function takes the following parameters:
//
//   - group to remove.
//
func (self *PreferencesPage) Remove(group *PreferencesGroup) {
	var _arg0 *C.AdwPreferencesPage  // out
	var _arg1 *C.AdwPreferencesGroup // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.AdwPreferencesGroup)(unsafe.Pointer(coreglib.InternObject(group).Native()))

	C.adw_preferences_page_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(group)
}

// ScrollToTop scrolls the scrolled window of self to the top.
func (self *PreferencesPage) ScrollToTop() {
	var _arg0 *C.AdwPreferencesPage // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.adw_preferences_page_scroll_to_top(_arg0)
	runtime.KeepAlive(self)
}

// SetDescription sets the description of self.
//
// The description is displayed at the top of the page.
//
// The function takes the following parameters:
//
//   - description: description.
//
func (self *PreferencesPage) SetDescription(description string) {
	var _arg0 *C.AdwPreferencesPage // out
	var _arg1 *C.char               // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(description)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_preferences_page_set_description(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(description)
}

// SetIconName sets the icon name for self.
//
// The function takes the following parameters:
//
//   - iconName (optional): icon name.
//
func (self *PreferencesPage) SetIconName(iconName string) {
	var _arg0 *C.AdwPreferencesPage // out
	var _arg1 *C.char               // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if iconName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_preferences_page_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)
}

// SetName sets the name of self.
//
// The function takes the following parameters:
//
//   - name (optional): name.
//
func (self *PreferencesPage) SetName(name string) {
	var _arg0 *C.AdwPreferencesPage // out
	var _arg1 *C.char               // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if name != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_preferences_page_set_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
}

// SetTitle sets the title of self.
//
// The function takes the following parameters:
//
//   - title: title.
//
func (self *PreferencesPage) SetTitle(title string) {
	var _arg0 *C.AdwPreferencesPage // out
	var _arg1 *C.char               // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_preferences_page_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}

// SetUseUnderline sets whether an embedded underline in the title indicates a
// mnemonic.
//
// The function takes the following parameters:
//
//   - useUnderline: TRUE if underlines in the text indicate mnemonics.
//
func (self *PreferencesPage) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.AdwPreferencesPage // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.adw_preferences_page_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(useUnderline)
}

// PreferencesPageClass: instance of this type is always passed by reference.
type PreferencesPageClass struct {
	*preferencesPageClass
}

// preferencesPageClass is the struct that's finalized.
type preferencesPageClass struct {
	native *C.AdwPreferencesPageClass
}

// ParentClass: parent class.
func (p *PreferencesPageClass) ParentClass() *gtk.WidgetClass {
	valptr := &p.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
