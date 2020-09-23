package go_dnd5eapi

// func (ac *apiClient) GetProficienciesForClassByIndex(classIndex string) ([]ProficiencyDetail, error) {
// 	class, err := ac.GetClassByIndex(classIndex)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var result []ProficiencyDetail
// 	for _, v := range class.Proficiencies {
// 		prof, err := ac.GetProficiencyByIndex(v.Index)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if prof != nil {
// 			result = append(result, *prof)
// 		} else {
// 			return nil, errors.New("No proficiency called " + v.Index + " could be found")
// 		}
// 	}
// 	return result, nil
// }

// func (ac *apiClient) GetProficiencyChoicesForClassByIndex(classIndex string) (map[string][]ProficiencyDetail, error) {
// 	class, err := ac.GetClassByIndex(classIndex)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var result = make(map[string][]ProficiencyDetail)
// 	for i, choice := range class.ProficiencyChoices {
// 		var profs []ProficiencyDetail
// 		for _, prof := range class.ProficiencyChoices[i].From {
// 			p, err := ac.GetProficiencyByIndex(prof.Index)
// 			if err != nil {
// 				return nil, err
// 			}
// 			if p != nil {
// 				profs = append(profs, *p)
// 			} else {
// 				return nil, errors.New("No proficiency called " + prof.Index + " could be found")
// 			}
// 		}
// 		result[choice.Type] = profs
// 	}
// 	return result, nil
// }
