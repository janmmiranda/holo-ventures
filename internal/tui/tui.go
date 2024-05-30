package tui

import "github.com/charmbracelet/bubbles/list"

type ConvertibleToListItem interface {
	list.Item
}

func convertToItems[T ConvertibleToListItem](elements []T) []list.Item {
	items := make([]list.Item, len(elements))
	for i, element := range elements {
		items[i] = element
	}
	return items
}
