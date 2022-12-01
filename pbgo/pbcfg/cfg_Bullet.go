package config

type BulletCfg struct {
	Id             int32
	BulletRes      string
	TrackType      int32
	HitSolt        string
	ShootSolt      string
	ForwardTime    float64
	FlyTime        float64
	ParabolaHeight float64
	EaseType       int32
}

type BulletTable struct {
	data map[int32]*BulletCfg
}

var BulletData = &BulletTable{
	data: map[int32]*BulletCfg{},
}

func (table *BulletTable) Get(id int32) *BulletCfg {
	return table.data[id]
}

func (table *BulletTable) GetAll() []int32 {
	return bulletKeys
}

var bulletValues = []*BulletCfg{
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
}

var bulletKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11,
}

func init() {
	BulletData.data = make(map[int32]*BulletCfg)
	for i := 0; i < len(bulletKeys); i++ {
		BulletData.data[bulletKeys[i]] = bulletValues[i]
	}
}
