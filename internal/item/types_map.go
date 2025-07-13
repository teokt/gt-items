package item

import (
	"fmt"

	"github.com/teokt/gt-items/internal/utils"
)

var itemTypeMap = map[ItemType]string{
	ItemTypeFist:                   "Fist",
	ItemTypeWrench:                 "Wrench",
	ItemTypeUserDoor:               "Userdoor",
	ItemTypeLock:                   "Lock",
	ItemTypeGems:                   "Gems",
	ItemTypeTreasure:               "Treasure",
	ItemTypeDeadly:                 "Deadly",
	ItemTypeTrampoline:             "Trampoline",
	ItemTypeConsumable:             "Consumable",
	ItemTypeGateway:                "Gateway",
	ItemTypeSign:                   "Sign",
	ItemTypeSfxWithExtraFrame:      "SfxWithExtraFrame",
	ItemTypeBoombox:                "Boombox",
	ItemTypeDoor:                   "Door",
	ItemTypePlatform:               "Platform",
	ItemTypeBedrock:                "Bedrock",
	ItemTypeLava:                   "Lava",
	ItemTypeNormal:                 "Normal",
	ItemTypeBackground:             "Background",
	ItemTypeSeed:                   "Seed",
	ItemTypeClothes:                "Clothes",
	ItemTypeNormalWithExtraFrame:   "NormalWithExtraFrame",
	ItemTypeBackgdSfxExtraFrame:    "BackgdSfxExtraFrame",
	ItemTypeBackBoombox:            "BackBoombox",
	ItemTypeBouncy:                 "Bouncy",
	ItemTypePointy:                 "Pointy",
	ItemTypePortal:                 "Portal",
	ItemTypeCheckpoint:             "Checkpoint",
	ItemTypeMusicnote:              "Musicnote",
	ItemTypeIce:                    "Ice",
	ItemTypeRaceFlag:               "RaceFlag",
	ItemTypeSwitcherroo:            "Switcherroo",
	ItemTypeChest:                  "Chest",
	ItemTypeMailbox:                "Mailbox",
	ItemTypeBulletin:               "Bulletin",
	ItemTypePinata:                 "Pinata",
	ItemTypeDice:                   "Dice",
	ItemTypeComponent:              "Component",
	ItemTypeProvider:               "Provider",
	ItemTypeLab:                    "Lab",
	ItemTypeAchievement:            "Achievement",
	ItemTypeWeatherMachine:         "WeatherMachine",
	ItemTypeScoreboard:             "Scoreboard",
	ItemTypeSungate:                "Sungate",
	ItemTypeProfile:                "Profile",
	ItemTypeDeadlyIfOn:             "DeadlyIfOn",
	ItemTypeHeartMonitor:           "HeartMonitor",
	ItemTypeDonationBox:            "DonationBox",
	ItemTypeToybox:                 "Toybox",
	ItemTypeMannequin:              "Mannequin",
	ItemTypeCamera:                 "Camera",
	ItemTypeMagicegg:               "Magicegg",
	ItemTypeTeam:                   "Team",
	ItemTypeGameGen:                "GameGen",
	ItemTypeXenonite:               "Xenonite",
	ItemTypeDressup:                "Dressup",
	ItemTypeCrystal:                "Crystal",
	ItemTypeBurglar:                "Burglar",
	ItemTypeCompactor:              "Compactor",
	ItemTypeSpotlight:              "Spotlight",
	ItemTypeWind:                   "Wind",
	ItemTypeDisplayBlock:           "DisplayBlock",
	ItemTypeVending:                "Vending",
	ItemTypeFishtank:               "Fishtank",
	ItemTypePetfish:                "Petfish",
	ItemTypeSolar:                  "Solar",
	ItemTypeForge:                  "Forge",
	ItemTypeGivingTree:             "GivingTree",
	ItemTypeGivingTreeStump:        "GivingTreeStump",
	ItemTypeSteampunk:              "Steampunk",
	ItemTypeSteamLavaIfOn:          "SteamLavaIfOn",
	ItemTypeSteamOrgan:             "SteamOrgan",
	ItemTypeTamagotchi:             "Tamagotchi",
	ItemTypeSewing:                 "Sewing",
	ItemTypeFlag:                   "Flag",
	ItemTypeLobsterTrap:            "LobsterTrap",
	ItemTypeArtcanvas:              "Artcanvas",
	ItemTypeBattleCage:             "BattleCage",
	ItemTypePetTrainer:             "PetTrainer",
	ItemTypeSteamEngine:            "SteamEngine",
	ItemTypeLockBot:                "LockBot",
	ItemTypeWeatherSpecial:         "WeatherSpecial",
	ItemTypeSpiritStorage:          "SpiritStorage",
	ItemTypeDisplayShelf:           "DisplayShelf",
	ItemTypeVipDoor:                "VipDoor",
	ItemTypeChalTimer:              "ChalTimer",
	ItemTypeChalFlag:               "ChalFlag",
	ItemTypeFishMount:              "FishMount",
	ItemTypePortrait:               "Portrait",
	ItemTypeWeatherSpecial2:        "WeatherSpecial2",
	ItemTypeFossil:                 "Fossil",
	ItemTypeFossilPrep:             "FossilPrep",
	ItemTypeDnaMachine:             "DnaMachine",
	ItemTypeBlaster:                "Blaster",
	ItemTypeValhowla:               "Valhowla",
	ItemTypeChemsynth:              "Chemsynth",
	ItemTypeChemtank:               "Chemtank",
	ItemTypeStorage:                "Storage",
	ItemTypeOven:                   "Oven",
	ItemTypeSuperMusic:             "SuperMusic",
	ItemTypeGeigercharge:           "Geigercharge",
	ItemTypeAdventureReset:         "AdventureReset",
	ItemTypeTombRobber:             "TombRobber",
	ItemTypeFaction:                "Faction",
	ItemTypeRedFaction:             "RedFaction",
	ItemTypeGreenFaction:           "GreenFaction",
	ItemTypeBlueFaction:            "BlueFaction",
	ItemTypeArtifact:               "Artifact",
	ItemTypeTrampolineMomentum:     "TrampolineMomentum",
	ItemTypeFishgotchiTank:         "FishgotchiTank",
	ItemTypeFishingBlock:           "FishingBlock",
	ItemTypeItemSucker:             "ItemSucker",
	ItemTypeItemPlanter:            "ItemPlanter",
	ItemTypeRobot:                  "Robot",
	ItemTypeCommand:                "Command",
	ItemTypeLuckyTicket:            "LuckyTicket",
	ItemTypeStatsBlock:             "StatsBlock",
	ItemTypeFieldNode:              "FieldNode",
	ItemTypeOuijaBoard:             "OuijaBoard",
	ItemTypeArchitectMachine:       "ArchitectMachine",
	ItemTypeStarship:               "Starship",
	ItemTypeAutodelete:             "Autodelete",
	ItemTypeBoombox2:               "Boombox2",
	ItemTypeAutoActionBreak:        "AutoActionBreak",
	ItemTypeAutoActionHarvest:      "AutoActionHarvest",
	ItemTypeAutoActionHarvestSuck:  "AutoActionHarvestSuck",
	ItemTypeLightningCloud:         "LightningCloud",
	ItemTypePhasedBlock:            "PhasedBlock",
	ItemTypeMud:                    "Mud",
	ItemTypeRootCutting:            "RootCutting",
	ItemTypePasswordStorage:        "PasswordStorage",
	ItemTypePhasedBlock2:           "PhasedBlock2",
	ItemTypeBomb:                   "Bomb",
	ItemTypePveNpc:                 "PveNpc",
	ItemTypeInfinityWeatherMachine: "InfinityWeatherMachine",
	ItemTypeSlime:                  "Slime",
	ItemTypeAcid:                   "Acid",
	ItemTypeCompletionist:          "Completionist",
	ItemTypePunchToggle:            "PunchToggle",
	ItemTypeAnzuBlock:              "AnzuBlock",
	ItemTypeFeedingBlock:           "FeedingBlock",
	ItemTypeKrankensBlock:          "KrankensBlock",
	ItemTypeFriendsEntrance:        "FriendsEntrance",
	ItemTypePearls:                 "Pearls",
}

var itemMaterialMap = map[ItemMaterial]string{
	ItemMaterialWood:  "Wood",
	ItemMaterialGlass: "Glass",
	ItemMaterialRock:  "Rock",
	ItemMaterialMetal: "Metal",
}

var tileVisualEffectMap = map[TileVisualEffect]string{
	TileVisualEffectNone:             "None",
	TileVisualEffectFlameLick:        "FlameLick",
	TileVisualEffectSmoking:          "Smoking",
	TileVisualEffectGlowTint1:        "GlowTint1",
	TileVisualEffectAnim:             "Anim",
	TileVisualEffectBubbles:          "Bubbles",
	TileVisualEffectPet:              "Pet",
	TileVisualEffectPetAnim:          "PetAnim",
	TileVisualEffectNoArms:           "NoArms",
	TileVisualEffectWavey:            "Wavey",
	TileVisualEffectWaveyAnim:        "WaveyAnim",
	TileVisualEffectBothArms:         "BothArms",
	TileVisualEffectLowHair:          "LowHair",
	TileVisualEffectUnderFace:        "UnderFace",
	TileVisualEffectSkinTint:         "SkinTint",
	TileVisualEffectMask:             "Mask",
	TileVisualEffectAnimMask:         "AnimMask",
	TileVisualEffectLowHairMask:      "LowHairMask",
	TileVisualEffectGhost:            "Ghost",
	TileVisualEffectPulse:            "Pulse",
	TileVisualEffectColorize:         "Colorize",
	TileVisualEffectColorizeToShirt:  "ColorizeToShirt",
	TileVisualEffectColorizeAnim:     "ColorizeAnim",
	TileVisualEffectHighFace:         "HighFace",
	TileVisualEffectHighFaceAnim:     "HighFaceAnim",
	TileVisualEffectRainbowShift:     "RainbowShift",
	TileVisualEffectBackFore:         "BackFore",
	TileVisualEffectColorizeWithSkin: "ColorizeWithSkin",
	TileVisualEffectNoRender:         "NoRender",
	TileVisualEffectSpin:             "Spin",
	TileVisualEffectOffhand:          "Offhand",
	TileVisualEffectWinged:           "Winged",
	TileVisualEffectSink:             "Sink",
	TileVisualEffectDarkness:         "Darkness",
	TileVisualEffectLightSource:      "LightSource",
	TileVisualEffectLightIfOn:        "LightIfOn",
	TileVisualEffectDiscolor:         "Discolor",
	TileVisualEffectStepSpin:         "StepSpin",
	TileVisualEffectPetColored:       "PetColored",
	TileVisualEffectSilkFoot:         "SilkFoot",
	TileVisualEffectTilty:            "Tilty",
	TileVisualEffectTiltyDark:        "TiltyDark",
	TileVisualEffectNextFrameIfOn:    "NextFrameIfOn",
	TileVisualEffectWobble:           "Wobble",
	TileVisualEffectScroll:           "Scroll",
	TileVisualEffectLightSourcePulse: "LightSourcePulse",
	TileVisualEffectBubbleMachine:    "BubbleMachine",
	TileVisualEffectVeryLowHair:      "VeryLowHair",
	TileVisualEffectVeryLowHairMask:  "VeryLowHairMask",
}
var tileStorageMap = map[TileStorage]string{
	TileStorageSingleFrameAlone:   "SingleFrameAlone",
	TileStorageSingleFrame:        "SingleFrame",
	TileStorageSmartEdge:          "SmartEdge",
	TileStorageSmartEdgeHoriz:     "SmartEdgeHoriz",
	TileStorageSmartCling:         "SmartCling",
	TileStorageSmartOuter:         "SmartOuter",
	TileStorageRandom:             "Random",
	TileStorageSmartEdgeVert:      "SmartEdgeVert",
	TileStorageSmartEdgeHorizCave: "SmartEdgeHorizCave",
	TileStorageSmartCling2:        "SmartCling2",
	TileStorageSmartEdgeDiagon:    "SmartEdgeDiagon",
}

var tileCollisionMap = map[TileCollision]string{
	TileCollisionNone:            "None",
	TileCollisionSolid:           "Solid",
	TileCollisionJumpThrough:     "JumpThrough",
	TileCollisionGateway:         "Gateway",
	TileCollisionIfOff:           "IfOff",
	TileCollisionOneWay:          "OneWay",
	TileCollisionVIP:             "VIP",
	TileCollisionJumpDown:        "JumpDown",
	TileCollisionAdventure:       "Adventure",
	TileCollisionIfOn:            "IfOn",
	TileCollisionFaction:         "Faction",
	TileCollisionGuild:           "Guild",
	TileCollisionCloud:           "Cloud",
	TileCollisionFriendsEntrance: "FriendsEntrance",
}

var bodyPartMap = map[BodyPart]string{
	BodyPartHat:          "Hat",
	BodyPartShirt:        "Shirt",
	BodyPartPants:        "Pants",
	BodyPartShoes:        "Shoes",
	BodyPartFaceItem:     "FaceItem",
	BodyPartHand:         "Hand",
	BodyPartBack:         "Back",
	BodyPartHair:         "Hair",
	BodyPartChestItem:    "ChestItem",
	BodyPartNumBodyParts: "NumBodyParts",
}

var avatarPartsMap = map[AvatarParts]string{
	AvatarPartsHead:     "Head",
	AvatarPartsFace:     "Shirt",
	AvatarPartsBody:     "Body",
	AvatarPartsFrontArm: "FrontArm",
	AvatarPartsBackArm:  "BackArm",
	AvatarPartsLegs:     "Legs",
}

var itemFlagsMap = map[ItemFlags]string{
	ItemFlagsFlippable:   "Flippable",
	ItemFlagsEditable:    "Editable",
	ItemFlagsSeedless:    "Seedless",
	ItemFlagsPermanent:   "Permanent",
	ItemFlagsDropless:    "Dropless",
	ItemFlagsNoSelf:      "NoSelf",
	ItemFlagsNoShadow:    "NoShadow",
	ItemFlagsWorldlocked: "Worldlocked",
	ItemFlagsBeta:        "Beta",
	ItemFlagsAutopickup:  "Autopickup",
	ItemFlagsMod:         "Mod",
	ItemFlagsRandomGrow:  "RandomGrow",
	ItemFlagsPublic:      "Public",
	ItemFlagsForeground:  "Foreground",
	ItemFlagsHoliday:     "Holiday",
	ItemFlagsUntradeable: "Untradeable",
}

var itemFlags2Map = map[ItemFlags2]string{
	ItemFlags2RobotDeadly:            "RobotDeadly",
	ItemFlags2RobotShootLeft:         "RobotShootLeft",
	ItemFlags2RobotShootRight:        "RobotShootRight",
	ItemFlags2RobotShootDown:         "RobotShootDown",
	ItemFlags2RobotShootUp:           "RobotShootUp",
	ItemFlags2RobotCanShoot:          "RobotCanShoot",
	ItemFlags2RobotLava:              "RobotLava",
	ItemFlags2RobotPointy:            "RobotPointy",
	ItemFlags2RobotShootDeadly:       "RobotShootDeadly",
	ItemFlags2GuildItem:              "GuildItem",
	ItemFlags2GuildFlag:              "GuildFlag",
	ItemFlags2StarshipHelm:           "StarshipHelm",
	ItemFlags2StarshipReactor:        "StarshipReactor",
	ItemFlags2StarshipViewscreen:     "StarshipViewscreen",
	ItemFlags2SMod:                   "SMod",
	ItemFlags2TileDeadlyIfOn:         "TileDeadlyIfOn",
	ItemFlags2LongHandItem64x32:      "LongHandItem64x32",
	ItemFlags2Gemless:                "Gemless",
	ItemFlags2ClothesTransmutable:    "ClothesTransmutable",
	ItemFlags2DungeonItem:            "DungeonItem",
	ItemFlags2PVEMelee:               "PVEMelee",
	ItemFlags2PVERanged:              "PVERanged",
	ItemFlags2PVEAutoAim:             "PVEAutoAim",
	ItemFlags2OneInWorld:             "OneInWorld",
	ItemFlags2OnlyForWorldOwner:      "OnlyForWorldOwner",
	ItemFlags2NoUpgrade:              "NoUpgrade",
	ItemFlags2ExtinguishFire:         "ExtinguishFire",
	ItemFlags2ExtinguishFireNoDamage: "ExtinguishFireNoDamage",
	ItemFlags2NeedReceptionDesk:      "NeedReceptionDesk",
	ItemFlags2UsePaint:               "UsePaint",
}

var itemFXFlagsMap = map[ItemFXFlags]string{
	ItemFXFlagsMultiAnim:          "MultiAnim",
	ItemFXFlagsPingPongAnim:       "PingPongAnim",
	ItemFXFlagsOverlayObject:      "OverlayObject",
	ItemFXFlagsOffsetUp:           "OffsetUp",
	ItemFXFlagsDualLayer:          "DualLayer",
	ItemFXFlagsMultiAnim2:         "MultiAnim2",
	ItemFXFlagsUnk0x40:            "Unk0x40",
	ItemFXFlagsUseSkinTint:        "UseSkinTint",
	ItemFXFlagsSeedTintLayer1:     "SeedTintLayer1",
	ItemFXFlagsSeedTintLayer2:     "SeedTintLayer2",
	ItemFXFlagsRainbowTintLayer1:  "RainbowTintLayer1",
	ItemFXFlagsRainbowTintLayer2:  "RainbowTintLayer2",
	ItemFXFlagsGlow:               "Glow",
	ItemFXFlagsNoArms:             "NoArms",
	ItemFXFlagsFrontArmPunch:      "FrontArmPunch",
	ItemFXFlagsRenderOffhand:      "RenderOffhand",
	ItemFXFlagsSlowfallObject:     "SlowfallObject",
	ItemFXFlagsReplacementSprite:  "ReplacementSprite",
	ItemFXFlagsOrbFloat:           "OrbFloat",
	ItemFXFlagsCapeForeground:     "CapeForeground",
	ItemFXFlagsRenderFXVariantVer: "RenderFXVariantVer",
}

func enumToString[T utils.Integer](value T, mp map[T]string) string {
	enum, ok := mp[value]
	if !ok {
		return fmt.Sprintf("Unknown<%d>", value)
	}
	return enum
}

func flagsToString[T utils.Integer](value T, mp map[T]string) string {
	if value == 0 {
		return "None"
	}

	str := ""
	for k, v := range mp {
		if value&k != 0 {
			if len(str) > 0 {
				str += ", "
			}
			str += v
		}
	}
	return str
}

func (ItemType) Map() map[ItemType]string {
	return itemTypeMap
}

func (v ItemType) String() string {
	return enumToString(v, v.Map())
}

func (ItemType) IsEnum() bool {
	return true
}

func (ItemMaterial) Map() map[ItemMaterial]string {
	return itemMaterialMap
}

func (v ItemMaterial) String() string {
	return enumToString(v, v.Map())
}

func (ItemMaterial) IsEnum() bool {
	return true
}

func (TileVisualEffect) Map() map[TileVisualEffect]string {
	return tileVisualEffectMap
}

func (v TileVisualEffect) String() string {
	return enumToString(v, v.Map())
}

func (TileVisualEffect) IsEnum() bool {
	return true
}

func (TileStorage) Map() map[TileStorage]string {
	return tileStorageMap
}

func (v TileStorage) String() string {
	return enumToString(v, v.Map())
}

func (TileStorage) IsEnum() bool {
	return true
}

func (TileCollision) Map() map[TileCollision]string {
	return tileCollisionMap
}

func (v TileCollision) String() string {
	return enumToString(v, v.Map())
}

func (TileCollision) IsEnum() bool {
	return true
}

func (BodyPart) Map() map[BodyPart]string {
	return bodyPartMap
}

func (v BodyPart) String() string {
	return enumToString(v, v.Map())
}

func (BodyPart) IsEnum() bool {
	return true
}

func (AvatarParts) Map() map[AvatarParts]string {
	return avatarPartsMap
}

func (v AvatarParts) String() string {
	return flagsToString(v, v.Map())
}

func (AvatarParts) IsFlags() bool {
	return true
}

func (ItemFlags) Map() map[ItemFlags]string {
	return itemFlagsMap
}

func (v ItemFlags) String() string {
	return flagsToString(v, v.Map())
}

func (ItemFlags) IsFlags() bool {
	return true
}

func (ItemFlags2) Map() map[ItemFlags2]string {
	return itemFlags2Map
}

func (v ItemFlags2) String() string {
	return flagsToString(v, v.Map())
}

func (ItemFlags2) IsFlags() bool {
	return true
}

func (ItemFXFlags) Map() map[ItemFXFlags]string {
	return itemFXFlagsMap
}

func (v ItemFXFlags) String() string {
	return flagsToString(v, v.Map())
}

func (ItemFXFlags) IsFlags() bool {
	return true
}
