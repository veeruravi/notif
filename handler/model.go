package handler

type SegmentPayload struct {
	Segment  int 		`form:"segment_id" binding:"required"`
	Notification int 	`form:"notification_id" binding:"required"`
	Template     int    `form:"template_id"`
}
