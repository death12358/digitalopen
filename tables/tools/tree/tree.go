package tree

import "fmt"

type TreeData struct {
	Node        *Node
	SymbolList  []string
	LineSetting [][]int
	PaySetting  [][]int
}

type Node struct {
	Value    string
	Site     int
	Children []*Node
}

type ResultTree struct {
	Count int
	Value []string
}

var (
	treeCount   = 100
	TreeChannel = make(chan *Node, treeCount)
	Tree        TreeData
	min         = 3
	row         = 5

	Wild    = "W"
	AllWild = "ALL Wild"

	lineSetting = [][]int{{1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {2, 2, 2, 2, 2}, {0, 1, 2, 1, 0}, {2, 1, 0, 1, 2}}
	paySetting  = [][]int{{0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {25, 20, 15, 10, 5, 1}, {30, 25, 20, 15, 10, 5}, {35, 30, 20, 15, 15, 10}}
	symbolList  = []string{"W", "A", "B", "C", "D"}
)

func Init() {
	newLine := make([][]int, len(lineSetting))
	for index, lineSet := range lineSetting {
		tmp := make([]int, len(lineSet))
		for i := 0; i < len(lineSet); i++ {
			tmp[i] = i*10 + lineSet[i]
		}
		newLine[index] = tmp
	}

	Tree = TreeData{
		LineSetting: newLine,
		PaySetting:  paySetting,
		SymbolList:  symbolList,
	}
	Tree.InitCreatTree()
}

// InitCreatTree 創建樹
func (t *TreeData) InitCreatTree() {

	for i := 0; i < treeCount; i++ {
		t.Node = &Node{Site: -1}
		for _, lineSet := range t.LineSetting {
			t.Node.CreateTree(lineSet)
		}
		TreeChannel <- t.Node
	}
}

// PayData 計算所有線得分
func (t *TreeData) PayData(treeInfo []ResultTree) {
	total := 0
	for _, tree := range treeInfo {
		tmp := t.payLine(tree)
		total += tmp
		fmt.Println(tmp)
	}
	fmt.Println("total:", total)
}

// pay 計算單線得分
func (t *TreeData) payLine(tree ResultTree) int {
	//var sum int
	var tmpPay int
	var minSame bool
	if tree.Value[0] != Wild {
		minSame = true
		goto checkPay
	}

	for i := 1; i < min; i++ {
		if tree.Value[i] != tree.Value[i-1] {
			minSame = true
			break
		}
	}

checkPay:
	if minSame {
		for i := 0; i < len(tree.Value); i++ {
			if tree.Value[i] == Wild {
				continue
			}
			for j := 0; j < len(t.SymbolList); j++ {
				if tree.Value[i] == t.SymbolList[j] {
					tmpPay = t.PaySetting[tree.Count-1][j]
				}
			}
		}
	} else {
		//最小長度都是W
		var wPay int
		for i := min; i < len(tree.Value); i++ {
			//一樣是wild
			if tree.Value[i] == Wild && len(tree.Value) != i+1 {
				continue
			}

			for j := 0; j < len(t.SymbolList); j++ {
				if t.SymbolList[j] == Wild {
					wPay = t.PaySetting[i-1][j]
				}
			}

			for j := 0; j < len(t.SymbolList); j++ {
				if tree.Value[i] == t.SymbolList[j] && tree.Value[i] != Wild {
					tmpPay = t.PaySetting[tree.Count-1][j]
				}
			}

			if wPay > tmpPay {
				tmpPay = wPay
			}

		}
	}

	return tmpPay
}

// CreateTree 創建樹
func (n *Node) CreateTree(values []int) *Node {
	if n == nil || len(values) == 0 {
		return nil
	}

	var noSite bool
reSearch:
	for _, child := range n.Children {
		if child.Site == values[0] {
			noSite = true
			child.CreateTree(values[1:])
			return n
		}
	}

	if !noSite {
		n.AddNode(values[0])
		goto reSearch
	}

	return nil
}

// ReplaceTreeNode 替換樹節點值
func (n *Node) ReplaceTreeNode(nodeMap map[int]string) {
	if n == nil {
		return
	}

	for site, value := range nodeMap {
		if n.Site == site {
			n.Value = value
		}

		for _, child := range n.Children {
			child.ReplaceNode(site, value)
		}
	}

}

// TraverseLengthTree 遍歷固定長度的樹並尋找目標值
func (n *Node) TraverseLengthTree(path []string, sitePath []int) []ResultTree {
	if n == nil {
		return nil
	}

	//去掉根節點
	if n.Site >= 0 {
		path = append(path, n.Value)
		sitePath = append(sitePath, n.Site)
	}

	//至少中min，進入多一輪才能判斷實際擊中數
	realLen := len(path) - 1
	if len(path) > 1 {

		if CheckWild(path, len(path)) != CheckWild(path, len(path)-1) && CheckWild(path, len(path)-1) != AllWild {
			if realLen >= min {
				//fmt.Println(realLen, "  ", path[:realLen])
				return []ResultTree{{Count: realLen, Value: path[0:realLen]}}
			} else {
				return nil
			}
		}
	}

	//全中
	if len(path) == row {
		//fmt.Println(len(path), "  ", path)
		return []ResultTree{{Count: len(path), Value: path}}
	}

	var res []ResultTree
	for _, child := range n.Children {
		res = append(res, child.TraverseLengthTree(path, sitePath)...)
	}
	return res
}

func (n *Node) ReplaceReel(reel [][]string) {

	if n == nil {
		return
	}

	//nodes := []*Node{}
	for i := 0; i < len(reel); i++ {
		//node:=&Node{}
		siteX := i * 10
		for j := 0; j < len(reel[i]); j++ {
			if n.Site == siteX+j {
				n.Value = reel[i][j]
			}
			for _, child := range n.Children {
				child.ReplaceNode(siteX+j, reel[i][j])
			}
		}
	}
}
