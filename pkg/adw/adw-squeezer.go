// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeSqueezerTransitionType = coreglib.Type(C.adw_squeezer_transition_type_get_type())
	GTypeSqueezer               = coreglib.Type(C.adw_squeezer_get_type())
	GTypeSqueezerPage           = coreglib.Type(C.adw_squeezer_page_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSqueezerTransitionType, F: marshalSqueezerTransitionType},
		coreglib.TypeMarshaler{T: GTypeSqueezer, F: marshalSqueezer},
		coreglib.TypeMarshaler{T: GTypeSqueezerPage, F: marshalSqueezerPage},
	})
}

// SqueezerTransitionType describes the possible transitions in a squeezer
// widget.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
type SqueezerTransitionType C.gint

const (
	// SqueezerTransitionTypeNone: no transition.
	SqueezerTransitionTypeNone SqueezerTransitionType = iota
	// SqueezerTransitionTypeCrossfade: cross-fade.
	SqueezerTransitionTypeCrossfade
)

func marshalSqueezerTransitionType(p uintptr) (interface{}, error) {
	return SqueezerTransitionType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SqueezerTransitionType.
func (s SqueezerTransitionType) String() string {
	switch s {
	case SqueezerTransitionTypeNone:
		return "None"
	case SqueezerTransitionTypeCrossfade:
		return "Crossfade"
	default:
		return fmt.Sprintf("SqueezerTransitionType(%d)", s)
	}
}

// SqueezerOverrides contains methods that are overridable.
type SqueezerOverrides struct {
}

func defaultSqueezerOverrides(v *Squeezer) SqueezerOverrides {
	return SqueezerOverrides{}
}

// Squeezer: best fit container.
//
// <picture> <source srcset="squeezer-wide-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="squeezer-wide.png"
// alt="squeezer-wide"> </picture> <picture> <source
// srcset="squeezer-narrow-dark.png" media="(prefers-color-scheme: dark)"> <img
// src="squeezer-narrow.png" alt="squeezer-narrow"> </picture>
//
// The AdwSqueezer widget is a container which only shows the first of its
// children that fits in the available size. It is convenient to offer different
// widgets to represent the same data with different levels of detail, making
// the widget seem to squeeze itself to fit in the available space.
//
// Transitions between children can be animated as fades. This can be controlled
// with squeezer:transition-type.
//
// # CSS nodes
//
// AdwSqueezer has a single CSS node with name squeezer.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
type Squeezer struct {
	_ [0]func() // equal guard
	gtk.Widget

	*coreglib.Object
	gtk.Orientable
}

var (
	_ gtk.Widgetter     = (*Squeezer)(nil)
	_ coreglib.Objector = (*Squeezer)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Squeezer, *SqueezerClass, SqueezerOverrides](
		GTypeSqueezer,
		initSqueezerClass,
		wrapSqueezer,
		defaultSqueezerOverrides,
	)
}

func initSqueezerClass(gclass unsafe.Pointer, overrides SqueezerOverrides, classInitFunc func(*SqueezerClass)) {
	if classInitFunc != nil {
		class := (*SqueezerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSqueezer(obj *coreglib.Object) *Squeezer {
	return &Squeezer{
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
		Object: obj,
		Orientable: gtk.Orientable{
			Object: obj,
		},
	}
}

func marshalSqueezer(p uintptr) (interface{}, error) {
	return wrapSqueezer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSqueezer creates a new AdwSqueezer.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function returns the following values:
//
//   - squeezer: newly created AdwSqueezer.
//
func NewSqueezer() *Squeezer {
	var _cret *C.GtkWidget // in

	_cret = C.adw_squeezer_new()

	var _squeezer *Squeezer // out

	_squeezer = wrapSqueezer(coreglib.Take(unsafe.Pointer(_cret)))

	return _squeezer
}

// Add adds a child to self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function takes the following parameters:
//
//   - child: widget to add.
//
// The function returns the following values:
//
//   - squeezerPage: squeezerpage for child.
//
func (self *Squeezer) Add(child gtk.Widgetter) *SqueezerPage {
	var _arg0 *C.AdwSqueezer     // out
	var _arg1 *C.GtkWidget       // out
	var _cret *C.AdwSqueezerPage // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_cret = C.adw_squeezer_add(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)

	var _squeezerPage *SqueezerPage // out

	_squeezerPage = wrapSqueezerPage(coreglib.Take(unsafe.Pointer(_cret)))

	return _squeezerPage
}

// AllowNone gets whether to allow squeezing beyond the last child's minimum
// size.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function returns the following values:
//
//   - ok: whether self allows squeezing beyond the last child.
//
func (self *Squeezer) AllowNone() bool {
	var _arg0 *C.AdwSqueezer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_squeezer_get_allow_none(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Homogeneous gets whether all children have the same size for the opposite
// orientation.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function returns the following values:
//
//   - ok: whether self is homogeneous.
//
func (self *Squeezer) Homogeneous() bool {
	var _arg0 *C.AdwSqueezer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_squeezer_get_homogeneous(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InterpolateSize gets whether self interpolates its size when changing the
// visible child.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function returns the following values:
//
//   - ok: whether the size is interpolated.
//
func (self *Squeezer) InterpolateSize() bool {
	var _arg0 *C.AdwSqueezer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_squeezer_get_interpolate_size(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Page returns the squeezerpage object for child.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function takes the following parameters:
//
//   - child of self.
//
// The function returns the following values:
//
//   - squeezerPage: page object for child.
//
func (self *Squeezer) Page(child gtk.Widgetter) *SqueezerPage {
	var _arg0 *C.AdwSqueezer     // out
	var _arg1 *C.GtkWidget       // out
	var _cret *C.AdwSqueezerPage // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_cret = C.adw_squeezer_get_page(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)

	var _squeezerPage *SqueezerPage // out

	_squeezerPage = wrapSqueezerPage(coreglib.Take(unsafe.Pointer(_cret)))

	return _squeezerPage
}

// Pages returns a gio.ListModel that contains the pages of self.
//
// This can be used to keep an up-to-date view. The model also implements
// gtk.SelectionModel and can be used to track the visible page.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function returns the following values:
//
//   - selectionModel: GtkSelectionModel for the squeezer's children.
//
func (self *Squeezer) Pages() *gtk.SelectionModel {
	var _arg0 *C.AdwSqueezer       // out
	var _cret *C.GtkSelectionModel // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_squeezer_get_pages(_arg0)
	runtime.KeepAlive(self)

	var _selectionModel *gtk.SelectionModel // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_selectionModel = &gtk.SelectionModel{
			ListModel: gio.ListModel{
				Object: obj,
			},
		}
	}

	return _selectionModel
}

// SwitchThresholdPolicy gets the switch threshold policy for self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function returns the following values:
//
//   - foldThresholdPolicy: fold threshold policy.
//
func (self *Squeezer) SwitchThresholdPolicy() FoldThresholdPolicy {
	var _arg0 *C.AdwSqueezer           // out
	var _cret C.AdwFoldThresholdPolicy // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_squeezer_get_switch_threshold_policy(_arg0)
	runtime.KeepAlive(self)

	var _foldThresholdPolicy FoldThresholdPolicy // out

	_foldThresholdPolicy = FoldThresholdPolicy(_cret)

	return _foldThresholdPolicy
}

// TransitionDuration gets the transition animation duration for self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function returns the following values:
//
//   - guint: transition duration, in milliseconds.
//
func (self *Squeezer) TransitionDuration() uint {
	var _arg0 *C.AdwSqueezer // out
	var _cret C.guint        // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_squeezer_get_transition_duration(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// TransitionRunning gets whether a transition is currently running for self.
//
// If a transition is impossible, the property value will be set to TRUE and
// then immediately to FALSE, so it's possible to rely on its notifications to
// know that a transition has happened.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function returns the following values:
//
//   - ok: whether a transition is currently running.
//
func (self *Squeezer) TransitionRunning() bool {
	var _arg0 *C.AdwSqueezer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_squeezer_get_transition_running(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TransitionType gets the type of animation used for transitions between
// children in self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function returns the following values:
//
//   - squeezerTransitionType: current transition type of self.
//
func (self *Squeezer) TransitionType() SqueezerTransitionType {
	var _arg0 *C.AdwSqueezer              // out
	var _cret C.AdwSqueezerTransitionType // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_squeezer_get_transition_type(_arg0)
	runtime.KeepAlive(self)

	var _squeezerTransitionType SqueezerTransitionType // out

	_squeezerTransitionType = SqueezerTransitionType(_cret)

	return _squeezerTransitionType
}

// VisibleChild gets the currently visible child of self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function returns the following values:
//
//   - widget (optional): visible child.
//
func (self *Squeezer) VisibleChild() gtk.Widgetter {
	var _arg0 *C.AdwSqueezer // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_squeezer_get_visible_child(_arg0)
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

// XAlign gets the horizontal alignment, from 0 (start) to 1 (end).
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function returns the following values:
//
//   - gfloat: alignment value.
//
func (self *Squeezer) XAlign() float32 {
	var _arg0 *C.AdwSqueezer // out
	var _cret C.float        // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_squeezer_get_xalign(_arg0)
	runtime.KeepAlive(self)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// YAlign gets the vertical alignment, from 0 (top) to 1 (bottom).
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function returns the following values:
//
//   - gfloat: alignment value.
//
func (self *Squeezer) YAlign() float32 {
	var _arg0 *C.AdwSqueezer // out
	var _cret C.float        // in

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_squeezer_get_yalign(_arg0)
	runtime.KeepAlive(self)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Remove removes a child widget from self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function takes the following parameters:
//
//   - child to remove.
//
func (self *Squeezer) Remove(child gtk.Widgetter) {
	var _arg0 *C.AdwSqueezer // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.adw_squeezer_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetAllowNone sets whether to allow squeezing beyond the last child's minimum
// size.
//
// If set to TRUE, the squeezer can shrink to the point where no child can be
// shown. This is functionally equivalent to appending a widget with 0×0 minimum
// size.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function takes the following parameters:
//
//   - allowNone: whether self allows squeezing beyond the last child.
//
func (self *Squeezer) SetAllowNone(allowNone bool) {
	var _arg0 *C.AdwSqueezer // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if allowNone {
		_arg1 = C.TRUE
	}

	C.adw_squeezer_set_allow_none(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(allowNone)
}

// SetHomogeneous sets whether all children have the same size for the opposite
// orientation.
//
// For example, if a squeezer is horizontal and is homogeneous, it will request
// the same height for all its children. If it isn't, the squeezer may change
// size when a different child becomes visible.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function takes the following parameters:
//
//   - homogeneous: whether self is homogeneous.
//
func (self *Squeezer) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.AdwSqueezer // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.adw_squeezer_set_homogeneous(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(homogeneous)
}

// SetInterpolateSize sets whether self interpolates its size when changing the
// visible child.
//
// If TRUE, the squeezer will interpolate its size between the one of the
// previous visible child and the one of the new visible child, according to
// the set transition duration and the orientation, e.g. if the squeezer is
// horizontal, it will interpolate the its height.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function takes the following parameters:
//
//   - interpolateSize: whether to interpolate the size.
//
func (self *Squeezer) SetInterpolateSize(interpolateSize bool) {
	var _arg0 *C.AdwSqueezer // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if interpolateSize {
		_arg1 = C.TRUE
	}

	C.adw_squeezer_set_interpolate_size(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(interpolateSize)
}

// SetSwitchThresholdPolicy sets the switch threshold policy for self.
//
// Determines when the squeezer will switch children.
//
// If set to ADW_FOLD_THRESHOLD_POLICY_MINIMUM, it will only switch when the
// visible child cannot fit anymore. With ADW_FOLD_THRESHOLD_POLICY_NATURAL,
// it will switch as soon as the visible child doesn't get their natural size.
//
// This can be useful if you have a long ellipsizing label and want to let it
// ellipsize instead of immediately switching.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function takes the following parameters:
//
//   - policy to use.
//
func (self *Squeezer) SetSwitchThresholdPolicy(policy FoldThresholdPolicy) {
	var _arg0 *C.AdwSqueezer           // out
	var _arg1 C.AdwFoldThresholdPolicy // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.AdwFoldThresholdPolicy(policy)

	C.adw_squeezer_set_switch_threshold_policy(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(policy)
}

// SetTransitionDuration sets the transition animation duration for self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function takes the following parameters:
//
//   - duration: new duration, in milliseconds.
//
func (self *Squeezer) SetTransitionDuration(duration uint) {
	var _arg0 *C.AdwSqueezer // out
	var _arg1 C.guint        // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(duration)

	C.adw_squeezer_set_transition_duration(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(duration)
}

// SetTransitionType sets the type of animation used for transitions between
// children in self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function takes the following parameters:
//
//   - transition: new transition type.
//
func (self *Squeezer) SetTransitionType(transition SqueezerTransitionType) {
	var _arg0 *C.AdwSqueezer              // out
	var _arg1 C.AdwSqueezerTransitionType // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.AdwSqueezerTransitionType(transition)

	C.adw_squeezer_set_transition_type(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(transition)
}

// SetXAlign sets the horizontal alignment, from 0 (start) to 1 (end).
//
// This affects the children allocation during transitions, when they exceed the
// size of the squeezer.
//
// For example, 0.5 means the child will be centered, 0 means it will keep the
// start side aligned and overflow the end side, and 1 means the opposite.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function takes the following parameters:
//
//   - xalign: new alignment value.
//
func (self *Squeezer) SetXAlign(xalign float32) {
	var _arg0 *C.AdwSqueezer // out
	var _arg1 C.float        // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.float(xalign)

	C.adw_squeezer_set_xalign(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(xalign)
}

// SetYAlign sets the vertical alignment, from 0 (top) to 1 (bottom).
//
// This affects the children allocation during transitions, when they exceed the
// size of the squeezer.
//
// For example, 0.5 means the child will be centered, 0 means it will keep the
// top side aligned and overflow the bottom side, and 1 means the opposite.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function takes the following parameters:
//
//   - yalign: new alignment value.
//
func (self *Squeezer) SetYAlign(yalign float32) {
	var _arg0 *C.AdwSqueezer // out
	var _arg1 C.float        // out

	_arg0 = (*C.AdwSqueezer)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.float(yalign)

	C.adw_squeezer_set_yalign(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(yalign)
}

// SqueezerPageOverrides contains methods that are overridable.
type SqueezerPageOverrides struct {
}

func defaultSqueezerPageOverrides(v *SqueezerPage) SqueezerPageOverrides {
	return SqueezerPageOverrides{}
}

// SqueezerPage: auxiliary class used by squeezer.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
type SqueezerPage struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*SqueezerPage)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*SqueezerPage, *SqueezerPageClass, SqueezerPageOverrides](
		GTypeSqueezerPage,
		initSqueezerPageClass,
		wrapSqueezerPage,
		defaultSqueezerPageOverrides,
	)
}

func initSqueezerPageClass(gclass unsafe.Pointer, overrides SqueezerPageOverrides, classInitFunc func(*SqueezerPageClass)) {
	if classInitFunc != nil {
		class := (*SqueezerPageClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSqueezerPage(obj *coreglib.Object) *SqueezerPage {
	return &SqueezerPage{
		Object: obj,
	}
}

func marshalSqueezerPage(p uintptr) (interface{}, error) {
	return wrapSqueezerPage(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Child returns the squeezer child to which self belongs.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function returns the following values:
//
//   - widget: child to which self belongs.
//
func (self *SqueezerPage) Child() gtk.Widgetter {
	var _arg0 *C.AdwSqueezerPage // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.AdwSqueezerPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_squeezer_page_get_child(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

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

	return _widget
}

// Enabled gets whether self is enabled.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function returns the following values:
//
//   - ok: whether self is enabled.
//
func (self *SqueezerPage) Enabled() bool {
	var _arg0 *C.AdwSqueezerPage // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwSqueezerPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_squeezer_page_get_enabled(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetEnabled sets whether self is enabled.
//
// If a child is disabled, it will be ignored when looking for the child fitting
// the available size best.
//
// This allows to programmatically and prematurely hide a child even if it fits
// in the available space.
//
// This can be used e.g. to ensure a certain child is hidden below a certain
// window width, or any other constraint you find suitable.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwsqueezer).
//
// The function takes the following parameters:
//
//   - enabled: whether self is enabled.
//
func (self *SqueezerPage) SetEnabled(enabled bool) {
	var _arg0 *C.AdwSqueezerPage // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.AdwSqueezerPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if enabled {
		_arg1 = C.TRUE
	}

	C.adw_squeezer_page_set_enabled(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(enabled)
}

// SqueezerClass: instance of this type is always passed by reference.
type SqueezerClass struct {
	*squeezerClass
}

// squeezerClass is the struct that's finalized.
type squeezerClass struct {
	native *C.AdwSqueezerClass
}

func (s *SqueezerClass) ParentClass() *gtk.WidgetClass {
	valptr := &s.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// SqueezerPageClass: instance of this type is always passed by reference.
type SqueezerPageClass struct {
	*squeezerPageClass
}

// squeezerPageClass is the struct that's finalized.
type squeezerPageClass struct {
	native *C.AdwSqueezerPageClass
}
