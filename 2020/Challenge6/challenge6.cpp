#include <iostream>
#include <fstream>
#include <string>
#include <vector>

using namespace std;

int main () {
	// Load the file.
	ifstream file("input.txt");

	// Initialize variables.
	string str;
	vector<int> v1 = {1, 1};
	vector<int> v2 = {3, 1};
	vector<int> v3 = {5, 1};
	vector<int> v4 = {7, 1};
	vector<int> v5 = {1, 2};
	vector<vector<int>> v6 = {v1, v2, v3, v4, v5};
	vector<string> v;
	int treesHit = 0;
	int segmentWidth;
	int xCounter = 0;
	int currDown;
	int currRight;

	// Read the input file into a vector of strings.
	while (getline(file, str)) {
		v.push_back(str);
	}

	// Find the width of the given tree map.
	segmentWidth = v[0].length();

	// Set the angle.
	for (int jj = 0; jj < v6.size(); jj++) {
		currRight = v6[jj][0];
		currDown = v6[jj][1];

		// Parse through each line and count trees hit with given angle taken.
		for (int ii = 0; ii < v.size(); ii += currDown) {
			if (v[ii][xCounter % segmentWidth] == '#') {
				treesHit++;
			}
			xCounter += currRight;
		}
		// Print each answer and zero out variables.
		cout << treesHit << '\n';
		treesHit = 0;
		xCounter = 0;
	}

	// Print final instructions.
	cout << 'Multiply the above numbers to get the answer to the puzzle!\n';
	return 0;
}
