package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// BTreeItem 表示 B+ 树中的一个节点
type BTreeItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// BTree 表示 B+ 树结构
type BTree struct {
	Root *Node `json:"root"`
}

// Node 表示 B+ 树中的一个节点
type Node struct {
	Items    []*BTreeItem `json:"items"`
	Children []*Node      `json:"children"`
}

// SearchFile 根据文件名搜索文件路径和文件内容
func (b *BTree) SearchFile(filename string) (string, string) {
	return b.searchFileHelper(b.Root, filename)
}

// searchFileHelper 是 SearchFile 的辅助函数
func (b *BTree) searchFileHelper(node *Node, filename string) (string, string) {
	if node == nil {
		return "", ""
	}
	for i, item := range node.Items {
		if item.Key == filename {
			return item.Key, item.Value
		} else if i == len(node.Items)-1 || item.Key > filename {
			// 递归搜索子节点
			return b.searchFileHelper(node.Children[i], filename)
		}
	}
	return b.searchFileHelper(node.Children[len(node.Children)-1], filename)
}

// SaveToFile 将 B+ 树保存到文件中
func (b *BTree) SaveToFile(filepath string) error {
	data, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath, data, 0644)
}

// LoadFromFile 从文件中加载 B+ 树
func LoadFromFile(filepath string) (*BTree, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var tree BTree
	err = json.Unmarshal(data, &tree)
	if err != nil {
		return nil, err
	}
	return &tree, nil
}

func main() {
	// 创建一个简单的 B+ 树结构
	root := &Node{
		Items: []*BTreeItem{
			{Key: "file1.txt", Value: "/path/to/file1.txt"},
			{Key: "file2.txt", Value: "/path/to/file2.txt"},
		},
		Children: []*Node{
			{Items: []*BTreeItem{{Key: "file3.txt", Value: "/path/to/file3.txt"}}},
		},
	}
	btree := &BTree{Root: root}

	// 保存 B+ 树到文件
	if err := btree.SaveToFile("btree.json"); err != nil {
		fmt.Println("Error saving B+ tree:", err)
		return
	}

	// 从文件加载 B+ 树
	loadedBTree, err := LoadFromFile("btree.json")
	if err != nil {
		fmt.Println("Error loading B+ tree:", err)
		return
	}

	// 搜索文件
	filename := "file2.txt"
	filePath, fileContent := loadedBTree.SearchFile(filename)
	if filePath != "" {
		fmt.Printf("File %s found at path: %s\n", filename, filePath)
		// 模拟读取文件内容
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println("File content:", string(content))
	} else {
		fmt.Printf("File %s not found\n", filename)
	}
}
