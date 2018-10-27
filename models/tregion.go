package models

type TreeRegion struct {
	RegionID   int
	RegionName string
	Children   []*TreeRegion
}

func (tr *TreeRegion) GetTreeRegionAccording(trs []*TreeRegion, regionID int) *TreeRegion {
	var item *TreeRegion
	var target *TreeRegion

	for i := 0; i < len(trs); i++ {
		item = trs[i]

		if item.RegionID == regionID {
			target = item
			break
		} else {
			var children = item.Children
			target = tr.GetTreeRegionAccording(children, regionID)

			if target != nil {
				break
			}
		}

	}

	return target
}
