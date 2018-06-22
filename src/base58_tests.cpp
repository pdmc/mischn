// Copyright (c) 2011-2017 The Bitcoin Core developers
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

#include <base58.h>
#include <utilstrencodings.h>

#include <boost/test/unit_test.hpp>

extern std::string EncodeBase58(const unsigned char* pbegin, const unsigned char* pend);

// Goal: test low-level base58 encoding functionality
int testEncode()
{
        //UniValue test = tests[idx];
		const std::string test = "123";
		cout << EncodeBase58(test, test + 3) << endl;
        //std::string strTest = test.write();
        //std::vector<unsigned char> sourcedata = ParseHex(test[0].get_str());
        //std::string base58string = test[1].get_str();
}

int main(){
	testEncode();
}
