package item

type ItemType uint8
type ItemMaterial uint8

type TileVisualEffect uint8
type TileStorage uint8
type TileCollision uint8

type BodyPart uint8
type AvatarParts uint8

type ItemFlags uint16
type ItemFlags2 uint32
type ItemFXFlags uint32

const (
	ItemTypeFist ItemType = iota
	ItemTypeWrench
	ItemTypeUserDoor
	ItemTypeLock
	ItemTypeGems
	ItemTypeTreasure
	ItemTypeDeadly
	ItemTypeTrampoline
	ItemTypeConsumable
	ItemTypeGateway
	ItemTypeSign
	ItemTypeSfxWithExtraFrame
	ItemTypeBoombox
	ItemTypeDoor
	ItemTypePlatform
	ItemTypeBedrock
	ItemTypeLava
	ItemTypeNormal
	ItemTypeBackground
	ItemTypeSeed
	ItemTypeClothes
	ItemTypeNormalWithExtraFrame
	ItemTypeBackgdSfxExtraFrame
	ItemTypeBackBoombox
	ItemTypeBouncy
	ItemTypePointy
	ItemTypePortal
	ItemTypeCheckpoint
	ItemTypeMusicnote
	ItemTypeIce
	ItemTypeRaceFlag
	ItemTypeSwitcherroo
	ItemTypeChest
	ItemTypeMailbox
	ItemTypeBulletin
	ItemTypePinata
	ItemTypeDice
	ItemTypeComponent
	ItemTypeProvider
	ItemTypeLab
	ItemTypeAchievement
	ItemTypeWeatherMachine
	ItemTypeScoreboard
	ItemTypeSungate
	ItemTypeProfile
	ItemTypeDeadlyIfOn
	ItemTypeHeartMonitor
	ItemTypeDonationBox
	ItemTypeToybox
	ItemTypeMannequin
	ItemTypeCamera
	ItemTypeMagicegg
	ItemTypeTeam
	ItemTypeGameGen
	ItemTypeXenonite
	ItemTypeDressup
	ItemTypeCrystal
	ItemTypeBurglar
	ItemTypeCompactor
	ItemTypeSpotlight
	ItemTypeWind
	ItemTypeDisplayBlock
	ItemTypeVending
	ItemTypeFishtank
	ItemTypePetfish
	ItemTypeSolar
	ItemTypeForge
	ItemTypeGivingTree
	ItemTypeGivingTreeStump
	ItemTypeSteampunk
	ItemTypeSteamLavaIfOn
	ItemTypeSteamOrgan
	ItemTypeTamagotchi
	ItemTypeSewing
	ItemTypeFlag
	ItemTypeLobsterTrap
	ItemTypeArtcanvas
	ItemTypeBattleCage
	ItemTypePetTrainer
	ItemTypeSteamEngine
	ItemTypeLockBot
	ItemTypeWeatherSpecial
	ItemTypeSpiritStorage
	ItemTypeDisplayShelf
	ItemTypeVipDoor
	ItemTypeChalTimer
	ItemTypeChalFlag
	ItemTypeFishMount
	ItemTypePortrait
	ItemTypeWeatherSpecial2
	ItemTypeFossil
	ItemTypeFossilPrep
	ItemTypeDnaMachine
	ItemTypeBlaster
	ItemTypeValhowla
	ItemTypeChemsynth
	ItemTypeChemtank
	ItemTypeStorage
	ItemTypeOven
	ItemTypeSuperMusic
	ItemTypeGeigercharge
	ItemTypeAdventureReset
	ItemTypeTombRobber
	ItemTypeFaction
	ItemTypeRedFaction
	ItemTypeGreenFaction
	ItemTypeBlueFaction
	ItemTypeArtifact
	ItemTypeTrampolineMomentum
	ItemTypeFishgotchiTank
	ItemTypeFishingBlock
	ItemTypeItemSucker
	ItemTypeItemPlanter
	ItemTypeRobot
	ItemTypeCommand
	ItemTypeLuckyTicket
	ItemTypeStatsBlock
	ItemTypeFieldNode
	ItemTypeOuijaBoard
	ItemTypeArchitectMachine
	ItemTypeStarship
	ItemTypeAutodelete
	ItemTypeBoombox2
	ItemTypeAutoActionBreak
	ItemTypeAutoActionHarvest
	ItemTypeAutoActionHarvestSuck
	ItemTypeLightningCloud
	ItemTypePhasedBlock
	ItemTypeMud
	ItemTypeRootCutting
	ItemTypePasswordStorage
	ItemTypePhasedBlock2
	ItemTypeBomb
	ItemTypePveNpc
	ItemTypeInfinityWeatherMachine
	ItemTypeSlime
	ItemTypeAcid
	ItemTypeCompletionist
	ItemTypePunchToggle
	ItemTypeAnzuBlock
	ItemTypeFeedingBlock
	ItemTypeKrankensBlock
	ItemTypeFriendsEntrance
	ItemTypePearls
)

const (
	ItemMaterialWood ItemMaterial = iota
	ItemMaterialGlass
	ItemMaterialRock
	ItemMaterialMetal
)

const (
	TileVisualEffectNone TileVisualEffect = iota
	TileVisualEffectFlameLick
	TileVisualEffectSmoking
	TileVisualEffectGlowTint1
	TileVisualEffectAnim
	TileVisualEffectBubbles
	TileVisualEffectPet
	TileVisualEffectPetAnim
	TileVisualEffectNoArms
	TileVisualEffectWavey
	TileVisualEffectWaveyAnim
	TileVisualEffectBothArms
	TileVisualEffectLowHair
	TileVisualEffectUnderFace
	TileVisualEffectSkinTint
	TileVisualEffectMask
	TileVisualEffectAnimMask
	TileVisualEffectLowHairMask
	TileVisualEffectGhost
	TileVisualEffectPulse
	TileVisualEffectColorize
	TileVisualEffectColorizeToShirt
	TileVisualEffectColorizeAnim
	TileVisualEffectHighFace
	TileVisualEffectHighFaceAnim
	TileVisualEffectRainbowShift
	TileVisualEffectBackFore
	TileVisualEffectColorizeWithSkin
	TileVisualEffectNoRender
	TileVisualEffectSpin
	TileVisualEffectOffhand
	TileVisualEffectWinged
	TileVisualEffectSink
	TileVisualEffectDarkness
	TileVisualEffectLightSource
	TileVisualEffectLightIfOn
	TileVisualEffectDiscolor
	TileVisualEffectStepSpin
	TileVisualEffectPetColored
	TileVisualEffectSilkFoot
	TileVisualEffectTilty
	TileVisualEffectTiltyDark
	TileVisualEffectNextFrameIfOn
	TileVisualEffectWobble
	TileVisualEffectScroll
	TileVisualEffectLightSourcePulse
	TileVisualEffectBubbleMachine
	TileVisualEffectVeryLowHair
	TileVisualEffectVeryLowHairMask
)

const (
	TileStorageSingleFrameAlone TileStorage = iota
	TileStorageSingleFrame
	TileStorageSmartEdge
	TileStorageSmartEdgeHoriz
	TileStorageSmartCling
	TileStorageSmartOuter
	TileStorageRandom
	TileStorageSmartEdgeVert
	TileStorageSmartEdgeHorizCave
	TileStorageSmartCling2
	TileStorageSmartEdgeDiagon
)

const (
	TileCollisionNone TileCollision = iota
	TileCollisionSolid
	TileCollisionJumpThrough
	TileCollisionGateway
	TileCollisionIfOff
	TileCollisionOneWay
	TileCollisionVIP
	TileCollisionJumpDown
	TileCollisionAdventure
	TileCollisionIfOn
	TileCollisionFaction
	TileCollisionGuild
	TileCollisionCloud
	TileCollisionFriendsEntrance
)

const (
	BodyPartHat BodyPart = iota
	BodyPartShirt
	BodyPartPants
	BodyPartShoes
	BodyPartFaceItem
	BodyPartHand
	BodyPartBack
	BodyPartHair
	BodyPartChestItem
	BodyPartNumBodyParts
)

const (
	AvatarPartsHead AvatarParts = 1 << iota
	AvatarPartsFace
	AvatarPartsBody
	AvatarPartsFrontArm
	AvatarPartsBackArm
	AvatarPartsLegs
)

const (
	ItemFlagsFlippable ItemFlags = 1 << iota
	ItemFlagsEditable
	ItemFlagsSeedless
	ItemFlagsPermanent
	ItemFlagsDropless
	ItemFlagsNoSelf
	ItemFlagsNoShadow
	ItemFlagsWorldlocked
	ItemFlagsBeta
	ItemFlagsAutopickup
	ItemFlagsMod
	ItemFlagsRandomGrow
	ItemFlagsPublic
	ItemFlagsForeground
	ItemFlagsHoliday
	ItemFlagsUntradeable
)

const (
	ItemFlags2RobotDeadly ItemFlags2 = 1 << iota
	ItemFlags2RobotShootLeft
	ItemFlags2RobotShootRight
	ItemFlags2RobotShootDown
	ItemFlags2RobotShootUp
	ItemFlags2RobotCanShoot
	ItemFlags2RobotLava
	ItemFlags2RobotPointy
	ItemFlags2RobotShootDeadly
	ItemFlags2GuildItem
	ItemFlags2GuildFlag
	ItemFlags2StarshipHelm
	ItemFlags2StarshipReactor
	ItemFlags2StarshipViewscreen
	ItemFlags2SMod
	ItemFlags2TileDeadlyIfOn
	ItemFlags2LongHandItem64x32
	ItemFlags2Gemless
	ItemFlags2ClothesTransmutable
	ItemFlags2DungeonItem
	ItemFlags2PVEMelee
	ItemFlags2PVERanged
	ItemFlags2PVEAutoAim
	ItemFlags2OneInWorld
	ItemFlags2OnlyForWorldOwner
	ItemFlags2NoUpgrade
	ItemFlags2ExtinguishFire
	ItemFlags2ExtinguishFireNoDamage
	ItemFlags2NeedReceptionDesk
)

const (
	ItemFXFlagsMultiAnim ItemFXFlags = 1 << iota
	ItemFXFlagsPingPongAnim
	ItemFXFlagsOverlayObject
	ItemFXFlagsOffsetUp
	ItemFXFlagsDualLayer
	ItemFXFlagsMultiAnim2
	ItemFXFlagsUnk0x40
	ItemFXFlagsUseSkinTint
	ItemFXFlagsSeedTintLayer1
	ItemFXFlagsSeedTintLayer2
	ItemFXFlagsRainbowTintLayer1
	ItemFXFlagsRainbowTintLayer2
	ItemFXFlagsGlow
	ItemFXFlagsNoArms
	ItemFXFlagsFrontArmPunch
	ItemFXFlagsRenderOffhand
	ItemFXFlagsSlowfallObject
	ItemFXFlagsReplacementSprite
	ItemFXFlagsOrbFloat
	ItemFXFlagsCapeForeground
	ItemFXFlagsRenderFXVariantVer
)
