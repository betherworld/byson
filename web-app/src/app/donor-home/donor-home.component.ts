import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-donor-home',
  templateUrl: './donor-home.component.html',
  styleUrls: ['./donor-home.component.scss']
})
export class DonorHomeComponent implements OnInit {

  LEVELS = [
    { name: 'Beginner', image: 'bioson_level_beginner.png' },
    { name: 'Intermediate', image: 'bioson_level_intermediate.png' },
    { name: 'Expert', image: 'bioson_level_pro.png' },
  ];

  data = {
    level: 1,
    donated: 7.806,
    bisonCollection: [
      { id: '1', image: 'token_bison1.png' },
    ],
    wildLifeCollection: [
      { id: '1000', image: 'token_bird.png' },
      { id: '1001', image: 'token_bear.png' },
      { id: '1002', image: 'token_fox.png' },
    ],
    statistics2019: {
      co2Reduction: '3T',
      wildLifeIncrease: '100'
    },
    donationAddress: '0xd12Cd8A37F074e7eAF618C996Ff82C666191bd',
  };

  constructor() { }

  ngOnInit() {
  }

}
