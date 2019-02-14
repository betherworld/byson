pragma solidity ^0.5.2;

import "./FundManagerInt.sol";
import "./administrable.sol";

contract Hunters is administrable{
	address fmContrAddress;
	FundManager fm = FundManager(fmContrAddress);

	struct activ {
		uint ETHReward;
		uint RTReward;
		bool active;
		bool claimed;
		bool completed;
		string description;
		address payable user;
	}

	activ[] public activities;
	uint activityCount = 0;
	uint[] proofHashes;
	uint proofCount = 0;

	mapping (uint => activ) ProofToActivity;

	event newActivity (uint activNum);
	event actBooked (uint activNum, address hunt);
	event proofSent (uint activNum, uint hash);
	event proofValidated (uint ETHrew, uint RTrew, uint activNum, address hunt);
	event proofRejected(uint ETHrew, uint RTrew, uint activNum, address hunt);

	constructor (address fundManagerAddr) public {
		fmContrAddress = fundManagerAddr;
	}

	function CreateActivity(uint _rew, uint _ranRew, string memory  _description) public restricted{
		activities.push(activ(_rew,_ranRew,true,false,false,_description,address(0)));
		activityCount++;
		emit newActivity(activityCount - 1);
	}

	function bookActivity(uint _index) public {
		require(activities[_index].active);
		activities[_index].user = msg.sender;
		activities[_index].active = false;
		emit actBooked(_index,activities[_index].user);
	}

	function sendProof(uint _index, uint _hash) public {
		require(activities[_index].user == msg.sender);
		proofHashes.push(_hash);
		proofCount++;
		ProofToActivity[proofCount-1] = activities[_index];
		activities[_index].claimed = true;
		emit proofSent(_index,_hash);
	}

	function validateProof(uint _index, bool approved) public restricted {
		require(!activities[_index].active && activities[_index].claimed);
		if(approved){
			activities[_index].completed = true;
			emit proofValidated(activities[_index].ETHReward,activities[_index].RTReward,_index,activities[_index].user);
			fm.redeemEther(activities[_index].user, activities[_index].ETHReward);
		}else{
			activities[_index].claimed = false;
			activities[_index].active = true;
			activities[_index].user = address(0);
			emit proofRejected(activities[_index].ETHReward, activities[_index].RTReward, _index, activities[_index].user);
		}
	}

	function displayActivity(uint id) public view returns(uint,uint,bool,bool,bool,string memory,address) {
		require(id < activityCount);
		return(activities[id].ETHReward,activities[id].RTReward,activities[id].active,activities[id].claimed,activities[id].completed,activities[id].description,activities[id].user);
	}

	function getActivityCount() public view returns(uint){
		return activityCount;
	}
}