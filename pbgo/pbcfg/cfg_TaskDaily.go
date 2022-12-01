package config

type TaskDailyCfg struct {
	Id          int32
	TaskOpening int32
	TaskList    int32
	TaskLoad    int32
	Reward      []*Reward
}

type TaskDailyTable struct {
	data                      map[int32]*TaskDailyCfg
	taskDailyTaskListIndexMap map[int32][]int32
}

var TaskDailyData = &TaskDailyTable{
	data: map[int32]*TaskDailyCfg{},
	taskDailyTaskListIndexMap: map[int32][]int32{
		1: {2},
		2: {3},
		3: {4},
		4: {5},
		5: {6},
		6: {7},
		7: {8},
		8: {1},
	},
}

func (table *TaskDailyTable) Get(id int32) *TaskDailyCfg {
	return table.data[id]
}

func (table *TaskDailyTable) GetAll() []int32 {
	return taskDailyKeys
}

var taskDailyValues = []*TaskDailyCfg{
	{
		Id:          1,
		TaskOpening: 2007,
		TaskList:    8,
		TaskLoad:    1,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    3,
				Count: 150,
			},
		},
	},
	{
		Id:          2,
		TaskOpening: 2007,
		TaskList:    1,
		TaskLoad:    1,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    51,
				Count: 1,
			},
		},
	},
	{
		Id:          3,
		TaskOpening: 2007,
		TaskList:    2,
		TaskLoad:    1,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 1,
			},
		},
	},
	{
		Id:          4,
		TaskOpening: 2007,
		TaskList:    3,
		TaskLoad:    1,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    3,
				Count: 30,
			},
		},
	},
	{
		Id:          5,
		TaskOpening: 2007,
		TaskList:    4,
		TaskLoad:    1,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    3,
				Count: 50,
			},
		},
	},
	{
		Id:          6,
		TaskOpening: 2007,
		TaskList:    5,
		TaskLoad:    1,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    41,
				Count: 2,
			},
		},
	},
	{
		Id:          7,
		TaskOpening: 2007,
		TaskList:    6,
		TaskLoad:    1,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 2,
			},
		},
	},
	{
		Id:          8,
		TaskOpening: 2007,
		TaskList:    7,
		TaskLoad:    1,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2002,
				Count: 1,
			},
		},
	},
}

var taskDailyKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8,
}

func init() {
	TaskDailyData.data = make(map[int32]*TaskDailyCfg)
	for i := 0; i < len(taskDailyKeys); i++ {
		TaskDailyData.data[taskDailyKeys[i]] = taskDailyValues[i]
	}
}

func (table *TaskDailyTable) GetByTaskList(TaskList int32) (res []*TaskDailyCfg) {
	for _, i := range table.taskDailyTaskListIndexMap[TaskList] {
		res = append(res, table.Get(i))
	}
	return
}
