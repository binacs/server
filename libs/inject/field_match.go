package inject

import "fmt"

const (
	FieldMatchStrategy_NameOnly     = "NameOnly"
	FieldMatchStrategy_TypeOnly     = "TypeOnly"
	FieldMatchStrategy_Com_NameType = "ComNameType"
	DEFAULT_STRATEGY                = FieldMatchStrategy_NameOnly
)

func init() {
	strategyMap = getStrategy()
	strategyMap[DEFAULT_STRATEGY] = &NameMatchStrategy{}
}

type NameMatchStrategy struct {
}

func (n *NameMatchStrategy) name() string {
	return FieldMatchStrategy_NameOnly
}

func (n *NameMatchStrategy) findObjByTFieldInfo(objMap map[string]*ObjInfo, fieldInfo *InjectFieldInfo) (*ObjInfo, error) {
	if n == nil {
		return nil, nil
	}
	return objMap[fieldInfo.tag.name], nil
}

type TypeMatchStrategy struct {
}

func (t *TypeMatchStrategy) name() string {
	return FieldMatchStrategy_TypeOnly
}

func (t *TypeMatchStrategy) findObjByTFieldInfo(objMap map[string]*ObjInfo, fieldInfo *InjectFieldInfo) (*ObjInfo, error) {
	var ret *ObjInfo
	for _, v := range objMap {
		if v.objDefination.reflectType.AssignableTo(fieldInfo.reflectType) {
			if ret != nil {
				return nil, fmt.Errorf("Multiple match %v %v", ret.name, v.name)
			}
			ret = v
		}
	}
	return ret, nil
}

type MatchStrategyComNameType struct {
}

func (m *MatchStrategyComNameType) name() string {
	return FieldMatchStrategy_Com_NameType
}

func (m *MatchStrategyComNameType) findObjByTFieldInfo(objMap map[string]*ObjInfo, fieldInfo *InjectFieldInfo) (*ObjInfo, error) {
	sty, err := strategyMap.getStrategy(FieldMatchStrategy_NameOnly).findObjByTFieldInfo(objMap, fieldInfo)
	if err != nil {
		return nil, err
	}
	if sty != nil {
		return sty, nil
	}
	sty, err = strategyMap.getStrategy(FieldMatchStrategy_TypeOnly).findObjByTFieldInfo(objMap, fieldInfo)
	if err != nil {
		return nil, err
	}
	return sty, err
}

type FieldMatchStrategy interface {
	name() string
	findObjByTFieldInfo(objMap map[string]*ObjInfo, fieldInfo *InjectFieldInfo) (*ObjInfo, error)
}

type FieldMatchStrategyMap map[string]FieldMatchStrategy

func (f FieldMatchStrategyMap) regist(fms FieldMatchStrategy) error {
	if _, ok := f[fms.name()]; ok {
		return fmt.Errorf("Strategy named [%v] already exist", fms.name())
	}
	f[fms.name()] = fms
	return nil
}

func (f FieldMatchStrategyMap) getStrategy(name string) FieldMatchStrategy {
	ret := f[name]
	if ret == nil {
		return f[DEFAULT_STRATEGY]
	} else {
		return ret
	}
}

func (f FieldMatchStrategyMap) hasStrategy(name string) bool {
	_, ok := f[name]
	return ok
}

var strategyMap FieldMatchStrategyMap

func getStrategy() FieldMatchStrategyMap {
	ret := make(FieldMatchStrategyMap)
	ret.regist(&NameMatchStrategy{})
	ret.regist(&TypeMatchStrategy{})
	ret.regist(&MatchStrategyComNameType{})
	return ret
}
