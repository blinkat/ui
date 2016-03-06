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
	WNDCLASSEX cls;
	UINT styles;
	gIcon icon;

	// icon = LoadImage(NULL, L"default.ico", IMAGE_ICON, 32, 32, LR_LOADFROMFILE);
	// if (icon == NULL)
	// {
	// 	MessageBox(NULL, L"err", L"", MB_OK);
	// }

	guid = (gCHAR)malloc(GUID_LENGTH * sizeof(wchar_t));
	newGUID(guid);
	hin = GetModuleHandle(NULL);

	cls.cbSize = sizeof(WNDCLASSEX);
	cls.style = CS_HREDRAW | CS_VREDRAW;
	cls.lpfnWndProc = (WNDPROC)WinProc;
	cls.cbClsExtra = 0;
	cls.cbWndExtra = 0;
	cls.hInstance = hin;
	cls.hIcon = _DEFAULT_ICON;
	cls.hCursor = LoadCursor(NULL, IDC_ARROW);
	cls.hbrBackground = (HBRUSH)GetStockObject(WHITE_BRUSH);
	cls.lpszMenuName = NULL;
	cls.lpszClassName = (LPCTSTR)guid;
	cls.hIconSm = _DEFAULT_ICON;

	if (!RegisterClassEx(&cls))
	{
		MessageBox(NULL, L"", L"", MB_OK);
		return NULL;
	}

	hwnd = CreateWindowEx(
	           WS_EX_CLIENTEDGE | WS_EX_CONTROLPARENT,	// ex style
	           (LPCTSTR)guid,
	           (LPCTSTR)title,
	           WS_CLIPCHILDREN | WS_CLIPSIBLINGS | WS_POPUP,
	           0,
	           0,
	           width,
	           height,
	           parent,
	           NULL,
	           hin,
	           0
	       );

	free(guid);
	return hwnd;
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

gIcon gLoadIcon(void* buffer, int width, int height)
{
	HBITMAP bit;
	gIcon icon;
	HANDLE img;

	bit = CreateBitmap(width, height, 1, 8, buffer);
	if (bit == NULL)
	{
		return NULL;
	}

	img = CopyImage(bit, IMAGE_ICON, 32, 32, LR_COPYDELETEORG | LR_COPYRETURNORG);
	if (img == NULL)
	{
		MessageBox(NULL, L"err", L"", MB_OK);
		return NULL;
	}

	return (gIcon)img;
}
#endif