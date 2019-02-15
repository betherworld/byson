pragma solidity ^0.5.2;

contract claim {

    struct activ {
        uint ETHReward;
        uint RTReward;
        bool active;
        bool claimed;
        bool completed;
        string description;
        address user;
    }

    activ[] public activities;
    address WWF;
    uint activityCount = 0;
    uint[] proofHashes;
    uint proofCount = 0;

    modifier contValidity(uint _time, bool _stat) {
        require(_stat == true && _time <= now);
        _;
    }

    modifier WWFOnly () {
        require(msg.sender == WWF);
        _;
    }

    mapping (uint => activ) ProofToActivity;

    event newActivity (uint activNum);
    event actBooked (uint activNum, address hunt);
    event proofSent (uint activNum, uint hash);
    event proofValidated (uint ETHrew, uint RTrew, uint activNum, address hunt);
    event proofRejected(uint ETHrew, uint RTrew, uint activNum, address hunt);

    function CreateActivity(uint _rew, uint _ranRew, string memory  _description) public {
        //require(WWF == msg.sender);
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

    function validateProof(uint _index, bool approved) public {
        require(!activities[_index].active && activities[_index].claimed);
        if(approved){
            activities[_index].completed = true;
            emit proofValidated(activities[_index].ETHReward,activities[_index].RTReward,_index,activities[_index].user);
            // call contract pay
        }else{
            activities[_index].claimed = false;
            activities[_index].active = true;
            activities[_index].user = address(0);
            emit proofRejected(activities[_index].ETHReward, activities[_index].RTReward, _index, activities[_index].user);
        }
    }

    function displayActivity(uint _i) public view returns(uint,uint,bool,bool,bool,string memory,address) {
        require(_i < activityCount);
        return(activities[_i].ETHReward,activities[_i].RTReward,activities[_i].active,activities[_i].claimed,activities[_i].completed,activities[_i].description,activities[_i].user);
    }

    function getActivityCount() public view returns(uint){
        return activityCount;
    }
}
