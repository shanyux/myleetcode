class Solution{
public:
	vector<int> twoSum(vector<int>& nums, int target){
		unordered_map<int,int> hash;
		vector<int> result;
		for (int i = 0; i < nums.size(); i++){
			if (hash.find(target - nums[i]) != nums.end()){
				result.push_back(hash[target - num[i]]);
				result.push_back(i);
				return result;
			}
			else
				hash[num[i]] = i;
		}
		return result;


	}
};


class Solution{
public:
	vector<int> twoSum(vector<int>& nums, int target){
		bool flag = false;
		vector<int> result;
		for (int i = 0; i < nums.size(); i++){
			for(int j = i; j < nums.size(); j++){
				if(nums[i] + nums[j] == target){
					flag = true;
					break;
				}				
			}
		}
		if(flag){
			result.push_back(i);
			result.push_back(j);
		}
		return result;
	}
};