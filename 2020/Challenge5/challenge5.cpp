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
	vector<string> v;
	int treesHit = 0;
	int segmentWidth;
	int xCounter = 0;

	// Read the input file into a vector of strings.
	while (getline(file, str)) {
		v.push_back(str);
	}

	// Find the width of the given tree map.
	segmentWidth = v[0].length();

	// Parse through each line and count each time going down 1, right 3 results in a tree being hit.
	for (int ii = 0; ii < v.size(); ii++) {
		if (v[ii][xCounter % segmentWidth] == '#') {
			treesHit++;
		}
		xCounter += 3;
	}

	// Print out final count.
	cout << treesHit << '\n';
	return 0;
}
