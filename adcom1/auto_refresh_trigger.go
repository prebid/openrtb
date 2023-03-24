package adcom1

// AutoRefreshTrigger represents a list of triggers support by a placement or required by an ad.
type AutoRefreshTrigger int8

// Triggers support by a placement or required by an ad.
const (
	AutoRefreshTriggerUnknown    AutoRefreshTrigger = 0
	AutoRefreshTriggerUserAction AutoRefreshTrigger = 1 // Refresh triggered by user-initiated action such as scrolling.
	AutoRefreshTriggerEvent      AutoRefreshTrigger = 2 // Event-driven content change. For example, ads refresh when the football game score changes on the page.
	AutoRefreshTriggerTime       AutoRefreshTrigger = 3 // Time-based refresh. Ads refresh on a predefined time interval even without user activity.
)
