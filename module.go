package ui

// all module base st
// type Module struct {
// 	handle Handle
// }

// func CreateModule(width, height, flags int, title string, parent Handle) (*Module, error) {
// 	h, err := ccreateWindow(width, height, title, flags, parent)
// 	if err != nil {
// 		return nil, err
// 	}

// 	m := &Module{
// 		handle: h,
// 	}

// 	return m
// }

type Module interface {
	Handle() Handle
}

func CreateModule(parent Handle) Module {
	return nil
}

// ========== module base ================
type module struct {
	handle Handle
}

func (m *module) Handle() Handle {
	return m.handle
}
