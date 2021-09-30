// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"fmt"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_centering_policy_get_type()), F: marshalCenteringPolicy},
		{T: externglib.Type(C.adw_header_bar_get_type()), F: marshalHeaderBarrer},
	})
}

// CenteringPolicy describes title centering behavior of a adw.HeaderBar widget.
type CenteringPolicy int

const (
	// CenteringPolicyLoose: keep the title centered when possible
	CenteringPolicyLoose CenteringPolicy = iota
	// CenteringPolicyStrict: keep the title centered at all cost
	CenteringPolicyStrict
)

func marshalCenteringPolicy(p uintptr) (interface{}, error) {
	return CenteringPolicy(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for CenteringPolicy.
func (c CenteringPolicy) String() string {
	switch c {
	case CenteringPolicyLoose:
		return "Loose"
	case CenteringPolicyStrict:
		return "Strict"
	default:
		return fmt.Sprintf("CenteringPolicy(%d)", c)
	}
}

// HeaderBar: title bar widget.
//
// AdwHeaderBar is similar to gtk.HeaderBar, but provides additional features
// compared to it. Refer to GtkHeaderBar for details.
//
// adw.HeaderBar:centering-policy allows to enforce strict centering of the
// title widget, this is useful for adw.ViewSwitcherTitle.
//
// adw.HeaderBar:show-start-title-buttons and
// adw.HeaderBar:show-end-title-buttons allow to easily create split header bar
// layouts using adw.Leaflet, as follows:
//
//    <object class="AdwLeaflet" id="leaflet">
//      <child>
//        <object class="GtkBox">
//          <property name="orientation">vertical</property>
//          <object class="AdwHeaderBar">
//            <binding name="show-end-title-buttons">
//              <lookup name="folded">leaflet</lookup>
//            </binding>
//          </object>
//          ...
//        </object>
//      </child>
//      ...
//      <child>
//        <object class="GtkBox">
//          <property name="orientation">vertical</property>
//          <object class="AdwHeaderBar">
//            <binding name="show-start-title-buttons">
//              <lookup name="folded">leaflet</lookup>
//            </binding>
//          </object>
//          ...
//        </object>
//      </child>
//    </object>
//
//
// CSS nodes
//
//    headerbar
//    ╰── windowhandle
//        ╰── box
//            ├── widget
//            │   ╰── box.start
//            │       ├── windowcontrols.start
//            │       ╰── [other children]
//            ├── [Title Widget]
//            ╰── widget
//                ╰── box.end
//                    ├── [other children]
//                    ╰── windowcontrols.end
//
//
// AdwHeaderBar's CSS node is called headerbar. It contains a windowhandle
// subnode, which contains a box subnode, which contains two widget subnodes at
// the start and end of the header bar, each of which contains a box subnode
// with the .start and .end style classes respectively, as well as a center node
// that represents the title.
//
// Each of the boxes contains a windowcontrols subnode, see gtk.WindowControls
// for details, as well as other children.
//
//
// Accessibility
//
// AdwHeaderBar uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type HeaderBar struct {
	gtk.Widget
}

func wrapHeaderBar(obj *externglib.Object) *HeaderBar {
	return &HeaderBar{
		Widget: gtk.Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: gtk.Accessible{
				Object: obj,
			},
			Buildable: gtk.Buildable{
				Object: obj,
			},
			ConstraintTarget: gtk.ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalHeaderBarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapHeaderBar(obj), nil
}

// NewHeaderBar creates a new AdwHeaderBar.
func NewHeaderBar() *HeaderBar {
	var _cret *C.GtkWidget // in

	_cret = C.adw_header_bar_new()

	var _headerBar *HeaderBar // out

	_headerBar = wrapHeaderBar(externglib.Take(unsafe.Pointer(_cret)))

	return _headerBar
}

// CenteringPolicy gets the policy for aligning the center widget.
func (self *HeaderBar) CenteringPolicy() CenteringPolicy {
	var _arg0 *C.AdwHeaderBar      // out
	var _cret C.AdwCenteringPolicy // in

	_arg0 = (*C.AdwHeaderBar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_header_bar_get_centering_policy(_arg0)
	runtime.KeepAlive(self)

	var _centeringPolicy CenteringPolicy // out

	_centeringPolicy = CenteringPolicy(_cret)

	return _centeringPolicy
}

// DecorationLayout gets the decoration layout for self.
func (self *HeaderBar) DecorationLayout() string {
	var _arg0 *C.AdwHeaderBar // out
	var _cret *C.char         // in

	_arg0 = (*C.AdwHeaderBar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_header_bar_get_decoration_layout(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ShowEndTitleButtons gets whether to show title buttons at the end of self.
func (self *HeaderBar) ShowEndTitleButtons() bool {
	var _arg0 *C.AdwHeaderBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AdwHeaderBar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_header_bar_get_show_end_title_buttons(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowStartTitleButtons gets whether to show title buttons at the start of
// self.
func (self *HeaderBar) ShowStartTitleButtons() bool {
	var _arg0 *C.AdwHeaderBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AdwHeaderBar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_header_bar_get_show_start_title_buttons(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TitleWidget gets the title widget widget of self.
func (self *HeaderBar) TitleWidget() gtk.Widgetter {
	var _arg0 *C.AdwHeaderBar // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.AdwHeaderBar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_header_bar_get_title_widget(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(gtk.Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// PackEnd adds child to self, packed with reference to the end of self.
func (self *HeaderBar) PackEnd(child gtk.Widgetter) {
	var _arg0 *C.AdwHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwHeaderBar)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.adw_header_bar_pack_end(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// PackStart adds child to self, packed with reference to the start of the self.
func (self *HeaderBar) PackStart(child gtk.Widgetter) {
	var _arg0 *C.AdwHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwHeaderBar)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.adw_header_bar_pack_start(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// Remove removes a child from self.
//
// The child must have been added with adw.HeaderBar.PackStart(),
// adw.HeaderBar.PackEnd() or adw.HeaderBar:title-widget.
func (self *HeaderBar) Remove(child gtk.Widgetter) {
	var _arg0 *C.AdwHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwHeaderBar)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.adw_header_bar_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetCenteringPolicy sets the policy for aligning the center widget.
func (self *HeaderBar) SetCenteringPolicy(centeringPolicy CenteringPolicy) {
	var _arg0 *C.AdwHeaderBar      // out
	var _arg1 C.AdwCenteringPolicy // out

	_arg0 = (*C.AdwHeaderBar)(unsafe.Pointer(self.Native()))
	_arg1 = C.AdwCenteringPolicy(centeringPolicy)

	C.adw_header_bar_set_centering_policy(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(centeringPolicy)
}

// SetDecorationLayout sets the decoration layout for self.
func (self *HeaderBar) SetDecorationLayout(layout string) {
	var _arg0 *C.AdwHeaderBar // out
	var _arg1 *C.char         // out

	_arg0 = (*C.AdwHeaderBar)(unsafe.Pointer(self.Native()))
	if layout != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(layout)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_header_bar_set_decoration_layout(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(layout)
}

// SetShowEndTitleButtons sets whether to show title buttons at the end of self.
func (self *HeaderBar) SetShowEndTitleButtons(setting bool) {
	var _arg0 *C.AdwHeaderBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.AdwHeaderBar)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.adw_header_bar_set_show_end_title_buttons(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetShowStartTitleButtons sets whether to show title buttons at the start of
// self.
func (self *HeaderBar) SetShowStartTitleButtons(setting bool) {
	var _arg0 *C.AdwHeaderBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.AdwHeaderBar)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.adw_header_bar_set_show_start_title_buttons(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetTitleWidget sets the title widget for self.
func (self *HeaderBar) SetTitleWidget(titleWidget gtk.Widgetter) {
	var _arg0 *C.AdwHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwHeaderBar)(unsafe.Pointer(self.Native()))
	if titleWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(titleWidget.Native()))
	}

	C.adw_header_bar_set_title_widget(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(titleWidget)
}
