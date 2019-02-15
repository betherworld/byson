pragma solidity ^0.5.2;

import "../openzeppelin-solidity/contracts/token/ERC721/ERC721.sol";
import "../openzeppelin-solidity/contracts/token/ERC721/ERC721Metadata.sol";
import "../openzeppelin-solidity/contracts/token/ERC721/ERC721Mintable.sol";
import "../openzeppelin-solidity/contracts/token/ERC721/ERC721Enumerable.sol";

contract WildlifeToken is ERC721, ERC721Enumerable, ERC721Metadata, ERC721Mintable{
	address fmContrAddress;

	mapping(uint => uint) private dna;
	uint nextID = 0;	// track which IDs have been used to keep them unique
	uint mod = 10 ** 4;	// number of digits required to encode dna

	// fund manager contract is the only allowed minter
	constructor(address fundManagerAddr) ERC721() ERC721Enumerable() ERC721Metadata("WildlifeToken", "WLT") ERC721Mintable() public {
		fmContrAddress = fundManagerAddr;
		addMinter(fmContrAddress);
	}

	// mint will fail if called by anyone other than fundmanager
	function createNewToken(address donor) public {
		mint(donor, nextID);
		dna[nextID] = _generateNewDna(donor);
	}

	// "random" DNA based on donor address and number of already owned tokens
	// dna encoding:
	// bit position | bit values | meaning
	// 3-2			| 00		 | bison
	// 3-2			| 01		 | bear
	// 3-2			| 10		 | fox
	// 3-2			| 11		 | bird
	// 1-0			| 00		 | variant1
	// 1-0			| 01		 | variant2
	// 1-0			| 10		 | variant3
	// 1-0			| 11		 | variant1
	function _generateNewDna(address donor) private view returns (uint){
		return uint(keccak256(abi.encodePacked(donor, balanceOf(donor)))) % mod;
	}

	function getDna(uint tokenID) public view returns (uint){
		require(tokenID < nextID, "Token doesn't exist!");
		return dna[tokenID];
	}
}