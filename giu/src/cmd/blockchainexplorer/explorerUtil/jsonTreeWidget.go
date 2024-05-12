package explorerutil

import (
	"encoding/json"
	"fmt"
	"sort"

	g "github.com/AllenDang/giu"
)

func DrawJsonTreeWidgetWithRoot(jsonString string, rootName string) []g.Widget {
	if jsonString == "" {
		return nil
	}

	var data interface{}
	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
		SetGlobalError(fmt.Errorf("failed to unmarshal json: %v", err))
		return nil
	}

	return []g.Widget{
		g.TreeNode(rootName).Layout(makeJsonTreeWidgetRecursively(data)...),
	}
}

func DrawJsonTreeWidget(jsonString string) []g.Widget {
	if jsonString == "" {
		return nil
	}

	var data interface{}
	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
		SetGlobalError(fmt.Errorf("failed to unmarshal json: %v", err))
		return nil
	}

	return makeJsonTreeWidgetRecursively(data)
}

func makeJsonTreeWidgetRecursively(data interface{}) []g.Widget {
	var widgets []g.Widget = make([]g.Widget, 0)

	switch v := data.(type) {
	case map[string]interface{}:
		mapp := v
		keys := make([]string, 0, len(mapp))
		for k := range mapp {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, key := range keys {
			value := mapp[key]
			widgets = append(widgets, g.Row(
				g.TreeNode(key).Layout(makeJsonTreeWidgetRecursively(value)...),
			))
		}
	case []interface{}:
		arr := v
		for i, value := range arr {
			widgets = append(widgets, g.Row(
				g.TreeNode(fmt.Sprintf("%d", i)).Layout(makeJsonTreeWidgetRecursively(value)...),
			))
		}
	default:
		widgets = append(widgets, g.Row(
			DrawSuperTextWidget(fmt.Sprintf("%v", data)),
		))
	}

	return widgets
}
