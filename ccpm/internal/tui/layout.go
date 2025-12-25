package tui

type LayoutSize int

const (
	LayoutTooSmall LayoutSize = iota
	LayoutSmall
	LayoutMedium
	LayoutLarge
)

const (
	MinWidth  = 80
	MinHeight = 24
)

type LayoutDimensions struct {
	Size LayoutSize

	Width  int
	Height int

	HeaderHeight  int
	FooterHeight  int
	ContentHeight int

	EpicCardWidth  int
	EpicCardHeight int

	TaskListWidth  int
	TaskListHeight int
	TaskItemCount  int

	ActivityWidth   int
	ActivityHeight  int
	ActivityVisible bool

	SideBySide           bool
	ProgressBarLen       int
	UseAbbreviatedLabels bool
}

func CalculateLayout(width, height int) LayoutDimensions {
	d := LayoutDimensions{
		Width:  width,
		Height: height,
	}

	if width < MinWidth || height < MinHeight {
		d.Size = LayoutTooSmall
		return d
	}

	switch {
	case width < 100:
		d.Size = LayoutSmall
	case width < 120:
		d.Size = LayoutMedium
	default:
		d.Size = LayoutLarge
	}

	d.HeaderHeight = 8
	d.FooterHeight = 2
	d.ContentHeight = height - d.HeaderHeight - d.FooterHeight

	switch d.Size {
	case LayoutSmall:
		d.calculateSmallLayout()
	case LayoutMedium:
		d.calculateMediumLayout()
	case LayoutLarge:
		d.calculateLargeLayout()
	}

	if height >= 40 {
		d.TaskItemCount += 4
		if d.ActivityVisible {
			d.ActivityHeight += 4
		}
	}

	return d
}

func (d *LayoutDimensions) calculateSmallLayout() {
	d.EpicCardWidth = d.Width - 4
	d.EpicCardHeight = 10

	d.TaskListWidth = d.Width - 4
	d.TaskListHeight = d.ContentHeight - d.EpicCardHeight - 2
	d.TaskItemCount = min(6, d.TaskListHeight-2)

	d.ActivityVisible = false
	d.ProgressBarLen = 30
	d.SideBySide = false
	d.UseAbbreviatedLabels = true
}

func (d *LayoutDimensions) calculateMediumLayout() {
	d.EpicCardWidth = d.Width - 4
	d.EpicCardHeight = 12

	d.TaskListWidth = d.Width - 4
	d.ActivityHeight = 6
	d.TaskListHeight = d.ContentHeight - d.EpicCardHeight - d.ActivityHeight - 4
	d.TaskItemCount = min(8, d.TaskListHeight-2)

	d.ActivityWidth = d.Width - 4
	d.ActivityVisible = true
	d.ProgressBarLen = 50
	d.SideBySide = false
	d.UseAbbreviatedLabels = false
}

func (d *LayoutDimensions) calculateLargeLayout() {
	d.EpicCardWidth = d.Width - 4
	d.EpicCardHeight = 12

	d.TaskListWidth = (d.Width - 8) / 2
	d.ActivityWidth = (d.Width - 8) / 2
	d.TaskListHeight = d.ContentHeight - d.EpicCardHeight - 2
	d.ActivityHeight = d.TaskListHeight
	d.TaskItemCount = min(12, d.TaskListHeight-2)

	d.ActivityVisible = true
	d.ProgressBarLen = 60
	d.SideBySide = true
	d.UseAbbreviatedLabels = false
}

func (d LayoutDimensions) IsTooSmall() bool {
	return d.Size == LayoutTooSmall
}

func (d LayoutDimensions) SizeString() string {
	switch d.Size {
	case LayoutTooSmall:
		return "too-small"
	case LayoutSmall:
		return "small"
	case LayoutMedium:
		return "medium"
	case LayoutLarge:
		return "large"
	default:
		return "unknown"
	}
}
