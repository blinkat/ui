/**
 * for windows
 */

#include "inc/ui.h"
#include <malloc.h>
#include <windowsx.h>
#include <shellapi.h>
#include <stdlib.h>

#ifdef _WIN32

// convert png jpg bmp to ico
gIcon converToIco(void* buffer, int width, height)
{
	HDC dc;
	HBITMAP bit, mask;
	gIcon icon;
	HANDLE img;
	BITMAPV5HEADER bi;
	ICONINFO ii;
	DWORD* target = 0;
	BYTE* source = (BYTE*)buffer;
	int i = 0;

	ZeroMemory(&bi, sizeof(bi));
	bi.bV5Size        = sizeof(BITMAPV5HEADER);
	bi.bV5Width       = width;
	bi.bV5Height      = -height;
	bi.bV5Planes      = 1;
	bi.bV5BitCount    = 32;
	bi.bV5Compression = BI_BITFIELDS;
	bi.bV5RedMask     = 0x00ff0000;
	bi.bV5GreenMask   = 0x0000ff00;
	bi.bV5BlueMask    = 0x000000ff;
	bi.bV5AlphaMask   = 0xff000000;

	dc = GetDC(NULL);
	bit = CreateDIBSection(dc, (BITMAPINFO*)&bi, DIB_RGB_COLORS, (void**)&target, NULL, (DWORD)0);
	ReleaseDC(NULL, dc);

	if (!bit) return NULL;

	mask = CreateBitmap(width, height, 1, 1, NULL);
	if (!mask)
	{
		DeleteObject(bit);
		return NULL;
	}

	for (i = 0; i < width * height; i++, target++, source += 4)
	{
		*target = (source[3] << 24) |
		          (source[0] << 16) |
		          (source[1] << 8) |
		          source[2];
	}

	ZeroMemory(&ii, sizeof(ii));
	ii.fIcon = 1;
	ii.hbmMask = mask;
	ii.hbmColor = bit;

	icon = CreateIconIndirect(&ii);

	DeleteObject(bit);
	DeleteObject(mask);
	return icon;
}

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
gHANDLE gCreateWindow(int width, int height, gCHAR title, int px, int py, gIcon icon, gHANDLE parent)
{
	wchar_t* guid;
	gHANDLE ret;
	HWND hwnd;
	HINSTANCE hin;
	WNDCLASSEX cls;
	UINT styles;
	gIcon icon;

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

// ========== dc ==========
void gAbortPath(gDC dc)
{
	AbortPath(dc);
}
void gBeginPath(gDC dc)
{
	BeginPath(dc);
}
void gEndPath(gDC)
{
	EndPath(dc);
}

// x, y, radius, start angle, end angle
void gArc(gDC dc, gPoint p, int r, float start, float end)
{
	AngleArc(p.x, p.y, r, start, end);
}

void gBezier(gDC dc, gPoint s, gPoint cs, gPoint ce, gPoint e)
{
	POINT a[4] = {s.x, s.y, cs.x, cs.y, ce.x, ce.y, e.x, e.y};
	PolyBezier(dc, a, 4);
}
// rect path
//
void gRect(gDC dc, gPoint p, int w, int h)
{
	POINT a[3] = {p.x, p.y, p.x + w, p.y, p.x + w, p.y + h };
	Polygon(dc, a, 3);
}

void gRectPath(gDC dc, Rectangle r)
{
	POINT a[3] = {r.x, r.y, r.x + r.width, r.y, r.x + r.width, r.y + r.height};
	Polygon(dc, a, 3);
}

void Fill(gDC dc)
{
	FillPath(dc);
}

void Stroke(gDC dc)
{
	StrokePath(dc);
}

void RoundRect(gDC dc, gPoint p, int width, int height, int radius)
{
	RoundRect(dc, p.x, p.y, p.x + w, p.y + h, radius, radius);
}
void RoundRectPath(gDC dc, Rectangle r)
{
	RoundRect(dc, r.x, r.y, r.x + r.width, r.y + r.height, radius, radius);
}
#endif