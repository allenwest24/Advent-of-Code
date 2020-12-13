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
	string curr;
	int top = 0;
	vector<int> seatIDs;

	// Read the input file into a vector of strings.
	while (getline(file, str)) {
		v.push_back(str);
	}

	// Go through the current seat and find its score.
	for (int ii = 0; ii < v.size(); ii++) {
		int uRow = 128;
		int lRow = 1;
		int uSeat = 8;
		int lSeat = 1;
		for (int jj = 0; jj < v[ii].length(); jj++) {
			curr = v[ii];
			if (curr[jj] == 'F') {
				uRow = lRow + ((uRow - lRow + 1) / 2) - 1;
			}
			else if (curr[jj] == 'B') {
				lRow = lRow + ((uRow - lRow + 1) / 2);
			}
			else if (curr[jj] == 'L') {
				uSeat = lSeat + ((uSeat - lSeat + 1) / 2) - 1;
			}
			else if (curr[jj] == 'R') {
				lSeat = lSeat + ((uSeat - lSeat + 1) / 2);
			}
		}
		
		// Add the seat to our list.
		seatIDs.push_back(((lRow - 1) * 8) + lSeat);
	}


	// Find our seat.
	int oneMore = 0;
	int twoMore = 0;
	for (int jj = 0; jj < seatIDs.size(); jj++) {
		oneMore = 0;
		twoMore = 0;
		for (int kk = 0; kk < seatIDs.size(); kk++) {
			if (seatIDs[jj] == seatIDs[kk] - 2) {
				twoMore = 1;
			}
			if (seatIDs[jj] == seatIDs[kk] - 1) {
				oneMore = 1;
			}
		}
		// Print out the seat not accounted for but surrounded by taken seats.
		if (twoMore && !oneMore) {
			cout << seatIDs[jj] << '\n';
		}
	}
	return 0;
}
