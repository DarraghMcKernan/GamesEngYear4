// Authors:
// Shane, Jack & Darragh
// Pair programed in brain waves whoever had an idea took the wheel (Jesus wasnt available)
// Description:
/*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.
Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.
Custom Judge:
The judge will test your solution with the following code:
int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
If all assertions pass, then your solution will be accepted.

Example 1:
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Constraints:
1 <= nums.length <= 3 * 104
-100 <= nums[i] <= 100
nums is sorted in non-decreasing order.
*/
package main

// non decreasing, remove in place, return unique quantity
func removeDupes(numArray []int) (uniqueNums int) {
	if len(numArray) < 1 || len(numArray) > 30000 { //1 <= nums.length <= 3 * 10^4
		return 0
	}

	var lastUnique = -999 //the most recent unique is set to be below -100 which is the minimum value stated to be valid
	uniqueNums = 0

	modifiedArray := make([]int, len(numArray))
	var modArrayIndex = 0

	for index := 0; index < len(numArray); index++ {
		if numArray[index] > 100 || numArray[index] < -100 { //-100 <= nums[i] <= 100   must be non decreasing
			return 0
		} else if index > 0 {
			if numArray[index] < numArray[index-1] { //make sure its non decreasing
				return 0
			}
		}
		if numArray[index] != lastUnique { //ensure the current number isnt the same as the most recent unique
			lastUnique = numArray[index] //update the most recent unique
			uniqueNums++
			modifiedArray[modArrayIndex] = numArray[index] //add current number to the new array without duplicates
			modArrayIndex++
		}
	}

	for index := 0; index < len(numArray); index++ { //modify array to so that the unique elements only are present and in order
		if index < uniqueNums {
			numArray[index] = modifiedArray[index] //repopulate the array to remove duplicates
		} else {
			numArray[index] = -999 //fill the rest of the array with filler
		}
	}
	//fmt.Print(numArray)//display the updated array for testing

	return uniqueNums
}

func main() {
	//arrayOfNums := [8]int{1, 1, 1, 1, 1, 1, 1, 1} //array is in order and has duplicates
	//uniques := removeDupes(arrayOfNums[:])        //give back number of unique elements
	//fmt.Print("\nAmount of unique numbers: ")
	//fmt.Print(uniques)
}
