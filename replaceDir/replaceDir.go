package replaceDir

type DirPath string

// 교체할 디렉토리
type ReplaceDir struct {
	TargetDir    DirPath
	NoneDirSlice []string
	NoneDirName  string
}

func NewReplaceDir(targetDir DirPath) *ReplaceDir {
	replaceDir := &ReplaceDir{
		TargetDir: targetDir,
	}

	return replaceDir
}
