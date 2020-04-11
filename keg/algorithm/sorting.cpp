// Example sorting functions
// Accept an array and output a sorted array

#include <iostream> 
#include <map>
#include <vector>
#include <cstdlib>
#include <ctime>
using namespace std;

void print(const std::vector<int>& nums);
void test();

std::vector<int> buble_sort(std::vector<int>& nums);
std::vector<int> selection_sort(std::vector<int>& nums);
std::vector<int> insertion_sort(std::vector<int>& nums);
std::vector<int> quick_sort(std::vector<int>& nums);
std::vector<int> counting_sort(std::vector<int>& nums);
std::vector<int> heap_sort(std::vector<int>& nums);
std::vector<int> merge_sort(std::vector<int>& nums);

int partition(std::vector<int>& nums, int start, int end);
void merge_sort(std::vector<int>& nums, int start, int end);
void merge(std::vector<int>& nums, int start, int mid, int end);
void quick_sort(std::vector<int>& nums, int start, int end);
void max_heapify(std::vector<int>& nums, int p, int N);



int main(){
	std::vector<int> arr{3,7,4,5,8,3,2,1,3,9,6,4,10};
	
	std::cout << "Buble sort:" << endl;
	print(buble_sort(arr));
	std::cout << "Selection sort:" << endl;
	print(selection_sort(arr));
	std::cout << "Insertion sort:" << endl;
	print(insertion_sort(arr));
	std::cout << "Merge sort:" << endl;
	print(merge_sort(arr));
	std::cout << "Quick sort:" << endl;
	print(quick_sort(arr));
	std::cout << "Heap sort:" << endl;
	print(heap_sort(arr));
	std::cout << "Counting sort:" << endl;
	print(counting_sort(arr));
	
	//test();
}

// Buble sort
std::vector<int> buble_sort(std::vector<int>& nums){
	std::vector<int> ret = nums;
	for(int j = ret.size()-1; 0 <= j; j--){
		for(int i = 0; i < j; i++){
			if(ret[i] > ret[i+1])
				std::swap(ret[i], ret[i+1]);
		}
	}
	return ret;
}

// Selection sort
std::vector<int> selection_sort(std::vector<int>& nums){
	std::vector<int> ret = nums;
	for(int i = 0; i < ret.size(); i++){
		int idx = i;
		for(int j = i; j < ret.size(); j++){
			if(ret[j] < ret[idx])
				idx = j;
		}
		std::swap(ret[i], ret[idx]);
	}
	return ret;
}

//Insertion sort
std::vector<int> insertion_sort(std::vector<int>& nums){
	std::vector<int> ret(nums.size(), 0);
	int end = 0;
	for(auto e : nums){
		int i = end-1;
		while(0 <= i && e < ret[i]){
			ret[i+1] = ret[i];
			i--;
		}
		ret[i+1] = e;
		end++;
	}
	return ret;
}

// Merge sort
std::vector<int> merge_sort(std::vector<int>& nums){
	std::vector<int> ret = nums;
	merge_sort(ret, 0, ret.size()-1);
	return ret;
}
void merge_sort(std::vector<int>& nums, int start, int end){
	if(start >= end)
		return;
	int mid = (start+end)/2;
	merge_sort(nums, start, mid);
	merge_sort(nums, mid+1, end);
	merge(nums, start, mid, end);
}
void merge(std::vector<int>& nums, int start, int mid, int end){
	std::vector<int> tmp;
	int i = start, j = mid+1;
	while(i<=mid || j<=end){
		if(i > mid)
			tmp.push_back(nums[j++]);
		else if(j > end)
			tmp.push_back(nums[i++]);
		else
			tmp.push_back((nums[i]<nums[j]) ? nums[i++] : nums[j++]);
	}
	std::copy(tmp.begin(), tmp.end(), nums.begin()+start);
}

// Quick sort
std::vector<int> quick_sort(std::vector<int>& nums){
	std::vector<int> ret = nums;
	quick_sort(ret, 0, ret.size()-1);
	return ret;
}
void quick_sort(std::vector<int>& nums, int start, int end){
	if(start >= end)
		return;
	int p = partition(nums, start, end);
	quick_sort(nums, start, p-1);
	quick_sort(nums, p+1, end);
}
int partition(std::vector<int>& nums, int start, int end){
	int i = start+1;
	for(int j=start+1; j <= end; j++){
		if(nums[j] < nums[start]){
			std::swap(nums[i], nums[j]);
			i++;
		}
	}
	std::swap(nums[start], nums[i-1]);
	return i-1;
}

// Heap sort
std::vector<int> heap_sort(std::vector<int>& nums){
	std::vector<int> ret = nums;
	for(int i = ret.size()/2; 0 <= i; i--)
		max_heapify(ret, i, ret.size());
	
	for(int i = ret.size()-1; 0 <= i; i--){
		std::swap(ret[0], ret[i]);
		max_heapify(ret, 0, i);
	}
	
	return ret;
}

void max_heapify(std::vector<int>& nums, int p, int N){
	if(p >= N/2)//leaf node
		return;
		
	int idx = p, left = p*2+1, right = p*2 + 2;
	if(left < N && nums[left] > nums[idx])
		idx = left;
	if(right < N && nums[right] > nums[idx])
		idx = right;
		
	if(idx != p){
		std::swap(nums[idx], nums[p]);
		max_heapify(nums, idx, N);
	}
}

// Counting sort
std::vector<int> counting_sort(std::vector<int>& nums){
	std::map<int, int> mapNumCount;
	for(auto e : nums)
		mapNumCount[e]++;
	std::vector<int> ret;
	for(auto iter = mapNumCount.begin(); iter != mapNumCount.end(); iter++){
		while(iter->second){
			ret.push_back(iter->first);
			iter->second -= 1;
		}
	}
	return ret;
}

// Print a vector
void print(const std::vector<int>& nums){
	for(auto e : nums)
		std::cout << e << " ";
	std::cout << endl;
}


// Test these sorting functions
void test(){
	std::srand(std::time(NULL));
	
	for(int size = 100; size < 10000; size+=200){
		std::vector<int> arr(size, 0);
		for(int i = 0; i < arr.size(); i++){
			arr[i] = rand();
		}
		std::vector<int> buble_ret = buble_sort(arr);
		std::vector<int> selection_ret = selection_sort(arr);
		std::vector<int> insertion_ret = insertion_sort(arr);
		std::vector<int> merge_ret = merge_sort(arr);
		std::vector<int> quick_ret = quick_sort(arr);
		std::vector<int> heap_ret = heap_sort(arr);
		std::vector<int> counting_ret = counting_sort(arr);
		
		if(buble_ret == selection_ret &&
		   buble_ret == insertion_ret &&
		   buble_ret == merge_ret &&
		   buble_ret == quick_ret &&
		   buble_ret == heap_ret &&
		   buble_ret == counting_ret){
			cout << "Ok" << endl;   	
		}else{
			cout << "Bug exist!" <<endl;
		}
	}
}
