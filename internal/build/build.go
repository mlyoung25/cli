package build

import "runtime/debug"

var Version string

// Gets the CLI release version, or the latest git commit shorthash if the release could not be found (for local builds)
// Based on https://github.com/cue-lang/cue/issues/1697#issuecomment-1122097477
func init() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		Version = "unknown"
	}

	mod := &info.Main
	if mod.Replace != nil {
		mod = mod.Replace
	}
	if mod.Version == "(devel)" {
		var vcsRevision string
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				vcsRevision = setting.Value
				if len(vcsRevision) > 12 {
					vcsRevision = vcsRevision[:12]
				}
			}
		}

		Version = vcsRevision
	}

	Version = mod.Version
}
