package service

import "hello-run/dao"

func GetHouseholdVo(household *dao.Household) *HouseholdVo {
	if household == nil {
		return nil
	}
	return &HouseholdVo{
		Id:         household.Id,
		Age:        household.Age,
		Height:     household.Height,
		Wheelchair: household.Wheelchair,
	}
}

func GetHouseholdVoList(householdList []dao.Household) []HouseholdVo {
	if householdList == nil {
		return nil
	}
	var householdVoList []HouseholdVo
	for _, household := range householdList {
		householdVoList = append(householdVoList, *GetHouseholdVo(&household))
	}
	return householdVoList
}

func DeleteHouseholdByHouseholdId(householdId int64) error {
	//first check if the household exists
	_, err := dao.GetHouseholdInfoById(householdId)
	if err != nil {
		return err
	}
	return dao.DeleteHouseholdById(householdId)
}

func CreateHousehold(userId string, roomId int64, age int, height int, wheelchair bool) (*HouseholdVo, error) {
	household, err := dao.CreateHousehold(userId, roomId, age, height, wheelchair)
	if err != nil {
		return nil, err
	}
	return GetHouseholdVo(household), nil
}

func UpdateHousehold(householdId int64, age int, height int, wheelchair bool) (*HouseholdVo, error) {
	//first test if the room exists
	_, err := dao.GetHouseholdInfoById(householdId)
	if err != nil {
		return nil, err
	}
	householdId, err = dao.UpdateHousehold(householdId, age, height, wheelchair)
	if err != nil {
		return nil, err
	}
	household, _ := dao.GetHouseholdInfoById(householdId)
	return GetHouseholdVo(household), nil
}
