/**
 * for windows
 */

#include "inc/ui.h"

#ifdef _WIN32
LRESULT CALLBACK WinProc(HWND hwnd, UINT message, WPARAM wParam, LPARAM lParam)
{
	PAINTSTRUCT ps;
	HDC hdc;
	switch (message) {
	// mouse
	case WM_MOUSEWHEEL:
		gMouseMBWheelEvent(hwnd, (int)(LOWORD(lParam)), (int)(HIWORD(lParam)), ((int)wParam) < 0 ? -1 : 1);
		return 0;
	case WM_MBUTTONDOWN:
		gMouseMBDownEvent(hwnd, (int)(LOWORD(lParam)), (int)(HIWORD(lParam)));
		return 0;
	case WM_MBUTTONUP:
		gMouseMBUpEvent(hwnd, (int)(LOWORD(lParam)), (int)(HIWORD(lParam)));
		return 0;
	case WM_MBUTTONDBLCLK:
		gMouseMBDoubleEvent(hwnd, (int)(LOWORD(lParam)), (int)(HIWORD(lParam)));
		return 0;

	case WM_RBUTTONDOWN:
		gMouseRBDownEvent(hwnd, (int)(LOWORD(lParam)), (int)(HIWORD(lParam)));
		return 0;
	case WM_RBUTTONUP:
		gMouseRBUpEvent(hwnd, (int)(LOWORD(lParam)), (int)(HIWORD(lParam)));
		return 0;
	case WM_RBUTTONDBLCLK:
		gMouseRBDoubleEvent(hwnd, (int)(LOWORD(lParam)), (int)(HIWORD(lParam)));
		return 0;

	case WM_LBUTTONDOWN:
		gMouseLBDownEvent(hwnd, (int)(LOWORD(lParam)), (int)(HIWORD(lParam)));
		return 0;
	case WM_LBUTTONUP:
		gMouseLBUpEvent(hwnd, (int)(LOWORD(lParam)), (int)(HIWORD(lParam)));
		return 0;
	case WM_LBUTTONDBLCLK:
		gMouseLBDoubleEvent(hwnd, (int)(LOWORD(lParam)), (int)(HIWORD(lParam)));
		return 0;
	// key
	case WM_KEYDOWN:
		gKeyDownEvent(hwnd, (int)(wParam));
		return 0;
	case WM_KEYUP:
		gKeyUpEvent(hwnd, (int)(wParam));
		return 0;

	case WM_CLOSE:
		gCloseEvent(hwnd);
		return 0;
	case WM_PAINT:
		hdc = BeginPaint(hwnd, &ps);
		gPaintEvent(hwnd, hdc);
		EndPaint(hwnd, &ps);
		return 0;
	case WM_SETFOCUS:
		gFocusEvent(hwnd);
		return 0;
	case WM_KILLFOCUS:
		gFocusOutEvent(hwnd);
		return 0;
	case WM_MOVE:
		gMoveEvent(hwnd, (int)(LOWORD(lParam)), (int)(HIWORD(lParam)));
		return 0;
	case WM_SIZE:
		gSizeEvent(hwnd, (int)(LOWORD(lParam)), (int)(HIWORD(lParam)));
		return 0;
	case WM_DESTROY:
		gDestoryEvent(hwnd);
		PostQuitMessage(0);
		return 0;
	}
	return DefWindowProc(hwnd, message, wParam, lParam);
}

// ===============================================================

/**
 * create window func
 */
gHANDLE gCreateWindow(int width, int height, gCHAR title, int flags, gHANDLE parent)
{
	wchar_t* guid;
	gHANDLE ret;
	HWND hwnd;
	HINSTANCE hin;
	WNDCLASS cls;
	UINT styles;

	guid = (gCHAR)malloc(GUID_LENGTH * sizeof(wchar_t));
	newGUID(guid);
	hin = GetModuleHandle(NULL);
	cls.style         = CS_HREDRAW | CS_VREDRAW;                         //窗口样式
	cls.lpszClassName = (LPCTSTR)guid;                               //窗口类名
	cls.lpszMenuName  = NULL;                                     //窗口菜单:无
	cls.hbrBackground = (HBRUSH) GetStockObject(WHITE_BRUSH);    //窗口背景颜色
	cls.lpfnWndProc   = WinProc;                                   //窗口处理函数
	cls.cbWndExtra    = 0;                                          //窗口实例扩展:无
	cls.cbClsExtra    = 0;                                          //窗口类扩展:无
	cls.hInstance     = hin;                                   //窗口实例句柄
	cls.hIcon         = LoadIcon( NULL, IDI_APPLICATION );               //窗口最小化图标:使用缺省图标
	cls.hCursor       = LoadCursor( NULL, IDC_ARROW );                 //窗口采用箭头光标

	styles = WS_VISIBLE;
	if (parent != NULL) styles |= WS_CHILD;

	hwnd = CreateWindow(
	           (LPCTSTR)guid,                 //窗口类名
	           (LPCTSTR)title,             //窗口标题
	           styles,       //窗口的风格
	           0,             			//窗口初始显示位置x:使用缺省值
	           0,             			//窗口初始显示位置y:使用缺省值
	           width,             //窗口的宽度:使用缺省值
	           height,             //窗口的高度:使用缺省值
	           parent,                    //父窗口:无
	           NULL,                      //子菜单:无
	           hin,                 //该窗口应用程序的实例句柄
	           NULL                       //
	       );

	ret = hwnd;

	free(guid);
	return ret;
}

void gShowWindow(gHANDLE h)
{
	MSG msg;

	ShowWindow(h, SW_SHOWNORMAL);
	UpdateWindow(h);

	while (GetMessage(&msg, NULL, 0, 0))
	{
		TranslateMessage(&msg);
		DispatchMessage(&msg);
	}
}
#endif