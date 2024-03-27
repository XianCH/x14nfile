package x14nfile

type FileSystem struct {
	RootDir string
	tree    Tree
}

func NewFileSystem(root string, tree Tree) *FileSystem {
	return &FileSystem{
		RootDir: root,
		tree:    tree,
	}
}

//search file by file name

//get file by name

//delete file by name

//display all file path
