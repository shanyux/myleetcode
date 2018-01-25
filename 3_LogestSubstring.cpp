/* class Solution {
public:
	int lengthOfLongestSubstring(string s) {
		unordered_set<char> charset;
		int longest = 0;
		int firstlong = 0;
		int lastlong = 0;
		pair<set<int>::iterator, bool> pr;
		for (int i = 0; i < s.size(); i++){
			pr = charset.insert(s[i]);
			if(!pr.second){
				firstlong = pr->first - charset.begin() + 1;
				lastlong = charset.end() - pr->first;
				if (fistlong > lastlong){
					longest = firstlong > longest ? firstlong : longest;
				 	charset.erase(pr->first, charset.end());
				}
				else{
					longest = lastlong > longest ? lastlong : longest;
					charset.erase(charset.begin(), pr->first);
				}
			}

		}
		longest = charset.size() > longest ? charset.size() : longest;
		return longest;
	}
}
unordered_set 迭代器是正向迭代器，只能p++，++p, 不可以相减
看链接：
http://blog.csdn.net/qq_23100787/article/details/51388163
*/

/*
class Solution {
public:
	int lengthOfLongestSubstring(string s) {
		vector<int> charvec(256, -1);
		vector<int>::iterator left = charvec.begin();
		vector<int>::iterator iter;
		int longest = 0;
		for (int i = 0; i != s.size(); i++) {
			iter = find(left, charvec.end(), s[i]);
			if (iter != charvec.end()){
				left = iter;
				longest = longest > iter - left + 1 ? longest : iter - left + 1;
			}
			charvec.push_back(s[i]);
		}
		longest = longest > charvec.end() - left ? longest : charvec.end() - left;
		return longest;
	}
};




class Solution {
public:
	int lengthOfLongestSubstring(string s) {
		vector<int> charvec;
        charvec.push_back(s[0]);
		vector<int>::iterator left = charvec.begin();
		vector<int>::iterator iter;
		int longest = 0;
		for (int i = 1; i != s.length(); i++) {
            iter = find(left, charvec.end(), s[i]);
			if (iter != charvec.end()){				
				longest = longest > (iter++ - left) ? longest : (iter++ - left);
                left = iter;
			}
			charvec.push_back(s[i]);
		}
		longest = longest > (charvec.end() - left) ? longest : (charvec.end() - left);
		return longest;
	}
};

*/

class Solution {
public:
	int lengthOfLongestSubstring(string s) {
		vector<int> dic(256,-1);//记录的是位置
		int maxlen = 0, start = -1;
		for (int i = 0; i != s.length(); i++) {
			if (dic[s[i]] > start) //这个字符的位置在start的后面，说明将会重复所以更换start
				start = dic[s[i]];
			dic[s[i]]= i;
			maxlen = max(maxlen, dic[s[i]] - start);
		}
		return maxlen;
	}
};



//另外看到一个用unordered_map的
class Solution {
public:
	int lengthOfLongestSubstring(string s) {
    	unordered_map<char, int> mymap;
	    int maxlen = 0;
	    int start = -1, loc = -1;
	    for (int p = 0; p != s.length(); p++)
	    {
	        loc = mymap.find(s[p]) == mymap.end()? -1 : mymap.find(s[p])->second;
	        start = loc < start? start : loc;//位置
	        mymap[s[p]] = p;
	        maxlen = max(maxlen, mymap[s[p]] - start);
	    }
	    return maxlen;
	}
};

//底下这个是o^2
class Solution {
public:
	int lengthOfLongestSubstring(string s) {
		vector<char> myset;
	    int maxlen = 0;
	    int left = 0, right = 0, p = 0;
	    for (int i = 0; i < s.length(); i++){
	    	for(p = left; p != myset.size(); p++){
	    		if (myset[p] == s[i]) {
	    			left = p;
	    		}
	    	}
			myset.push_back(s[i]);
			right = myset.size() - 1;
			maxlen = max(maxlen, right - left);
	    }
    	return maxlen;
	}
};