// C++ implmentation of the approach 

#include <bits/stdc++.h> 
using namespace std; 

// Binary tree node consists of data, a 
// pointer to the left child and a 
// pointer to the right child 
struct node { 
	int data; 
	struct node* left; 
	struct node* right; 
}; 

// Function to create new Binary Tree node 
struct node* newNode(int data) 
{ 
	struct node* temp = new struct node; 
	temp->data = data; 
	temp->left = nullptr; 
	temp->right = nullptr; 
	return temp; 
}; 

// Function to return the sum of all 
// the nodes at Kth level using 
// level order traversal 
int sumUtil(struct node* root, int level,int currLevel ,int* currResul){
    
    if (root == nullptr ){
        return 0;
    }
    if (level == currLevel){
        cout << "Level Match = "<< root->data<<endl;
        *currResul += root->data; 
        return 0;
    }
    cout << "Level="<<level<<",curr Level=" << currLevel << ", Curr Res=" <<*currResul<<endl;
    if (root->left != nullptr) {
        sumUtil(root->left,level,currLevel+1,currResul);
    }
    if (root->right != nullptr) {
        sumUtil(root->right,level,currLevel+1,currResul);
    }
    return 0;
}

int sumOfNodesAtNthLevel(struct node* root, int level) {
    int out = 0;
    sumUtil(root,level,0,&out);
    return out;
    
}

// Driver code 
int main() 
{ 
	struct node* root = new struct node; 

	// Tree Construction 
	root = newNode(50); 
	root->left = newNode(30); 
	root->right = newNode(70); 
	root->left->left = newNode(20); 
	root->left->right = newNode(40); 
	root->right->left = newNode(60); 
	int level = 2; 
	int result = sumOfNodesAtNthLevel(root, level); 

	// Printing the result 
	cout << result; 

	return 0; 
} 
