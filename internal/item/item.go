package item

type Item struct {
	ID                         uint32
	Flags                      uint16
	Type                       uint8
	Material                   uint8
	Name                       string `encrypt:"true"` // only encrypted in versions >= 3
	Texture                    string
	TextureHash                uint32
	VisualEffect               uint8
	Cook                       int32
	TextureX                   uint8
	TextureY                   uint8
	Storage                    uint8
	Layer                      int8
	Collision                  uint8
	HitsToDestroy              uint8
	HealTime                   int32
	BodyPart                   uint8
	Rarity                     uint16
	MaxCanHold                 uint8
	ExtraFile                  string
	ExtraFileHash              uint32
	AnimMS                     int32
	PetName                    string `version:"4"`
	PetSubname                 string `version:"4"`
	PetEndname                 string `version:"4"`
	PetAbility                 string `version:"5"`
	SeedBG                     uint8
	SeedFG                     uint8
	TreeBG                     uint8
	TreeFG                     uint8
	ColorBG                    uint32
	ColorFG                    uint32
	Seed1                      uint16
	Seed2                      uint16
	SecondsToBloom             uint32
	FXFlags                    uint32    `version:"7"`
	MultiAnimData              string    `version:"7"`
	OverlayObjectTexture       string    `version:"8"`
	MultiAnim2Data             string    `version:"8"`
	DualLayer                  uint64    `version:"8"`
	Flags2                     uint32    `version:"9"`
	ClientData                 [15]int32 `version:"9"`
	TileRange                  uint32    `version:"10"`
	PileSize                   uint32    `version:"10"`
	PunchParams                string    `version:"11"`
	ExtraSlotCount             int32     `version:"12"`
	ExtraSlots                 [9]uint8  `version:"12"`
	LightIntensity             int32     `version:"13"`
	VariantItem                int32     `version:"14"`
	CustomChair                uint8     `version:"15"`
	ChairPlayerOffsetX         int32     `version:"15"`
	ChairPlayerOffsetY         int32     `version:"15"`
	ChairOverlayTextureX       int32     `version:"15"`
	ChairOverlayTextureY       int32     `version:"15"`
	ChairOverlayTextureOffsetX int32     `version:"15"`
	ChairOverlayTextureOffsetY int32     `version:"15"`
	ChairOverlayTexture        string    `version:"15"`
	ConfigName                 string    `version:"16"`
	OtherPlayerHitParticle     int32     `version:"17"`
	ConfigHash                 uint32    `version:"18"`
	Unknown1                   uint8     `version:"19"`
	Unknown2                   int32     `version:"19"`
	Unknown3                   float32   `version:"19"`
	Unknown4                   uint8     `version:"20"`
	Unknown5                   uint8     `version:"21"`
}
