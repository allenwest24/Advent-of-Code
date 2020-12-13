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
	int valid = 0;
	string buff = "";
	vector<string> v;
	int total = 0;
	int peopleInGroup = 0;
	
	// Read the input file into a vector of strings.
	while (getline(file, str)) {
		if (str.empty()) {
			buff += to_string(peopleInGroup);
			v.push_back(buff);
			buff = "";
			peopleInGroup = 0;
		}
		else {
			buff += str;
			peopleInGroup++;
		}
	}

	// Last line will not be an empty line so we have to add it.
	buff += to_string(peopleInGroup);
	v.push_back(buff);
	buff = "";

	// Count unique answers by group.
	string curr;
	int numInGroup = 0;
	int groupSize;
	string tempS = "";
	// Each group.
	for (int ii = 0; ii < v.size(); ii++) {
		tempS = "";
		tempS += v[ii][v[ii].size() - 1];
		groupSize = stoi(tempS);
		// Each answer.
		for (int jj = 0; jj < v[ii].size(); jj++) {
			curr = v[ii][jj];
			// Check if everyone answered.
			for (int kk = jj; kk < v[ii].size(); kk++) {
				if (v[ii][jj] == v[ii][kk]) {
					numInGroup++;
				}
			}
			if (numInGroup == groupSize) {
				total++;
			}
			numInGroup = 0;
		}
	}

	// Print results.
	cout << total << '\n';
	return 0;
}
