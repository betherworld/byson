pragma solidity ^0.5.2;

import "./administrable.sol";
import "./FundManagerInt.sol";
import "../openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";
import "../openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";
import "../openzeppelin-solidity/contracts/token/ERC20/ERC20Mintable.sol";
import "../openzeppelin-solidity/contracts/token/ERC20/ERC20Burnable.sol";

contract HunterToken is administrable, ERC20, ERC20Detailed, ERC20Mintable, ERC20Burnable{
	address huntContrAddress;	// address of contract responsible for hunter/gatherer interactions
	address fmContrAddress;		// address of contract holding all funds and responsible for distribution
	FundManager fm;

	uint proposalsCount = 0;
	uint winningProposalId;
	Proposal[] proposals;

	struct Proposal{
		string name;
		string description;
		uint ballot;
	}

	constructor(address fundManagerAddr) ERC20Burnable() ERC20Mintable() ERC20Detailed("HunterToken", "HTT", 0) ERC20() public {
		fmContrAddress = fundManagerAddr;
		fm = FundManager(fmContrAddress);
	}

	// admins can add new project proposals for community development
	function generateProposal(string memory name, string memory description) public restricted{
		proposals.push(Proposal(name, description, 0));
		proposalsCount++;
	}

	// use this only to iterate through all proposals, proposal ids are not static!
	function getProposal(uint256 propId) public view returns(string memory, string memory, uint) {
		return (proposals[propId].name, proposals[propId].description, proposals[propId].ballot);
	}

	function getPropCount() public view returns(uint) {
		return proposalsCount;
	}

	// hunters who have earned hunter tokens can place them in the corresponding ballot,
	// multiple vote disbursments are possible, if there are not enough tokens an error is raised
	function vote(uint propId, uint HTnum) public {
		require(propId < proposalsCount);
		burn(HTnum);
		proposals[propId].ballot += HTnum;
		if(proposals[propId].ballot > proposals[winningProposalId].ballot){
			winningProposalId = propId;
		}
	}

	// admins can determine the end of the voting period, the projected project costs are disbursed
	// to specified account
	function endVoting(address payable projAccount, uint projCost) public restricted returns (string memory, string memory, uint){
		string memory n = proposals[winningProposalId].name;
		string memory d = proposals[winningProposalId].description;
		uint b = proposals[winningProposalId].ballot;

		//delete winner
		proposals[winningProposalId].name = proposals[proposalsCount-1].name;
		proposals[winningProposalId].description = proposals[proposalsCount-1].description;
		proposals[winningProposalId].ballot = proposals[proposalsCount-1].ballot;
		delete proposals[proposalsCount-1];
		proposalsCount--;

		fm.payCommProj(projAccount, projCost, n);

		return (n, d, b);
	}

	// one-time executable functions to set priviliged contract addresses
	function setHunterContrAddr(address addr) public {
		require(huntContrAddress == address(0));
		huntContrAddress = addr;
		addMinter(addr);
	}
}