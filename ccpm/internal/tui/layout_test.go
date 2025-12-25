package tui

import "testing"

func TestCalculateLayout(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		height   int
		expected LayoutSize
	}{
		{"too small width", 79, 24, LayoutTooSmall},
		{"too small height", 80, 23, LayoutTooSmall},
		{"minimum small", 80, 24, LayoutSmall},
		{"upper small", 99, 30, LayoutSmall},
		{"minimum medium", 100, 30, LayoutMedium},
		{"upper medium", 119, 35, LayoutMedium},
		{"minimum large", 120, 40, LayoutLarge},
		{"very large", 200, 50, LayoutLarge},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateLayout(tt.width, tt.height)
			if result.Size != tt.expected {
				t.Errorf("For %dx%d: expected %v, got %v",
					tt.width, tt.height, tt.expected, result.Size)
			}
		})
	}
}

func TestLayoutTooSmall(t *testing.T) {
	layout := CalculateLayout(70, 20)

	if !layout.IsTooSmall() {
		t.Error("Expected IsTooSmall to return true")
	}

	if layout.SizeString() != "too-small" {
		t.Errorf("Expected SizeString 'too-small', got '%s'", layout.SizeString())
	}
}

func TestLayoutSmallHidesActivity(t *testing.T) {
	layout := CalculateLayout(80, 24)

	if layout.ActivityVisible {
		t.Error("Small layout should hide activity")
	}

	if !layout.UseAbbreviatedLabels {
		t.Error("Small layout should use abbreviated labels")
	}

	if layout.SideBySide {
		t.Error("Small layout should not be side-by-side")
	}

	if layout.ProgressBarLen != 30 {
		t.Errorf("Expected progress bar len 30, got %d", layout.ProgressBarLen)
	}
}

func TestLayoutMediumShowsActivity(t *testing.T) {
	layout := CalculateLayout(100, 30)

	if !layout.ActivityVisible {
		t.Error("Medium layout should show activity")
	}

	if layout.UseAbbreviatedLabels {
		t.Error("Medium layout should not use abbreviated labels")
	}

	if layout.SideBySide {
		t.Error("Medium layout should not be side-by-side")
	}

	if layout.ProgressBarLen != 50 {
		t.Errorf("Expected progress bar len 50, got %d", layout.ProgressBarLen)
	}
}

func TestLayoutLargeEnablesSideBySide(t *testing.T) {
	layout := CalculateLayout(120, 40)

	if !layout.SideBySide {
		t.Error("Large layout should enable side-by-side")
	}

	if !layout.ActivityVisible {
		t.Error("Large layout should show activity")
	}

	if layout.UseAbbreviatedLabels {
		t.Error("Large layout should not use abbreviated labels")
	}

	if layout.ProgressBarLen != 60 {
		t.Errorf("Expected progress bar len 60, got %d", layout.ProgressBarLen)
	}
}

func TestLayoutExtraTallAddsItems(t *testing.T) {
	normalLayout := CalculateLayout(100, 30)
	tallLayout := CalculateLayout(100, 45)

	if tallLayout.TaskItemCount <= normalLayout.TaskItemCount {
		t.Error("Extra tall layout should have more task items")
	}
}

func TestLayoutDimensionsCalculated(t *testing.T) {
	layout := CalculateLayout(120, 40)

	if layout.Width != 120 {
		t.Errorf("Expected width 120, got %d", layout.Width)
	}

	if layout.Height != 40 {
		t.Errorf("Expected height 40, got %d", layout.Height)
	}

	if layout.HeaderHeight != 8 {
		t.Errorf("Expected header height 8, got %d", layout.HeaderHeight)
	}

	if layout.FooterHeight != 2 {
		t.Errorf("Expected footer height 2, got %d", layout.FooterHeight)
	}

	expectedContent := 40 - 8 - 2
	if layout.ContentHeight != expectedContent {
		t.Errorf("Expected content height %d, got %d", expectedContent, layout.ContentHeight)
	}
}

func TestLayoutSizeString(t *testing.T) {
	tests := []struct {
		width    int
		height   int
		expected string
	}{
		{70, 20, "too-small"},
		{80, 24, "small"},
		{100, 30, "medium"},
		{120, 40, "large"},
	}

	for _, tt := range tests {
		layout := CalculateLayout(tt.width, tt.height)
		if layout.SizeString() != tt.expected {
			t.Errorf("For %dx%d: expected '%s', got '%s'",
				tt.width, tt.height, tt.expected, layout.SizeString())
		}
	}
}
