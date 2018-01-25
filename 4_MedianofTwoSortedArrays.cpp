class Solution {
public:
	double findMedianSortedArrays(vector<int>& num1, vector<int>& num2) {
		int lennum1 = num1.size();
		int lennum2 = num2.size();
		double mid = 0;
		vector<int> mergevec;
		vector<int> i = 0, j = 0;
		while (i < lennum1 && j < lennum2){
			if (num1[i] < num2[j]){
				mergevec.push_back(num1[i]);
				i++;
			}
			else{
				mergevec.push_back(num2[j]);
				j++;
			}
		}
		if (i == lennum1){
			mergevec.insert(mergevec.end(), num2.begin()+j, num2.end());

		}
		else
			mergevec.insert(mergevec.end(), num1.begin()+i, num1.end());
		if ((lennum1 + lennum2) % 2)
			mid = mergvec[(lennum1 + lennum2 + 1)/2];
		else
			mid = (mergvec[(lennum1 + lennum2)/2] + mergvec[(lennum1 + lennum2)/2 + 1])/2 + (mergvec[(lennum1 + lennum2)/2] + mergvec[(lennum1 + lennum2)/2 + 1])%2;
		return mid;
};




//来自leetcode官方解答

class Solution {
public:
	double median(vector<int>& num1, int l1, vector<int>& num2, int l2)
	{
	    int imin = 0, imax = l1, halflen = (l1 + l2 + 1) / 2;
	    while (imin <= imax)
	    {
	        int i = (imin + imax) / 2;
	        int j = halflen - i;
	        if(i < imax && num2[j-1] > num1[i])
	        {
	            imin++;
	        }
	        else if (i > imin && num1[i-1] > num2[j])
	        {
	            imax--;
	        }
	        else
	        {
	            int maxleft = 0;
	            if(i == 0){maxleft = num2[j-1];}
	            else if(j ==0){maxleft =num1[i-1];}
	            else {maxleft =max(num1[i-1], num2[j-1]);}
	            if ((l1+l2) %2 == 1)
	                return maxleft;
	            int minright = 0;
	            if(i == l1){minright = num2[j];}
	            else if(j ==l2){minright =num1[i];}
	            else {minright =min(num1[i], num2[j]);}
	            return (maxleft+minright)/2.0;

	        }
	    }
	}
	double findMedianSortedArrays(vector<int>& num1, vector<int>& num2)
	{
	    int m = num1.size();
	    int n = num2.size();
	    return m < n ? median(num1, m, num2, n) : median(num2, n, num1, m);
	}
};