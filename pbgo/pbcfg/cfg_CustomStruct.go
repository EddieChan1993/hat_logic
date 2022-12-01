package config

type Reward struct{
	Type int32
	Id int32
	Count int64
}

type Consume struct{
	Type int32
	Id int32
	Count int32
}

type BasicProp struct{
	Id int32
	Count int32
}

type AddProp struct{
	Id int32
	Count int32
}

type Attribute struct{
	Id int32
	Count int32
}

type AddAttribute struct{
	Type int32
	Id int32
	Count int32
}

type Status struct{
	Id int32
	Count int32
}

type Suit struct{
	Id int32
	Count int32
}

type Unlock struct{
	Location int32
	SkillLevel int32
}

type Fx struct{
	Bind string
	FxName string
}

type Item struct{
	Id int32
	Num int32
}

type RandomReward struct{
	Type int32
	Id int32
	Num int32
	Ratio int32
}

type DmgChange struct{
	Type int32
	Count int32
}

type SpecialCondition struct{
	Skill int32
	SkillLevel int32
}

type RandomRewardPer struct{
	Type int32
	Id int32
	Num int32
	Ratio int32
}

type AddPropRan struct{
	Id int32
	Ratio int32
}

