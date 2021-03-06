package main

import "github.com/libgit2/git2go"

var gitErrorClassNames = map[git.ErrorClass]string{
	git.ErrClassNone:       "ErrClassNone",
	git.ErrClassNoMemory:   "ErrClassNoMemory",
	git.ErrClassOs:         "ErrClassOs",
	git.ErrClassInvalid:    "ErrClassInvalid",
	git.ErrClassReference:  "ErrClassReference",
	git.ErrClassZlib:       "ErrClassZlib",
	git.ErrClassRepository: "ErrClassRepository",
	git.ErrClassConfig:     "ErrClassConfig",
	git.ErrClassRegex:      "ErrClassRegex",
	git.ErrClassOdb:        "ErrClassOdb",
	git.ErrClassIndex:      "ErrClassIndex",
	git.ErrClassObject:     "ErrClassObject",
	git.ErrClassNet:        "ErrClassNet",
	git.ErrClassTag:        "ErrClassTag",
	git.ErrClassTree:       "ErrClassTree",
	git.ErrClassIndexer:    "ErrClassIndexer",
	git.ErrClassSSL:        "ErrClassSSL",
	git.ErrClassSubmodule:  "ErrClassSubmodule",
	git.ErrClassThread:     "ErrClassThread",
	git.ErrClassStash:      "ErrClassStash",
	git.ErrClassCheckout:   "ErrClassCheckout",
	git.ErrClassFetchHead:  "ErrClassFetchHead",
	git.ErrClassMerge:      "ErrClassMerge",
	git.ErrClassSsh:        "ErrClassSsh",
	git.ErrClassFilter:     "ErrClassFilter",
	git.ErrClassRevert:     "ErrClassRevert",
	git.ErrClassCallback:   "ErrClassCallback",
}

var gitErrorCodeNames = map[git.ErrorCode]string{
	git.ErrOk:        "No error",
	git.ErrGeneric:   "Generic error",
	git.ErrNotFound:  "Requested object could not be found",
	git.ErrExists:    "Object exists, preventing operation",
	git.ErrAmbigious: "More that one object matches",
	git.ErrBuffs:     "Output buffer too short to hold data",
	// this is pretty weird, this is what git2go has to say for itself:
	// 	GIT_EUSER is a special error that is never generated by libgit2
	// 	code.  You can return it from a callback (e.g to stop an iteration)
	// 	to know that it was generated by the callback and not by libgit2.
	git.ErrUser:           "Special non libgit2 callback user error",
	git.ErrBareRepo:       "Operation not allowed on bare repository",
	git.ErrUnbornBranch:   "HEAD refers to branch with no commits",
	git.ErrUnmerged:       "Merge in progress prevented operation",
	git.ErrNonFastForward: "Reference was not fast-forwardable",
	git.ErrInvalidSpec:    "Name/ref spec was not in a valid format",
	git.ErrConflict:       "Checkout conflicts prevented operation",
	git.ErrLocked:         "Lock file prevented operation",
	git.ErrModified:       "Reference value does not match expected",
	git.ErrPassthrough:    "Unknown internal git2go error",
	git.ErrIterOver:       "Signal of end of iteration with iterator",
	git.ErrAuth:           "Authentication failed",
}
