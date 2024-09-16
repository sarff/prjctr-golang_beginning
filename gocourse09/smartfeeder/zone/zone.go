package zone

type Zone struct {
	ZoneChecker ZoneChecker
}

type ZoneChecker interface {
	CheckZone()
}

func (z Zone) CheckZone() {}
