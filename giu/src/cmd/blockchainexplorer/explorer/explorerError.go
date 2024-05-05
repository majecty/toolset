package explorer

import g "github.com/AllenDang/giu"

var (
	GlobalError error
)

func SetGlobalError(err error) {
	GlobalError = err
}

func DrawGlobalErrorWidgets() []g.Widget {
	if GlobalError != nil {
		return []g.Widget{
			g.Label("Error: " + GlobalError.Error()),
			g.Button("Reset").OnClick(func() {
				GlobalError = nil
			}),
		}
	}

	return nil
}
