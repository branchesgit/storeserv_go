package models

import "github.com/astaxie/beego/orm"

type Region struct {
	Id       int
	RegionID int `orm:"column(RegionID)"`
	Name     string
	Pid      int
}

func init() {
	orm.RegisterModel(new(Region))
}

const ROOTID = 1

// 将数据转换成树状结构进行输出；
func (r *Region) ChangeRegionsToTreeRegions(rs []*Region, num int) []*TreeRegion {
	var regions = make([]*TreeRegion, 0)
	var re *Region
	var pid int
	var idx = 0

	for len(rs) != 0 {
		// 需要处理的地区对象，应该放到哪个 TreeRegion集合里去。
		// 遍历完循环的时候，再重新来过，
		if idx == len(rs) {
			idx = 0
		}

		re = rs[idx]
		tr := new(TreeRegion)
		tr.RegionID = re.RegionID
		tr.RegionName = re.Name
		tr.Children = make([]*TreeRegion, 0)
		pid = re.Pid

		if pid == ROOTID {
			regions = append(regions, tr)
			copy(rs[idx:], rs[idx+1:])
			rs = rs[:len(rs)-1]
			idx--
		} else {
			// 找到父节点，那么可以把他挂进父节点；
			var target = tr.GetTreeRegionAccording(regions, pid)
			if target != nil {
				var children = target.Children
				children = append(children, tr)
				target.Children = children
				copy(rs[idx:], rs[idx+1:])
				rs = rs[:len(rs)-1]
				idx--
			}
		}
		idx++
	}

	return regions
}
