#include <iostream>
#include <fstream>
#include <string>
#include <vector>

using namespace std;

int main ()
{
	// Load the file.
	ifstream file("input.txt");

	// Initialize variables.
	string str;
	vector<int> v;
	int num;

	// Read the input file into a vector of integers.
	while (getline(file, str)) {
		num = stoi(str);
		v.push_back(num);
	}

	// Look for the three numbers that add to 2020 and print their product.
	for (int ii = 0; ii < v.size(); ii++) {
		for (int jj = ii + 1; jj < v.size(); jj++) {
			for (int kk = jj + 1; kk < v.size(); kk++) {
				if (v[ii] + v[jj] + v[kk] == 2020) {
					cout << (v[ii] * v[jj] * v[kk]) << '\n';
					return 0;
				}
			}
		}
	}
	return 0;
}
