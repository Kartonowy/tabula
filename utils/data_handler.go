package utils

type Theme struct {
	Title    string
	Cover    string
	Platform string
	Link     string
	Author   string
	Path     string
}

func newTheme(title string, path string) Theme {
	return Theme{
		Title:    title,
		Path:     path,
		Cover:    "",
		Platform: "",
		Link:     "",
	}
}

func GetThemes() []Theme {
	themes := [6]Theme{
		newTheme("Fight", "/theme/fight"),
		newTheme("Mystery", "/theme/mystery"),
		newTheme("Exploration", "/theme/exploration"),
		newTheme("Swamp", "/theme/swamp"),
		newTheme("Castle", "/theme/castle"),
		newTheme("Candlekeep", "/theme/candlekeep"),
	}

	return themes[:]
}