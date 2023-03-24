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
