package activitymonitor

import (
	"fmt"
	"smartfeeder/feeder"
	"smartfeeder/zone"
)

type ActivityMonitor struct {
	ZoneCheck                 zone.ZoneChecker
	DistributionAnimalsInZone feeder.DistributionAnimalsInZone
}

func (a *ActivityMonitor) FeedDelivery() {
	fmt.Println("test")
}
