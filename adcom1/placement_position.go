package adcom1

// PlacementPosition represents placement positions as a relative measure of visibility or prominence.
// This table has values derived from the TAG Inventory Quality Guidelines (IQG).
type PlacementPosition int8

// Placement positions.
const (
	PositionUnknown                 PlacementPosition = 0  // Unknown
	PositionAboveFold               PlacementPosition = 1  // Above The Fold
	PositionLocked                  PlacementPosition = 2  // Locked (i.e., fixed position)
	PositionBelowFold               PlacementPosition = 3  // Below The Fold
	PositionHeader                  PlacementPosition = 4  // Header
	PositionFooter                  PlacementPosition = 5  // Footer
	PositionSideBar                 PlacementPosition = 6  // Sidebar
	PositionFullScreen              PlacementPosition = 7  // Fullscreen
	PositionPartial                 PlacementPosition = 8  // Partial Screen
	PositionTopLeft                 PlacementPosition = 9  // Top Left
	PositionTopRight                PlacementPosition = 10 // Top Right
	PositionFrame                   PlacementPosition = 11 // Frame the content
	PositionDoubleBox               PlacementPosition = 12 // Double box
	PositionDoubleBoxWithBackground PlacementPosition = 13 // Double box with background
	PositionBottomLeft              PlacementPosition = 14 // Bottom Left
	PositionBottomRight             PlacementPosition = 15 // Bottom Right
	PositionLShape                  PlacementPosition = 16 // L shape
	PositionReverseLShape           PlacementPosition = 17 // Reversed L shape
)

// Ptr returns pointer to own value.
func (p PlacementPosition) Ptr() *PlacementPosition {
	return &p
}

// Val safely dereferences pointer, returning default value (PositionUnknown) for nil.
func (p *PlacementPosition) Val() PlacementPosition {
	if p == nil {
		return PositionUnknown
	}
	return *p
}
