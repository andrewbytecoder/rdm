package services

import (
	"sync"

	"github.com/andrewbytecoder/wrdm/backend/storage"
	"github.com/andrewbytecoder/wrdm/backend/types"
)

type PreferencesService struct {
	pref *storage.PreferencesStorage
}

var preferences *PreferencesService
var oncePreferences sync.Once

func Preferences() *PreferencesService {
	if preferences == nil {
		oncePreferences.Do(func() {
			preferences = &PreferencesService{
				pref: storage.NewPreferencesStorage(),
			}
		})
	}
	return preferences
}

func (p *PreferencesService) GetPreferences() (resp types.JSResp) {
	resp.Data = p.pref.GetPreferences()
	resp.Success = true
	return
}

func (p *PreferencesService) SetPreferences(values map[string]any) (resp types.JSResp) {
	err := p.pref.SetPreferencesN(values)
	if err != nil {
		resp.Msg = err.Error()
		return
	}

	resp.Success = true
	return
}

func (p *PreferencesService) RestorePreferences() (resp types.JSResp) {
	defaultPref := p.pref.RestoreDefault()
	resp.Data = map[string]any{
		"pref": defaultPref,
	}
	resp.Success = true
	return
}
