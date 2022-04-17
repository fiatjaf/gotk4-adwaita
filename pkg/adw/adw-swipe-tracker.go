// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
// extern void _gotk4_adw1_SwipeTracker_ConnectBeginSwipe(gpointer, guintptr);
// extern void _gotk4_adw1_SwipeTracker_ConnectEndSwipe(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_adw1_SwipeTracker_ConnectPrepare(gpointer, AdwNavigationDirection, guintptr);
// extern void _gotk4_adw1_SwipeTracker_ConnectUpdateSwipe(gpointer, gdouble, guintptr);
import "C"

// glib.Type values for adw-swipe-tracker.go.
var GTypeSwipeTracker = externglib.Type(C.adw_swipe_tracker_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeSwipeTracker, F: marshalSwipeTracker},
	})
}

// SwipeTrackerOverrider contains methods that are overridable.
type SwipeTrackerOverrider interface {
}

// SwipeTracker: swipe tracker used in carousel, flap and leaflet.
//
// The AdwSwipeTracker object can be used for implementing widgets with swipe
// gestures. It supports touch-based swipes, pointer dragging, and touchpad
// scrolling.
//
// The widgets will probably want to expose the swipetracker:enabled property.
// If they expect to use horizontal orientation, swipetracker:reversed can be
// used for supporting RTL text direction.
type SwipeTracker struct {
	_ [0]func() // equal guard
	*externglib.Object

	gtk.Orientable
}

var (
	_ externglib.Objector = (*SwipeTracker)(nil)
)

func classInitSwipeTrackerer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSwipeTracker(obj *externglib.Object) *SwipeTracker {
	return &SwipeTracker{
		Object: obj,
		Orientable: gtk.Orientable{
			Object: obj,
		},
	}
}

func marshalSwipeTracker(p uintptr) (interface{}, error) {
	return wrapSwipeTracker(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_adw1_SwipeTracker_ConnectBeginSwipe
func _gotk4_adw1_SwipeTracker_ConnectBeginSwipe(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectBeginSwipe: this signal is emitted right before a swipe will be
// started, after the drag threshold has been passed.
func (self *SwipeTracker) ConnectBeginSwipe(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "begin-swipe", false, unsafe.Pointer(C._gotk4_adw1_SwipeTracker_ConnectBeginSwipe), f)
}

//export _gotk4_adw1_SwipeTracker_ConnectEndSwipe
func _gotk4_adw1_SwipeTracker_ConnectEndSwipe(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(velocity, to float64)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(velocity, to float64))
	}

	var _velocity float64 // out
	var _to float64       // out

	_velocity = float64(arg1)
	_to = float64(arg2)

	f(_velocity, _to)
}

// ConnectEndSwipe: this signal is emitted as soon as the gesture has stopped.
//
// The user is expected to animate the deceleration from the current progress
// value to to with an animation using velocity as the initial velocity,
// provided in pixels per second. springanimation is usually a good fit for
// this.
func (self *SwipeTracker) ConnectEndSwipe(f func(velocity, to float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "end-swipe", false, unsafe.Pointer(C._gotk4_adw1_SwipeTracker_ConnectEndSwipe), f)
}

//export _gotk4_adw1_SwipeTracker_ConnectPrepare
func _gotk4_adw1_SwipeTracker_ConnectPrepare(arg0 C.gpointer, arg1 C.AdwNavigationDirection, arg2 C.guintptr) {
	var f func(direction NavigationDirection)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(direction NavigationDirection))
	}

	var _direction NavigationDirection // out

	_direction = NavigationDirection(arg1)

	f(_direction)
}

// ConnectPrepare: this signal is emitted when a possible swipe is detected.
//
// The direction value can be used to restrict the swipe to a certain direction.
func (self *SwipeTracker) ConnectPrepare(f func(direction NavigationDirection)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "prepare", false, unsafe.Pointer(C._gotk4_adw1_SwipeTracker_ConnectPrepare), f)
}

//export _gotk4_adw1_SwipeTracker_ConnectUpdateSwipe
func _gotk4_adw1_SwipeTracker_ConnectUpdateSwipe(arg0 C.gpointer, arg1 C.gdouble, arg2 C.guintptr) {
	var f func(progress float64)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(progress float64))
	}

	var _progress float64 // out

	_progress = float64(arg1)

	f(_progress)
}

// ConnectUpdateSwipe: this signal is emitted every time the progress value
// changes.
func (self *SwipeTracker) ConnectUpdateSwipe(f func(progress float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "update-swipe", false, unsafe.Pointer(C._gotk4_adw1_SwipeTracker_ConnectUpdateSwipe), f)
}

// NewSwipeTracker creates a new AdwSwipeTracker for widget.
//
// The function takes the following parameters:
//
//    - swipeable: widget to add the tracker on.
//
// The function returns the following values:
//
//    - swipeTracker: newly created AdwSwipeTracker.
//
func NewSwipeTracker(swipeable Swipeabler) *SwipeTracker {
	var _arg1 *C.AdwSwipeable    // out
	var _cret *C.AdwSwipeTracker // in

	_arg1 = (*C.AdwSwipeable)(unsafe.Pointer(externglib.InternObject(swipeable).Native()))

	_cret = C.adw_swipe_tracker_new(_arg1)
	runtime.KeepAlive(swipeable)

	var _swipeTracker *SwipeTracker // out

	_swipeTracker = wrapSwipeTracker(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _swipeTracker
}

// AllowLongSwipes gets whether to allow swiping for more than one snap point at
// a time.
//
// The function returns the following values:
//
//    - ok: whether long swipes are allowed.
//
func (self *SwipeTracker) AllowLongSwipes() bool {
	var _arg0 *C.AdwSwipeTracker // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_swipe_tracker_get_allow_long_swipes(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AllowMouseDrag gets whether self can be dragged with mouse pointer.
//
// The function returns the following values:
//
//    - ok: whether mouse dragging is allowed.
//
func (self *SwipeTracker) AllowMouseDrag() bool {
	var _arg0 *C.AdwSwipeTracker // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_swipe_tracker_get_allow_mouse_drag(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Enabled gets whether self is enabled.
//
// The function returns the following values:
//
//    - ok: whether self is enabled.
//
func (self *SwipeTracker) Enabled() bool {
	var _arg0 *C.AdwSwipeTracker // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_swipe_tracker_get_enabled(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Reversed gets whether self is reversing the swipe direction.
//
// The function returns the following values:
//
//    - ok: whether the direction is reversed.
//
func (self *SwipeTracker) Reversed() bool {
	var _arg0 *C.AdwSwipeTracker // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_swipe_tracker_get_reversed(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Swipeable: get the widget self is attached to.
//
// The function returns the following values:
//
//    - swipeable widget.
//
func (self *SwipeTracker) Swipeable() *Swipeable {
	var _arg0 *C.AdwSwipeTracker // out
	var _cret *C.AdwSwipeable    // in

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_swipe_tracker_get_swipeable(_arg0)
	runtime.KeepAlive(self)

	var _swipeable *Swipeable // out

	_swipeable = wrapSwipeable(externglib.Take(unsafe.Pointer(_cret)))

	return _swipeable
}

// SetAllowLongSwipes sets whether to allow swiping for more than one snap point
// at a time.
//
// The function takes the following parameters:
//
//    - allowLongSwipes: whether to allow long swipes.
//
func (self *SwipeTracker) SetAllowLongSwipes(allowLongSwipes bool) {
	var _arg0 *C.AdwSwipeTracker // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if allowLongSwipes {
		_arg1 = C.TRUE
	}

	C.adw_swipe_tracker_set_allow_long_swipes(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(allowLongSwipes)
}

// SetAllowMouseDrag sets whether self can be dragged with mouse pointer.
//
// The function takes the following parameters:
//
//    - allowMouseDrag: whether to allow mouse dragging.
//
func (self *SwipeTracker) SetAllowMouseDrag(allowMouseDrag bool) {
	var _arg0 *C.AdwSwipeTracker // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if allowMouseDrag {
		_arg1 = C.TRUE
	}

	C.adw_swipe_tracker_set_allow_mouse_drag(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(allowMouseDrag)
}

// SetEnabled sets whether self is enabled.
//
// The function takes the following parameters:
//
//    - enabled: whether self is enabled.
//
func (self *SwipeTracker) SetEnabled(enabled bool) {
	var _arg0 *C.AdwSwipeTracker // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if enabled {
		_arg1 = C.TRUE
	}

	C.adw_swipe_tracker_set_enabled(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(enabled)
}

// SetReversed sets whether to reverse the swipe direction.
//
// The function takes the following parameters:
//
//    - reversed: whether to reverse the swipe direction.
//
func (self *SwipeTracker) SetReversed(reversed bool) {
	var _arg0 *C.AdwSwipeTracker // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if reversed {
		_arg1 = C.TRUE
	}

	C.adw_swipe_tracker_set_reversed(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(reversed)
}

// ShiftPosition moves the current progress value by delta.
//
// This can be used to adjust the current position if snap points move during
// the gesture.
//
// The function takes the following parameters:
//
//    - delta: position delta.
//
func (self *SwipeTracker) ShiftPosition(delta float64) {
	var _arg0 *C.AdwSwipeTracker // out
	var _arg1 C.double           // out

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.double(delta)

	C.adw_swipe_tracker_shift_position(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(delta)
}
