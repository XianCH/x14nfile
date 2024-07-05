package x14nfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var FileDir *FileNode

type FileNode struct {
	FileName string      `json:"fileName"` //File name
	FileSize int64       `json:"fileSize"`
	IsDir    bool        `json:"isDIr"`
	Children []*FileNode `json:"children"`
	Path     string      `json:"path"`
}

func NewFileNode(path string, info os.FileInfo) *FileNode {
	return &FileNode{
		FileName: info.Name(),
		FileSize: info.Size(),
		IsDir:    info.IsDir(),
		Children: []*FileNode{},
		Path:     path,
	}
}

// BuildDirectoryTree builds a tree of FileNodes for the given directory
func BuildDirectoryTree(root string) (*FileNode, error) {
	info, err := os.Stat(root)
	if err != nil {
		return nil, err
	}

	rootNode := NewFileNode(root, info)

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == root {
			return nil
		}
		relativePath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		addNode(rootNode, relativePath, info)
		return nil
	})
	return rootNode, err
}

// PrintTree prints the tree of FileNodes to stdout
func PrintTree(node *FileNode, indent string) {
	fmt.Println(indent + node.FileName)
	if node.IsDir {
		for _, child := range node.Children {
			PrintTree(child, indent+"  ")
		}
	}
}

// addNode adds a new node to the tree of FileNodes
func addNode(root *FileNode, path string, info os.FileInfo) {
	parts := strings.Split(path, string(os.PathSeparator))
	current := root
	for _, part := range parts {
		found := false
		for i, child := range current.Children {
			if child.FileName == part {
				current = current.Children[i]
				found = true
				break
			}
		}
		if !found {
			newNode := NewFileNode(filepath.Join(current.Path, part), info)
			current.Children = append(current.Children, newNode)
			if info.IsDir() {
				current = newNode
			}
		}
	}
}

// GetFileNodeByPath returns the FileNode for the given path
func GetFileNodeByPath(root *FileNode, path string) *FileNode {
	parts := strings.Split(path, string(os.PathSeparator))
	current := root
	for _, part := range parts {
		found := false
		for i, child := range current.Children {
			if child.FileName == part {
				current = current.Children[i]
				found = true
				break
			}
		}
		if !found {
			return nil
		}
	}
	return current
}

// RemoveFileNodeByPath removes the FileNode for the given path
func RemoveFileNodeByPath(root *FileNode, path string) bool {
	parts := strings.Split(path, string(os.PathSeparator))
	current := root
	for i, part := range parts {
		if i == len(parts)-1 {
			// We're at the parent of the node to remove
			for j, child := range current.Children {
				if child.FileName == part {
					// Remove the node
					current.Children = append(current.Children[:j], current.Children[j+1:]...)
					return true
				}
			}
		} else {
			found := false
			for _, child := range current.Children {
				if child.FileName == part {
					current = child
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return false
}
