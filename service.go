package locale

var (
	//disable to remove error messages.
	DevMode   = true
	//map of string localizations
	LocaleMap map[string]map[string]string
)

//Localize a string
func Str(localename, lookup string) (localized string) {
	var hasTrans bool
	if !DevMode {
		localized = lookup
	} else {
		localized = "Localization not found!"
	}

	if _, ok := LocaleMap[localename]; ok {
		lmap := LocaleMap[localename]
		if localized, hasTrans = lmap[lookup]; !hasTrans {
			localized = "Could not find localized version of text."
		}
	}

	return
}

//Add localization of string
func Add(localename, base, localized string) {
	var localemap map[string]string
	var ok bool
	if LocaleMap == nil {
		LocaleMap = make(map[string]map[string]string)
	}

	if localemap, ok = LocaleMap[localename]; !ok {
		localemap = make(map[string]string)
	}

	localemap[base] = localized
	delete(LocaleMap, localename)
	LocaleMap[localename] = localemap

}


//Load locale with settings set in gorrilla session.
func StrSession(locs interface{}, str string) string {
	var localname = locs.(string)
	return Str(localname, str)
}
