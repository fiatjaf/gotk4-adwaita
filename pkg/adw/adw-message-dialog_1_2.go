// Code generated by girgen. DO NOT EDIT.

package adw

import (
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
// extern void _gotk4_adw1_MessageDialog_ConnectResponse(gpointer, gchar*, guintptr);
// extern void _gotk4_adw1_MessageDialogClass_response(AdwMessageDialog*, char*);
// void _gotk4_adw1_MessageDialog_virtual_response(void* fnptr, AdwMessageDialog* arg0, char* arg1) {
//   ((void (*)(AdwMessageDialog*, char*))(fnptr))(arg0, arg1);
// };
import "C"

// GType values.
var (
	GTypeMessageDialog = coreglib.Type(C.adw_message_dialog_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMessageDialog, F: marshalMessageDialog},
	})
}

// MessageDialogOverrides contains methods that are overridable.
type MessageDialogOverrides struct {
	// Response emits the messagedialog::response signal with the given response
	// ID.
	//
	// Used to indicate that the user has responded to the dialog in some way.
	//
	// The function takes the following parameters:
	//
	//   - response ID.
	//
	Response func(response string)
}

func defaultMessageDialogOverrides(v *MessageDialog) MessageDialogOverrides {
	return MessageDialogOverrides{
		Response: v.response,
	}
}

// MessageDialog: dialog presenting a message or a question.
//
// <picture> <source srcset="message-dialog-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="message-dialog.png"
// alt="message-dialog"> </picture>
//
// Message dialogs have a heading, a body, an optional child widget, and one or
// multiple responses, each presented as a button.
//
// Each response has a unique string ID, and a button label. Additionally, each
// response can be enabled or disabled, and can have a suggested or destructive
// appearance.
//
// When one of the responses is activated, or the dialog is closed, the
// messagedialog::response signal will be emitted. This signal is detailed,
// and the detail, as well as the response parameter will be set to the ID of
// the activated response, or to the value of the messagedialog:close-response
// property if the dialog had been closed without activating any of the
// responses.
//
// Response buttons can be presented horizontally or vertically depending on
// available space.
//
// When a response is activated, AdwMessageDialog is closed automatically.
//
// An example of using a message dialog:
//
//    GtkWidget *dialog;
//
//    dialog = adw_message_dialog_new (parent, _("Replace File?"), NULL);
//
//    adw_message_dialog_format_body (ADW_MESSAGE_DIALOG (dialog),
//                                    _("A file named “s” already exists. Do you want to replace it?"),
//                                    filename);
//
//    adw_message_dialog_add_responses (ADW_MESSAGE_DIALOG (dialog),
//                                      "cancel",  _("_Cancel"),
//                                      "replace", _("_Replace"),
//                                      NULL);
//
//    adw_message_dialog_set_response_appearance (ADW_MESSAGE_DIALOG (dialog), "replace", ADW_RESPONSE_DESTRUCTIVE);
//
//    adw_message_dialog_set_default_response (ADW_MESSAGE_DIALOG (dialog), "cancel");
//    adw_message_dialog_set_close_response (ADW_MESSAGE_DIALOG (dialog), "cancel");
//
//    g_signal_connect (dialog, "response", G_CALLBACK (response_cb), self);
//
//    gtk_window_present (GTK_WINDOW (dialog));
//
// # Async API
//
// AdwMessageDialog can also be used via the messagedialog.Choose method.
// This API follows the GIO async pattern, and the result can be obtained by
// calling messagedialog.ChooseFinish, for example:
//
//    static void
//    dialog_cb (AdwMessageDialog *dialog,
//               GAsyncResult     *result,
//               MyWindow         *self)
//    {
//      const char *response = adw_message_dialog_choose_finish (dialog, result);
//
//      // ...
//    }
//
//    static void
//    show_dialog (MyWindow *self)
//    {
//      GtkWidget *dialog;
//
//      dialog = adw_message_dialog_new (GTK_WINDOW (self), _("Replace File?"), NULL);
//
//      adw_message_dialog_format_body (ADW_MESSAGE_DIALOG (dialog),
//                                      _("A file named “s” already exists. Do you want to replace it?"),
//                                      filename);
//
//      adw_message_dialog_add_responses (ADW_MESSAGE_DIALOG (dialog),
//                                        "cancel",  _("_Cancel"),
//                                        "replace", _("_Replace"),
//                                        NULL);
//
//      adw_message_dialog_set_response_appearance (ADW_MESSAGE_DIALOG (dialog), "replace", ADW_RESPONSE_DESTRUCTIVE);
//
//      adw_message_dialog_set_default_response (ADW_MESSAGE_DIALOG (dialog), "cancel");
//      adw_message_dialog_set_close_response (ADW_MESSAGE_DIALOG (dialog), "cancel");
//
//      adw_message_dialog_choose (ADW_MESSAGE_DIALOG (dialog), NULL, (GAsyncReadyCallback) dialog_cb, self);
//    }
//
// # AdwMessageDialog as GtkBuildable
//
// AdwMessageDialog supports adding responses in UI definitions by via the
// <responses> element that may contain multiple <response> elements, each
// respresenting a response.
//
// Each of the <response> elements must have the id attribute specifying the
// response ID. The contents of the element are used as the response label.
//
// Response labels can be translated with the usual translatable, context and
// comments attributes.
//
// The <response> elements can also have enabled and/or appearance attributes.
// See messagedialog.SetResponseEnabled and messagedialog.SetResponseAppearance
// for details.
//
// Example of an AdwMessageDialog UI definition:
//
//    <object class="AdwMessageDialog" id="dialog">
//      <property name="heading" translatable="yes">Save Changes?</property>
//      <property name="body" translatable="yes">Open documents contain unsaved changes. Changes which are not saved will be permanently lost.</property>
//      <property name="default-response">save</property>
//      <property name="close-response">cancel</property>
//      <signal name="response" handler="response_cb"/>
//      <responses>
//        <response id="cancel" translatable="yes">_Cancel</response>
//        <response id="discard" translatable="yes" appearance="destructive">_Discard</response>
//        <response id="save" translatable="yes" appearance="suggested" enabled="false">_Save</response>
//      </responses>
//    </object>
//
// # Accessibility
//
// AdwMessageDialog uses the GTK_ACCESSIBLE_ROLE_DIALOG role.
type MessageDialog struct {
	_ [0]func() // equal guard
	gtk.Window
}

var (
	_ gtk.Widgetter     = (*MessageDialog)(nil)
	_ coreglib.Objector = (*MessageDialog)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*MessageDialog, *MessageDialogClass, MessageDialogOverrides](
		GTypeMessageDialog,
		initMessageDialogClass,
		wrapMessageDialog,
		defaultMessageDialogOverrides,
	)
}

func initMessageDialogClass(gclass unsafe.Pointer, overrides MessageDialogOverrides, classInitFunc func(*MessageDialogClass)) {
	pclass := (*C.AdwMessageDialogClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeMessageDialog))))

	if overrides.Response != nil {
		pclass.response = (*[0]byte)(C._gotk4_adw1_MessageDialogClass_response)
	}

	if classInitFunc != nil {
		class := (*MessageDialogClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapMessageDialog(obj *coreglib.Object) *MessageDialog {
	return &MessageDialog{
		Window: gtk.Window{
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
			Root: gtk.Root{
				NativeSurface: gtk.NativeSurface{
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
				},
			},
			ShortcutManager: gtk.ShortcutManager{
				Object: obj,
			},
		},
	}
}

func marshalMessageDialog(p uintptr) (interface{}, error) {
	return wrapMessageDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectResponse: this signal is emitted when the dialog is closed.
//
// response will be set to the response ID of the button that had been
// activated.
//
// if the dialog was closed by pressing <kbd>Escape</kbd> or with a system
// action, response will be set to the value of messagedialog:close-response.
func (self *MessageDialog) ConnectResponse(f func(response string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "response", false, unsafe.Pointer(C._gotk4_adw1_MessageDialog_ConnectResponse), f)
}

// NewMessageDialog creates a new AdwMessageDialog.
//
// heading and body can be set to NULL. This can be useful if they need
// to be formatted or use markup. In that case, set them to NULL and call
// messagedialog.FormatBody or similar methods afterwards:
//
//    GtkWidget *dialog;
//
//    dialog = adw_message_dialog_new (parent, _("Replace File?"), NULL);
//    adw_message_dialog_format_body (ADW_MESSAGE_DIALOG (dialog),
//                                    _("A file named “s” already exists.  Do you want to replace it?"),
//                                    filename);.
//
// The function takes the following parameters:
//
//   - parent (optional): transient parent.
//   - heading (optional): heading.
//   - body (optional) text.
//
// The function returns the following values:
//
//   - messageDialog: newly created AdwMessageDialog.
//
func NewMessageDialog(parent *gtk.Window, heading, body string) *MessageDialog {
	var _arg1 *C.GtkWindow // out
	var _arg2 *C.char      // out
	var _arg3 *C.char      // out
	var _cret *C.GtkWidget // in

	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	}
	if heading != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(heading)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	if body != "" {
		_arg3 = (*C.char)(unsafe.Pointer(C.CString(body)))
		defer C.free(unsafe.Pointer(_arg3))
	}

	_cret = C.adw_message_dialog_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(heading)
	runtime.KeepAlive(body)

	var _messageDialog *MessageDialog // out

	_messageDialog = wrapMessageDialog(coreglib.Take(unsafe.Pointer(_cret)))

	return _messageDialog
}

// AddResponse adds a response with id and label to self.
//
// Responses are represented as buttons in the dialog.
//
// Response ID must be unique. It will be used in messagedialog::response to
// tell which response had been activated, as well as to inspect and modify the
// response later.
//
// An embedded underline in label indicates a mnemonic.
//
// messagedialog.SetResponseLabel can be used to change the response label after
// it had been added.
//
// messagedialog.SetResponseEnabled and messagedialog.SetResponseAppearance can
// be used to customize the responses further.
//
// The function takes the following parameters:
//
//   - id: response ID.
//   - label: response label.
//
func (self *MessageDialog) AddResponse(id, label string) {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 *C.char             // out
	var _arg2 *C.char             // out

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))

	C.adw_message_dialog_add_response(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(id)
	runtime.KeepAlive(label)
}

// ChooseFinish finishes the messagedialog.Choose call and returns the response
// ID.
//
// The function takes the following parameters:
//
//   - result: GAsyncResult.
//
// The function returns the following values:
//
//   - utf8: ID of the response that was selected, or
//     messagedialog:close-response if the call was cancelled.
//
func (self *MessageDialog) ChooseFinish(result gio.AsyncResulter) string {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 *C.GAsyncResult     // out
	var _cret *C.char             // in

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_cret = C.adw_message_dialog_choose_finish(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(result)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Body gets the body text of self.
//
// The function returns the following values:
//
//   - utf8: body of self.
//
func (self *MessageDialog) Body() string {
	var _arg0 *C.AdwMessageDialog // out
	var _cret *C.char             // in

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_message_dialog_get_body(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// BodyUseMarkup gets whether the body text of self includes Pango markup.
//
// The function returns the following values:
//
//   - ok: whether self uses markup for body text.
//
func (self *MessageDialog) BodyUseMarkup() bool {
	var _arg0 *C.AdwMessageDialog // out
	var _cret C.gboolean          // in

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_message_dialog_get_body_use_markup(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CloseResponse gets the ID of the close response of self.
//
// The function returns the following values:
//
//   - utf8: close response ID.
//
func (self *MessageDialog) CloseResponse() string {
	var _arg0 *C.AdwMessageDialog // out
	var _cret *C.char             // in

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_message_dialog_get_close_response(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// DefaultResponse gets the ID of the default response of self.
//
// The function returns the following values:
//
//   - utf8 (optional): default response ID.
//
func (self *MessageDialog) DefaultResponse() string {
	var _arg0 *C.AdwMessageDialog // out
	var _cret *C.char             // in

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_message_dialog_get_default_response(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ExtraChild gets the child widget of self.
//
// The function returns the following values:
//
//   - widget (optional): child widget of self.
//
func (self *MessageDialog) ExtraChild() gtk.Widgetter {
	var _arg0 *C.AdwMessageDialog // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_message_dialog_get_extra_child(_arg0)
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

// Heading gets the heading of self.
//
// The function returns the following values:
//
//   - utf8 (optional): heading of self.
//
func (self *MessageDialog) Heading() string {
	var _arg0 *C.AdwMessageDialog // out
	var _cret *C.char             // in

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_message_dialog_get_heading(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// HeadingUseMarkup gets whether the heading of self includes Pango markup.
//
// The function returns the following values:
//
//   - ok: whether self uses markup for heading.
//
func (self *MessageDialog) HeadingUseMarkup() bool {
	var _arg0 *C.AdwMessageDialog // out
	var _cret C.gboolean          // in

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_message_dialog_get_heading_use_markup(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ResponseAppearance gets the appearance of response.
//
// See messagedialog.SetResponseAppearance.
//
// The function takes the following parameters:
//
//   - response ID.
//
// The function returns the following values:
//
//   - responseAppearance: appearance of response.
//
func (self *MessageDialog) ResponseAppearance(response string) ResponseAppearance {
	var _arg0 *C.AdwMessageDialog     // out
	var _arg1 *C.char                 // out
	var _cret C.AdwResponseAppearance // in

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(response)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.adw_message_dialog_get_response_appearance(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(response)

	var _responseAppearance ResponseAppearance // out

	_responseAppearance = ResponseAppearance(_cret)

	return _responseAppearance
}

// ResponseEnabled gets whether response is enabled.
//
// See messagedialog.SetResponseEnabled.
//
// The function takes the following parameters:
//
//   - response ID.
//
// The function returns the following values:
//
//   - ok: whether response is enabled.
//
func (self *MessageDialog) ResponseEnabled(response string) bool {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 *C.char             // out
	var _cret C.gboolean          // in

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(response)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.adw_message_dialog_get_response_enabled(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(response)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ResponseLabel gets the label of response.
//
// See messagedialog.SetResponseLabel.
//
// The function takes the following parameters:
//
//   - response ID.
//
// The function returns the following values:
//
//   - utf8: label of response.
//
func (self *MessageDialog) ResponseLabel(response string) string {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 *C.char             // out
	var _cret *C.char             // in

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(response)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.adw_message_dialog_get_response_label(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(response)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// HasResponse gets whether self has a response with the ID response.
//
// The function takes the following parameters:
//
//   - response ID.
//
// The function returns the following values:
//
//   - ok: whether self has a response with the ID response.
//
func (self *MessageDialog) HasResponse(response string) bool {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 *C.char             // out
	var _cret C.gboolean          // in

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(response)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.adw_message_dialog_has_response(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(response)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RemoveResponse removes a response from self.
//
// The function takes the following parameters:
//
//   - id: response ID.
//
func (self *MessageDialog) RemoveResponse(id string) {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 *C.char             // out

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_message_dialog_remove_response(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(id)
}

// Response emits the messagedialog::response signal with the given response ID.
//
// Used to indicate that the user has responded to the dialog in some way.
//
// The function takes the following parameters:
//
//   - response ID.
//
func (self *MessageDialog) Response(response string) {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 *C.char             // out

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(response)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_message_dialog_response(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(response)
}

// SetBody sets the body text of self.
//
// The function takes the following parameters:
//
//   - body of self.
//
func (self *MessageDialog) SetBody(body string) {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 *C.char             // out

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(body)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_message_dialog_set_body(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(body)
}

// SetBodyUseMarkup sets whether the body text of self includes Pango markup.
//
// See pango.ParseMarkup().
//
// The function takes the following parameters:
//
//   - useMarkup: whether to use markup for body text.
//
func (self *MessageDialog) SetBodyUseMarkup(useMarkup bool) {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if useMarkup {
		_arg1 = C.TRUE
	}

	C.adw_message_dialog_set_body_use_markup(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(useMarkup)
}

// SetCloseResponse sets the ID of the close response of self.
//
// It will be passed to messagedialog::response if the window is closed by
// pressing <kbd>Escape</kbd> or with a system action.
//
// It doesn't have to correspond to any of the responses in the dialog.
//
// The default close response is close.
//
// The function takes the following parameters:
//
//   - response: close response ID.
//
func (self *MessageDialog) SetCloseResponse(response string) {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 *C.char             // out

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(response)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_message_dialog_set_close_response(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(response)
}

// SetDefaultResponse sets the ID of the default response of self.
//
// If set, pressing <kbd>Enter</kbd> will activate the corresponding button.
//
// If set to NULL or to a non-existent response ID, pressing <kbd>Enter</kbd>
// will do nothing.
//
// The function takes the following parameters:
//
//   - response (optional): default response ID.
//
func (self *MessageDialog) SetDefaultResponse(response string) {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 *C.char             // out

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if response != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(response)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_message_dialog_set_default_response(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(response)
}

// SetExtraChild sets the child widget of self.
//
// The child widget is displayed below the heading and body.
//
// The function takes the following parameters:
//
//   - child (optional) widget.
//
func (self *MessageDialog) SetExtraChild(child gtk.Widgetter) {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 *C.GtkWidget        // out

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	C.adw_message_dialog_set_extra_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetHeading sets the heading of self.
//
// The function takes the following parameters:
//
//   - heading (optional) of self.
//
func (self *MessageDialog) SetHeading(heading string) {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 *C.char             // out

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if heading != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(heading)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_message_dialog_set_heading(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(heading)
}

// SetHeadingUseMarkup sets whether the heading of self includes Pango markup.
//
// See pango.ParseMarkup().
//
// The function takes the following parameters:
//
//   - useMarkup: whether to use markup for heading.
//
func (self *MessageDialog) SetHeadingUseMarkup(useMarkup bool) {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if useMarkup {
		_arg1 = C.TRUE
	}

	C.adw_message_dialog_set_heading_use_markup(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(useMarkup)
}

// SetResponseAppearance sets the appearance for response.
//
// <picture> <source srcset="message-dialog-appearance-dark.png"
// media="(prefers-color-scheme: dark)"> <img
// src="message-dialog-appearance.png" alt="message-dialog-appearance">
// </picture>
//
// Use ADW_RESPONSE_SUGGESTED to mark important responses such as the
// affirmative action, like the Save button in the example.
//
// Use ADW_RESPONSE_DESTRUCTIVE to draw attention to the potentially damaging
// consequences of using response. This appearance acts as a warning to the
// user. The Discard button in the example is using this appearance.
//
// The default appearance is ADW_RESPONSE_DEFAULT.
//
// Negative responses like Cancel or Close should use the default appearance.
//
// The function takes the following parameters:
//
//   - response ID.
//   - appearance for response.
//
func (self *MessageDialog) SetResponseAppearance(response string, appearance ResponseAppearance) {
	var _arg0 *C.AdwMessageDialog     // out
	var _arg1 *C.char                 // out
	var _arg2 C.AdwResponseAppearance // out

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(response)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.AdwResponseAppearance(appearance)

	C.adw_message_dialog_set_response_appearance(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(response)
	runtime.KeepAlive(appearance)
}

// SetResponseEnabled sets whether response is enabled.
//
// If response is not enabled, the corresponding button will have
// gtk.Widget:sensitive set to FALSE and it can't be activated as a default
// response.
//
// response can still be used as messagedialog:close-response while it's not
// enabled.
//
// Responses are enabled by default.
//
// The function takes the following parameters:
//
//   - response ID.
//   - enabled: whether to enable response.
//
func (self *MessageDialog) SetResponseEnabled(response string, enabled bool) {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 *C.char             // out
	var _arg2 C.gboolean          // out

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(response)))
	defer C.free(unsafe.Pointer(_arg1))
	if enabled {
		_arg2 = C.TRUE
	}

	C.adw_message_dialog_set_response_enabled(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(response)
	runtime.KeepAlive(enabled)
}

// SetResponseLabel sets the label of response to label.
//
// Labels are displayed on the dialog buttons. An embedded underline in label
// indicates a mnemonic.
//
// The function takes the following parameters:
//
//   - response ID.
//   - label of response.
//
func (self *MessageDialog) SetResponseLabel(response, label string) {
	var _arg0 *C.AdwMessageDialog // out
	var _arg1 *C.char             // out
	var _arg2 *C.char             // out

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(response)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))

	C.adw_message_dialog_set_response_label(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(response)
	runtime.KeepAlive(label)
}

// Response emits the messagedialog::response signal with the given response ID.
//
// Used to indicate that the user has responded to the dialog in some way.
//
// The function takes the following parameters:
//
//   - response ID.
//
func (self *MessageDialog) response(response string) {
	gclass := (*C.AdwMessageDialogClass)(coreglib.PeekParentClass(self))
	fnarg := gclass.response

	var _arg0 *C.AdwMessageDialog // out
	var _arg1 *C.char             // out

	_arg0 = (*C.AdwMessageDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(response)))
	defer C.free(unsafe.Pointer(_arg1))

	C._gotk4_adw1_MessageDialog_virtual_response(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(response)
}
