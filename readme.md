# go-locale

![](https://talks.golang.org/2014/go4gophers/gopherflag.png)

Go Server (GopherSauce) localization extension.

## Install
Add the following tag within your `.gxml` file :

	<import src="github.com/cheikhshift/locale/gos.xml" />

### Download package
Use the following command to download this package :

	go get github.com/cheikhshift/locale
	
### How it works?
The setup of this package is in Go. Feel free to use a go file to specify the translation of a particular phrase. You may add localizations of your strings within your `.gxml`'s main tag as well.


### Add localization
Use the following function to add a new localized phrase.

	func Add(local, base, localized string)

For example to add a new localization of string `FOO` to local `en-US` the call would be :

	locale.Add("en-US", "FOO", "Huh") 

### Get Localization
Use the following function to get a localized phrase.

	Str(localename, lookup string) (localized string) 

For example to get the localization of string `FOO` with local `en-US` the call would be :
	
	locale.Str("en-US", "Foo")

### Set local service
Adjust your user's local with the following endpoint :

	$(HOST_WITH_PROTOCOL):$(port)/set_local/$(desired_locale)

For example to set your website to `wf-SN` (Wolof senegal) to request would be (To your application...) :

	GET /set_local/wf-SN

### Localize templates
Here is a list of methods used to localize and update localization settings :

#### Update local 
	{{ SetLoc $arg1 $arg2 }}

##### Argument information
- `$arg1` : Valid `sessions.Session` object. With templates in your `web/` folder use the  field `.Session`. 
- `$arg2` : Local to set for user.

#### Localize string with session settings
	{{ Loc $arg1 $arg2 }}

##### Argument information
- `$arg1` : Valid `sessions.Session` object. With templates in your `web/` folder use the  field `.Session`. 
- `$arg2` :  String to localize with current settings

##### Example usage :
The following template snippet will localize string `"Hello"` to the current user setting
		
			{{  $session := .Session }}
			{{ with "Hello" }}
				{{ Loc $session . }}
			{{ end }}
			{{ with "World" }}
				{{ Loc $session . }}
			{{ end }}

#### Localize with Custom string

### Package variables

	var (
		//disable to remove error messages.
		DevMode   = true
		//map of string localizations
		LocaleMap map[string]map[string]string
	)