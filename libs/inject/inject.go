package inject

import (
	"container/list"
	"fmt"
	"reflect"
	"sync"

	"github.com/BinacsLee/server/libs/log"
)

const (
	InjectTag_Name      = "inject-name"
	InjectTag_Required  = "inject-required"
	InjectTag_Strategy  = "inject-strategy"
	InjectTagFlag_True  = "true"
	InjectTagFlag_False = "flase"
)

var (
	globalContainer *InjectContainer
	logger          log.Logger
)

func init() {
	globalContainer = &InjectContainer{
		objMap:                make(map[string]*ObjInfo, 16),
		registList:            list.New(),
		fieldMatchStrategyMap: strategyMap,
	}
}

type InjectBeforeVisitor interface {
	BeforeInject()
}
type InjectAfterVisitor interface {
	AfterInject() error
}

type Properties interface {
	getValue(key string) string
}
type ObjFactory interface {
	createObj() interface{}
}

func Regist(name string, obj interface{}) {
	globalContainer.regist(name, obj)
}
func DoInject() error {
	return globalContainer.inject()
}
func InjectReport() string {
	return globalContainer.InjectReport()
}

func SetLogger(l log.Logger) {
	logger = l
}

func isPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr
}
func isStructPtr(t reflect.Type) bool {
	return isPtr(t) && t.Elem().Kind() == reflect.Struct
}
func realType(t reflect.Type) reflect.Type {
	if isStructPtr(t) {
		return t.Elem()
	}
	return t
}
func injectFieldCheck(targetType, instanceType reflect.Type) bool {
	if targetType.Kind() == reflect.Interface {
		if instanceType.AssignableTo(targetType) {
			return true
		}
		fmt.Printf("%v not AssignableTo %v\n", instanceType.Name(), targetType.Name())
		return false
	}
	if isStructPtr(targetType) {
		targetRealType := realType(targetType)
		instanceRealType := realType(instanceType)
		if instanceRealType.AssignableTo(targetRealType) {
			return true
		}
		fmt.Printf("%v not AssignableTo %v (isStructPtr)\n", instanceType.Name(), targetType.Name())
		return false
	}
	fmt.Printf("injectFieldCheck unknow\n")
	return false
}
func newInjectFieldInfo(tag reflect.StructTag) (ret InjectFieldInfo, ok bool) {
	injectTag, ok := tagParse(tag)
	ret.tag = injectTag
	return ret, ok
}
func tagParse(tag reflect.StructTag) (InjectTag, bool) {
	ok := false
	var foundtag bool
	ret := newDefaultTag()
	if ret.name, foundtag = tag.Lookup(InjectTag_Name); foundtag {
		ok = true
	} else {
		return ret, false
	}
	if tagValue, foundtag := tag.Lookup(InjectTag_Required); foundtag {
		if tagValue == InjectTagFlag_False {
			ret.required = false
		}
	}
	if tagValue, foundtag := tag.Lookup(InjectTag_Strategy); foundtag {
		ret.strategy = tagValue
	}
	return ret, ok
}
func doInject(fieldinfo *InjectFieldInfo, objInfo *ObjInfo) {
	fmt.Printf("doInject fieldName = %s %v %v %v\n", fieldinfo.fieldName, fieldinfo.reflectValue.CanSet(), fieldinfo.reflectValue.Kind(), objInfo.objDefination.reflectType.AssignableTo(fieldinfo.reflectType))
	fieldinfo.reflectValue.Set(objInfo.objDefination.reflectValue)
	fieldinfo.objInfo = objInfo
}

type InjectTag struct {
	name     string
	required bool
	strategy string
}

func newDefaultTag() InjectTag {
	return InjectTag{
		name:     "",
		required: true,
		strategy: FieldMatchStrategy_Com_NameType,
	}
}

type InjectFieldInfo struct {
	fieldName    string
	tag          InjectTag
	reflectType  reflect.Type
	reflectValue reflect.Value
	objInfo      *ObjInfo
}

type ObjDefination struct {
	reflectType  reflect.Type
	reflectValue reflect.Value
	injectList   []InjectFieldInfo
}

func (def *ObjDefination) defination(obj interface{}) error {
	reflectType := reflect.TypeOf(obj)
	if !isStructPtr(reflectType) {
		return fmt.Errorf("just support struct ptr")
	}
	def.reflectType = reflectType
	def.reflectValue = reflect.ValueOf(obj)
	def.injectList = make([]InjectFieldInfo, 0)

	fieldNum := def.reflectType.Elem().NumField()
	for i := 0; i < fieldNum; i++ {
		fieldtag := def.reflectType.Elem().Field(i).Tag
		if fieldInfo, ok := newInjectFieldInfo(fieldtag); ok {
			fieldInfo.reflectValue = def.reflectValue.Elem().Field(i)
			fieldInfo.reflectType = fieldInfo.reflectValue.Type()
			fieldInfo.fieldName = def.reflectType.Elem().Field(i).Name
			def.injectList = append(def.injectList, fieldInfo)
		}
	}
	return nil
}

type ObjInfo struct {
	name           string
	order          int32
	objDefination  ObjDefination
	injectComplete bool
	instance       interface{}
}

func newObjInfo(name string, order int32, obj interface{}) (*ObjInfo, error) {
	newObjInfo := ObjInfo{name: name, instance: obj, order: order}
	if err := newObjInfo.objDefination.defination(obj); err != nil {
		return nil, err
	}
	return &newObjInfo, nil
}

type InjectContainer struct {
	objMap                map[string]*ObjInfo
	definationMap         map[reflect.Type]*ObjDefination
	fieldMatchStrategyMap FieldMatchStrategyMap
	registList            *list.List
	order                 int32
	mutex                 sync.Mutex
}

func (ic *InjectContainer) regist(name string, obj interface{}) error {
	if len(name) <= 0 || obj == nil {
		return fmt.Errorf("invalid Param")
	}
	if err := ic.canRegisterCheck(obj); err != nil {
		fmt.Printf("Object [%s] Regist fail: %s\n", name, err.Error())
		return err
	}
	ic.mutex.Lock()
	defer ic.mutex.Unlock()
	if _, ok := ic.objMap[name]; ok {
		return fmt.Errorf("object name %s already exist", name)
	}
	ic.order += 1
	newObj, err := newObjInfo(name, ic.order, obj)
	if err != nil {
		return err
	}
	ic.registList.PushBack(newObj)
	ic.objMap[name] = newObj
	return nil
}

func (ic *InjectContainer) canRegisterCheck(obj interface{}) error {
	if !isPtr(reflect.TypeOf(obj)) {
		return fmt.Errorf("Regist Just support type ptr")
	}
	return nil
}

func (ic *InjectContainer) inject() error {
	ic.mutex.Lock()
	defer ic.mutex.Unlock()
	pedding := make([]*ObjInfo, 0)
	elem := ic.registList.Front()
	for elem != nil {
		objInfo := elem.Value.(*ObjInfo)
		if objInfo.injectComplete {
			continue
		}

		ic.invokeInjectBefore(objInfo)
		if err := ic.injectFields(objInfo); err != nil {
			fmt.Printf("inject fail, objname=%v\n", objInfo.name)
			return err
		}
		pedding = append(pedding, objInfo)
		elem = elem.Next()
	}
	return ic.invokeInjectAfter(pedding)
}

func (ic *InjectContainer) invokeInjectBefore(obj *ObjInfo) {
	if visitor, ok := obj.instance.(InjectBeforeVisitor); ok {
		visitor.BeforeInject()
	}
}
func (ic *InjectContainer) invokeInjectAfter(objList []*ObjInfo) error {
	for _, obj := range objList {
		if visitor, ok := obj.instance.(InjectAfterVisitor); ok {
			err := visitor.AfterInject()
			if err != nil {
				return err
			}
		}
		obj.injectComplete = true
	}
	return nil
}
func (ic *InjectContainer) injectFields(objInfo *ObjInfo) error {
	if objInfo.instance == nil {
		return nil
	}
	listlen := len(objInfo.objDefination.injectList)
	for i := 0; i < listlen; i++ {
		injectObj, err := ic.findObjByTFieldInfo(&(objInfo.objDefination.injectList[i]))
		if err != nil {
			return err
		}
		if injectObj != nil {
			fmt.Printf("Object [%s] Inject field[%s] -> %s\n", objInfo.name, objInfo.objDefination.injectList[i].fieldName, injectObj.name)
			doInject(&(objInfo.objDefination.injectList[i]), injectObj)
		} else {
			if objInfo.objDefination.injectList[i].tag.required {
				fmt.Printf("Object [%s] Inject fail field[%s] -> not found\n", objInfo.name, objInfo.objDefination.injectList[i].fieldName)
				return fmt.Errorf("Object [%s] Inject fail field[%s] -> not found\n", objInfo.name, objInfo.objDefination.injectList[i].fieldName)
			} else {
				fmt.Printf("Object [%s] Inject fail field[%s] -> not found (not required)\n", objInfo.name, objInfo.objDefination.injectList[i].fieldName)
			}
		}
	}
	return nil
}
func (ic *InjectContainer) findObjByTFieldInfo(fieldinfo *InjectFieldInfo) (*ObjInfo, error) {
	depObj, err := ic.fieldMatchStrategyMap.getStrategy(fieldinfo.tag.strategy).findObjByTFieldInfo(ic.objMap, fieldinfo)
	if err != nil {
		return nil, err
	}
	if depObj == nil {
		return nil, nil
	}
	if !injectFieldCheck(fieldinfo.reflectType, depObj.objDefination.reflectType) {
		fmt.Printf("injectFieldCheck fail %v can not inject to %v\n", depObj.objDefination.reflectType, fieldinfo.reflectType)
		return nil, fmt.Errorf("injectFieldCheck fail %v can not inject to %v", depObj.objDefination.reflectType, fieldinfo.reflectType)
	}
	return depObj, nil
}

func (ic *InjectContainer) InjectReport() string {
	return ""
}
