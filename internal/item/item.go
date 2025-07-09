package item

import (
	"reflect"
	"strconv"

	"github.com/teokt/gt-items/internal/math"
	"github.com/teokt/gt-items/internal/memory"
)

type Item struct {
	ID                      uint32
	Flags                   ItemFlags
	Type                    ItemType
	Material                ItemMaterial
	Name                    string
	Texture                 string
	TextureHash             uint32
	VisualEffect            TileVisualEffect
	Cook                    int32
	TexturePos              math.Vec2[uint8]
	Storage                 TileStorage
	Layer                   int8
	Collision               TileCollision
	HitsToDestroy           uint8
	HealTime                int32
	BodyPart                BodyPart
	Rarity                  uint16
	MaxCanHold              uint8
	ExtraFile               string
	ExtraFileHash           uint32
	AnimMS                  int32
	PetName                 string `version:"4"`
	PetSubname              string `version:"4"`
	PetEndname              string `version:"4"`
	PetAbility              string `version:"5"`
	SeedBG                  uint8
	SeedFG                  uint8
	TreeBG                  uint8
	TreeFG                  uint8
	ColorBG                 uint32
	ColorFG                 uint32
	Seed1                   uint16
	Seed2                   uint16
	SecondsToBloom          uint32
	FXFlags                 ItemFXFlags                 `version:"7"`
	MultiAnimData           string                      `version:"7"`
	OverlayObjectTexture    string                      `version:"8"`
	MultiAnim2Data          string                      `version:"8"`
	DualLayer               math.Vec2[uint32]           `version:"8"`
	Flags2                  ItemFlags2                  `version:"9"`
	ClientData              ItemClientData              `version:"9"`
	TileRange               uint32                      `version:"10"`
	StorageSize             uint32                      `version:"10"`
	PunchParameters         string                      `version:"11"`
	ExtraSlots              ItemExtraSlots              `version:"12"`
	LightSourceMod          int32                       `version:"13"`
	VariantItem             int32                       `version:"14"`
	CustomChair             ItemCustomChair             `version:"15"`
	ConfigName              string                      `version:"16"`
	OtherPlayerHitParticle  int32                       `version:"17"`
	ConfigHash              uint32                      `version:"18"`
	RandomSpriteReplacement ItemRandomSpriteReplacement `version:"19"`
	Unknown                 uint8                       `version:"20"`
	IsTransform             uint8                       `version:"21"`
}

type ItemClientData struct {
	Collision          int32
	StartFrame         int32
	NumWalkCycle       int32
	NumShootCycle      int32
	NumProjectileCycle int32
	ProjectileOffsetLR math.Vec2[int32]
	ProjectileOffsetU  math.Vec2[int32]
	ProjectileOffsetD  math.Vec2[int32]
	Unknown            [4]int32
}

type ItemExtraSlots struct {
	Count     int32
	BodyParts [9]uint8
}

type ItemCustomChair struct {
	Enabled          uint8
	PlayerOffset     math.Vec2[int32]
	ArmTexturePos    math.Vec2[int32]
	ArmTextureOffset math.Vec2[int32]
	ArmTexture       string
}

type ItemRandomSpriteReplacement struct {
	Enabled uint8
	Offset  int32
	Chance  float32
}

func (i *Item) Deserialize(r *memory.Reader, version int) error {
	v := reflect.ValueOf(i).Elem()
	t := v.Type()

	for idx := range v.NumField() {
		field := v.Field(idx)
		if !field.CanSet() {
			continue
		}

		fieldType := t.Field(idx)
		versionTag := fieldType.Tag.Get("version")
		fieldVersion, _ := strconv.Atoi(versionTag)

		if fieldVersion > version {
			continue
		}

		fieldPtr := field.Addr().Interface()

		// TODO: some deserializer class where can put custom handlers for specific fields like this
		if fieldType.Name == "Name" && version >= 3 {
			if err := r.ReadEncryptedString(fieldPtr.(*string), int(i.ID), "PBG892FXX982ABC*"); err != nil {
				return err
			}
			continue
		}

		if err := r.Read(fieldPtr); err != nil {
			return err
		}
	}

	return nil
}
