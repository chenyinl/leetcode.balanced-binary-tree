/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    re := true;
    defer func(){
        if err:= recover(); err!=nil{
            fmt.Println("panic", err);
            re = false;
        }
    }()
    
    h:=treeH(root);
    fmt.Println(h);
    return re;
    
}

func treeH(root *TreeNode) int {
    if(root == nil){
        return 0;
    }
    l:= treeH(root.Left);
    r:= treeH(root.Right);
    if(l-r > 1 || l-r < -1){
        panic(nil);
    }
    return max(l,r)+1;
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
