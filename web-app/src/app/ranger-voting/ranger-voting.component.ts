import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-ranger-bonus',
  templateUrl: './ranger-voting.component.html',
  styleUrls: ['./ranger-voting.component.scss']
})
export class RangerVotingComponent implements OnInit {

  projects = [
    { title: 'Build a new well', tokens: 0 },
    { title: 'Extend the local school', tokens: 0 },
  ];

  constructor() { }

  ngOnInit() {
  }

  onVoteClicked() {
    // TODO Send to server
    console.log(this.projects);
  }
}
