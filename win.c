/**
 * for windows
 */

#include "inc/ui.h"
#include <malloc.h>
#include <windowsx.h>
#include <shellapi.h>
#include <stdlib.h>
#include <wingdi.h>

#ifdef _WIN32

#define WND_NAME TEXT("golang window")

// convert png jpg bmp to ico
HICON converToIco(void* buffer, int width, int height)
{
	HDC dc;
	HBITMAP bit, mask;
	HICON icon;
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

void callPaint(HWND hwnd)
{
	PAINTSTRUCT ps;
	HDC hdc, mdc;
	RECT rect;
	HBITMAP bit;
	HGDIOBJ gdi;
	int width, height;

	hdc = BeginPaint(hwnd, &ps);
	mdc = CreateCompatibleDC(hdc);
	GetWindowRect(hwnd, &rect);
	width = rect.right - rect.left;
	height = rect.bottom - rect.top;
	bit = CreateCompatibleBitmap(hdc, width, height);
	gdi = SelectObject(mdc, bit);
	SetMapMode (mdc, MM_TEXT);

	gPaintEvent(hwnd, mdc);

	BitBlt(hdc, 0, 0, width, height, mdc, 0, 0, SRCCOPY);
	SelectObject(hdc, gdi);
	DeleteDC(mdc);
	DeleteDC(hdc);
	DeleteObject(bit);
	DeleteObject(gdi);
	EndPaint(hwnd, &ps);
}

LRESULT CALLBACK WinProc(HWND hwnd, UINT message, WPARAM wParam, LPARAM lParam)
{
	switch (message) {
	case WM_MOUSEMOVE:
		gMouseMove(hwnd, (int)(LOWORD(lParam)), (int)(HIWORD(lParam)));
		return 0;
	case WM_SHOWWINDOW:
		gShowEvent(hwnd);
		return 0;
	// case WM_CREATE:
	// 	gCreatedEvent(hwnd);
	// 	return 0;
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
		callPaint(hwnd);
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

// ************ init *************

int gInit()
{
	HINSTANCE hin;
	WNDCLASSEX cls;

	hin = GetModuleHandle(NULL);
	cls.cbSize = sizeof(WNDCLASSEX);
	cls.style = CS_HREDRAW | CS_VREDRAW;
	cls.lpfnWndProc = (WNDPROC)WinProc;
	cls.cbClsExtra = 0;
	cls.cbWndExtra = 0;
	cls.hInstance = hin;
	cls.hIcon = NULL;
	cls.hCursor = LoadCursor(NULL, IDC_ARROW);
	cls.hbrBackground = (HBRUSH)GetStockObject(BLACK_BRUSH);
	cls.lpszMenuName = NULL;
	cls.lpszClassName = WND_NAME;
	cls.hIconSm = NULL;

	if (!RegisterClassEx(&cls))
	{
		return 0;
	}

	return 1;
}

HFONT create_font(gCHAR face, int style, int size);

/**
 * create window func
 */
gHANDLE gCreateWindow(int w, int h, int px, int py, int style, gHANDLE parent)
{
	gHANDLE ret;
	HWND hwnd;
	HINSTANCE hin;
	UINT st = 0;
	UINT exst = 0;

	if ((style & gWS_DEFAULT) == gWS_DEFAULT)
	{
		st = WS_OVERLAPPEDWINDOW;
	}
	else
	{
		st = WS_OVERLAPPED | WS_POPUP;
		if ((style & gWS_CHILD) == gWS_CHILD) st |= WS_CHILD;
	}

	if ((style & gWS_TOOL) == gWS_TOOL) exst = WS_EX_TOOLWINDOW;

	hin = GetModuleHandle(NULL);
	hwnd = CreateWindowEx(
	           //WS_EX_CLIENTEDGE | WS_EX_CONTROLPARENT,	// ex style
	           exst,
	           WND_NAME,
	           TEXT(""),
	           st,
	           px,
	           py,
	           w,
	           h,
	           parent,
	           NULL,
	           hin,
	           0
	       );

	if (hwnd == NULL) return NULL;
	// SetLayeredWindowAttributes(hwnd, RGB(0, 0, 0), 0, LWA_ALPHA | LWA_COLORKEY);

	return hwnd;
}

void gShowWindow(gHANDLE hwnd)
{
	ShowWindow(hwnd, SW_SHOWNORMAL);
	UpdateWindow(hwnd);
}

int gGetMessage(gHANDLE hwnd)
{
	MSG msg;
	int ret;
	ret = GetMessage(&msg, NULL, 0, 0);
	TranslateMessage(&msg);
	DispatchMessage(&msg);
	return ret;
}

void gDestroyWindow(gHANDLE h)
{
	DestroyWindow(h);
}

void gGetSize(gHANDLE h, int* width, int* height)
{
	RECT r;

	if (!GetWindowRect(h, &r)) {
		*width = 0;
		*height = 0;
	} else {
		*width = r.right - r.left;
		*height = r.bottom - r.top;
	}
}

void gSetOpacity(gHANDLE hwnd, gBYTE opa)
{
	SetLayeredWindowAttributes(hwnd, RGB(0, 0, 0), opa, LWA_ALPHA | LWA_COLORKEY);
}

gBYTE gGetOpacity(gHANDLE hwnd)
{
	gBYTE op;
	COLORREF pcr;
	DWORD flags;
	GetLayeredWindowAttributes(hwnd, &pcr, &op, &flags);
	return op;
}

// gIcon gLoadIcon(void* buffer, int width, int height)
// {
// 	return converToIco(buffer, width, height);
// }

void gDestoryIcon(gIcon ico)
{
	DeleteObject(ico);
}

void gSetIcon(gHANDLE hwnd, void* buffer, int width, int height)
{
	HICON ico;
	ico = converToIco(buffer, width, height);

	SendMessage(hwnd, WM_SETICON, ICON_BIG, (LPARAM)ico);
	SendMessage(hwnd, WM_SETICON, ICON_SMALL, (LPARAM)ico);

	DeleteObject(ico);
}

void gGetLocation(gHANDLE hwnd, int* x, int* y)
{
	RECT r;
	if (GetWindowRect(hwnd, &r))
	{
		*x = (int)r.left;
		*y = (int)r.top;
	} else {
		*x = 0;
		*y = 0;
	}
}

void gSetLocation(gHANDLE hwnd, int x, int y)
{
	SetWindowPos(hwnd, NULL, x, y, 0, 0, SWP_NOZORDER | SWP_NOSIZE);
}

void gSetSize(gHANDLE hwnd, int width, int height)
{
	SetWindowPos(hwnd, NULL, 0, 0, width, height, SWP_NOZORDER | SWP_NOMOVE);
}

void gMoveTop(gHANDLE hwnd)
{
	SetWindowPos(hwnd, HWND_TOP, 0, 0, 0, 0, SWP_NOMOVE | SWP_NOSIZE);
}

void gMoveBottom(gHANDLE hwnd)
{
	SetWindowPos(hwnd, HWND_BOTTOM, 0, 0, 0, 0, SWP_NOMOVE | SWP_NOSIZE);
}

void gRepaint(gHANDLE hwnd)
{
	//PostMessage(hwnd, WM_PAINT, 0, 0);
	RECT r;
	GetWindowRect(hwnd, &r);
	RedrawWindow(hwnd, &r, NULL, RDW_ERASE | RDW_INVALIDATE);
}

void gGetRect(gHANDLE hwnd, int* left, int* top, int* right, int* bottom)
{
	RECT r;
	if (GetWindowRect(hwnd, &r))
	{
		*left = (int)r.left;
		*top = (int)r.top;
		*right = (int)r.right;
		*bottom = (int)r.bottom;
	}
}

void gSetRect(gHANDLE hwnd, int x, int y, int width, int height)
{
	MoveWindow(hwnd, x, y, width, height, TRUE);
}

// *************** drawer *******************
// void init_dc(gDC dc, gFillStyle *style, gColor *cpen, int pwidth, HBRUSH *brush, HPEN *pen)
// {
// 	gColor* col;

// 	LPTSTR tmp = malloc(sizeof(TCHAR) * 1024 );
// 	// HBRUSH brush;
// 	// HPEN pen;

// 	*pen = CreatePen(PS_SOLID, pwidth, RGB(cpen->R, cpen->G, cpen->B));
// 	if (style->Length == 0)
// 	{
// 		col = style->Colors;
// 		*brush = CreateSolidBrush(RGB(col->R, col->G, col->B));
// 	}

// 	SelectObject(dc, *brush);
// 	SelectObject(dc, *pen);
// }

HBRUSH create_brush(gFillStyle *style)
{
	gColor *col;
	HBRUSH brush;

	if (style->Length == 0)
	{
		col = style->Colors;
		brush = CreateSolidBrush(RGB(col->R, col->G, col->B));
	}

	return brush;
}

void gFill(gDC dc, gFillStyle *c)
{
	HBRUSH brush = create_brush(c);
	SelectObject(dc, brush);
	FillPath(dc);
	DeleteObject(brush);
}

void gBeginPath(gDC dc, int px, int py)
{
	BeginPath(dc);
	MoveToEx(dc, px, py, NULL);
}

void gEndPath(gDC dc)
{
	EndPath(dc);
}

void gLineTo(gDC dc, int px, int py)
{
	LineTo(dc, px, py);
}

void gLineMoveTo(gDC dc, int px, int py)
{
	MoveToEx(dc, px, py, NULL);
}

void gStroke(gDC dc, gColor* col, int width)
{
	HPEN hpen = CreatePen(PS_SOLID, width, RGB(col->R, col->G, col->B));
	SelectObject(dc, hpen);
	StrokePath(dc);
	DeleteObject(hpen);
}

void gDrawText(gDC dc, int px, int py, int size, gCHAR text, gColor* color, gCHAR face, int style, int len)
{
	HFONT font;
	font = create_font(face, style, size);

	SelectObject(dc, font);
	SetBkMode(dc, TRANSPARENT);
	SetTextColor(dc, RGB(color->R, color->G, color->B));
	TextOut(dc, px, py, text, len);

	DeleteObject(font);
}

void gGetPix(gDC dc, int px, int py, gBYTE* r, gBYTE* g, gBYTE* b)
{
	COLORREF c = GetPixel(dc, px, py);
	*r = GetRValue(c);
	*g = GetGValue(c);
	*b = GetBValue(c);
}

void gSetPix(gDC dc, int px, int py, gBYTE r, gBYTE g, gBYTE b)
{
	SetPixel(dc, px, py, RGB(r, g, b));
}

// *************** global methods *****************
void gGetCursorPos(int* px, int* py)
{
	POINT p;
	GetCursorPos(&p);
	*px = p.x;
	*py = p.y;
}

int gAddFontResource(gCHAR path)
{
	return AddFontResource(path);
}

int gRemoveFontResource(gCHAR path)
{
	return RemoveFontResource(path);
}

// *************** font *****************
HFONT create_font(gCHAR face, int style, int size)
{
	HFONT font;
	int is_bold, is_italic, is_under, is_strike;


	if ((style & gFONT_BOLD) == gFONT_BOLD) is_bold = 700;
	else is_bold = 0;

	if ((style & gFONT_ITALIC) == gFONT_ITALIC) is_italic = 1;
	else is_italic = 0;


	if ((style & gFONT_UNDERLINE) == gFONT_UNDERLINE) is_under = 1;
	else is_under = 0;

	if ((style & gFONT_STRIKEOUT) == gFONT_STRIKEOUT) is_strike = 1;
	else is_strike = 0;


	// font = CreateFontIndirect(&tag);
	font = CreateFont(size, 0, 0, 0,
	                  is_bold, is_italic, is_under, is_strike,
	                  DEFAULT_CHARSET, OUT_DEFAULT_PRECIS, OUT_DEFAULT_PRECIS,
	                  DEFAULT_QUALITY, DEFAULT_PITCH,
	                  face
	                 );
	return font;
}

// void gDestoryBrush(gBrush h)
// {
// 	DeleteObject(h);
// }

// gBrush gCreateSolidBrush(gBYTE r, gBYTE g, gBYTE b)
// {
// 	HBRUSH h;
// 	h = CreateSolidBrush(RGB(r, g, b));
// 	return (gBrush)h;
// }

// // ***************** pen ********************
// void gDestoryPen(gPen h)
// {
// 	DeleteObject(h);
// }

// gPen gCreatePen(int style, gBYTE r, gBYTE g, gBYTE b, int width)
// {
// 	(gPen)CreatePen(style, width, RGB(r, g, b));
// }

// // ========== dc ==========
// void gFillRect(gDC dc, int left, int top, int right, int bottom, gBrush brush)
// {
// 	RECT r = {
// 		.left = left,
// 		.top = top,
// 		.right = right,
// 		.bottom = bottom
// 	};

// 	FillRect(dc, &r, brush);
// }

// void gClearBackground(gHANDLE hwnd, gDC dc)
// {
// 	RECT r, dr;
// 	HBITMAP bit, old;
// 	HDC mdc;
// 	int width, height;
// 	HBRUSH brush;
// 	BLENDFUNCTION bf;

// 	GetWindowRect(hwnd, &r);
// 	width = r.right - r.left;
// 	height = r.bottom - r.top;

// 	mdc = CreateCompatibleDC(dc);
// 	bit = CreateCompatibleBitmap(dc, width, height);
// 	old = SelectObject(mdc, bit);

// 	brush = CreateSolidBrush(RGB(0, 0, 0));
// 	dr.left = 0;
// 	dr.top = 0;
// 	dr.right = width;
// 	dr.bottom = height;

// 	FillRect(mdc, &dr, brush);

// 	memset(&bf, 0, sizeof(bf));
// 	bf.SourceConstantAlpha = 0x3f;
// 	bf.BlendOp = AC_SRC_OVER;

// 	AlphaBlend(dc, 0, 0, width, height, mdc, 0, 0, width, height, bf);
// 	SelectObject(mdc, old);

// }

// void gStrokeRect(gDC dc, int left, int top, int right, int bottom, gPen pen)
// {
// 	SelectObject(dc, pen);
// 	// MoveToEx(dc, left, top, NULL);
// 	// LineTo(dc, right, top);
// 	// LineTo(dc, right, bottom);
// 	// LineTo(dc, left, bottom);
// 	// LineTo(dc, left, top);
// 	Rectangle(dc, left, top, right, bottom);
// 	StrokePath(dc);
// 	// DeleteObject(dc, pen)
// }

// void gRePaint(gHANDLE h)
// {
// 	PostMessage(h, WM_PAINT, 0, 0);
// }
#endif