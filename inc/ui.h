/**====================================
 * cross phatom for golang ui library =
 * ====================================
 * GO_HANDLE window handler. expressed as a window.
 */

#ifndef GO_UI_H
#define GO_UI_H

#include "stdio.h"

#define UNICODE
#define _UNICODE
#define GUID_LENGTH 38
typedef wchar_t* gCHAR;
// generate guid
extern void newGUID(gCHAR str);
extern int gInit();
extern void gSetDefaultIcon(void* buf);

struct WinStyle;


// ======== include API ==========
// == win
#ifdef _WIN32
#define WIN32_MEAN_AND_LEAN
#include "windows.h"

// define type for window
typedef HWND gHANDLE;
typedef HDC gDC;			// paint device context
typedef HICON gIcon;

// define key code
#define gKEY_BACK 		VK_BACK
#define gKEY_TAB 		VK_TAB
#define gKEY_ENTER 		VK_RETURN
#define gKEY_SHIFT 		VK_SHIFT
#define gKEY_CTRL 		VK_CONTROL
#define gKEY_ALT 		VK_MENU
#define gKEY_PAUSE 		VK_PAUSE
#define gKEY_CAPS 		VK_CAPITAL
#define gKEY_ESC 		VK_ESCAPE
#define gKEY_SPACE 		VK_SPACE

#define gKEY_PAGEUP 	VK_PRIOR
#define gKEY_PAGEDOWN 	VK_NEXT
#define gKEY_END 		VK_END
#define gKEY_HOME 		VK_HOME

#define gKEY_LEFT 		VK_LEFT
#define gKEY_UP 		VK_UP
#define gKEY_RIGHT 		VK_RIGHT
#define gKEY_DOWN 		VK_DOWN

//#define gKey_select 	VK_SELECT
#define gKEY_PRINT 		VK_SNAPSHOT
#define gKEY_INSERT 	VK_INSERT
#define gKEY_DELETE 	VK_DELETE
#define gKEY_HELP 		VK_HELP

#define gKEY_NUM0 		VK_NUMPAD0
#define gKEY_NUM1 		VK_NUMPAD1
#define gKEY_NUM2 		VK_NUMPAD2
#define gKEY_NUM3 		VK_NUMPAD3
#define gKEY_NUM4 		VK_NUMPAD4
#define gKEY_NUM5 		VK_NUMPAD5
#define gKEY_NUM6 		VK_NUMPAD6
#define gKEY_NUM7 		VK_NUMPAD7
#define gKEY_NUM8 		VK_NUMPAD8
#define gKEY_NUM9 		VK_NUMPAD9

#define gKEY_F1 		VK_F1
#define gKEY_F2 		VK_F2
#define gKEY_F3 		VK_F3
#define gKEY_F4 		VK_F4
#define gKEY_F5 		VK_F5
#define gKEY_F6 		VK_F6
#define gKEY_F7 		VK_F7
#define gKEY_F8 		VK_F8
#define gKEY_F9 		VK_F9
#define gKEY_F10 		VK_F10
#define gKEY_F11 		VK_F11
#define gKEY_F12 		VK_F12

#define gKEY_NUMLOCK 	VK_NUMLOCK
#define gKEY_SCROLLLOCK VK_SCROLLLOCK
#define gKEY_LSHIFT 	VK_LSHIFT
#define gKEY_RSHIFT 	VK_RSHIFT
#define gKEY_LCTRL 		VK_LCONTROL
#define gKEY_RCTRL 		VK_RCONTROL
#define gKEY_LALT 		VK_LMENU
#define gKEY_RALT 		VK_RMENU

#define gKEY_A 		65
#define gKEY_B 		66
#define gKEY_C 		67
#define gKEY_D 		68
#define gKEY_E 		69
#define gKEY_F 		70
#define gKEY_G 		71
#define gKEY_H 		72
#define gKEY_I 		73
#define gKEY_J 		74
#define gKEY_K 		75
#define gKEY_L 		76
#define gKEY_M 		77
#define gKEY_N 		78
#define gKEY_O 		79
#define gKEY_P 		80
#define gKEY_Q 		81
#define gKEY_R 		82
#define gKEY_S 		83
#define gKEY_T 		84
#define gKEY_U 		85
#define gKEY_V 		86
#define gKEY_W 		87
#define gKEY_X 		88
#define gKEY_Y 		89
#define gKEY_Z 		90


#define gKEY_0 		48
#define gKEY_1 		49
#define gKEY_2 		50
#define gKEY_3 		51
#define gKEY_4 		52
#define gKEY_5 		53
#define gKEY_6 		54
#define gKEY_7 		55
#define gKEY_8 		56
#define gKEY_9 		57

#define gKEY_MINUS_UNDERLINE 	189 // -_
#define gKEY_ADDA_ND			187 // +=
#define gKEY_BACKLASH			220 // |\

#define gKEY_TILDE 				192 // `~
#define gKEY_SEMICONLON			186 // :;
#define gKEY_QUOTES				222 // '"
#define gKEY_COMMA				188 // <,
#define gKEY_PERIOD 			190 // .>
#define gKEY_FORWARDSLASH		191 // /?
#define gKEY_BRACKETLEFT		219 // [{
#define gKEY_BRACKETRIGHT		221 // ]}

#define gKEY_NUM_FORWARDSLASH	111
#define gKEY_NUM_ASTERISK		106
#define gKEY_NUM_MINUS			109
#define gKEY_NUM_ADD			107
#define gKEY_NUM_DOT			110

#define gFLAGS_DEFAULT WS_OVERLAPPEDWINDOW
#endif
//== end win


// ======== define flags ========
// #define gFLAGS_NO_FRAME 1;
// #define gFLAGS_CHILD    2;
// #define gFLAGS_NO_TAB   4;

// // default windows flags
// #define gFLAGS_MODULE gFLAGS_NO_FRAME | gFLAGS_CHILD | gFLAGS_NO_TAB;

// ======== event callback ========
// size int width, int height
extern void gSizeEvent(gHANDLE, int, int);	//size change
// int x int y
extern void gMoveEvent(gHANDLE, int, int);	//win move
extern void gFocusEvent(gHANDLE);	// win focus
extern void gFocusOutEvent(gHANDLE);// win focus out
extern void gPaintEvent(gHANDLE, gDC);	// win paint
extern void gCloseEvent(gHANDLE);	// win close
extern void gDestoryEvent(gHANDLE);
// key event
// key event params
// gHandle window handle
// int key enum
// int press down alt/shift/ctrl
extern void gKeyDownEvent(gHANDLE, int);
extern void gKeyUpEvent(gHANDLE, int);

// mouse event
// int x
// int y

// mouse left
extern void gMouseLBDownEvent(gHANDLE, int, int);
extern void gMouseLBUpEvent(gHANDLE, int, int);
extern void gMouseLBDoubleEvent(gHANDLE, int, int);
// mouse right
extern void gMouseRBUpEvent(gHANDLE, int, int);
extern void gMouseRBDownEvent(gHANDLE, int, int);
extern void gMouseRBDoubleEvent(gHANDLE, int, int);
// mouse middle button
extern void gMouseMBUpEvent(gHANDLE, int, int);
extern void gMouseMBDownEvent(gHANDLE, int, int);
extern void gMouseMBDoubleEvent(gHANDLE, int, int);
// last int is wheel 1 or -1
// wheel = -1 down
// wheel = 1 up
extern void gMouseMBWheelEvent(gHANDLE, int, int, int);

//===============================
// extern functions
/**
 * @param width
 * @param height
 * @param title
 * @param flags
 * @param parent
 * @return
 */
extern gHANDLE gCreateWindow(int, int, gCHAR, int, gHANDLE);
extern void gShowWindow(gHANDLE h);
extern gIcon gLoadIcon(void*, int, int);

extern gIcon _DEFAULT_ICON;
// ============= create windows structure ============
struct WinStyle {
	int x, y;
	int width, height;
	gCHAR icon;
	gCHAR title;
	gHANDLE parent;
	int style;
};

#define gWS_DEFAULT 0
#define gWS_NO_BORDER 1
#define gWS_CHILD 2

/**
 * generate default window structure
 * @result
 */
// extern void GenDefaultWinStyle(WinStyle *s);

#endif