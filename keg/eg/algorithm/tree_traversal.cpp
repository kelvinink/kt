#include <iostream>
#include <stack>
using namespace std;

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
};

// 前序遍历和中序遍历都是一直往左走，走到无路可走再返回并且切换路径

// 前序遍历
void preorderTraversal(TreeNode* root){
    std::stack<TreeNode*> mstack;
    TreeNode* curr = root;
    while(curr || !mstack.empty()){
        while(curr){
            //访问节点...
            mstack.push(curr);
            curr = curr->left;
        }
        curr = mstack.top()->right;
        mstack.pop();
    }
}

// 中序遍历
void inorderTraversal(TreeNode* root){
    std::stack<TreeNode*> mstack;
    TreeNode* curr = root;
    while(curr || !mstack.empty()){
        while(curr){
            mstack.push(curr);
            curr = curr->left;
        }
        curr = mstack.top()->right;
        mstack.pop();
        //访问节点...
    }
}

// 后续遍历
