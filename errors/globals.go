package errors

// These are globals that can be used to standardize error handling in all my code.

var (
	eMrpkgengine map[string]string = map[string]string{
		// DownloadPackage error messages
		"dwnerr1": "could not create the file",
		"dwnerr2": "could not get data from server",
		"dwnerr3": "server did not find file",
		"dwnerr4": "server did not allow permission to access the resource",
		"dwnerr5": "user is not authorized to access the resource",
		"dwnerr6": "server encountered an internal error",
		"dwnerr7": "server is currently unavailable",
		"dwnerr8": "could not write to the file",
		// BuildPackage error messages
		"blderr1": "no configuartion file found",
		"blderr2": "build failed due to an intenal error",
		// InstallPackage error messages
		"insterr1": "package could not be downloaded due to an internal error",
		"insterr2": "package could not be unzipped",
		"insterr3": "package could not be built due to an internal error",
	}
	eMItypes map[string]string = map[string]string{
		// cast error messages
		"csterr1": "could not cast to string",
		"csterr2": "could not cast to int",
		"csterr3": "could not cast to float32",
		"csterr4": "could not cast to float64",
		"csterr5": "unknown type",
	}
	eMIerrors map[string]string = map[string]string{
		// ErrChecker error messages
		"tsterr1": "test error 1",
	}
	EM map[string]map[string]string = map[string]map[string]string{
		"eMre":  eMrpkgengine,
		"eMIty": eMItypes,
		"eMIer": eMIerrors,
	}
)
