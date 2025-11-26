package xlsx

import (
	"testing"
)

const (
	RoleDataXlsx = "D:\\project\\go\\BattleRabbit\\Server\\GameServer\\bin\\xlsx\\RoleData.xlsx"
)

var (
// PublicAttribute    = new(Attribute)
// AttributeFieldInfo = getFieldInfos(reflect.TypeOf(PublicAttribute).Elem())
// RoleDataConfig     = make(map[string]*RoleData)
)

//type RoleData struct {
//	Id         string    `xlsx:"id"`
//	Name       string    `xlsx:"name"`
//	Describe   string    `xlsx:"describe"`
//	Relation   string    `xlsx:"relation"`
//	Star_level int       `xlsx:"star_level"`
//	Level      int       `xlsx:"level"`
//	Exp        int       `xlsx:"exp"`
//	Attribute  Attribute `xlsx:"attribute"`
//}
//
//type Attribute struct {
//	Life_a          int     `xlsx:"life_a"`
//	Life_c          float32 `xlsx:"life_c"`
//	Defense_a       int     `xlsx:"defense_a"`
//	Defense_c       float32 `xlsx:"defense_c"`
//	Attack_a        int     `xlsx:"attack_a"`
//	Attack_c        float32 `xlsx:"attack_c"`
//	Attack_speed_a  int     `xlsx:"attack_speed_a"`
//	Attack_speed_c  float32 `xlsx:"attack_speed_c"`
//	Move_speed_a    int     `xlsx:"move_speed_a"`
//	Move_speed_c    float32 `xlsx:"move_speed_c"`
//	Dodge_b         int     `xlsx:"dodge_b"`
//	Dodge_c         int     `xlsx:"dodge_c"`
//	Critical_b      int     `xlsx:"critical_b"`
//	Critical_c      int     `xlsx:"critical_c"`
//	Buff_time_a     float32 `xlsx:"buff_time_a"`
//	Buff_time_c     float32 `xlsx:"buff_time_c"`
//	Boss_hurt_a     float32 `xlsx:"boss_hurt_a"`
//	Boss_hurt_c     float32 `xlsx:"boss_hurt_c"`
//	Critical_hurt_a float32 `xlsx:"critical_hurt_a"`
//	Critical_hurt_c float32 `xlsx:"critical_hurt_c"`
//	Gold_add_c      float32 `xlsx:"gold_add_c"`
//}

func TestLoadXlsxConfig(t *testing.T) {

	//var iface  = new(RoleData)
	//var iface  = &RoleData{}
	//err := LoadXlsxFile(RoleDataXlsx,iface,"Sheet2", func(v any) {
	//	val := v.(*RoleData)
	//	RoleDataConfig[val.Id] = val
	//})
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//m,err  := jsoniter.Marshal(RoleDataConfig)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Debug(string(m))

}
