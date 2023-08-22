package config

// Packages wrapper, all the package from '/internal/pkg' should be
// registered here, there is two main reason why the package
// had to be registered, first is it would make the dev
// easier to maintain and lookup for specific package
// and these packages could be called by other
// package's service or repository
//type Packages struct {
//	StringConv *stringconv.Package
//}

// NewPackages responsible to register all the package from '/internal/pkg'.
//func NewPackages(engine *Engine, db *Database, logger *Logger) {
//	pkg := new(Packages)
//	pkg.StringConv = stringconv.New("string-conv", engine, db, logger)
//}
